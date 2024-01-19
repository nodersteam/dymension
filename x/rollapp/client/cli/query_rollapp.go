package cli

import (
	"context"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/dymensionxyz/dymension/x/rollapp/types"
	gType "github.com/gogo/protobuf/types"
)

func CmdListRollapp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Query all rollapps currently registered in the hub",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			flagSet := cmd.Flags()
			argHeight, _ := flagSet.GetInt64(flags.FlagHeight)
			val := &gType.UInt64Value{Value: uint64(argHeight)}
			params := &types.QueryAllRollappRequest{
				Pagination: pageReq,
				Height:     val,
			}

			res, err := queryClient.RollappAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowRollapp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show [rollapp-id]",
		Short: "Query the rollapp associated with the specified rollapp-id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			argRollappId := args[0]

			params := &types.QueryGetRollappRequest{
				RollappId: argRollappId,
			}

			res, err := queryClient.Rollapp(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
