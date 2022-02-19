package types

import (
	"fmt"
	"github.com/okxtuta/anatha-cli/config"
	sdk "github.com/okxtuta/go-anatha/types"
	"github.com/okxtuta/go-anatha/x/params"
)

var (
	DefaultFeePercentage = sdk.NewDecWithPrec(2, 3)
	DefaultMinimumFee    = sdk.NewCoins(sdk.NewInt64Coin(config.DefaultDenom, 200))       // 200pin
	DefaultMaximumFee    = sdk.NewCoins(sdk.NewInt64Coin(config.DefaultDenom, 100000000)) // 1 anatha

	DefaultFeeExcludedMessages = []string{
		"treasury/disburse",
		"treasury/disburse_to_escrow",
		"treasury/disburse_from_escrow",
		"treasury/revert_from_escrow",
		"governance/submit_proposal",
		"governance/vote",
		"governance/expedite",
	}

	KeyFeePercentage = []byte("FeePercentage")
	KeyMinimumFee    = []byte("MinimumFee")
	KeyMaximumFee    = []byte("MaximumFee")
)

func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

type Params struct {
	FeePercentage sdk.Dec   `json:"fee_percentage" yaml:"fee_percentage"`
	MinimumFee    sdk.Coins `json:"minimum_fee" yaml:"minimum_fee"`
	MaximumFee    sdk.Coins `json:"maximum_fee" yaml:"maximum_fee"`
}

func NewParams(feePercentage sdk.Dec, minimumFee sdk.Coins, maximumFee sdk.Coins) Params {
	return Params{
		FeePercentage: feePercentage,
		MinimumFee:    minimumFee,
		MaximumFee:    maximumFee,
	}
}

// String implements the stringer interface for Params
func (p Params) String() string {
	return fmt.Sprintf(`Params:
`)
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(KeyFeePercentage, &p.FeePercentage, validateFeePercentage),
		params.NewParamSetPair(KeyMinimumFee, &p.MinimumFee, validateFee),
		params.NewParamSetPair(KeyMaximumFee, &p.MaximumFee, validateFee),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultFeePercentage,
		DefaultMinimumFee,
		DefaultMaximumFee,
	)
}

func (p Params) Validate() error {
	if err := validateFeePercentage(p.FeePercentage); err != nil {
		return err
	}

	if err := validateFee(p.MinimumFee); err != nil {
		return err
	}

	if err := validateFee(p.MaximumFee); err != nil {
		return err
	}

	return nil
}

func validateFee(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if !v.IsValid() {
		return fmt.Errorf("invalid name info fee: %s", v)
	}
	if !v.AmountOf(config.DefaultDenom).IsPositive() {
		return fmt.Errorf("invalid fee denomination. expected: %s", config.DefaultDenom)
	}
	return nil
}

func validateFeePercentage(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.GT(sdk.OneDec()) {
		return fmt.Errorf("fee percentage must be less than 100%: %s", v)
	}
	if v.IsNegative() {
		return fmt.Errorf("fee percentage must be positive: %s", v)
	}

	return nil
}
