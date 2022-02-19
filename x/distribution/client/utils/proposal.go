package utils

import (
	"github.com/okxtuta/anatha-cli/x/distribution/internal/types"
	"github.com/okxtuta/go-anatha/codec"
	sdk "github.com/okxtuta/go-anatha/types"
	"io/ioutil"
)

type DevelopmentFundDistributionProposalJSON struct {
	Title       string         `json:"title" yaml:"title"`
	Description string         `json:"description" yaml:"description"`
	Amount      sdk.Coins      `json:"amount" yaml:"amount"`
	Recipient   sdk.AccAddress `json:"recipient" yaml:"recipient"`
}

func ParseDevelopmentFundDistributionProposalJSON(cdc *codec.Codec, proposalFile string) (DevelopmentFundDistributionProposalJSON, error) {
	proposal := DevelopmentFundDistributionProposalJSON{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}

type SecurityTokenFundDistributionProposalJSON struct {
	Title       string             `json:"title" yaml:"title"`
	Description string             `json:"description" yaml:"description"`
	Recipients  []types.Recipients `json:"recipients" yaml:"recipients"`
}

func ParseSecurityTokenFundDistributionProposalJSON(cdc *codec.Codec, proposalFile string) (SecurityTokenFundDistributionProposalJSON, error) {
	proposal := SecurityTokenFundDistributionProposalJSON{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
