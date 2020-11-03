from __future__ import print_function

import grpc

# Note : Pylance underline the *_pb2 and *_pb2_grpc imports
# if you don't open the current directory as root folder in vscode

import auth_pb2_grpc
import user_pb2_grpc
import chatroom_pb2_grpc

from auth_pb2_grpc import AuthServiceStub # typing
from user_pb2_grpc import UserServiceStub # typing
from chatroom_pb2_grpc import ChatroomServiceStub # typing

from auth_pb2 import LoginUserRequest, LoginUserResponse
from auth_pb2 import CreateNewUserRequest, CreateNewUserResponse

from user_pb2 import SearchUserByUsernameRequest, SearchUserByUsernameResponse

from chatroom_pb2 import CreateChatroomRequest, CreateChatroomResponse
from chatroom_pb2 import DeleteChatroomRequest, DeleteChatroomResponse
from chatroom_pb2 import LeaveChatroomRequest, LeaveChatroomResponse


def register(stub: AuthServiceStub, new_user: CreateNewUserRequest):
    try:
        stub.CreateNewUser(new_user)
        print(f'Successfully registered user {new_user.username}.')
    except Exception as e:
        # print(e)
        print(f'Account {new_user.username} (probably) already exists.')


def login(stub: AuthServiceStub, username, password):
    resp: LoginUserResponse = stub.LoginUser(
        LoginUserRequest(
            username=username,
            password=password,
        )
    )
    return resp.result.AccessToken


def search_user_by_username(stub: UserServiceStub, search_str, metadata):
    return stub.SearchUserByUsername(
        SearchUserByUsernameRequest(username=search_str),
        metadata=metadata
    )


def create_chatroom(stub: ChatroomServiceStub, name, users, metadata):
    pass


bhubbard = CreateNewUserRequest(
    first_name="bernard",
    last_name="hubbard",
    username="bhubbard",
    password="pass",
    email="bh@mail.com",
)

jval = CreateNewUserRequest(
    first_name="jean",
    last_name="valjean",
    username="jval",
    password="pass",
    email="jv@mail.com",
)

carbin = CreateNewUserRequest(
    first_name="paul",
    last_name="carbinet",
    username="carbi,",
    password="pass",
    email="cb@mail.com",
)


def run():

    users = [bhubbard, jval, carbin]
    user_tokens = {}
    users_metadata = {}
    user_ids = {}

    with grpc.insecure_channel('localhost:50054') as channel:
        stub = auth_pb2_grpc.AuthServiceStub(channel)

        for user in users:
            register(stub, user)

        user_tokens = {
            user.username: login(stub, user.username, user.password)
            for user in users
        }

        users_metadata = {
            user.username: [('authorization', user_tokens[user.username])]
            for user in users
        }

    with grpc.insecure_channel('localhost:50051') as channel:
        stub = user_pb2_grpc.UserServiceStub(channel)

        user_ids = {
            'bhubbard': search_user_by_username(stub, 'bhubbard', users_metadata['jval']).results[0].uuid,
            'jval': search_user_by_username(stub, 'jval', users_metadata['bhubbard']).results[0].uuid,
            'carbi': search_user_by_username(stub, 'carbin', users_metadata['bhubbard']).results[0].uuid,
        }

        print(f'{user_ids=}')


if __name__ == "__main__":
    run()
