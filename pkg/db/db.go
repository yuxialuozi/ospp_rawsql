package db

import (
	"errors"
	"log"
	"os"
	"ospp_rawsql/pkg/db/codegen"
	"ospp_rawsql/pkg/db/plugin"

	"ospp_rawsql/config"
	"ospp_rawsql/pkg/const"
	"ospp_rawsql/pkg/utils"
	"path/filepath"
)

func Db(c *config.DbArgument) error {
	if err := check(c); err != nil {
		return err
	}

	if c.IdlPath != "" {
		if err := plugin.MysqlTriggerPlugin(c); err != nil {
			return err
		}

		ModelDir := "D:\\project\\go\\cwgo_test\\biz\\doc\\model" // 这里替换成实际获取的值

		// 删除目录
		if err := os.RemoveAll(ModelDir); err != nil {
			log.Fatalf("删除 biz 目录失败: %s", err.Error())
		}

		log.Printf("成功删除目录: %s", ModelDir)
	}

	codegen.HandleCodegen(c)

	utils.ReplaceThriftVersion()

	return nil
}

func check(c *config.DbArgument) (err error) {

	if c.Name == "" {
		c.Name = consts.MySQL
	}
	if c.Name != consts.MySQL {
		return errors.New("db name not supported")
	}

	if c.IdlPath == "" {
		c.IfIdl = false
		c.IdlPath = "D:/project/go/ospp_rawsql/pkg/db/thrift/idl/test.thrift"
	}

	c.IdlType, err = utils.GetIdlType(c.IdlPath)
	if err != nil {
		return err
	}

	c.OutDir, err = filepath.Abs(c.OutDir)
	if err != nil {
		return err
	}

	if c.ModelDir == "" {
		c.ModelDir = consts.DefaultDocModelOutDir
	}
	c.ModelDir, err = filepath.Abs(filepath.Join(c.OutDir, c.ModelDir))
	if err != nil {
		return err
	}

	return nil
}
