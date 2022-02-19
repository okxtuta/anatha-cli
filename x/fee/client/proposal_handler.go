package client

import (
	"github.com/okxtuta/anatha-cli/x/fee/client/cli"
	govclient "github.com/okxtuta/anatha-cli/x/governance/client"
)

var (
	AddFeeExcludedMessageProposalHandler    = govclient.NewProposalHandler(cli.GetCmdSubmitAddFeeExcludedMessageProposal)
	RemoveFeeExcludedMessageProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitRemoveFeeExcludedMessageProposal)
)
