#!/bin/bash


# when editing this file on windows, be sure to set the end of line sequence to LF,
# or it will cause an error in the likes of :
# bash: ./runTests.sh: /bin/bash^M: bad interpreter: No such file or directory

go mod download
go test ./auth/auth_server
go test ./user/user_server
go test ./chatroom/chatroom_server
go test ./message/message_server
