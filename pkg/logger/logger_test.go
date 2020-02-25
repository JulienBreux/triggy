package logger_test

import (
	"testing"

	"github.com/julienbreux/triggy/pkg/logger"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestParseLevel(t *testing.T) {
	level := "info"

	assert.Equal(t, logger.ParseLevel(level), zerolog.InfoLevel)
}
