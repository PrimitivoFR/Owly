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


go mod download

go test ./auth/auth_server
check-error
go test ./user/user_server 
check-error
go test ./chatroom/chatroom_server 
check-error
go test ./message/message_server
check-error

exit $error