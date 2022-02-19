package client

import (
	"github.com/okxtuta/anatha-cli/x/distribution/client/cli"
	govclient "github.com/okxtuta/anatha-cli/x/governance/client"
)

var DevelopmentFundDistributionProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitDevelopmentFundDistributionProposal)
var SecurityTokenFundDistributionProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitSecurityTokenFundDistributionProposal)
