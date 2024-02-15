package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"

	"github.com/UptickNetwork/uptick/x/cw721/types"
)

// GetTokenPairs - get all registered token tokenPairs
func (k Keeper) GetTokenPairs(ctx sdk.Context) []types.TokenPair {
	tokenPairs := []types.TokenPair{}

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefixTokenPair)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var tokenPair types.TokenPair
		k.cdc.MustUnmarshal(iterator.Value(), &tokenPair)

		tokenPairs = append(tokenPairs, tokenPair)
	}

	return tokenPairs
}

// GetTokenPair - get registered token pair from the identifier
func (k Keeper) GetTokenPair(ctx sdk.Context, id []byte) (types.TokenPair, bool) {
	if id == nil {
		return types.TokenPair{}, false
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPair)
	var tokenPair types.TokenPair
	bz := store.Get(id)
	if len(bz) == 0 {
		return types.TokenPair{}, false
	}

	k.cdc.MustUnmarshal(bz, &tokenPair)
	return tokenPair, true
}

// SetTokenPair stores a token pair
func (k Keeper) SetTokenPair(ctx sdk.Context, tokenPair types.TokenPair) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPair)
	key := tokenPair.GetID()
	bz := k.cdc.MustMarshal(&tokenPair)
	store.Set(key, bz)
}

// SetWasmCode stores a WasmCode
func (k Keeper) SetWasmCode(ctx sdk.Context, account string, codeId uint64) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixWasmCode)
	store.Set([]byte(account), []byte(strconv.FormatUint(codeId, 10)))
}

// GetWasmCode stores a WasmCode
func (k Keeper) GetWasmCode(ctx sdk.Context, account string) []byte {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixWasmCode)
	return store.Get([]byte(account))
}

// DeleteTokenPair removes a token pair.
func (k Keeper) DeleteTokenPair(ctx sdk.Context, tokenPair types.TokenPair) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPair)
	key := tokenPair.GetID()
	store.Delete(key)
}

// SetCW721Map sets the token pair id for the given address
func (k Keeper) SetCW721Map(ctx sdk.Context, cw721 string, id []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByCW721)
	store.Set([]byte(cw721), id)
}

// GetCW721Map returns the token pair id for the given address
func (k Keeper) GetCW721Map(ctx sdk.Context, cw721 string) []byte {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByCW721)
	return store.Get([]byte(cw721))
}

// GetClassMap returns the token pair id for the given class
func (k Keeper) GetClassMap(ctx sdk.Context, classID string) []byte {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByClass)
	return store.Get([]byte(classID))
}

// DeleteCW721Map deletes the token pair id for the given address
func (k Keeper) DeleteCW721Map(ctx sdk.Context, cw721 common.Address) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByCW721)
	store.Delete(cw721.Bytes())
}

// SetClassMap sets the token pair id for the classID
func (k Keeper) SetClassMap(ctx sdk.Context, classID string, id []byte) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByClass)
	store.Set([]byte(classID), id)
}

// IsTokenPairRegistered - check if registered token tokenPair is registered
func (k Keeper) IsTokenPairRegistered(ctx sdk.Context, id []byte) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPair)
	return store.Has(id)
}

// IsCW721Registered check if registered CW721 token is registered
func (k Keeper) IsCW721Registered(ctx sdk.Context, cw721 string) bool {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByCW721)
	return store.Has([]byte(cw721))
}

// IsClassRegistered check if registered nft class is registered
func (k Keeper) IsClassRegistered(ctx sdk.Context, classID string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixTokenPairByClass)
	return store.Has([]byte(classID))
}

func (k Keeper) SetNFTPairs(ctx sdk.Context, contractAddress string, tokenID string, classID string, nftID string) {

	// save nft pair
	if len(k.GetNFTPairByContractTokenID(ctx, contractAddress, tokenID)) == 0 {
		k.SetNFTPairByContractTokenID(ctx, contractAddress, tokenID, classID, nftID)
	}

	if len(k.GetNFTPairByClassNFTID(ctx, classID, nftID)) == 0 {
		k.SetNFTPairByClassNFTID(ctx, classID, nftID, contractAddress, tokenID)
	}
}

