#!/bin/bash

base_path=$(cd `dirname $0`; pwd)

echo "0 10,14,19 * * *    cd $base_path && sh empty-recycle-bin.sh"
