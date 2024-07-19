package static

import (
	"github.com/urfave/cli/v2"
	"ospp_rawsql/config"
	"ospp_rawsql/meta"
	consts "ospp_rawsql/pkg/const"
	"ospp_rawsql/pkg/db"
)

func Init() *cli.App {
	globalArgs := config.GetGlobalArgs()
	verboseFlag := cli.BoolFlag{Name: "verbose,vv", Usage: "turn on verbose mode"}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = meta.Name
	app.Usage = AppUsage
	app.Version = meta.Version
	// The default separator for multiple parameters is modified to ";"
	app.SliceFlagSeparator = consts.Comma

	// global flags
	app.Flags = []cli.Flag{
		&verboseFlag,
	}

	// Commands
	app.Commands = []*cli.Command{
		{
			Name:  DbName,
			Usage: DbUsage,
			Flags: dbFlags(),
			Action: func(c *cli.Context) error {
				if err := globalArgs.DbArgument.ParseCli(c); err != nil {
					return err
				}
				return db.Db(globalArgs.DbArgument)
			},
		},
	}
	return app
}

const (
	AppUsage = "All in one tools for CloudWeGo"

	DbName = "db"

	DbUsage = `generate db model

Examples:
  # Generate db model code
  cwgo db --name mysql --idl {{path/to/IDL_file.thrift}} `
)
