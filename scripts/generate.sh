#!/bin/sh

projectDir=$(dirname $(pwd))
echo $projectDir
protoc  $projectDir/api/*.proto --proto_path=$projectDir --go-grpc_out=$projectDir --go_out=$projectDir