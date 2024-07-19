package plugin

import (
	"fmt"
	"github.com/cloudwego/thriftgo/parser"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/cloudwego/hertz/cmd/hz/meta"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	thriftgo_plugin "github.com/cloudwego/thriftgo/plugin"
	"ospp_rawsql/config"
	"ospp_rawsql/pkg/db/thrift"
)

type thriftGoPlugin struct {
	req    *thriftgo_plugin.Request
	dbargs *config.DbArgument
}

// ThriftPluginRun 是插件的入口函数
func ThriftPluginRun() int {
	plu := &thriftGoPlugin{}

	// 处理请求
	err := plu.handleRequest()
	if err != nil {
		logs.Errorf("处理请求失败: %s", err.Error())
		return meta.PluginError
	}

	// 解析参数
	err = plu.parseArgs()
	if err != nil {
		logs.Errorf("解析参数失败: %s", err.Error())
		return meta.PluginError
	}

	// 生成 SQL 内容
	schemaGenerated := generateSQLContent(plu.dbargs.IdlPath, plu.req.AST)

	querisGenerated := generateSQLQuerisContent(plu.dbargs.IdlPath, plu.req.AST)

	// 构造响应对象
	res := &thriftgo_plugin.Response{
		Contents: append(schemaGenerated, querisGenerated...),
	}

	// 发送响应
	if err := response(res); err != nil {
		logs.Error(err.Error())
		return meta.PluginError
	}

	return 0
}

// handleRequest 读取标准输入中的数据并反序列化请求对象
func (p *thriftGoPlugin) handleRequest() error {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("读取请求失败: %s", err.Error())
	}

	req, err := thriftgo_plugin.UnmarshalRequest(data)
	if err != nil {
		return fmt.Errorf("反序列化请求失败: %s", err.Error())
	}
	p.req = req

	//todo:remove
	buf, err := thriftgo_plugin.MarshalRequest(req)
	err = os.MkdirAll("D:\\project\\go\\ospp_rawsql\\pkg\\db\\thrift\\idl\\test.out", os.FileMode(0755))
	err = os.WriteFile("D:\\project\\go\\ospp_rawsql\\pkg\\db\\thrift\\idl\\test.out", buf, os.FileMode(0755))

	return nil
}

// parseArgs 从请求中解析插件参数
func (plu *thriftGoPlugin) parseArgs() error {
	if plu.req == nil {
		return fmt.Errorf("请求为空")
	}
	args := new(config.DbArgument)
	if err := args.Unpack(plu.req.PluginParameters); err != nil {
		logs.Errorf("解析参数失败: %s", err.Error())
		return err
	}
	plu.dbargs = args
	return nil
}

// stringPtr 返回字符串指针
func stringPtr(s string) *string {
	return &s
}

// response 将响应对象序列化并写入标准输出
func response(res *thriftgo_plugin.Response) error {
	data, err := thriftgo_plugin.MarshalResponse(res)
	if err != nil {
		return fmt.Errorf("marshal response failed: %s", err.Error())
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		return fmt.Errorf("write response failed: %s", err.Error())
	}
	return nil
}

// generateSQLContent 生成建表 SQL 内容并返回一个包含生成结果的切片
func generateSQLContent(idlPath string, ast *parser.Thrift) []*thriftgo_plugin.Generated {
	var sqlContent string
	var err error

	if idlPath != "" {
		sqlContent, err = thrift.GenerateSQLFromAST(ast)
		if err != nil {
			log.Printf("生成SQL代码失败: %s", err.Error())
			// 这里可以根据实际情况处理错误，比如记录日志或返回特定的错误结构
		}
	}

	// 生成到 sql 文件夹
	filePath := filepath.Join("sql", "generated_file.sql")
	generated := []*thriftgo_plugin.Generated{
		{
			Content: sqlContent,
			Name:    stringPtr(filePath),
		},
	}

	return generated
}

func generateSQLQuerisContent(idlPath string, ast *parser.Thrift) []*thriftgo_plugin.Generated {
	var sqlContent string
	var err error

	if idlPath != "" {
		sqlContent, err = thrift.GenerateSQLQuerisFromAST(ast)
		if err != nil {
			log.Printf("生成SQL代码失败: %s", err.Error())
			// 这里可以根据实际情况处理错误，比如记录日志或返回特定的错误结构
		}
	}

	// 生成到 sql 文件夹
	filePath := filepath.Join("sql", "user_queris.sql")
	generated := []*thriftgo_plugin.Generated{
		{
			Content: sqlContent,
			Name:    stringPtr(filePath),
		},
	}

	return generated
}
