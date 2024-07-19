package static

import (
	"github.com/urfave/cli/v2"
	consts "ospp_rawsql/pkg/const"
)

func dbFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{Name: consts.IDLPath, Usage: "Specify the IDL file path. (.thrift or .proto)"},
		&cli.StringFlag{Name: consts.SchemaPath, Usage: "Specify the SQL_schema file path."},
		&cli.StringFlag{Name: consts.QuerisPath, Usage: "Specify the SQL_queris file path."},
		&cli.StringFlag{Name: consts.Module, Aliases: []string{"mod"}, Usage: "Specify the Go module name to generate go.mod."},
		&cli.StringFlag{Name: consts.Name, Usage: "Specifies the name of the relational database used to generate the code. The default is mysql. (Currently, only mysql is supported)"},
	}
}
