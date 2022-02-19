package client

import (
	govclient "github.com/okxtuta/anatha-cli/x/governance/client"
	"github.com/okxtuta/anatha-cli/x/hra/client/cli"
)

var (
	RegisterBlockchainIdProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitRegisterBlockchainIdProposal)
	RemoveBlockchainIdProposalHandler   = govclient.NewProposalHandler(cli.GetCmdSubmitRemoveBlockchainIdProposal)
)
