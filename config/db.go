package config

import (
	"fmt"
	"github.com/cloudwego/hertz/cmd/hz/util"
	"github.com/urfave/cli/v2"
	consts "ospp_rawsql/pkg/const"
	"strings"
)

type DbArgument struct {
	IdlPath       string
	IdlType       string
	IfIdl         bool
	Schema        string
	Queris        string
	GoMod         string
	Name          string
	Verbose       bool
	PackagePrefix string
	ThriftOptions []string
	ModelDir      string
	OutDir        string
}

func NewDbArgument() *DbArgument {
	return &DbArgument{}
}

func (d *DbArgument) ParseCli(ctx *cli.Context) error {
	d.IdlPath = ctx.String(consts.IDLPath)
	d.Schema = ctx.String(consts.SchemaPath)
	d.Queris = ctx.String(consts.QuerisPath)
	d.GoMod = ctx.String(consts.Module)
	d.Name = ctx.String(consts.Name)
	return nil
}

func (d *DbArgument) Unpack(data []string) error {
	err := util.UnpackArgs(data, d)
	if err != nil {
		return fmt.Errorf("unpack argument failed: %s", err)
	}
	return nil
}

func (d *DbArgument) Pack() ([]string, error) {
	data, err := util.PackArgs(d)
	if err != nil {
		return nil, fmt.Errorf("pack argument failed: %s", err)
	}
	return data, nil
}

func (d *DbArgument) GetThriftgoOptions(prefix string) (string, error) {
	d.ThriftOptions = append(d.ThriftOptions, "package_prefix="+prefix)
	gas := "go:" + strings.Join(d.ThriftOptions, ",")
	return gas, nil
}
