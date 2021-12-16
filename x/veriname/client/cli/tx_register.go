package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/dmolik/veriname/x/veriname/types"
)

var _ = strconv.Itoa(0)

func CmdRegister() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register [alias] [user] [kind] [target] [payload]",
		Short: "Broadcast message register",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAlias := args[0]
			argUser := args[1]
			argKind := args[2]
			argTarget := args[3]
			argPayload := args[4]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRegister(
				clientCtx.GetFromAddress().String(),
				argAlias,
				argUser,
				argKind,
				argTarget,
				argPayload,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
