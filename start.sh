#!/bin/bash
rm -f pid
rm -f site.log

nohup go run site.go > site.log 2>&1&
echo $! > pid