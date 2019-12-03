package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringTypeDescription_Config(t *testing.T) {
	sd := StringTypeDescription{}

	c := sd.Config()

	assert.Equal(t, 0, len(c))
}

func TestStringTypeDescription_New(t *testing.T) {
	sd := StringTypeDescription{}

	s, e := sd.New(nil, []string{})

	assert.NotNil(t, s)
	assert.NoError(t, e)
}

func TestStringType_ValidateEmptyString(t *testing.T) {
	s := stringType{}

	e := s.Validate(nil, "")

	assert.NoError(t, e)
}
