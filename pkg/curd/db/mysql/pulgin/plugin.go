package pulgin

import (
	"fmt"
	"github.com/cloudwego/hertz/cmd/hz/meta"
	"os"
	"os/exec"
	"ospp_rawsql/config"
	"ospp_rawsql/pkg/common/utils"
	"ospp_rawsql/pkg/consts"
	"strings"
)

func MysqlTriggerPlugin(c *config.ModelArgument) error {
	cmd, err := buildPluginCmd(c)
}

func buildPluginCmd(args *config.ModelArgument) (*exec.Cmd, error) {
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
		return nil, err
	}
	cmd := &exec.Cmd{
		Path: path,
	}

	if args.IdlType == meta.IdlThrift {
		os.Setenv(consts.CwgoDocPluginMode, consts.ThriftCwgoDocPluginName)

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
			"-r",
			args.IdlPath,
		)
	} else {
		cmd.Args = append(cmd.Args, meta.TpCompilerProto)

		var isFindIdl bool

		var importPaths []string

		for _, inc := range args.ProtoSearchPath {
			idlParser := parser.NewProtoParser()

			if !isFindIdl {
				_, importPaths, err = idlParser.GetDependentFilePaths(inc, args.IdlPath)
				if err == nil {
					isFindIdl = true
				}

			}

			cmd.Args = append(cmd.Args, "-I", inc)
		}

		cmd.Args = append(cmd.Args, "--go_out="+args.ModelDir)
		for _, kv := range args.ProtocOptions {
			cmd.Args = append(cmd.Args, "--"+kv)
		}

		cmd.Args = append(cmd.Args, importPaths...)
		cmd.Args = append(cmd.Args, args.IdlPath)
	}

	return cmd, err
}
