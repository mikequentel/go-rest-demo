HTML_DOCS_DIR=htmldoc
SERVER_KEY=certs/key.pem
SERVER_CERT=certs/cert.pem
SERVER_PORT=9999

.PHONY: all default fmt lint clean build doc start stop test

default: all

fmt: 
	go fmt .

lint:
	golint .

clean:
	go clean .

build:
	go build

doc:
	mkdir -p ${HTML_DOCS_DIR}
	godoc .
	mv *.html ${HTML_DOCS_DIR}/

start:
	export SHELL=/bin/bash
	export SERVER_KEY=${SERVER_KEY}
	export SERVER_CERT=${SERVER_CERT}
	export SERVER_PORT=${SERVER_PORT}
	export DB_USER=${DB_USER}
	export DB_PASSWD=${DB_PASSWD}
	export DB_SID=${DB_SID}
	 
	#@if [ -z "${DB_PASSWD}" ]; then export DB_PASSWD=`source set_db_passwd.sh`; else export DB_PASSWD=${DB_PASSWD};fi
	
	go run main.go

stop:
	# 

test:
	go test .
all: fmt lint clean doc start test
