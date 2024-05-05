.PHONY: gen-proto

gen-proto:
	mkdir -p ./addressbookpb
	protoc --go_opt=paths=source_relative --go_out=./addressbookpb ./addressbook.proto

test:
	go test ./cmd/*
