package main

import (
	"context"
	"strings"

	"github.com/urfave/cli"
)

func Quick(c *cli.Context) error {
	client := GetClient(c)

	if !c.Args().Present() {
		return CommandFailed
	}

	client.QuickCommand(
		context.Background(),
		strings.Join(c.Args(), " "),
	)

	return Sync(c)
}
