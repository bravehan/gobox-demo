#!/bin/bash

curPath=`dirname $0`
cd $curPath
prjHome=`pwd`

/usr/local/bin/rigger -rconfDir=$prjHome/conf/rigger/ -logLevel=2 prjHome=$prjHome
