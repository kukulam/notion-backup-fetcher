package nbf

import (
	"fmt"

	"github.com/urfave/cli"
)


func FetchCommand(c *cli.Context, logger *Logger) error {
	rootPageUrl := c.String("url")
	logger.debug(fmt.Sprintf("starting root url: %s", rootPageUrl))

	return nil
}
