#!/bin/bash

## Script is used to initialize the current node as a consumer chain to the
## provider chain. Script assumes this runs with provider-chain docker image.
## This script will perform the following steps:
## - create address on provider for ics initialization # use provider faucet
## - submit proposal on provider
## - pass proposal on provider
## - fetch ccv specific info from
