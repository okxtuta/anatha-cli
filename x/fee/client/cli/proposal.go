package cli

import (
	"bufio"
	govutils "github.com/okxtuta/anatha-cli/x/fee/client/utils"
	"github.com/okxtuta/anatha-cli/x/fee/internal/types"
	govtypes "github.com/okxtuta/anatha-cli/x/governance"
	"github.com/okxtuta/go-anatha/client/context"
	"github.com/okxtuta/go-anatha/codec"
	sdk "github.com/okxtuta/go-anatha/types"
	"github.com/okxtuta/go-anatha/x/auth"
	"github.com/okxtuta/go-anatha/x/auth/client/utils"
	"github.com/spf13/cobra"
)

func GetCmdSubmitAddFeeExcludedMessageProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-fee-excluded-message [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to exclude a message type from fees",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := govutils.ParseFeeExclusionProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()

			content := types.NewAddFeeExcludedMessageProposal(proposal.Title, proposal.Description, proposal.MessageType)

			msg := govtypes.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdSubmitRemoveFeeExcludedMessageProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-fee-excluded-message [proposal-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Submit a proposal to remove a message from fee exclusion",
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContextWithInput(inBuf).WithCodec(cdc)

			proposal, err := govutils.ParseFeeExclusionProposalJSON(cdc, args[0])
			if err != nil {
				return err
			}

			from := cliCtx.GetFromAddress()

			content := types.NewRemoveFeeExcludedMessageProposal(proposal.Title, proposal.Description, proposal.MessageType)

			msg := govtypes.NewMsgSubmitProposal(content, from)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
