# kubeshuttle
A k8s based testing setup framework built with helm

## Next steps
* 12/10/2022: 
  * Run ibc txns from outside the system

## Improvements
* Need not build all the docker images, the docker requirements are to have jq and bash in alpine
* Can look into using strangelove-ventures/heighliner for docker images creation instead of having self hosted
* Key initialization and recovery takes the most amount of time, could see if we can do this
  * in parallel
  * or precompute the keyring-test directory itself for each of the cases

## How to test upgrades
* Run chain in current state, need to install cosmovisor on all nodes
  * Might neeed to binary pre-installed for both upgrades
  * Can fetch binary via a wget in init-containers
  * Can build the code in init container as well
* run some pre-upgrade txns state
* create upgrade proposal
* cosmovisor will do the upgrade automatically
