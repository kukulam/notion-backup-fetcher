package nbf

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"
)

var logger = NewLogger()

func TestFetchCommand_DebugLogs(t *testing.T) {
	// given
	c := &cli.Context{}

	// when
	err := FetchCommand(c, logger)
	
	// then
	require.Nil(t, err)
}
