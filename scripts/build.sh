#!/bin/bash
set -e

go build -o reporeport .

cp reporeport ./dist/reporeport 

rm -f /usr/bin/reporeport

rm -f /usr/local/bin/reporeport

sudo mv reporeport /usr/local/bin