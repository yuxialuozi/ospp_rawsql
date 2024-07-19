package plugin

import (
	"fmt"
	"github.com/cloudwego/cwgo/pkg/common/utils"
	"github.com/cloudwego/hertz/cmd/hz/meta"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"os"
	"os/exec"
	"ospp_rawsql/config"
	consts "ospp_rawsql/pkg/const"
	"strings"
)

func MysqlTriggerPlugin(c *config.DbArgument) error {
	cmd, err := buildPluginCmd(c)
	if err != nil {
		return fmt.Errorf("build plugin command failed: %v", err)
	}

	buf, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("plugin cwgo-db returns error: %v, output: %s", err, string(buf))
	}

	// If len(buf) != 0, the plugin returned the log.
	if len(buf) != 0 {
		fmt.Println(string(buf))
	}

	return nil
}

func buildPluginCmd(args *config.DbArgument) (*exec.Cmd, error) {
	exe, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("failed to detect current executable, err: %v", err)
	}

	argPacks, err := args.Pack()
	if err != nil {
		return nil, err
	}
	kas := strings.Join(argPacks, ",")

	path, err := utils.LookupTool(args.IdlType)
	if err != nil {
		return nil, fmt.Errorf("lookup tool failed: %v", err)
	}
	cmd := &exec.Cmd{
		Path: path,
	}

	if args.IdlType == meta.IdlThrift {
		os.Setenv(consts.CwgoDbPluginMode, consts.ThriftCwgoDbPluginName)

		thriftOpt, err := args.GetThriftgoOptions(args.PackagePrefix)
		if err != nil {
			return nil, err
		}

		cmd.Args = append(cmd.Args, meta.TpCompilerThrift)
		if args.Verbose {
			cmd.Args = append(cmd.Args, "-v")
		}
		cmd.Args = append(cmd.Args,
			"-o", args.ModelDir,
			"-p", "cwgo-db="+exe+":"+kas,
			"-g", thriftOpt,
			"-r", args.IdlPath,
		)
	} else {
		// 处理其他 IDL 类型，如 Protobuf
		// 设置环境变量和插件模式
		os.Setenv(consts.CwgoDbPluginMode, consts.ProtoCwgoDbPluginName)

		// 添加编译器命令
		cmd.Args = append(cmd.Args, meta.TpCompilerProto)

		// 添加protoc插件相关参数
		cmd.Args = append(cmd.Args,
			"--plugin=protoc-gen-cwgo="+exe,
			"--cwgo_out="+args.OutDir,
			"--cwgo_opt="+kas,
		)

		// 处理插件参数
		for _, p := range args.ThriftOptions {
			pluginParams := strings.Split(p, ":")
			if len(pluginParams) != 3 {
				logs.Warnf("Failed to get the correct protoc plugin parameters for %s. "+
					"Please specify the protoc plugin in the form of \"plugin_name:options:out_dir\"", p)
				os.Exit(1)
			}
			// pluginParams[0] -> plugin name, pluginParams[1] -> plugin options, pluginParams[2] -> out_dir
			cmd.Args = append(cmd.Args,
				fmt.Sprintf("--%s_out=%s", pluginParams[0], pluginParams[2]),
				fmt.Sprintf("--%s_opt=%s", pluginParams[0], pluginParams[1]),
			)
		}

	}

	return cmd, err
}

func MysqlPluginMode() {
	mode := os.Getenv(consts.CwgoDbPluginMode)
	if len(os.Args) <= 1 && mode != "" {
		switch mode {
		case consts.ThriftCwgoDbPluginName:

			os.Exit(ThriftPluginRun())
		case consts.ProtoCwgoDbPluginName:
			os.Exit(ProtoPluginRun())

		}
	}
}
