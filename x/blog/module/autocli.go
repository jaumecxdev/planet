package blog

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "planet/api/planet/blog"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "PostAll",
					Use:       "list-post",
					Short:     "List all post",
				},
				{
					RpcMethod:      "Post",
					Use:            "show-post [id]",
					Short:          "Shows a post by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "SentPostAll",
					Use:       "list-sent-post",
					Short:     "List all sentPost",
				},
				{
					RpcMethod:      "SentPost",
					Use:            "show-sent-post [id]",
					Short:          "Shows a sentPost by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "TimedoutPostAll",
					Use:       "list-timedout-post",
					Short:     "List all timedoutPost",
				},
				{
					RpcMethod:      "TimedoutPost",
					Use:            "show-timedout-post [id]",
					Short:          "Shows a timedoutPost by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
