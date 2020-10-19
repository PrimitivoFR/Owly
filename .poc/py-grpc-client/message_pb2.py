# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: message.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='message.proto',
  package='message',
  syntax='proto3',
  serialized_options=b'Z\021message/messagepb',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\rmessage.proto\x12\x07message\"/\n\x0b\x41ttachments\x12\x10\n\x08\x66ileName\x18\x01 \x01(\t\x12\x0e\n\x06\x62ucket\x18\x02 \x01(\t\"\xd9\x01\n\x07Message\x12\x12\n\nauthorUUID\x18\x01 \x01(\t\x12\x12\n\nchatroomID\x18\x02 \x01(\t\x12\x12\n\nauthorNAME\x18\x03 \x01(\t\x12\x0f\n\x07\x63ontent\x18\x04 \x01(\t\x12\x11\n\ttimestamp\x18\x05 \x01(\x03\x12\x17\n\x0fhasFileAttached\x18\x06 \x01(\x08\x12$\n\x06\x61ttach\x18\x07 \x03(\x0b\x32\x14.message.Attachments\x12\x10\n\x08isAnswer\x18\x08 \x01(\x08\x12\x11\n\tanswersTo\x18\t \x01(\t\x12\n\n\x02id\x18\n \x01(\t\"5\n\x1fStreamMessagesByChatroomRequest\x12\x12\n\nchatroomID\x18\x01 \x01(\t\"E\n StreamMessagesByChatroomResponse\x12!\n\x07message\x18\x01 \x01(\x0b\x32\x10.message.Message\"7\n\x12SendMessageRequest\x12!\n\x07message\x18\x01 \x01(\x0b\x32\x10.message.Message\"&\n\x13SendMessageResponse\x12\x0f\n\x07success\x18\x01 \x01(\x08\"2\n\x1cGetMessagesByChatroomRequest\x12\x12\n\nchatroomID\x18\x01 \x01(\t\"C\n\x1dGetMessagesByChatroomResponse\x12\"\n\x08messages\x18\x01 \x03(\x0b\x32\x10.message.Message2\xbb\x02\n\x0eMessageService\x12s\n\x18StreamMessagesByChatroom\x12(.message.StreamMessagesByChatroomRequest\x1a).message.StreamMessagesByChatroomResponse\"\x00\x30\x01\x12J\n\x0bSendMessage\x12\x1b.message.SendMessageRequest\x1a\x1c.message.SendMessageResponse\"\x00\x12h\n\x15GetMessagesByChatroom\x12%.message.GetMessagesByChatroomRequest\x1a&.message.GetMessagesByChatroomResponse\"\x00\x42\x13Z\x11message/messagepbb\x06proto3'
)




_ATTACHMENTS = _descriptor.Descriptor(
  name='Attachments',
  full_name='message.Attachments',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='fileName', full_name='message.Attachments.fileName', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='bucket', full_name='message.Attachments.bucket', index=1,
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
  serialized_start=26,
  serialized_end=73,
)


_MESSAGE = _descriptor.Descriptor(
  name='Message',
  full_name='message.Message',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='authorUUID', full_name='message.Message.authorUUID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='chatroomID', full_name='message.Message.chatroomID', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='authorNAME', full_name='message.Message.authorNAME', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='content', full_name='message.Message.content', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='message.Message.timestamp', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='hasFileAttached', full_name='message.Message.hasFileAttached', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='attach', full_name='message.Message.attach', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='isAnswer', full_name='message.Message.isAnswer', index=7,
      number=8, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='answersTo', full_name='message.Message.answersTo', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='id', full_name='message.Message.id', index=9,
      number=10, type=9, cpp_type=9, label=1,
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
  serialized_start=76,
  serialized_end=293,
)


_STREAMMESSAGESBYCHATROOMREQUEST = _descriptor.Descriptor(
  name='StreamMessagesByChatroomRequest',
  full_name='message.StreamMessagesByChatroomRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='chatroomID', full_name='message.StreamMessagesByChatroomRequest.chatroomID', index=0,
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
  serialized_start=295,
  serialized_end=348,
)


_STREAMMESSAGESBYCHATROOMRESPONSE = _descriptor.Descriptor(
  name='StreamMessagesByChatroomResponse',
  full_name='message.StreamMessagesByChatroomResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='message', full_name='message.StreamMessagesByChatroomResponse.message', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=350,
  serialized_end=419,
)


_SENDMESSAGEREQUEST = _descriptor.Descriptor(
  name='SendMessageRequest',
  full_name='message.SendMessageRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='message', full_name='message.SendMessageRequest.message', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
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
  serialized_start=421,
  serialized_end=476,
)


_SENDMESSAGERESPONSE = _descriptor.Descriptor(
  name='SendMessageResponse',
  full_name='message.SendMessageResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='success', full_name='message.SendMessageResponse.success', index=0,
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
  serialized_start=478,
  serialized_end=516,
)


_GETMESSAGESBYCHATROOMREQUEST = _descriptor.Descriptor(
  name='GetMessagesByChatroomRequest',
  full_name='message.GetMessagesByChatroomRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='chatroomID', full_name='message.GetMessagesByChatroomRequest.chatroomID', index=0,
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
  serialized_start=518,
  serialized_end=568,
)


