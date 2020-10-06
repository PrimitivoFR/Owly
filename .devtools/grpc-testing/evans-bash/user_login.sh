#!/bin/bash
user=$1
pass=$2

if [ $# -lt 2 ] ; then
    echo "This script needs two arguments, user and password."
    exit 1
fi

echo "{\"username\": \"$user\", \"password\": \"$pass\" }" | evans cli -r -p 50054 call LoginUser