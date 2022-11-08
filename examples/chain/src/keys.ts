const fs = require("fs")

// TODO: Need to move this to an api call, can be part of chain registry
export function getMnemonic() {
    const jsonString = fs.readFileSync("../../../charts/devnet/configs/keys.json")
    const keys = JSON.parse(jsonString)

    return keys["keys"][0]["mnemonic"];
}
