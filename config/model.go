/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"fmt"
	"github.com/cloudwego/hertz/cmd/hz/util"
	"github.com/urfave/cli/v2"
	"ospp_rawsql/pkg/consts"
	"strings"
)

type ModelArgument struct {
	DSN               string
	Type              string
	Tables            []string
	ExcludeTables     []string
	OnlyModel         bool
	OutPath           string
	OutFile           string
	WithUnitTest      bool
	ModelPkgName      string
	FieldNullable     bool
	FieldSignable     bool
	FieldWithIndexTag bool
	FieldWithTypeTag  bool
	SQLDir            string
	GenBase           bool
	ThriftOptions     []string
}

func NewModelArgument() *ModelArgument {
	return &ModelArgument{
		OutPath: consts.DefaultDbOutDir,
		OutFile: consts.DefaultDbOutFile,
	}
}

func (c *ModelArgument) ParseCli(ctx *cli.Context) error {
	c.DSN = ctx.String(consts.DSN)
	c.Type = strings.ToLower(ctx.String(consts.DBType))
	c.Tables = ctx.StringSlice(consts.Tables)
	c.ExcludeTables = ctx.StringSlice(consts.ExcludeTables)
	c.OnlyModel = ctx.Bool(consts.OnlyModel)
	c.OutPath = ctx.String(consts.OutDir)
	c.OutFile = ctx.String(consts.OutFile)
	c.WithUnitTest = ctx.Bool(consts.UnitTest)
	c.ModelPkgName = ctx.String(consts.ModelPkgName)
	c.FieldNullable = ctx.Bool(consts.Nullable)
	c.FieldSignable = ctx.Bool(consts.Signable)
	c.FieldWithIndexTag = ctx.Bool(consts.IndexTag)
	c.FieldWithTypeTag = ctx.Bool(consts.TypeTag)
	c.SQLDir = ctx.String(consts.SQLDir)
	c.IdlPath = ctx.String(consts.IDLPath)
	c.ThriftOptions = ctx.StringSlice(consts.ThriftGo)
	return nil
}

func (d *ModelArgument) Unpack(data []string) error {
	err := util.UnpackArgs(data, d)
	if err != nil {
		return fmt.Errorf("unpack argument failed: %s", err)
	}
	return nil
}

func (d *ModelArgument) Pack() ([]string, error) {
	data, err := util.PackArgs(d)
	if err != nil {
		return nil, fmt.Errorf("pack argument failed: %s", err)
	}
	return data, nil
}

func (d *ModelArgument) GetThriftgoOptions(prefix string) (string, error) {
	d.ThriftOptions = append(d.ThriftOptions, "package_prefix="+prefix)
	gas := "go:" + strings.Join(d.ThriftOptions, ",")
	return gas, nil
}

func (d *DocArgument) modelUnpack(data []string) error {
	err := util.UnpackArgs(data, d)
	if err != nil {
		return fmt.Errorf("unpack argument failed: %s", err)
	}
	return nil
}
