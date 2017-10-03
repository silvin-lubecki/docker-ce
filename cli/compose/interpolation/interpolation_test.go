package interpolation

import (
	"testing"

	"strconv"

	"github.com/gotestyourself/gotestyourself/env"
	"github.com/stretchr/testify/assert"
)

var defaults = map[string]string{
	"USER":  "jenny",
	"FOO":   "bar",
	"count": "5",
}

func defaultMapping(name string) (string, bool) {
	val, ok := defaults[name]
	return val, ok
}

func TestInterpolate(t *testing.T) {
	services := map[string]interface{}{
		"servicea": map[string]interface{}{
			"image":   "example:${USER}",
			"volumes": []interface{}{"$FOO:/target"},
			"logging": map[string]interface{}{
				"driver": "${FOO}",
				"options": map[string]interface{}{
					"user": "$USER",
				},
			},
		},
	}
	expected := map[string]interface{}{
		"servicea": map[string]interface{}{
			"image":   "example:jenny",
			"volumes": []interface{}{"bar:/target"},
			"logging": map[string]interface{}{
				"driver": "bar",
				"options": map[string]interface{}{
					"user": "jenny",
				},
			},
		},
	}
	result, err := Interpolate(services, Options{
		SectionName: "service",
		LookupValue: defaultMapping,
	})
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestInvalidInterpolation(t *testing.T) {
	services := map[string]interface{}{
		"servicea": map[string]interface{}{
			"image": "${",
		},
	}
	_, err := Interpolate(services, Options{
		SectionName: "service",
		LookupValue: defaultMapping,
	})
	assert.EqualError(t, err, `Invalid interpolation format for "image" option in service "servicea": "${". You may need to escape any $ with another $.`)
}

func TestInterpolateWithDefaults(t *testing.T) {
	defer env.Patch(t, "FOO", "BARZ")()

	config := map[string]interface{}{
		"networks": map[string]interface{}{
			"foo": "thing_${FOO}",
		},
	}
	expected := map[string]interface{}{
		"networks": map[string]interface{}{
			"foo": "thing_BARZ",
		},
	}
	result, err := Interpolate(config, Options{})
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
}

func TestInterpolateWithCast(t *testing.T) {
	config := map[string]interface{}{
		"foo": map[string]interface{}{
			"replicas": "$count",
		},
	}
	toInt := func(value string) (interface{}, error) {
		return strconv.Atoi(value)
	}
	result, err := Interpolate(config, Options{
		LookupValue:     defaultMapping,
		TypeCastMapping: map[Path]Cast{NewPath("*", "replicas"): toInt},
	})
	assert.NoError(t, err)
	expected := map[string]interface{}{
		"foo": map[string]interface{}{
			"replicas": 5,
		},
	}
	assert.Equal(t, expected, result)
}

func TestPathMatches(t *testing.T) {
	var testcases = []struct {
		doc      string
		path     Path
		pattern  Path
		expected bool
	}{
		{
			doc:     "pattern too short",
			path:    NewPath("one", "two", "three"),
			pattern: NewPath("one", "two"),
		},
		{
			doc:     "pattern too long",
			path:    NewPath("one", "two"),
			pattern: NewPath("one", "two", "three"),
		},
		{
			doc:     "pattern mismatch",
			path:    NewPath("one", "three", "two"),
			pattern: NewPath("one", "two", "three"),
		},
		{
			doc:     "pattern mismatch with match-all part",
			path:    NewPath("one", "three", "two"),
			pattern: NewPath(PathMatchAll, "two", "three"),
		},
		{
			doc:      "pattern match with match-all part",
			path:     NewPath("one", "two", "three"),
			pattern:  NewPath("one", "*", "three"),
			expected: true,
		},
		{
			doc:      "pattern match",
			path:     NewPath("one", "two", "three"),
			pattern:  NewPath("one", "two", "three"),
			expected: true,
		},
	}
	for _, testcase := range testcases {
		assert.Equal(t, testcase.expected, testcase.path.matches(testcase.pattern))
	}
}
