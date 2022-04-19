package nbf

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli"
)

func TestFetchCommand_DebugLogs(t *testing.T) {
	// given
	c := &cli.Context{}

	// when
	err := FetchCommand(c)
	
	// then
	require.Nil(t, err)
}
