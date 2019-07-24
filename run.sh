#!/bin/bash
SERVER_NAME=$1
SERVER_PORT=$2

if [ -z "${SERVER_NAME}" ]; then
  SERVER_NAME="`uname -n`"
fi

if [ -z "${SERVER_PORT}" ]; then
  SERVER_PORT=9999
fi

export SERVER_NAME
export SERVER_PORT
export SERVER_KEY="/etc/pki/ca-trust/source/anchors/${SERVER_NAME}.key.pem"
export SERVER_CERT="/etc/pki/ca-trust/source/anchors/${SERVER_NAME}.cert.pem"

go run main.go
