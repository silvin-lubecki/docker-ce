// +build linux

package signal

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func buildTestBinary(t *testing.T, tmpdir string, prefix string) (string, string) {
	tmpDir, err := ioutil.TempDir(tmpdir, prefix)
	require.NoError(t, err)
	exePath := tmpDir + "/" + prefix
	wd, _ := os.Getwd()
	testHelperCode := wd + "/testfiles/main.go"
	cmd := exec.Command("go", "build", "-o", exePath, testHelperCode)
	err = cmd.Run()
	require.NoError(t, err)
	return exePath, tmpDir
}

func TestTrap(t *testing.T) {
	var sigmap = []struct {
		name     string
		signal   os.Signal
		multiple bool
	}{
		{"TERM", syscall.SIGTERM, false},
		{"QUIT", syscall.SIGQUIT, true},
		{"INT", os.Interrupt, false},
		{"TERM", syscall.SIGTERM, true},
		{"INT", os.Interrupt, true},
	}
	exePath, tmpDir := buildTestBinary(t, "", "main")
	defer os.RemoveAll(tmpDir)

	for _, v := range sigmap {
		cmd := exec.Command(exePath)
		cmd.Env = append(os.Environ(), fmt.Sprintf("SIGNAL_TYPE=%s", v.name))
		if v.multiple {
			cmd.Env = append(cmd.Env, "IF_MULTIPLE=1")
		}
		err := cmd.Start()
		require.NoError(t, err)
		err = cmd.Wait()
		if e, ok := err.(*exec.ExitError); ok {
			code := e.Sys().(syscall.WaitStatus).ExitStatus()
			if v.multiple {
				assert.Equal(t, 128+int(v.signal.(syscall.Signal)), code)
			} else {
				assert.Equal(t, 99, code)
			}
			continue
		}
		t.Fatal("process didn't end with any error")
	}

}

func TestDumpStacks(t *testing.T) {
	directory, err := ioutil.TempDir("", "test-dump-tasks")
	assert.NoError(t, err)
	defer os.RemoveAll(directory)
	dumpPath, err := DumpStacks(directory)
	assert.NoError(t, err)
	readFile, _ := ioutil.ReadFile(dumpPath)
	fileData := string(readFile)
	assert.Contains(t, fileData, "goroutine")
}

func TestDumpStacksWithEmptyInput(t *testing.T) {
	path, err := DumpStacks("")
	assert.NoError(t, err)
	assert.Equal(t, os.Stderr.Name(), path)
}
