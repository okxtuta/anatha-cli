package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/okxtuta/anatha-cli/x/mint/internal/types"
	"github.com/okxtuta/go-anatha/client"
	"github.com/okxtuta/go-anatha/client/context"
	"github.com/okxtuta/go-anatha/client/flags"
	"github.com/okxtuta/go-anatha/codec"
)

// GetQueryCmd returns the cli query commands for the minting module.
func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	mintingQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the minting module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	mintingQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdQueryParams(cdc),
			GetCmdQueryMinter(cdc),
		)...,
	)

	return mintingQueryCmd
}

func GetCmdQueryParams(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Query the current minting parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParameters)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var params types.Params
			if err := cdc.UnmarshalJSON(res, &params); err != nil {
				return err
			}

			return cliCtx.PrintOutput(params)
		},
	}
}

func GetCmdQueryMinter(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "minter",
		Short: "Query the current minter parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryMinter)
			res, _, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var params types.Minter
			if err := cdc.UnmarshalJSON(res, &params); err != nil {
				return err
			}

			return cliCtx.PrintOutput(params)
		},
	}
}
