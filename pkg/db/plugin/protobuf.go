package plugin

import (
	"google.golang.org/protobuf/compiler/protogen"
	"ospp_rawsql/config"
)

type Plugin struct {
	*protogen.Plugin
	dbargs *config.DbArgument
}

func ProtoPluginRun() int {

	return 0
}
