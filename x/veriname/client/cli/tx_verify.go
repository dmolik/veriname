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

func CmdVerify() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "verify [alias] [user] [kind] [target]",
		Short: "Broadcast message verify",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argAlias := args[0]
			argUser := args[1]
			argKind := args[2]
			argTarget := args[3]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgVerify(
				clientCtx.GetFromAddress().String(),
				argAlias,
				argUser,
				argKind,
				argTarget,
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
