package main

import (
	cmdcfg "github.com/UptickNetwork/wasm-nft-convert/testing/simapp/simd/cmd/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/UptickNetwork/wasm-nft-convert/testing/simapp/simd/cmd"
)

func main() {

	setupConfig()
	cmdcfg.RegisterDenoms()

	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, "simd", ""); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}

func setupConfig() {
	// set the address prefixes
	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	cmdcfg.SetBip44CoinType(config)
	config.Seal()
}
