package app

import (
	"github.com/okxtuta/anatha-cli/x/distribution"
	"github.com/okxtuta/anatha-cli/x/fee"
	"github.com/okxtuta/go-anatha/x/auth"
	"github.com/okxtuta/go-anatha/x/auth/ante"

	sdk "github.com/okxtuta/go-anatha/types"
)

func (app *AnathaApp) NewAnteHandler() sdk.AnteHandler {
	return sdk.ChainAnteDecorators(
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewValidateMemoDecorator(app.accountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(app.accountKeeper),
		ante.NewSetPubKeyDecorator(app.accountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(app.accountKeeper),
		ante.NewDeductFeeDecorator(app.accountKeeper, app.supplyKeeper),
		ante.NewSigGasConsumeDecorator(app.accountKeeper, auth.DefaultSigVerificationGasConsumer),
		ante.NewSigVerificationDecorator(app.accountKeeper),
		fee.NewFeeDecorator(app.feeKeeper, app.bankKeeper, app.hraKeeper, app.supplyKeeper, distribution.AmcModuleName),
		ante.NewIncrementSequenceDecorator(app.accountKeeper), // innermost AnteDecorator
	)
}
