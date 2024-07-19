package plugin

import (
	thriftgo_plugin "github.com/cloudwego/thriftgo/plugin"
	"io/ioutil"
	"ospp_rawsql/pkg/db/thrift"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	data, err := ioutil.ReadFile("D:\\project\\go\\ospp_rawsql\\pkg\\db\\thrift\\idl\\test.out")
	if err != nil {
		t.Fatal(err)
	}

	req, err := thriftgo_plugin.UnmarshalRequest(data)
	if err != nil {
		t.Fatal(err)
	}

	plu := new(thriftGoPlugin)
	plu.req = req

	// 解析参数
	err = plu.parseArgs()
	if err != nil {
		t.Fatalf("解析参数失败: %s", err.Error())
	}

	req.OutputPath = ""

	sqlContent, err := thrift.GenerateSQLFromAST(plu.req.AST)
	if err != nil {
		t.Fatalf("生成SQL代码失败: %s", err.Error())
	}

	// 生成到 sql 文件夹
	filePath := filepath.Join("sql", "generated_file.sql")

	res := &thriftgo_plugin.Response{
		Contents: []*thriftgo_plugin.Generated{
			{
				Content: sqlContent,
				Name:    stringPtr(filePath),
			},
		},
	}

	// 序列化并写出响应
	thriftgo_plugin.MarshalResponse(res)

}
