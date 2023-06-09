
build-all:
	cd checkout && GOOS=linux make build
	cd loms && GOOS=linux make build
	cd notifications && GOOS=linux make build

run-all: build-all
	sudo docker compose up --force-recreate --build

precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit

vendor-proto:
		mkdir -p vendor-proto
		@if [ ! -d vendor-proto/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor-proto/googleapis &&\
			mkdir -p  vendor-proto/google/ &&\
			mv vendor-proto/googleapis/google/api vendor-proto/google &&\
			rm -rf vendor-proto/googleapis ;\
		fi
		@if [ ! -d vendor-proto/google/protobuf ]; then\
			git clone https://github.com/protocolbuffers/protobuf vendor-proto/protobuf &&\
			mkdir -p  vendor-proto/google/protobuf &&\
			mv vendor-proto/protobuf/src/google/protobuf/*.proto vendor-proto/google/protobuf &&\
			rm -rf vendor-proto/protobuf ;\
		fi
		@if [ ! -d vendor-proto/github.com/envoyproxy ]; then \
			mkdir -p vendor-proto/validate &&\
			git clone https://github.com/envoyproxy/protoc-gen-validate vendor-proto/protoc-gen-validate &&\
			mv vendor-proto/protoc-gen-validate/validate/*.proto vendor-proto/validate &&\
			rm -r -f vendor-proto/protoc-gen-validate ;\
		fi

generate-all:
	cd checkout && GOOS=linux make generate
	cd loms && GOOS=linux make generate