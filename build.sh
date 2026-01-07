#!/bin/bash
set -e
go build main.go && mv main reporeport && cp reporeport ./dist/reporeport  && sudo mv reporeport /usr/bin