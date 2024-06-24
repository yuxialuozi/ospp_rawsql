package pulgin

import (
	"fmt"
	"github.com/cloudwego/cwgo/pkg/curd/template"
	"github.com/cloudwego/hertz/cmd/hz/meta"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/thriftgo/plugin"
	"io"
	"os"
	"ospp_rawsql/config"
	"ospp_rawsql/pkg/curd/db/mysql/codegen"
	"ospp_rawsql/pkg/curd/extract"
	"ospp_rawsql/pkg/curd/parse"
)

type thriftGoPlugin struct {
	req       *plugin.Request
	ModelArgs *config.ModelArgument
}

func thriftPluginRun() int {
	plu := &thriftGoPlugin{}

	if err := plu.handleRequest(); err != nil {
		logs.Errorf("handle request failed: %s", err.Error())
		return meta.PluginError
	}

	if err := plu.parseArgs(); err != nil {
		logs.Errorf("parse args failed: %s", err.Error())
		return meta.PluginError
	}

	tfUsedInfo := &extract.ThriftUsedInfo{
		Req:       plu.req,
		ModelArgs: plu.ModelArgs,
	}
	rawStructs, err := tfUsedInfo.ParseThriftIdl()
	if err != nil {
		logs.Errorf("parse thrift idl failed: %s", err.Error())
		return meta.PluginError
	}

	operations, err := parse.HandleOperations(rawStructs)
	if err != nil {
		logs.Error(err.Error())
		return meta.PluginError
	}

	methodRenders := codegen.HandleCodegen(operations)

	generated, err := plu.buildResponse(rawStructs, methodRenders, tfUsedInfo)
	if err != nil {
		logs.Error(err.Error())
		return meta.PluginError
	}

	res := &plugin.Response{
		Contents: generated,
	}
	if err = response(res); err != nil {
		logs.Error(err.Error())
		return meta.PluginError
	}

	if plu.ModelArgs.GenBase {
		if err = generateBaseMysqlFile(plu.ModelArgs.DaoDir, tfUsedInfo.ImportPaths, codegen.HandleBaseCodegen()); err != nil {
			return meta.PluginError
		}
	}

	return 0
}

func (plu *thriftGoPlugin) handleRequest() error {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("read request failed: %s", err.Error())
	}

	req, err := plugin.UnmarshalRequest(data)
	if err != nil {
		return fmt.Errorf("unmarshal request failed: %s", err.Error())
	}

	plu.req = req
	return nil
}

func (plu *thriftGoPlugin) buildResponse(structs []*extract.IdlExtractStruct, methodRenders [][]*template.MethodRender,
	info *extract.ThriftUsedInfo,
) (result []*plugin.Generated, err error) {
	result = make([]*plugin.Generated, len(structs))
	return
}

func (plu *thriftGoPlugin) parseArgs() error {
	if plu.req == nil {
		return fmt.Errorf("request is nil")
	}
	args := new(config.ModelArgument)
	if err := args.Unpack(plu.req.PluginParameters); err != nil {
		logs.Errorf("unpack args failed: %s", err.Error())
		return err
	}
	plu.ModelArgs = args
	return nil
}

func response(res *plugin.Response) error {
	data, err := plugin.MarshalResponse(res)
	if err != nil {
		return fmt.Errorf("marshal response failed: %s", err.Error())
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		return fmt.Errorf("write response failed: %s", err.Error())
	}
	return nil
}
