# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import auth_pb2 as auth__pb2


class AuthServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateNewUser = channel.unary_unary(
                '/auth.AuthService/CreateNewUser',
                request_serializer=auth__pb2.CreateNewUserRequest.SerializeToString,
                response_deserializer=auth__pb2.CreateNewUserResponse.FromString,
                )
        self.LoginUser = channel.unary_unary(
                '/auth.AuthService/LoginUser',
                request_serializer=auth__pb2.LoginUserRequest.SerializeToString,
                response_deserializer=auth__pb2.LoginUserResponse.FromString,
                )


class AuthServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateNewUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def LoginUser(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AuthServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateNewUser': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateNewUser,
                    request_deserializer=auth__pb2.CreateNewUserRequest.FromString,
                    response_serializer=auth__pb2.CreateNewUserResponse.SerializeToString,
            ),
            'LoginUser': grpc.unary_unary_rpc_method_handler(
                    servicer.LoginUser,
                    request_deserializer=auth__pb2.LoginUserRequest.FromString,
                    response_serializer=auth__pb2.LoginUserResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'auth.AuthService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AuthService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateNewUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/auth.AuthService/CreateNewUser',
            auth__pb2.CreateNewUserRequest.SerializeToString,
            auth__pb2.CreateNewUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def LoginUser(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/auth.AuthService/LoginUser',
            auth__pb2.LoginUserRequest.SerializeToString,
            auth__pb2.LoginUserResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