// GetTokenPairID returns the pair id from either of the registered tokens.
func (k Keeper) GetTokenPairID(ctx sdk.Context, token string) []byte {

	//if common.IsHexAddress(token) {
	//	return k.GetCW721Map(ctx, token)
	//}
	if len(k.GetCW721Map(ctx, token)) == 0 {
		return k.GetClassMap(ctx, token)
	} else {
		return k.GetCW721Map(ctx, token)
	}
}

func (k Keeper) SetNFTPairByContractTokenID(ctx sdk.Context, contractAddress string, tokenID string, classID string, nftID string) {

	tokenUID := types.CreateTokenUID(contractAddress, tokenID)
	nftUID := types.CreateNFTUID(classID, nftID)

	k.SetNFTUIDPairByTokenUID(ctx, tokenUID, nftUID)
}

func (k Keeper) SetNFTUIDPairByTokenUID(ctx sdk.Context, tokenUID string, nftUID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixNFTUIDPairByTokenUID)
	store.Set([]byte(tokenUID), []byte(nftUID))
}

func (k Keeper) GetNFTPairByContractTokenID(ctx sdk.Context, contractAddress string, tokenID string) []byte {
	tokenUID := types.CreateTokenUID(contractAddress, tokenID)
	return k.GetNFTUIDPairByTokenUID(ctx, tokenUID)
}

func (k Keeper) GetNFTUIDPairByTokenUID(ctx sdk.Context, tokenUID string) []byte {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixNFTUIDPairByTokenUID)
	return store.Get([]byte(tokenUID))
}

func (k Keeper) DeleteNFTPairByTokenID(ctx sdk.Context, contractAddress string, tokenID string) {
	tokenUID := types.CreateTokenUID(contractAddress, tokenID)
	k.DeleteNFTUIDPairByTokenUID(ctx, tokenUID)
}

func (k Keeper) DeleteNFTUIDPairByTokenUID(ctx sdk.Context, tokenUID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixNFTUIDPairByTokenUID)
	store.Delete([]byte(tokenUID))
}

func (k Keeper) SetNFTPairByClassNFTID(ctx sdk.Context, classID string, nftID string, contractAddress string, tokenID string) {

	nftUID := types.CreateNFTUID(classID, nftID)
	tokenUID := types.CreateTokenUID(contractAddress, tokenID)

	k.SetNFTUIDPairByNFTUID(ctx, nftUID, tokenUID)
}

func (k Keeper) SetNFTUIDPairByNFTUID(ctx sdk.Context, nftUID string, tokenUID string) {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixNFTUIDPairByNFTUID)
	store.Set([]byte(nftUID), []byte(tokenUID))
}

func (k Keeper) GetNFTPairByClassNFTID(ctx sdk.Context, classID string, nftID string) []byte {
	nftUID := types.CreateNFTUID(classID, nftID)
	return k.GetTokenUIDPairByNFTUID(ctx, nftUID)
}

func (k Keeper) GetTokenUIDPairByNFTUID(ctx sdk.Context, nftUID string) []byte {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixNFTUIDPairByNFTUID)
	return store.Get([]byte(nftUID))
}

func (k Keeper) DeleteNFTPairByNFTID(ctx sdk.Context, classID string, nftID string) {
	nftUID := types.CreateNFTUID(classID, nftID)
	k.DeleteNFTUIDPairByNFTUID(ctx, nftUID)
}

func (k Keeper) DeleteNFTUIDPairByNFTUID(ctx sdk.Context, nftUID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixNFTUIDPairByNFTUID)
	store.Delete([]byte(nftUID))
}

// SetCwAddressByContractTokenId
func (k Keeper) SetCwAddressByContractTokenId(ctx sdk.Context, evmContractAddress string, evmTokenId string, evmAddress string) {

	contractAndTokenId := evmContractAddress + evmTokenId
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixCwAddressByContractTokenId)
	store.Set([]byte(contractAndTokenId), []byte(evmAddress))
}

func (k Keeper) GetCwAddressByContractTokenId(ctx sdk.Context, evmContractAddress string, evmTokenId string) []byte {

	contractAndTokenId := evmContractAddress + evmTokenId
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixCwAddressByContractTokenId)
	return store.Get([]byte(contractAndTokenId))
}

func (k Keeper) DeleteCwAddressByContractTokenId(ctx sdk.Context, evmContractAddress string, evmTokenId string) {

	contractAndTokenId := evmContractAddress + evmTokenId
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixCwAddressByContractTokenId)
	store.Delete([]byte(contractAndTokenId))
}
