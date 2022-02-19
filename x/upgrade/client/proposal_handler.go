package client

import (
	govclient "github.com/okxtuta/anatha-cli/x/governance/client"
	"github.com/okxtuta/anatha-cli/x/upgrade/client/cli"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitSoftwareUpgradeProposal)
