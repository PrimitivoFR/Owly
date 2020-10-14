# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chatroom.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='chatroom.proto',
  package='chatroom',
  syntax='proto3',
  serialized_options=b'Z\023chatroom/chatroompb',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x0e\x63hatroom.proto\x12\x08\x63hatroom\"3\n\x08\x43hatroom\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\r\n\x05owner\x18\x03 \x01(\t\"4\n\x15\x43reateChatroomRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05users\x18\x02 \x03(\t\"5\n\x16\x43reateChatroomResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\n\n\x02ID\x18\x02 \x01(\t\"\x1b\n\x19GetChatroomsByUserRequest\"c\n\x1aGetChatroomsByUserResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x12\r\n\x05\x63ount\x18\x02 \x01(\x03\x12%\n\tchatrooms\x18\x03 \x03(\x0b\x32\x12.chatroom.Chatroom\"#\n\x15\x44\x65leteChatroomRequest\x12\n\n\x02id\x18\x01 \x01(\t\")\n\x16\x44\x65leteChatroomResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\"\"\n\x14LeaveChatroomRequest\x12\n\n\x02id\x18\x01 \x01(\t\"(\n\x15LeaveChatroomResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\x32\xf6\x02\n\x0f\x43hatroomService\x12U\n\x0e\x43reateChatroom\x12\x1f.chatroom.CreateChatroomRequest\x1a .chatroom.CreateChatroomResponse\"\x00\x12\x61\n\x12GetChatroomsByUser\x12#.chatroom.GetChatroomsByUserRequest\x1a$.chatroom.GetChatroomsByUserResponse\"\x00\x12U\n\x0e\x44\x65leteChatroom\x12\x1f.chatroom.DeleteChatroomRequest\x1a .chatroom.DeleteChatroomResponse\"\x00\x12R\n\rLeaveChatroom\x12\x1e.chatroom.LeaveChatroomRequest\x1a\x1f.chatroom.LeaveChatroomResponse\"\x00\x42\x15Z\x13\x63hatroom/chatroompbb\x06proto3'
)




_CHATROOM = _descriptor.Descriptor(
  name='Chatroom',
  full_name='chatroom.Chatroom',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='chatroom.Chatroom.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='chatroom.Chatroom.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='owner', full_name='chatroom.Chatroom.owner', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=28,
  serialized_end=79,
)


_CREATECHATROOMREQUEST = _descriptor.Descriptor(
  name='CreateChatroomRequest',
  full_name='chatroom.CreateChatroomRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='chatroom.CreateChatroomRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='users', full_name='chatroom.CreateChatroomRequest.users', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=81,
  serialized_end=133,
)


_CREATECHATROOMRESPONSE = _descriptor.Descriptor(
  name='CreateChatroomResponse',
  full_name='chatroom.CreateChatroomResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='success', full_name='chatroom.CreateChatroomResponse.success', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ID', full_name='chatroom.CreateChatroomResponse.ID', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=135,
  serialized_end=188,
)


_GETCHATROOMSBYUSERREQUEST = _descriptor.Descriptor(
  name='GetChatroomsByUserRequest',
  full_name='chatroom.GetChatroomsByUserRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=190,
  serialized_end=217,
)


_GETCHATROOMSBYUSERRESPONSE = _descriptor.Descriptor(
  name='GetChatroomsByUserResponse',
  full_name='chatroom.GetChatroomsByUserResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='success', full_name='chatroom.GetChatroomsByUserResponse.success', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='count', full_name='chatroom.GetChatroomsByUserResponse.count', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='chatrooms', full_name='chatroom.GetChatroomsByUserResponse.chatrooms', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=219,
  serialized_end=318,
)


_DELETECHATROOMREQUEST = _descriptor.Descriptor(
  name='DeleteChatroomRequest',
  full_name='chatroom.DeleteChatroomRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='chatroom.DeleteChatroomRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=320,
  serialized_end=355,
)


_DELETECHATROOMRESPONSE = _descriptor.Descriptor(
  name='DeleteChatroomResponse',
  full_name='chatroom.DeleteChatroomResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='success', full_name='chatroom.DeleteChatroomResponse.success', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=357,
  serialized_end=398,
)


_LEAVECHATROOMREQUEST = _descriptor.Descriptor(
  name='LeaveChatroomRequest',
  full_name='chatroom.LeaveChatroomRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='chatroom.LeaveChatroomRequest.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=400,
  serialized_end=434,
)


_LEAVECHATROOMRESPONSE = _descriptor.Descriptor(
  name='LeaveChatroomResponse',
  full_name='chatroom.LeaveChatroomResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='success', full_name='chatroom.LeaveChatroomResponse.success', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=436,
  serialized_end=476,
)

_GETCHATROOMSBYUSERRESPONSE.fields_by_name['chatrooms'].message_type = _CHATROOM
DESCRIPTOR.message_types_by_name['Chatroom'] = _CHATROOM
DESCRIPTOR.message_types_by_name['CreateChatroomRequest'] = _CREATECHATROOMREQUEST
DESCRIPTOR.message_types_by_name['CreateChatroomResponse'] = _CREATECHATROOMRESPONSE
DESCRIPTOR.message_types_by_name['GetChatroomsByUserRequest'] = _GETCHATROOMSBYUSERREQUEST
DESCRIPTOR.message_types_by_name['GetChatroomsByUserResponse'] = _GETCHATROOMSBYUSERRESPONSE
DESCRIPTOR.message_types_by_name['DeleteChatroomRequest'] = _DELETECHATROOMREQUEST
DESCRIPTOR.message_types_by_name['DeleteChatroomResponse'] = _DELETECHATROOMRESPONSE
DESCRIPTOR.message_types_by_name['LeaveChatroomRequest'] = _LEAVECHATROOMREQUEST
DESCRIPTOR.message_types_by_name['LeaveChatroomResponse'] = _LEAVECHATROOMRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Chatroom = _reflection.GeneratedProtocolMessageType('Chatroom', (_message.Message,), {
  'DESCRIPTOR' : _CHATROOM,
  '__module__' : 'chatroom_pb2'
  # @@protoc_insertion_point(class_scope:chatroom.Chatroom)
  })
_sym_db.RegisterMessage(Chatroom)

CreateChatroomRequest = _reflection.GeneratedProtocolMessageType('CreateChatroomRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATECHATROOMREQUEST,
  '__module__' : 'chatroom_pb2'
  # @@protoc_insertion_point(class_scope:chatroom.CreateChatroomRequest)
  })
_sym_db.RegisterMessage(CreateChatroomRequest)

CreateChatroomResponse = _reflection.GeneratedProtocolMessageType('CreateChatroomResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATECHATROOMRESPONSE,
  '__module__' : 'chatroom_pb2'
  # @@protoc_insertion_point(class_scope:chatroom.CreateChatroomResponse)
  })
_sym_db.RegisterMessage(CreateChatroomResponse)

GetChatroomsByUserRequest = _reflection.GeneratedProtocolMessageType('GetChatroomsByUserRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETCHATROOMSBYUSERREQUEST,
  '__module__' : 'chatroom_pb2'
  # @@protoc_insertion_point(class_scope:chatroom.GetChatroomsByUserRequest)
  })
_sym_db.RegisterMessage(GetChatroomsByUserRequest)

GetChatroomsByUserResponse = _reflection.GeneratedProtocolMessageType('GetChatroomsByUserResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETCHATROOMSBYUSERRESPONSE,
  '__module__' : 'chatroom_pb2'
  # @@protoc_insertion_point(class_scope:chatroom.GetChatroomsByUserResponse)
  })
_sym_db.RegisterMessage(GetChatroomsByUserResponse)

DeleteChatroomRequest = _reflection.GeneratedProtocolMessageType('DeleteChatroomRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETECHATROOMREQUEST,
  '__module__' : 'chatroom_pb2'
  # @@protoc_insertion_point(class_scope:chatroom.DeleteChatroomRequest)
  })
_sym_db.RegisterMessage(DeleteChatroomRequest)

DeleteChatroomResponse = _reflection.GeneratedProtocolMessageType('DeleteChatroomResponse', (_message.Message,), {
  'DESCRIPTOR' : _DELETECHATROOMRESPONSE,
  '__module__' : 'chatroom_pb2'
  # @@protoc_insertion_point(class_scope:chatroom.DeleteChatroomResponse)
  })
_sym_db.RegisterMessage(DeleteChatroomResponse)

LeaveChatroomRequest = _reflection.GeneratedProtocolMessageType('LeaveChatroomRequest', (_message.Message,), {
  'DESCRIPTOR' : _LEAVECHATROOMREQUEST,
  '__module__' : 'chatroom_pb2'
  # @@protoc_insertion_point(class_scope:chatroom.LeaveChatroomRequest)
  })
_sym_db.RegisterMessage(LeaveChatroomRequest)

LeaveChatroomResponse = _reflection.GeneratedProtocolMessageType('LeaveChatroomResponse', (_message.Message,), {
  'DESCRIPTOR' : _LEAVECHATROOMRESPONSE,
  '__module__' : 'chatroom_pb2'
  # @@protoc_insertion_point(class_scope:chatroom.LeaveChatroomResponse)
  })
_sym_db.RegisterMessage(LeaveChatroomResponse)


DESCRIPTOR._options = None

_CHATROOMSERVICE = _descriptor.ServiceDescriptor(
  name='ChatroomService',
  full_name='chatroom.ChatroomService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=479,
  serialized_end=853,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateChatroom',
    full_name='chatroom.ChatroomService.CreateChatroom',
    index=0,
    containing_service=None,
    input_type=_CREATECHATROOMREQUEST,
    output_type=_CREATECHATROOMRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetChatroomsByUser',
    full_name='chatroom.ChatroomService.GetChatroomsByUser',
    index=1,
    containing_service=None,
    input_type=_GETCHATROOMSBYUSERREQUEST,
    output_type=_GETCHATROOMSBYUSERRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='DeleteChatroom',
    full_name='chatroom.ChatroomService.DeleteChatroom',
    index=2,
    containing_service=None,
    input_type=_DELETECHATROOMREQUEST,
    output_type=_DELETECHATROOMRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='LeaveChatroom',
    full_name='chatroom.ChatroomService.LeaveChatroom',
    index=3,
    containing_service=None,
    input_type=_LEAVECHATROOMREQUEST,
    output_type=_LEAVECHATROOMRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_CHATROOMSERVICE)

DESCRIPTOR.services_by_name['ChatroomService'] = _CHATROOMSERVICE

# @@protoc_insertion_point(module_scope)
