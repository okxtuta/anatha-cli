package upgrade

import (
	"github.com/okxtuta/anatha-cli/x/upgrade/internal/types"
	sdk "github.com/okxtuta/go-anatha/types"
	sdkerrors "github.com/okxtuta/go-anatha/types/errors"
	govtypes "github.com/okxtuta/go-anatha/x/gov/types"
	"strconv"
)

// NewSoftwareUpgradeProposalHandler creates a governance handler to manage new proposal types.
// It enables SoftwareUpgradeProposal to propose an Upgrade, and CancelSoftwareUpgradeProposal
// to abort a previously voted upgrade.
func NewSoftwareUpgradeProposalHandler(k Keeper) govtypes.Handler {
	return func(ctx sdk.Context, content govtypes.Content) error {
		switch c := content.(type) {
		case SoftwareUpgradeProposal:
			return handleSoftwareUpgradeProposal(ctx, k, c)

		case CancelSoftwareUpgradeProposal:
			return handleCancelSoftwareUpgradeProposal(ctx, k, c)

		default:
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized software upgrade proposal content type: %T", c)
		}
	}
}

func handleSoftwareUpgradeProposal(ctx sdk.Context, k Keeper, p SoftwareUpgradeProposal) error {
	err := k.ScheduleUpgrade(ctx, p.Plan)

	if err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpgrade,
			sdk.NewAttribute(types.AttributeKeyTitle, p.Title),
			sdk.NewAttribute(types.AttributeKeyDescription, p.Description),
			sdk.NewAttribute(types.AttributeKeyPlanName, p.Plan.Name),
			sdk.NewAttribute(types.AttributeKeyPlanHeight, strconv.Itoa(int(p.Plan.Height))),
			sdk.NewAttribute(types.AttributeKeyPlanInfo, p.Plan.Info),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueModule),
		),
	)

	return nil
}

func handleCancelSoftwareUpgradeProposal(ctx sdk.Context, k Keeper, p CancelSoftwareUpgradeProposal) error {
	k.ClearUpgradePlan(ctx)
	return nil
}
