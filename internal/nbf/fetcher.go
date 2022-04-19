package nbf

import (
	"fmt"

	"github.com/urfave/cli"
)

var logger = NewLogger()

func FetchCommand(c *cli.Context) error {
	rootPageUrl := c.String("url")
	logger.debug(fmt.Sprintf("fetching root url: %s", rootPageUrl))

	return nil
}
