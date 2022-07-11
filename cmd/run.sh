#!/bin/bash

pid = `ps -ef | grep mingyuan.site | grep -v grep | awk '{print $2}'`

kill -9 $pid

/services/mingyuan.site/mingyuan.site