#! /usr/bin/env bash
set -uvx
set -e
which gcc
go build -buildmode=c-archive exportgo.go
#g++ -pthread -o main.exe main.cpp exportgo.a -lWinMM -lntdll -lWS2_32 -static
rm -rf ./main.exe
pro ./main.pro
mingwx ./main.exe
