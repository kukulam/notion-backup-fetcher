package nbf

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebug(t *testing.T) {
	// given
	buf := new(bytes.Buffer)
	logger := NewLogger(WithStream(buf))

	expectedLog := "[debug] testing msg"

	// when
	logger.debug("testing msg")

	// then
	log := buf.String()
	assert.Equal(t, expectedLog, log)
}

func TestFatal(t *testing.T) {
	// given
	buf := new(bytes.Buffer)
	logger := NewLogger(WithStream(buf))

	expectedLog := "[fatal] testing msg"

	// when
	logger.fatal("testing msg")

	// then
	log := buf.String()
	assert.Equal(t, expectedLog, log)
}