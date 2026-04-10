package types_test

import (
	"testing"

	wasmtypes "github.com/UptickNetwork/wasm-nft-convert/types"
	evmkeeper "github.com/evmos/ethermint/x/evm/keeper"
)

// Compile-time guard: ethermint Keeper must satisfy project CWKeeper interface.
// This test catches upstream signature drift early (e.g. tracer type changes).
var _ wasmtypes.CWKeeper = (*evmkeeper.Keeper)(nil)

func TestCWKeeperInterfaceCompatibility(t *testing.T) {
	t.Helper()
}
