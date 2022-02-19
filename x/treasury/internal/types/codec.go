package types

import (
	"github.com/okxtuta/go-anatha/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgAddOperator{}, "treasury/AddOperator", nil)
	cdc.RegisterConcrete(MsgRemoveOperator{}, "treasury/RemoveOperator", nil)
	cdc.RegisterConcrete(MsgDisburse{}, "treasury/Disburse", nil)
	cdc.RegisterConcrete(MsgDisburseToEscrow{}, "treasury/DisburseToEscrow", nil)
	cdc.RegisterConcrete(MsgDisburseFromEscrow{}, "treasury/DisburseFromEscrow", nil)
	cdc.RegisterConcrete(MsgRevertFromEscrow{}, "treasury/RevertFromEscrow", nil)
	cdc.RegisterConcrete(MsgCancelDisbursement{}, "treasury/CancelDisbursement", nil)
	cdc.RegisterConcrete(MsgCreateSellOrder{}, "treasury/CreateSellOrder", nil)
	cdc.RegisterConcrete(MsgCreateBuyOrder{}, "treasury/CreateBuyOrder", nil)
	cdc.RegisterConcrete(MsgSwap{}, "treasury/Swap", nil)

	cdc.RegisterConcrete(AddBuyBackLiquidityProposal{}, "treasury/AddBuyBackLiquidityProposal", nil)
	cdc.RegisterConcrete(RemoveBuyBackLiquidityProposal{}, "treasury/RemoveBuyBackLiquidityProposal", nil)
	cdc.RegisterConcrete(BurnDistributionProfitsProposal{}, "treasury/BurnDistributionProfitsProposal", nil)
	cdc.RegisterConcrete(TransferFromDistributionProfitsToBuyBackLiquidityProposal{}, "treasury/TransferFromDistributionProfitsToBuyBackLiquidityProposal", nil)
	cdc.RegisterConcrete(TransferFromTreasuryToSwapEscrowProposal{}, "treasury/TransferFromTreasuryToSwapEscrowProposal", nil)
	cdc.RegisterConcrete(TransferFromSwapEscrowToBuyBackProposal{}, "treasury/TransferFromSwapEscrowToBuyBackProposal", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)

	codec.RegisterCrypto(ModuleCdc)

	ModuleCdc.Seal()
}
