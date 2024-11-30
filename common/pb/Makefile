.DEFAULT_GOAL = help

all: setup build_domain build_app build_console


setup:
	mkdir -p $(APP_HOME)/server/console/ui/src/domain/entity
	mkdir -p $(APP_HOME)/server/console/ui/src/dto
	mkdir -p $(APP_HOME)/server/common/pkg/domain
	mkdir -p $(APP_HOME)/server/common/pkg/service

build_domain:
	protoc -I./ \
		--go_out=$(APP_HOME)/server/common/pkg/domain \
		domain/entity/*.proto
	protoc -I./ \
		--dart_out=grpc:$(APP_HOME)/app/common/lib \
		domain/entity/*.proto
	pbjs \
		-t static-module \
		-w commonjs \
		-o ./.tmp/entity.js \
		-p . \
		domain/entity/*.proto
	pbts \
		-o $(APP_HOME)/server/console/ui/src/domain/entity/entity.d.ts \
		./.tmp/entity.js

build_app:
	protoc -I./ \
		--go_out=$(APP_HOME)/server/common/pkg/service \
		--go-grpc_out=$(APP_HOME)/server/common/pkg/service \
		--go-grpc_opt=require_unimplemented_servers=false \
		service/app/*.proto
	protoc -I./ \
		--dart_out=grpc:$(APP_HOME)/app/common/lib \
		service/app/*.proto

build_console:
	protoc -I./ \
		-I/$(GOPATH)/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis \
		--go_out=$(APP_HOME)/server/common/pkg/service \
		--go-grpc_out=$(APP_HOME)/server/common/pkg/service \
		--go-grpc_opt=require_unimplemented_servers=false \
		--grpc-gateway_out $(APP_HOME)/server/common/pkg/service \
		--grpc-gateway_opt logtostderr=true \
		service/console/*.proto
	pbjs \
		-t static-module \
		-w commonjs \
		-o $(APP_HOME)/server/console/ui/src/dto/dto.js \
		-p . \
		service/console/*_dto.proto
	pbts \
		-o $(APP_HOME)/server/console/ui/src/dto/dto.d.ts \
		$(APP_HOME)/server/console/ui/src/dto/dto.js

install_go:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway

install:
	brew install protobuf
	
	go get github.com/grpc-ecosystem/grpc-gateway/v2
	npm install protobufjs -g
	npm install protobufjs-cli -g
	dart pub global activate protoc_plugin

help:
	@echo "help"
