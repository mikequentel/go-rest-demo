#!/bin/bash

SERVER=$1
if [ -z "${SERVER}" ]; then
  SERVER=`uname -n`
fi

SUBJECT="/C=CA/O=Go Rest Demo/CN=${SERVER}"

openssl req -x509 -newkey rsa:4096 -nodes -keyout ${SERVER}.key.pem -out ${SERVER}.cert.pem -subj "${SUBJECT}" -days 365

sudo cp ${SERVER}.cert.pem /etc/pki/ca-trust/source/anchors/
sudo cp ${SERVER}.key.pem /etc/pki/ca-trust/source/anchors/
