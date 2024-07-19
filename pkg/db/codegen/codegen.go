package codegen

import (
	"ospp_rawsql/config"
)

func HandleCodegen(c *config.DbArgument) {

	processSQLStatements(c.Queris)

}