_GETMESSAGESBYCHATROOMRESPONSE = _descriptor.Descriptor(
  name='GetMessagesByChatroomResponse',
  full_name='message.GetMessagesByChatroomResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='messages', full_name='message.GetMessagesByChatroomResponse.messages', index=0,
      number=1, type=11, cpp_type=10, label=3,
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
  serialized_start=570,
  serialized_end=637,
)

_MESSAGE.fields_by_name['attach'].message_type = _ATTACHMENTS
_STREAMMESSAGESBYCHATROOMRESPONSE.fields_by_name['message'].message_type = _MESSAGE
_SENDMESSAGEREQUEST.fields_by_name['message'].message_type = _MESSAGE
_GETMESSAGESBYCHATROOMRESPONSE.fields_by_name['messages'].message_type = _MESSAGE
DESCRIPTOR.message_types_by_name['Attachments'] = _ATTACHMENTS
DESCRIPTOR.message_types_by_name['Message'] = _MESSAGE
DESCRIPTOR.message_types_by_name['StreamMessagesByChatroomRequest'] = _STREAMMESSAGESBYCHATROOMREQUEST
DESCRIPTOR.message_types_by_name['StreamMessagesByChatroomResponse'] = _STREAMMESSAGESBYCHATROOMRESPONSE
DESCRIPTOR.message_types_by_name['SendMessageRequest'] = _SENDMESSAGEREQUEST
DESCRIPTOR.message_types_by_name['SendMessageResponse'] = _SENDMESSAGERESPONSE
DESCRIPTOR.message_types_by_name['GetMessagesByChatroomRequest'] = _GETMESSAGESBYCHATROOMREQUEST
DESCRIPTOR.message_types_by_name['GetMessagesByChatroomResponse'] = _GETMESSAGESBYCHATROOMRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Attachments = _reflection.GeneratedProtocolMessageType('Attachments', (_message.Message,), {
  'DESCRIPTOR' : _ATTACHMENTS,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:message.Attachments)
  })
_sym_db.RegisterMessage(Attachments)

Message = _reflection.GeneratedProtocolMessageType('Message', (_message.Message,), {
  'DESCRIPTOR' : _MESSAGE,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:message.Message)
  })
_sym_db.RegisterMessage(Message)

StreamMessagesByChatroomRequest = _reflection.GeneratedProtocolMessageType('StreamMessagesByChatroomRequest', (_message.Message,), {
  'DESCRIPTOR' : _STREAMMESSAGESBYCHATROOMREQUEST,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:message.StreamMessagesByChatroomRequest)
  })
_sym_db.RegisterMessage(StreamMessagesByChatroomRequest)

StreamMessagesByChatroomResponse = _reflection.GeneratedProtocolMessageType('StreamMessagesByChatroomResponse', (_message.Message,), {
  'DESCRIPTOR' : _STREAMMESSAGESBYCHATROOMRESPONSE,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:message.StreamMessagesByChatroomResponse)
  })
_sym_db.RegisterMessage(StreamMessagesByChatroomResponse)

SendMessageRequest = _reflection.GeneratedProtocolMessageType('SendMessageRequest', (_message.Message,), {
  'DESCRIPTOR' : _SENDMESSAGEREQUEST,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:message.SendMessageRequest)
  })
_sym_db.RegisterMessage(SendMessageRequest)

SendMessageResponse = _reflection.GeneratedProtocolMessageType('SendMessageResponse', (_message.Message,), {
  'DESCRIPTOR' : _SENDMESSAGERESPONSE,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:message.SendMessageResponse)
  })
_sym_db.RegisterMessage(SendMessageResponse)

GetMessagesByChatroomRequest = _reflection.GeneratedProtocolMessageType('GetMessagesByChatroomRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETMESSAGESBYCHATROOMREQUEST,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:message.GetMessagesByChatroomRequest)
  })
_sym_db.RegisterMessage(GetMessagesByChatroomRequest)

GetMessagesByChatroomResponse = _reflection.GeneratedProtocolMessageType('GetMessagesByChatroomResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETMESSAGESBYCHATROOMRESPONSE,
  '__module__' : 'message_pb2'
  # @@protoc_insertion_point(class_scope:message.GetMessagesByChatroomResponse)
  })
_sym_db.RegisterMessage(GetMessagesByChatroomResponse)


DESCRIPTOR._options = None

_MESSAGESERVICE = _descriptor.ServiceDescriptor(
  name='MessageService',
  full_name='message.MessageService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=640,
  serialized_end=955,
  methods=[
  _descriptor.MethodDescriptor(
    name='StreamMessagesByChatroom',
    full_name='message.MessageService.StreamMessagesByChatroom',
    index=0,
    containing_service=None,
    input_type=_STREAMMESSAGESBYCHATROOMREQUEST,
    output_type=_STREAMMESSAGESBYCHATROOMRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='SendMessage',
    full_name='message.MessageService.SendMessage',
    index=1,
    containing_service=None,
    input_type=_SENDMESSAGEREQUEST,
    output_type=_SENDMESSAGERESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetMessagesByChatroom',
    full_name='message.MessageService.GetMessagesByChatroom',
    index=2,
    containing_service=None,
    input_type=_GETMESSAGESBYCHATROOMREQUEST,
    output_type=_GETMESSAGESBYCHATROOMRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_MESSAGESERVICE)

DESCRIPTOR.services_by_name['MessageService'] = _MESSAGESERVICE

# @@protoc_insertion_point(module_scope)
