#!/bin/bash


# when editing this file on windows, be sure to set the end of line sequence to LF,
# or it will cause an error in the likes of :
# bash: ./runTests.sh: /bin/bash^M: bad interpreter: No such file or directory

error=0

check-error()
{
        if [ "$?" -ne "0" ]; then
                error=1
        fi
}


cd main/ && go run main.go & # Start all servers, for context creator
cd auth/ && go mod download && go test ./auth_server
check-error
cd ../user && go mod download && go test ./user_server
check-error
cd ../chatroom && go mod download && go test ./chatroom_server
check-error
cd ../message && go mod download && go test ./message_server
check-error


exit $error