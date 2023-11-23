#!/bin/zsh

rm -f ./**/*_templ.go && templ generate && go run cmd/noestabien.go
