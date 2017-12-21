package template

import (
	"reflect"
	"testing"

	"github.com/gotestyourself/gotestyourself/assert"
)

var defaults = map[string]string{
	"FOO": "first",
	"BAR": "",
}

func defaultMapping(name string) (string, bool) {
	val, ok := defaults[name]
	return val, ok
}

func TestEscaped(t *testing.T) {
	result, err := Substitute("$${foo}", defaultMapping)
	assert.Check(t, err)
	assert.Check(t, is.Equal("${foo}", result))
}

func TestInvalid(t *testing.T) {
	invalidTemplates := []string{
		"${",
		"$}",
		"${}",
		"${ }",
		"${ foo}",
		"${foo }",
		"${foo!}",
	}

	for _, template := range invalidTemplates {
		_, err := Substitute(template, defaultMapping)
		assert.ErrorContains(t, err, "Invalid template")
	}
}

func TestNoValueNoDefault(t *testing.T) {
	for _, template := range []string{"This ${missing} var", "This ${BAR} var"} {
		result, err := Substitute(template, defaultMapping)
		assert.Check(t, err)
		assert.Check(t, is.Equal("This  var", result))
	}
}

func TestValueNoDefault(t *testing.T) {
	for _, template := range []string{"This $FOO var", "This ${FOO} var"} {
		result, err := Substitute(template, defaultMapping)
		assert.Check(t, err)
		assert.Check(t, is.Equal("This first var", result))
	}
}

func TestNoValueWithDefault(t *testing.T) {
	for _, template := range []string{"ok ${missing:-def}", "ok ${missing-def}"} {
		result, err := Substitute(template, defaultMapping)
		assert.Check(t, err)
		assert.Check(t, is.Equal("ok def", result))
	}
}

func TestEmptyValueWithSoftDefault(t *testing.T) {
	result, err := Substitute("ok ${BAR:-def}", defaultMapping)
	assert.Check(t, err)
	assert.Check(t, is.Equal("ok def", result))
}

func TestEmptyValueWithHardDefault(t *testing.T) {
	result, err := Substitute("ok ${BAR-def}", defaultMapping)
	assert.Check(t, err)
	assert.Check(t, is.Equal("ok ", result))
}

func TestNonAlphanumericDefault(t *testing.T) {
	result, err := Substitute("ok ${BAR:-/non:-alphanumeric}", defaultMapping)
	assert.Check(t, err)
	assert.Check(t, is.Equal("ok /non:-alphanumeric", result))
}
