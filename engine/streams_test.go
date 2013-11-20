package engine

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestOutputAddString(t *testing.T) {
	var testInputs = [][2]string{
		{
			"hello, world!",
			"hello, world!",
		},

		{
			"One\nTwo\nThree",
			"Three",
		},

		{
			"",
			"",
		},

		{
			"A line\nThen another nl-terminated line\n",
			"Then another nl-terminated line",
		},

		{
			"A line followed by an empty line\n\n",
			"",
		},
	}
	for _, testData := range testInputs {
		input := testData[0]
		expectedOutput := testData[1]
		o := NewOutput()
		var output string
		if err := o.AddString(&output); err != nil {
			t.Error(err)
		}
		if n, err := o.Write([]byte(input)); err != nil {
			t.Error(err)
		} else if n != len(input) {
			t.Errorf("Expected %d, got %d", len(input), n)
		}
		o.Close()
		if output != expectedOutput {
			t.Errorf("Last line is not stored as return string.\nInput:   '%s'\nExpected: '%s'\nGot:       '%s'", input, expectedOutput, output)
		}
	}
}

type sentinelWriteCloser struct {
	calledWrite bool
	calledClose bool
}

func (w *sentinelWriteCloser) Write(p []byte) (int, error) {
	w.calledWrite = true
	return len(p), nil
}

func (w *sentinelWriteCloser) Close() error {
	w.calledClose = true
	return nil
}

func TestOutputAddClose(t *testing.T) {
	o := NewOutput()
	var s sentinelWriteCloser
	if err := o.Add(&s); err != nil {
		t.Fatal(err)
	}
	if err := o.Close(); err != nil {
		t.Fatal(err)
	}
	// Write data after the output is closed.
	// Write should succeed, but no destination should receive it.
	if _, err := o.Write([]byte("foo bar")); err != nil {
		t.Fatal(err)
	}
	if !s.calledClose {
		t.Fatal("Output.Close() didn't close the destination")
	}
}

func TestOutputAddPipe(t *testing.T) {
	var testInputs = []string{
		"hello, world!",
		"One\nTwo\nThree",
		"",
		"A line\nThen another nl-terminated line\n",
		"A line followed by an empty line\n\n",
	}
	for _, input := range testInputs {
		expectedOutput := input
		o := NewOutput()
		r, err := o.AddPipe()
		if err != nil {
			t.Fatal(err)
		}
		go func(o *Output) {
			if n, err := o.Write([]byte(input)); err != nil {
				t.Error(err)
			} else if n != len(input) {
				t.Errorf("Expected %d, got %d", len(input), n)
			}
			if err := o.Close(); err != nil {
				t.Error(err)
			}
		}(o)
		output, err := ioutil.ReadAll(r)
		if err != nil {
			t.Fatal(err)
		}
		if string(output) != expectedOutput {
			t.Errorf("Last line is not stored as return string.\nExpected: '%s'\nGot:       '%s'", expectedOutput, output)
		}
	}
}

func TestTail(t *testing.T) {
	var tests = make(map[string][][]string)
	tests["hello, world!"] = [][]string{
		{},
		{"hello, world!"},
		{"hello, world!"},
		{"hello, world!"},
	}
	tests["One\nTwo\nThree"] = [][]string{
		{},
		{"Three"},
		{"Two", "Three"},
		{"One", "Two", "Three"},
	}
	for input, outputs := range tests {
		for n, expectedOutput := range outputs {
			var output []string
			Tail(strings.NewReader(input), n, &output)
			if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expectedOutput) {
				t.Errorf("Tail n=%d returned wrong result.\nExpected: '%s'\nGot     : '%s'", expectedOutput, output)
			}
		}
	}
}

func TestOutputAddTail(t *testing.T) {
	var tests = make(map[string][][]string)
	tests["hello, world!"] = [][]string{
		{},
		{"hello, world!"},
		{"hello, world!"},
		{"hello, world!"},
	}
	tests["One\nTwo\nThree"] = [][]string{
		{},
		{"Three"},
		{"Two", "Three"},
		{"One", "Two", "Three"},
	}
	for input, outputs := range tests {
		for n, expectedOutput := range outputs {
			o := NewOutput()
			var output []string
			if err := o.AddTail(&output, n); err != nil {
				t.Error(err)
			}
			if n, err := o.Write([]byte(input)); err != nil {
				t.Error(err)
			} else if n != len(input) {
				t.Errorf("Expected %d, got %d", len(input), n)
			}
			o.Close()
			if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", expectedOutput) {
				t.Errorf("Tail(%d) returned wrong result.\nExpected: %v\nGot:      %v", n, expectedOutput, output)
			}
		}
	}
}

func lastLine(txt string) string {
	scanner := bufio.NewScanner(strings.NewReader(txt))
	var lastLine string
	for scanner.Scan() {
		lastLine = scanner.Text()
	}
	return lastLine
}

func TestOutputAdd(t *testing.T) {
	o := NewOutput()
	b := &bytes.Buffer{}
	o.Add(b)
	input := "hello, world!"
	if n, err := o.Write([]byte(input)); err != nil {
		t.Fatal(err)
	} else if n != len(input) {
		t.Fatalf("Expected %d, got %d", len(input), n)
	}
	if output := b.String(); output != input {
		t.Fatal("Received wrong data from Add.\nExpected: '%s'\nGot:     '%s'", input, output)
	}
}

func TestInputAddEmpty(t *testing.T) {
	i := NewInput()
	var b bytes.Buffer
	if err := i.Add(&b); err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadAll(i)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) > 0 {
		t.Fatalf("Read from empty input shoul yield no data")
	}
}

func TestInputAddTwo(t *testing.T) {
	i := NewInput()
	var b1 bytes.Buffer
	// First add should succeed
	if err := i.Add(&b1); err != nil {
		t.Fatal(err)
	}
	var b2 bytes.Buffer
	// Second add should fail
	if err := i.Add(&b2); err == nil {
		t.Fatalf("Adding a second source should return an error")
	}
}

func TestInputAddNotEmpty(t *testing.T) {
	i := NewInput()
	b := bytes.NewBufferString("hello world\nabc")
	expectedResult := b.String()
	i.Add(b)
	result, err := ioutil.ReadAll(i)
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != expectedResult {
		t.Fatalf("Expected: %v\nReceived: %v", expectedResult, result)
	}
}
