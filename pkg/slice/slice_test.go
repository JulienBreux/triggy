package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/julienbreux/triggy/pkg/slice"
)

func TestContains(t *testing.T) {
	tests := []struct {
		s      string
		ss     []string
		result bool
	}{
		{"rabbitmq", []string{"rabbitmq"}, true},
		{"redis", []string{"rabbitmq"}, false},
	}

	for _, test := range tests {
		assert.True(t, test.result == slice.Contains(test.s, test.ss))
	}
}
