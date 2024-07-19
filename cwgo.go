package main

import (
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"os"
	"ospp_rawsql/cmd/static"
	pluginDb "ospp_rawsql/pkg/db/plugin"
)

func main() {

	// run cwgo as mysql plugin mode
	pluginDb.MysqlPluginMode()

	cli := static.Init()
	err := cli.Run(os.Args)
	if err != nil {
		logs.Errorf("%v\n", err)
	}

}
