package jsonrpc

import (
	"github.com/AndrienkoAleksandr/che-theia-terminal-plugin/api/model"
	"github.com/eclipse/che/agents/go-agents/core/jsonrpc"
)

// Constants that represent RPC methods identifiers.
const (
	CreateMethod = "create"
	CheckMethod    = "check"
	ResizeMethod = "resize"
)

// RPCRoutes defines process jsonrpc routes.
var RPCRoutes = jsonrpc.RoutesGroup{
	Name: "Json-rpc MachineExec Routes",
	Items: []jsonrpc.Route{
		{
			Method: CreateMethod,
			Decode: jsonrpc.FactoryDec(func() interface{} { return &model.MachineExec{} }),
			Handle: jsonRpcCreateExec,
		},
		{
			Method: CheckMethod,
			Decode: jsonrpc.FactoryDec(func() interface{} { return &IdParam{} }),
			Handle: jsonRpcCheckExec,
		},
		{
			Method: ResizeMethod,
			Decode: jsonrpc.FactoryDec(func() interface{} { return &ResizeParam{} }),
			Handle: jsonrpc.HandleRet(jsonRpcResizeExec),
		},
	},
}

