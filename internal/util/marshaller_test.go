package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func TestMarshal(t *testing.T) {
	test := testStruct{
		Name:  "test",
		Count: 1,
	}

	output, err := Marshal(test)
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"test","count":1}`, string(output))
}

func TestUnmarshal(t *testing.T) {
	input := []byte(`{"name":"test","count":1}`)

	output, err := Unmarshal(input)
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"name": "test", "count": 1.0}, output)
}
