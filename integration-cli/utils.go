package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/docker/docker/vendor/src/code.google.com/p/go/src/pkg/archive/tar"
)

func getExitCode(err error) (int, error) {
	exitCode := 0
	if exiterr, ok := err.(*exec.ExitError); ok {
		if procExit := exiterr.Sys().(syscall.WaitStatus); ok {
			return procExit.ExitStatus(), nil
		}
	}
	return exitCode, fmt.Errorf("failed to get exit code")
}

func runCommandWithOutput(cmd *exec.Cmd) (output string, exitCode int, err error) {
	exitCode = 0
	out, err := cmd.CombinedOutput()
	if err != nil {
		var exiterr error
		if exitCode, exiterr = getExitCode(err); exiterr != nil {
			// TODO: Fix this so we check the error's text.
			// we've failed to retrieve exit code, so we set it to 127
			exitCode = 127
		}
	}
	output = string(out)
	return
}

func runCommandWithStdoutStderr(cmd *exec.Cmd) (stdout string, stderr string, exitCode int, err error) {
	var (
		stderrBuffer, stdoutBuffer bytes.Buffer
	)
	exitCode = 0
	cmd.Stderr = &stderrBuffer
	cmd.Stdout = &stdoutBuffer
	err = cmd.Run()

	if err != nil {
		var exiterr error
		if exitCode, exiterr = getExitCode(err); exiterr != nil {
			// TODO: Fix this so we check the error's text.
			// we've failed to retrieve exit code, so we set it to 127
			exitCode = 127
		}
	}
	stdout = stdoutBuffer.String()
	stderr = stderrBuffer.String()
	return
}

func runCommand(cmd *exec.Cmd) (exitCode int, err error) {
	exitCode = 0
	err = cmd.Run()
	if err != nil {
		var exiterr error
		if exitCode, exiterr = getExitCode(err); exiterr != nil {
			// TODO: Fix this so we check the error's text.
			// we've failed to retrieve exit code, so we set it to 127
			exitCode = 127
		}
	}
	return
}

func startCommand(cmd *exec.Cmd) (exitCode int, err error) {
	exitCode = 0
	err = cmd.Start()
	if err != nil {
		var exiterr error
		if exitCode, exiterr = getExitCode(err); exiterr != nil {
			// TODO: Fix this so we check the error's text.
			// we've failed to retrieve exit code, so we set it to 127
			exitCode = 127
		}
	}
	return
}

func logDone(message string) {
	fmt.Printf("[PASSED]: %s\n", message)
}

func stripTrailingCharacters(target string) string {
	target = strings.Trim(target, "\n")
	target = strings.Trim(target, " ")
	return target
}

func errorOut(err error, t *testing.T, message string) {
	if err != nil {
		t.Fatal(message)
	}
}

func errorOutOnNonNilError(err error, t *testing.T, message string) {
	if err == nil {
		t.Fatalf(message)
	}
}

func nLines(s string) int {
	return strings.Count(s, "\n")
}

func unmarshalJSON(data []byte, result interface{}) error {
	err := json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	return nil
}

func deepEqual(expected interface{}, result interface{}) bool {
	return reflect.DeepEqual(result, expected)
}

func convertSliceOfStringsToMap(input []string) map[string]struct{} {
	output := make(map[string]struct{})
	for _, v := range input {
		output[v] = struct{}{}
	}
	return output
}

func waitForContainer(contId string, args ...string) error {
	args = append([]string{"run", "--name", contId}, args...)
	cmd := exec.Command(dockerBinary, args...)
	if _, err := runCommand(cmd); err != nil {
		return err
	}

	if err := waitRun(contId); err != nil {
		return err
	}

	return nil
}

func waitRun(contId string) error {
	after := time.After(5 * time.Second)

	for {
		cmd := exec.Command(dockerBinary, "inspect", "-f", "{{.State.Running}}", contId)
		out, _, err := runCommandWithOutput(cmd)
		if err != nil {
			return fmt.Errorf("error executing docker inspect: %v", err)
		}

		if strings.Contains(out, "true") {
			break
		}

		select {
		case <-after:
			return fmt.Errorf("container did not come up in time")
		default:
		}

		time.Sleep(100 * time.Millisecond)
	}

	return nil
}

func compareDirectoryEntries(e1 []os.FileInfo, e2 []os.FileInfo) error {
	var (
		e1Entries = make(map[string]struct{})
		e2Entries = make(map[string]struct{})
	)
	for _, e := range e1 {
		e1Entries[e.Name()] = struct{}{}
	}
	for _, e := range e2 {
		e2Entries[e.Name()] = struct{}{}
	}
	if !reflect.DeepEqual(e1Entries, e2Entries) {
		return fmt.Errorf("entries differ")
	}
	return nil
}

func ListTar(f io.Reader) ([]string, error) {
	tr := tar.NewReader(f)
	var entries []string

	for {
		th, err := tr.Next()
		if err == io.EOF {
			// end of tar archive
			return entries, nil
		}
		if err != nil {
			return entries, err
		}
		entries = append(entries, th.Name)
	}
}

type FileServer struct {
	*httptest.Server
}

func fileServer(files map[string]string) (*FileServer, error) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		if filePath, found := files[r.URL.Path]; found {
			http.ServeFile(w, r, filePath)
		} else {
			http.Error(w, http.StatusText(404), 404)
		}
	}

	for _, file := range files {
		if _, err := os.Stat(file); err != nil {
			return nil, err
		}
	}
	server := httptest.NewServer(handler)
	return &FileServer{
		Server: server,
	}, nil
}

func copyWithCP(source, target string) error {
	copyCmd := exec.Command("cp", "-rp", source, target)
	out, exitCode, err := runCommandWithOutput(copyCmd)
	if err != nil || exitCode != 0 {
		return fmt.Errorf("failed to copy: error: %q ,output: %q", err, out)
	}
	return nil
}
