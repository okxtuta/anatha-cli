package cli

import (
	"bufio"
	"github.com/okxtuta/anatha-cli/x/governance"
	treasuryutils "github.com/okxtuta/anatha-cli/x/treasury/client/utils"
	"github.com/okxtuta/anatha-cli/x/treasury/internal/types"
	"github.com/okxtuta/go-anatha/client/context"
	"github.com/okxtuta/go-anatha/codec"
	sdk "github.com/okxtuta/go-anatha/types"
	"github.com/okxtuta/go-anatha/x/auth"
	"github.com/okxtuta/go-anatha/x/auth/client/utils"
	"github.com/spf13/cobra"
)

func GetCmdSubmitAddBuyBackLiquidityProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-buyback-liquidity [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Add BuyBack Liquidity",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := treasuryutils.ParseGenericTreasuryProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := types.NewAddBuyBackLiquidityProposal(
				proposal.Title,
				proposal.Description,
				proposal.Amount,
			)

			msg := governance.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSubmitRemoveBuyBackLiquidityProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-buyback-liquidity [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Remove BuyBack Liquidity",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := treasuryutils.ParseGenericTreasuryProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := types.NewRemoveBuyBackLiquidityProposal(
				proposal.Title,
				proposal.Description,
				proposal.Amount,
			)

			msg := governance.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSubmitBurnDistributionProfitsProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-distribution-profits [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Burn Distribution Profits",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := treasuryutils.ParseGenericTreasuryProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := types.NewBurnDistributionProfitsProposal(
				proposal.Title,
				proposal.Description,
				proposal.Amount,
			)

			msg := governance.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSubmitTransferFromDistributionProfitsToBuyBackLiquidityProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-from-distribution-profits-to-buyback-liquidity [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Transfer From Distribution Profits To BuyBack Liquidity",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := treasuryutils.ParseGenericTreasuryProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := types.NewTransferFromDistributionProfitsToBuyBackLiquidityProposal(
				proposal.Title,
				proposal.Description,
				proposal.Amount,
			)

			msg := governance.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSubmitTransferFromTreasuryToSwapEscrowProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-from-treasury-to-swap-escrow [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Transfer From Treasury To Swap Escrow",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := treasuryutils.ParseGenericTreasuryProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := types.NewTransferFromTreasuryToSwapEscrowProposal(
				proposal.Title,
				proposal.Description,
				proposal.Amount,
			)

			msg := governance.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSubmitTransferFromSwapEscrowToBuyBackProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-from-swap-escrow-to-buyback [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Transfer From Swap Escrow To BuyBack",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := treasuryutils.ParseGenericTreasuryProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()
			content := types.NewTransferFromSwapEscrowToBuyBackProposal(
				proposal.Title,
				proposal.Description,
				proposal.Amount,
			)

			msg := governance.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
