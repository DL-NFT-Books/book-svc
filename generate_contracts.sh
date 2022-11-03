#!/bin/bash
GENERATED_PATH="./solidity/generated"
SOLIDITY_CONTRACTS="./solidity/contracts"

contracts=$(ls $SOLIDITY_CONTRACTS | grep '.sol')
for contract in $contracts
do
  pkg="${contract%.*}"
  lowercase=$(echo "$pkg" | awk '{print tolower($0)}' )

  mkdir -p $GENERATED_PATH/"$lowercase"
  abigen --abi $SOLIDITY_CONTRACTS/build/"${contract%.*}".abi --out=$GENERATED_PATH/"$lowercase"/main.go --pkg="$lowercase"
done;
