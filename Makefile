all:
	thriftgo -o D:\project\go\cwgo_test -p db=D:\drive\go-respository\bin\ospp_rawsql.exe:IdlPath=D:\project\go\ospp_rawsql\pkg\db\thrift\idl\User.thrift,IdlType=thrift,OutDir=D:\project\go\cwgo_test -g go:package_prefix= -r D:\project\go\ospp_rawsql\pkg\db\thrift\idl\User.thrift
