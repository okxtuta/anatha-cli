package cli

import (
	"bufio"
	"fmt"
	denom "github.com/okxtuta/anatha-cli/utils"
	"github.com/okxtuta/anatha-cli/x/distribution/internal/types"
	"github.com/okxtuta/go-anatha/client"
	"github.com/okxtuta/go-anatha/client/context"
	"github.com/okxtuta/go-anatha/client/flags"
	"github.com/okxtuta/go-anatha/codec"
	sdk "github.com/okxtuta/go-anatha/types"
	"github.com/okxtuta/go-anatha/x/auth"
	"github.com/okxtuta/go-anatha/x/auth/client/utils"
	"github.com/spf13/cobra"
)

func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	distributionTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%S transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	distributionTxCmd.AddCommand(flags.PostCommands(
		GetCmdWithdrawNameReward(cdc),
		GetCmdWithdrawValidatorReward(cdc),
		GetCmdDepositSavings(cdc),
		GetCmdWithdrawSavings(cdc),
		GetCmdWithdrawSavingsInterest(cdc),
	)...)

	return distributionTxCmd
}

func GetCmdWithdrawNameReward(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw-name-reward",
		Short: "Withdraw pending Name rewards",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgWithdrawNameReward(cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdWithdrawValidatorReward(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw-validator-reward",
		Short: "Withdraw pending Name rewards",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgWithdrawValidatorReward(sdk.ValAddress(cliCtx.GetFromAddress()))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDepositSavings(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deposit-savings [amount]",
		Short: "Deposit savings",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			amount, err := denom.ParseAndConvertCoins(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgDepositSavings(cliCtx.GetFromAddress(), amount)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdWithdrawSavings(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw-savings",
		Short: "Withdraw savings deposit",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgWithdrawSavings(cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdWithdrawSavingsInterest(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw-savings-interest",
		Short: "Withdraw savings interest",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgWithdrawSavingsInterest(cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
