package cli

import (
	"bufio"
	"fmt"
	"github.com/okxtuta/go-anatha/client/context"
	"github.com/okxtuta/go-anatha/x/auth"
	"github.com/okxtuta/go-anatha/x/auth/client/utils"

	"github.com/spf13/cobra"

	"github.com/okxtuta/anatha-cli/x/treasury/internal/types"
	"github.com/okxtuta/go-anatha/client"
	"github.com/okxtuta/go-anatha/client/flags"
	"github.com/okxtuta/go-anatha/codec"

	sdk "github.com/okxtuta/go-anatha/types"
)

func GetCmdOperator(cdc *codec.Codec) *cobra.Command {
	operatorTxCmd := &cobra.Command{
		Use:                        "operator",
		Short:                      fmt.Sprintf("%s transactions subcommands", "operator"),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	operatorTxCmd.AddCommand(flags.PostCommands(
		GetCmdAddOpeator(cdc),
		GetCmdRemoveOpeator(cdc),
	)...)

	return operatorTxCmd
}

func GetCmdAddOpeator(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "add [address]",
		Short: "Add a Treasury Distribution Operator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			operator, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgAddOperator(cliCtx.GetFromAddress(), operator)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdRemoveOpeator(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "remove [address]",
		Short: "Remove a Treasury Distribution Operator",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			operator, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveOperator(cliCtx.GetFromAddress(), operator)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
