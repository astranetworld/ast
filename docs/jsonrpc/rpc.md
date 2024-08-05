# `rpc` Namespace

The `rpc` API provides methods to get information about the RPC server itself, such as the enabled namespaces.

## `rpc_modules`

Lists the enabled RPC namespaces and the versions of each.

| Client | Method invocation                         |
|--------|-------------------------------------------|
| RPC    | `{"method": "rpc_modules", "params": []}` |

### Example

```js
// > {"jsonrpc":"2.0","id":1,"method":"rpc_modules","params":[]}
{"jsonrpc":"2.0","id":1,"result":{"txpool":"1.0","eth":"1.0","rpc":"1.0"}}
```

## Handling Responses During Syncing

When interacting with the RPC server while it is still syncing, some RPC requests may return an empty or null response, while others return the expected results. This behavior can be observed due to the asynchronous nature of the syncing process and the availability of required data. Notably, endpoints that rely on specific stages of the syncing process, such as the execution stage, might not be available until those stages are complete.

It's important to understand that during pipeline sync, some endpoints may not be accessible until the necessary data is fully synchronized. For instance, the `eth_getBlockReceipts` endpoint is only expected to return valid data after the execution stage, where receipts are generated, has completed. As a result, certain RPC requests may return empty or null responses until the respective stages are finished.

This behavior is intrinsic to how the syncing mechanism works and is not indicative of an issue or bug. If you encounter such responses while the node is still syncing, it's recommended to wait until the sync process is complete to ensure accurate and expected RPC responses.



## The POA consensus algorithm is used in AST

Proof of Authority (PoA) is a consensus algorithm used in blockchain systems where a limited number of nodes, known as validators, are authorized to validate transactions and create new blocks. PoA is particularly suitable for private or consortium blockchains where trust is established among participants.

Validators: Validators in PoA are pre-selected and trusted entities whose identities are known. They are responsible for validating transactions and adding new blocks to the blockchain. This can include organizations or individuals with established reputations.

Security and Trust: Security in PoA is based on the reputations of the validators. Since their identities are public and their actions are transparent, validators are incentivized to behave honestly to maintain their status. Misconduct can lead to loss of their authority status.

One of the main criticisms of PoA is its centralization. The control is in the hands of a few validators, which can lead to concerns about censorship and single points of failure. However, this can be mitigated in a consortium blockchain where multiple stakeholders have validator roles.

 PoA is particularly useful for enterprise applications, supply chain management, and any scenario where a trusted environment is necessary. It is also used in test networks, like Ethereum's Rinkeby testnet, to simulate conditions similar to mainnet but with higher efficiency.