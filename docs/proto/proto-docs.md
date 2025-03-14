# Protocol Documentation

# uptick/cw721/v1/cw721.proto




## TokenPair
TokenPair defines an instance that records a pairing consisting of a native
Cosmos Coin and an CW721 token address.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cw721_address | string |  | address of CW721 contract token |
| class_id | string |  | cosmos nft class ID to be mapped to |




## UIDPair
defines the unique id of nft asset


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cw721_did | string |  | address of CW721 contract token + tokenId |
| class_did | string |  | cosmos nft class ID to be mapped to + nftId |






## Owner
Owner enumerates the ownership of a CW721 contract.

| Name | Number | Description |
| ---- | ------ | ----------- |
| OWNER_UNSPECIFIED | 0 | OWNER_UNSPECIFIED defines an invalid/undefined owner. |
| OWNER_MODULE | 1 | OWNER_MODULE cw721 is owned by the cw721 module account. |
| OWNER_EXTERNAL | 2 | EXTERNAL cw721 is owned by an external account. |









# uptick/cw721/v1/genesis.proto




## GenesisState
GenesisState defines the module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | Params |  | module parameters |
| token_pairs | TokenPair | repeated | registered token pairs |




## Params
Params defines the cw721 module params


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enable_cw721 | bool |  | parameter to enable the conversion of Cosmos nft <--> CW721 tokens. |
| enable_evm_hook | bool |  | parameter to enable the EVM hook that converts an CW721 token to a Cosmos
NFT by transferring the Tokens through a MsgEthereumTx to the
ModuleAddress Ethereum address. |












# uptick/cw721/v1/query.proto




## QueryParamsRequest
QueryParamsRequest is the request type for the Query/Params RPC method




## QueryParamsResponse
QueryParamsResponse is the response type for the Query/Params RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | Params |  |  |




## QueryTokenPairRequest
QueryTokenPairRequest is the request type for the Query/TokenPair RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | string |  | token identifier can be either the hex contract address of the CW721 or
the Cosmos nft class_id |




## QueryTokenPairResponse
QueryTokenPairResponse is the response type for the Query/TokenPair RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token_pair | TokenPair |  |  |




## QueryTokenPairsRequest
QueryTokenPairsRequest is the request type for the Query/TokenPairs RPC
method.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pagination | cosmos.base.query.v1beta1.PageRequest |  | pagination defines an optional pagination for the request |




## QueryTokenPairsResponse
QueryTokenPairsResponse is the response type for the Query/TokenPairs RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token_pairs | TokenPair | repeated |  |
| pagination | cosmos.base.query.v1beta1.PageResponse |  | pagination defines the pagination in the response |




## QueryWasmAddressRequest
QueryWasmAddressRequest is the request type for the Query/WasmContract RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| port | string |  | port identifier for the IBC connection |
| channel | string |  | channel identifier for the IBC connection |
| class_id | string |  | class_id is the NFT class identifier |




## QueryWasmContractResponse
QueryWasmContractResponse is the response type for the Query/WasmContract RPC method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token_pair | TokenPair |  |  |










## Query
Query defines the gRPC querier service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| TokenPairs | QueryTokenPairsRequest | QueryTokenPairsResponse | TokenPairs retrieves registered token pairs |
| TokenPair | QueryTokenPairRequest | QueryTokenPairResponse | TokenPair retrieves a registered token pair |
| WasmContract | QueryWasmAddressRequest | QueryWasmContractResponse | WasmContract retrieves a registered wasm contract |
| Params | QueryParamsRequest | QueryParamsResponse | Params retrieves the cw721 module params |




# uptick/cw721/v1/tx.proto




## MsgConvertCW721
MsgConvertCW721 defines a Msg to convert a CW721 token to a native Cosmos nft


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | string |  | CW721 token contract address registered in a token pair |
| token_ids | string | repeated | token_ids to convert |
| receiver | string |  | bech32 address to receive native Cosmos coins |
| sender | string |  | sender hex address from the owner of the given CW721 tokens |
| class_id | string |  | class_id to convert to CW721 |
| nft_ids | string | repeated | nft_ids to convert to CW721 |




## MsgConvertCW721Response
MsgConvertCW721Response returns no fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contract_address | string |  | CW721 token contract address registered in a token pair |
| token_ids | string | repeated | token_ids to convert |
| receiver | string |  | bech32 address to receive native Cosmos coins |
| sender | string |  | sender hex address from the owner of the given CW721 tokens |
| class_id | string |  | class_id to convert to CW721 |
| nft_ids | string | repeated | nft_ids to convert to CW721 |




## MsgConvertNFT
MsgConvertNFT defines a Msg to convert a native Cosmos nft to a CW721 token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| class_id | string |  | class_id to convert to CW721 |
| nft_ids | string | repeated | nft_ids to convert to CW721 |
| receiver | string |  | recipient hex address to receive CW721 token |
| sender | string |  | cosmos bech32 address from the owner of the given Cosmos coins |
| contract_address | string |  | CW721 token contract address registered in a token pair |
| token_ids | string | repeated | CW721 token id registered in a token pair |




## MsgConvertNFTResponse
MsgConvertNFTResponse returns no fields




## MsgTransferCW721
MsgTransferCW721 defines a message for transferring CW721 tokens through IBC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cw_contract_address | string |  | cw_contract_address is the CW721 token contract address |
| cw_token_ids | string | repeated | cw_token_ids are the token IDs to transfer |
| source_port | string |  | source_port is the port on which the packet will be sent |
| source_channel | string |  | source_channel is the channel by which the packet will be sent |
| class_id | string |  | class_id is the class ID of tokens to be transferred |
| cosmos_token_ids | string | repeated | cosmos_token_ids are the non fungible tokens to be transferred |
| cw_sender | string |  | cw_sender is the sender address |
| cosmos_receiver | string |  | cosmos_receiver is the recipient address on the destination chain |
| timeout_height | ibc.core.client.v1.Height |  | timeout_height is the timeout height relative to the current block height
The timeout is disabled when set to 0 |
| timeout_timestamp | uint64 |  | timeout_timestamp is the timeout timestamp in absolute nanoseconds since unix epoch
The timeout is disabled when set to 0 |
| memo | string |  | memo is an optional memo field |




## MsgTransferCW721Response
MsgTransferCW721Response defines the response type for TransferCW721 RPC










## Msg
Msg defines the cw721 Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ConvertNFT | MsgConvertNFT | MsgConvertNFTResponse | ConvertNFT mints a CW721 representation of the native Cosmos nft
that is registered on the token mapping. |
| ConvertCW721 | MsgConvertCW721 | MsgConvertCW721Response | ConvertCW721 mints a native Cosmos coin representation of the CW721 token
contract that is registered on the token mapping. |
| TransferCW721 | MsgTransferCW721 | MsgTransferCW721Response | TransferCW721 transfers a CW721 token from one chain to another chain through IBC |



 