#!/bin/bash
kill -9 `ps -ef | grep go | grep -v grep | awk '{ print $2 }'`