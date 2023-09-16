#!/bin/sh

CompileDaemon --build="go build main.go" --command="./main"

go run migrations/migrations.go