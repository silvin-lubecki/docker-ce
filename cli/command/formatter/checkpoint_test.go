package formatter

import (
	"bytes"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/gotestyourself/gotestyourself/assert"
	is "github.com/gotestyourself/gotestyourself/assert/cmp"
)

func TestCheckpointContextFormatWrite(t *testing.T) {
	cases := []struct {
		context  Context
		expected string
	}{
		{
			Context{Format: NewCheckpointFormat(defaultCheckpointFormat)},
			`CHECKPOINT NAME
checkpoint-1
checkpoint-2
checkpoint-3
`,
		},
		{
			Context{Format: NewCheckpointFormat("{{.Name}}")},
			`checkpoint-1
checkpoint-2
checkpoint-3
`,
		},
		{
			Context{Format: NewCheckpointFormat("{{.Name}}:")},
			`checkpoint-1:
checkpoint-2:
checkpoint-3:
`,
		},
	}

	checkpoints := []types.Checkpoint{
		{"checkpoint-1"},
		{"checkpoint-2"},
		{"checkpoint-3"},
	}
	for _, testcase := range cases {
		out := bytes.NewBufferString("")
		testcase.context.Output = out
		err := CheckpointWrite(testcase.context, checkpoints)
		if err != nil {
			assert.Check(t, is.ErrorContains(err, ""), testcase.expected)
		} else {
			assert.Check(t, is.Equal(out.String(), testcase.expected))
		}
	}
}
