package command

import (
	"github.com/urfave/cli/v2"

	"Component-Manager/module"
)

func Version(ctx *cli.Context) error {
	cli.ShowVersion(ctx)
	var _, err = module.CheckRemoteVersion(ctx, false)
	if err != nil {
		return err
	}
	return nil
}
