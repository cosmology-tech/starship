import { SigningCosmWasmClient } from "@cosmjs/cosmwasm-stargate";
import { DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
import { makeCosmoshubPath } from "@cosmjs/amino";

import { getChainInfo } from "./network";
import { getMnemonic } from "./keys";

export async function connect(name: string) {
    const network = getChainInfo(name);

    const { prefix, gasPrice, feeToken, rpcEndpoint } = network;
    const hdPath = makeCosmoshubPath(0);
    const mnemonic = getMnemonic();

    // Setup signer
    const offlineSigner = await DirectSecp256k1HdWallet.fromMnemonic(mnemonic, { prefix, hdPaths: [hdPath] });
    const { address } = (await offlineSigner.getAccounts())[0];
    console.log(`Connected to ${address}`);

    // Init SigningCosmWasmClient client
    const client = await SigningCosmWasmClient.connectWithSigner(rpcEndpoint, offlineSigner, {
        prefix,
        gasPrice,
    });
    const balance = await client.getBalance(address, feeToken);
    console.log(`Balance: ${balance.amount} ${balance.denom}`);

    const chainId = await client.getChainId();

    if (chainId !== network.chainId) {
        throw Error("Given ChainId doesn't match the clients ChainID!");
    }

    return { client, address };
}
