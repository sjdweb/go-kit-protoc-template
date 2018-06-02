#!/usr/bin/env bash
TARGET_SERVICE=$1

if [ "$TARGET_SERVICE" = "" ]; then
  echo "missing service"
  exit -1
fi

cd $TARGET_SERVICE
protoc --gotemplate_out=destination_dir=./gen,template_dir=../../templates/{{.File.Package}}/gen:./gen ./$TARGET_SERVICE.proto
protoc --go_out=plugins=grpc:./gen/pb ./$TARGET_SERVICE.proto