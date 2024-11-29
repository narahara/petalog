.DEFAULT_GOAL = help

build:
	mkdir -p $(APP_HOME)/server/console/ui/src/domain/entity
	protoc -I./ --go_out=$(APP_HOME)/server/common/pkg/domain domain/entity/*.proto
	protoc -I./ --dart_out=grpc:$(APP_HOME)/app/common/lib domain/entity/*.proto
	pbjs -t static-module -w commonjs -o ./.tmp/entity.js -p . domain/entity/*.proto
	pbts -o $(APP_HOME)/server/console/ui/src/domain/entity/entity.d.ts ./.tmp/entity.js

build2:
	protoc -I$(APP_HOME)/common/pb/service/app --go_out=./tmp service/app/*.proto

install:
	brew install protobuf
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	npm install protobufjs -g
	npm install protobufjs-cli -g
	dart pub global activate protoc_plugin

help:
	@echo "help"
