#!/bin/bash

DIR="/home/contract_docker_go/target/release/"

if [ ! -d $DIR ]; then
  mkdir -p $DIR
fi

echo "please input contract name, contract name should be same as name in tx: "
read contract_name
echo "please input zip file: "
read zip_file
go build -o $contract_name
7z a $zip_file $contract_name
# mv $zip_file.7z $DIR