#!/bin/sh
echo "\n===== LOADING DATABASE ====="
go run tools/sqlite/main.go

echo "===== RUNNING TESTS ====="
go test  ./... -count=1 -coverprofile=coverage.out

echo "\n===== COVERAGE ====="
go tool cover -func=coverage.out

echo "\n===== APPLICATION ====="
go run cmd/streetmarket/main.go
