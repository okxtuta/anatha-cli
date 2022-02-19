package types

import (
	"github.com/okxtuta/go-anatha/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(Plan{}, "upgrade/Plan", nil)
	cdc.RegisterConcrete(SoftwareUpgradeProposal{}, "upgrade/SoftwareUpgradeProposal", nil)
	cdc.RegisterConcrete(CancelSoftwareUpgradeProposal{}, "upgrade/CancelSoftwareUpgradeProposal", nil)
}

// module codec
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)

	codec.RegisterCrypto(ModuleCdc)

	ModuleCdc.Seal()
}
