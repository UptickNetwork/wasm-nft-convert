---
title: Non-Fungible Token Convert between coswasm and cosmos native nft
category: IBC/APP
requires: 25, 26,721
kind: instantiation
author: uptsmart <uptsmart@163.com>
created: 2022-03-04
modified: 2025-03-06
---

## Abstract

Coswasm and Cosmos are two different blockchain technologies, and conversion between them involves different protocols and bridging mechanisms. In this case, the conversion from Coswasm to Cosmos native NFT can realize the transfer of NFT issued on the Ethereum blockchain to the Cosmos blockchain for circulation and transactions. This conversion function can help NFT holders transfer their assets from one blockchain to another for better market liquidity and wider usage scenarios.

When performing this conversion, a corresponding bridge protocol or cross-chain interaction protocol needs to be used to establish the connection between the Coswasm and the Cosmos blockchain. Once this connection is established, NFT holders can transfer their assets between the two blockchains through specific operations. This conversion function can help expand the coverage of the NFT market, promote cross-chain transactions and interoperability, and improve the liquidity and use value of NFT.


### Desired Properties

- Preservation of non-fungibility (i.e., only one instance of any token is *live* across all the IBC-connected blockchains).
- Permissionless token transfers, no need to whitelist connections, modules, or `classId`s.
- Symmetric (all chains implement the same logic, no in-protocol differentiation of hubs & zones).

## Technical Specification

### Data Structures

#### Convert from cosmos NFT to Coswasm CW721 asset

```typescript
// MsgConvertNFT defines a Msg to convert a native Cosmos nft to a CW721 token
message MsgConvertNFT {
  // nft classID to cnvert to CW721
  string class_id = 1;
  // nftID to cnvert to CW721
  repeated string nft_ids = 2;
  // recipient hex address to receive CW721 token
  string receiver = 3;
  // cosmos bech32 address from the owner of the given Cosmos coins
  string sender = 4;
  // CW721 token contract address registered in a token pair
  string contract_address = 5;
  // CW721 token id registered in a token pair
  repeated string token_ids = 6;
}
```
`classId` is a required field that MUST never be empty, it uniquely identifies the class/collection/contract which the tokens being transferred belong to in the sending chain. In the case of an ERC-1155 compliant smart contract, for example, this could be a string representation of the top 128 bits of the token ID.

`nft_ids` array is a required field that MUST have a size greater than zero and hold non-empty entries that uniquely identify tokens (of the given class) that are being transferred. In the case of an ERC-1155 compliant smart contract, for example, a `nft_ids` could be a string representation of the bottom 128 bits of the token ID.

`receiver` recipient hex address to receive CW721 token.

`sender` CW721 token contract address registered in a token pair.

`contract_address` CW721 token contract address registered in a token pair.

`token_ids` CW721 token id registered in a token pair.

#### Convert from  Coswasm CW721 to cosmos NFT asset

```typescript
// MsgConvertCW721 defines a Msg to convert a CW721 token to a native Cosmos
// nft.
message MsgConvertCW721 {
  // CW721 token contract address registered in a token pair
  string contract_address = 1;
  // tokenID to convert
  repeated string token_ids = 2;
  // bech32 address to receive native Cosmos coins
  string receiver = 3;
  // sender hex address from the owner of the given CW721 tokens
  string sender = 4;
  // nft classID to cnvert to CW721
  string class_id = 5;
  // nftID to cnvert to CW721
  repeated string nft_ids = 6;
}
```
`contract_address` CW721 token contract address registered in a token pair.

`token_ids` CW721 token contract address registered in a token pair.

`receiver` bech32 address to receive native Cosmos coins.

`sender` sender hex address from the owner of the given CW721 tokens.

`class_id` is a required field that MUST never be empty, it uniquely identifies the class/collection/contract which the tokens being transferred belong to in the sending chain. In the case of an ERC-1155 compliant smart contract, for example, this could be a string representation of the top 128 bits of the token ID.

`nft_ids` array is a required field that MUST have a size greater than zero and hold non-empty entries that uniquely identify tokens (of the given class) that are being transferred. In the case of an ERC-1155 compliant smart contract, for example, a `cosmos_token_ids` could be a string representation of the bottom 128 bits of the token ID.


#### related SDK
- Implementation of SDK in nodejs can be found in [uptick-chain-sdk.js](https://github.com/UptickNetwork/uptick-chain-sdk.js).


#### technical documentation
- Related technical documentation can be found in [NFT Conversion Document](https://app.gitbook.com/o/uPfC9w7sfZt6S3IXypqI/s/6zlFzpQT9NAx43Tcz0mE/).

## Backwards Compatibility
Not applicable.


## Copyright
All content herein is licensed under [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0).
