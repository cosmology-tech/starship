## Mesh Security Devnet

> NOTE: Host and smart contract address are subject to change on redeployment of the system

### Provider Chain

Chain-id: `provider`

Host: `a97f533488fe84821bb5a9272031b169-1957337260.ap-southeast-1.elb.amazonaws.com`

Ports:
* RPC: [26657](http://a97f533488fe84821bb5a9272031b169-1957337260.ap-southeast-1.elb.amazonaws.com:26657/status)
* Rest: [1317](http://a97f533488fe84821bb5a9272031b169-1957337260.ap-southeast-1.elb.amazonaws.com:1317)
* Faucet: [8000](http://a97f533488fe84821bb5a9272031b169-1957337260.ap-southeast-1.elb.amazonaws.com:8000)

Get tokens from faucet
```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"denom":"ustake","address":"mesh1yre6ac7qfgyfgvh58ph0rgw627rhw766y430qq"}' \
  http://http://a97f533488fe84821bb5a9272031b169-1957337260.ap-southeast-1.elb.amazonaws.com:8000/credit
```

```
Provider Contracts:
  valut: mesh14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sysl6kf
  externalStaking: mesh1zwv6feuzhy6a9wekh96cd57lsarmqlwxdypdsplw6zhfncqw6ftqsqwra5
  nativeStaking: mesh1qg5ega6dykkxc307y25pecuufrjkxkaggkkxh7nad0vhyhtuhw3stmd2jl
```

### Consumer Chain

Chain-id: `consumer`
Denom: `ustake`
Prefix: `mesh`

Host: `ad655b3f02d52426184e4580d1069e8a-833461481.ap-southeast-1.elb.amazonaws.com`

Ports:
* RPC: 26657
* Rest: 1317
* Faucet: 8000

Get tokens from faucet
```bash
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{"denom":"ustake","address":"mesh1yre6ac7qfgyfgvh58ph0rgw627rhw766y430qq"}' \
  http://ad655b3f02d52426184e4580d1069e8a-833461481.ap-southeast-1.elb.amazonaws.com:8000/credit
```

```
Consumer Contracts:
  staking: mesh1xr3rq8yvd7qplsw5yx90ftsr2zdhg4e9z60h5duusgxpv72hud3syz4y6d
  priceFeed: mesh14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sysl6kf
  converter: mesh1qg5ega6dykkxc307y25pecuufrjkxkaggkkxh7nad0vhyhtuhw3stmd2jl
```

## Chain Registry

Host: `ac7cbb684e7fd41eca0866ec659b4de2-1861589237.ap-southeast-1.elb.amazonaws.com`
Port: [8080](http://ac7cbb684e7fd41eca0866ec659b4de2-1861589237.ap-southeast-1.elb.amazonaws.com:8080/chains)

Endpoints:
* Chains: `/chains/{chain-id}`
* IBC: `/ibc/{chain-1}/{chain-2}`
* Mnemonics: `/chains/{chain-id}/keys`


