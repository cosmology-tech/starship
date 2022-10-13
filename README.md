# kubeshuttle
A k8s based testing setup framework built with helm

## Next steps
* 12/10/2022: 
  * Run ibc txns from outside the system
  * Cleanup values.yaml interface
  * Add default values for all chains
* 13/10/2022:
  * Depriotize multi-relayer setup, single relayer working

## Work done
* 13/10/2022:
  * Values.yaml fixed for relayer to have seperate mnemonics for scaling
  * Need to make decision about running validators on same connections vs others

## Issues
* Relayer, when we have multi-node setup, then how to setup the inital connection
  * Can we use the same connection betwwen different chains? What is connection and client connection
  * Might not scale with multi-node setups

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
