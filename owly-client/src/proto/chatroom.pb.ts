/* tslint:disable */
/* eslint-disable */
//
// THIS IS A GENERATED FILE
// DO NOT MODIFY IT! YOUR CHANGES WILL BE LOST
import { GrpcMessage, RecursivePartial } from '@ngx-grpc/common';
import { BinaryReader, BinaryWriter, ByteSource } from 'google-protobuf';

/**
 * Message implementation for chatroom.Chatroom
 */
export class Chatroom implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new Chatroom();
    Chatroom.deserializeBinaryFromReader(instance, new BinaryReader(bytes));
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: Chatroom) {
    _instance.name = _instance.name || '';
    _instance.id = _instance.id || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: Chatroom,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.name = _reader.readString();
          break;
        case 2:
          _instance.id = _reader.readString();
          break;
        default:
          _reader.skipField();
      }
    }

    Chatroom.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(_instance: Chatroom, _writer: BinaryWriter) {
    if (_instance.name) {
      _writer.writeString(1, _instance.name);
    }
    if (_instance.id) {
      _writer.writeString(2, _instance.id);
    }
  }

  private _name?: string;
  private _id?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of Chatroom to deeply clone from
   */
  constructor(_value?: RecursivePartial<Chatroom>) {
    _value = _value || {};
    this.name = _value.name;
    this.id = _value.id;
    Chatroom.refineValues(this);
  }
  get name(): string | undefined {
    return this._name;
  }
  set name(value: string | undefined) {
    this._name = value;
  }
  get id(): string | undefined {
    return this._id;
  }
  set id(value: string | undefined) {
    this._id = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    Chatroom.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): Chatroom.AsObject {
    return {
      name: this.name,
      id: this.id
    };
  }

  /**
   * JSON serializer
   * Only intended to be used by `JSON.stringify` function. If you want to cast message to standard JavaScript object, use `toObject()` instead
   */
  toJSON() {
    return this.toObject();
  }
}
export module Chatroom {
  /**
   * Standard JavaScript object representation for Chatroom
   */
  export interface AsObject {
    name?: string;
    id?: string;
  }
}

/**
 * Message implementation for chatroom.CreateChatroomRequest
 */
export class CreateChatroomRequest implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new CreateChatroomRequest();
    CreateChatroomRequest.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: CreateChatroomRequest) {
    _instance.name = _instance.name || '';
    _instance.users = _instance.users || [];
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: CreateChatroomRequest,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.name = _reader.readString();
          break;
        case 2:
          (_instance.users = _instance.users || []).push(_reader.readString());
          break;
        default:
          _reader.skipField();
      }
    }

    CreateChatroomRequest.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: CreateChatroomRequest,
    _writer: BinaryWriter
  ) {
    if (_instance.name) {
      _writer.writeString(1, _instance.name);
    }
    if (_instance.users && _instance.users.length) {
      _writer.writeRepeatedString(2, _instance.users);
    }
  }

  private _name?: string;
  private _users?: string[];

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of CreateChatroomRequest to deeply clone from
   */
  constructor(_value?: RecursivePartial<CreateChatroomRequest>) {
    _value = _value || {};
    this.name = _value.name;
    this.users = (_value.users || []).slice();
    CreateChatroomRequest.refineValues(this);
  }
  get name(): string | undefined {
    return this._name;
  }
  set name(value: string | undefined) {
    this._name = value;
  }
  get users(): string[] | undefined {
    return this._users;
  }
  set users(value: string[] | undefined) {
    this._users = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    CreateChatroomRequest.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): CreateChatroomRequest.AsObject {
    return {
      name: this.name,
      users: (this.users || []).slice()
    };
  }

  /**
   * JSON serializer
   * Only intended to be used by `JSON.stringify` function. If you want to cast message to standard JavaScript object, use `toObject()` instead
   */
  toJSON() {
    return this.toObject();
  }
}
export module CreateChatroomRequest {
  /**
   * Standard JavaScript object representation for CreateChatroomRequest
   */
  export interface AsObject {
    name?: string;
    users?: string[];
  }
}

/**
 * Message implementation for chatroom.CreateChatroomResponse
 */
export class CreateChatroomResponse implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new CreateChatroomResponse();
    CreateChatroomResponse.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: CreateChatroomResponse) {
    _instance.success = _instance.success || false;
    _instance.id = _instance.id || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: CreateChatroomResponse,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.success = _reader.readBool();
          break;
        case 2:
          _instance.id = _reader.readString();
          break;
        default:
          _reader.skipField();
      }
    }

    CreateChatroomResponse.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: CreateChatroomResponse,
    _writer: BinaryWriter
  ) {
    if (_instance.success) {
      _writer.writeBool(1, _instance.success);
    }
    if (_instance.id) {
      _writer.writeString(2, _instance.id);
    }
  }

  private _success?: boolean;
  private _id?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of CreateChatroomResponse to deeply clone from
   */
  constructor(_value?: RecursivePartial<CreateChatroomResponse>) {
    _value = _value || {};
    this.success = _value.success;
    this.id = _value.id;
    CreateChatroomResponse.refineValues(this);
  }
  get success(): boolean | undefined {
    return this._success;
  }
  set success(value: boolean | undefined) {
    this._success = value;
  }
  get id(): string | undefined {
    return this._id;
  }
  set id(value: string | undefined) {
    this._id = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    CreateChatroomResponse.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): CreateChatroomResponse.AsObject {
    return {
      success: this.success,
      id: this.id
    };
  }

  /**
   * JSON serializer
   * Only intended to be used by `JSON.stringify` function. If you want to cast message to standard JavaScript object, use `toObject()` instead
   */
  toJSON() {
    return this.toObject();
  }
}
export module CreateChatroomResponse {
  /**
   * Standard JavaScript object representation for CreateChatroomResponse
   */
  export interface AsObject {
    success?: boolean;
    id?: string;
  }
}

/**
 * Message implementation for chatroom.GetChatroomsByUserRequest
 */
export class GetChatroomsByUserRequest implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new GetChatroomsByUserRequest();
    GetChatroomsByUserRequest.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: GetChatroomsByUserRequest) {}

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: GetChatroomsByUserRequest,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        default:
          _reader.skipField();
      }
    }

    GetChatroomsByUserRequest.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: GetChatroomsByUserRequest,
    _writer: BinaryWriter
  ) {}

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of GetChatroomsByUserRequest to deeply clone from
   */
  constructor(_value?: RecursivePartial<GetChatroomsByUserRequest>) {
    _value = _value || {};
    GetChatroomsByUserRequest.refineValues(this);
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    GetChatroomsByUserRequest.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): GetChatroomsByUserRequest.AsObject {
    return {};
  }

  /**
   * JSON serializer
   * Only intended to be used by `JSON.stringify` function. If you want to cast message to standard JavaScript object, use `toObject()` instead
   */
  toJSON() {
    return this.toObject();
  }
}
export module GetChatroomsByUserRequest {
  /**
   * Standard JavaScript object representation for GetChatroomsByUserRequest
   */
  export interface AsObject {}
}

/**
 * Message implementation for chatroom.GetChatroomsByUserResponse
 */
export class GetChatroomsByUserResponse implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new GetChatroomsByUserResponse();
    GetChatroomsByUserResponse.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: GetChatroomsByUserResponse) {
    _instance.success = _instance.success || false;
    _instance.count = _instance.count || '0';
    _instance.chatrooms = _instance.chatrooms || [];
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: GetChatroomsByUserResponse,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.success = _reader.readBool();
          break;
        case 2:
          _instance.count = _reader.readInt64String();
          break;
        case 3:
          const messageInitializer3 = new Chatroom();
          _reader.readMessage(
            messageInitializer3,
            Chatroom.deserializeBinaryFromReader
          );
          (_instance.chatrooms = _instance.chatrooms || []).push(
            messageInitializer3
          );
          break;
        default:
          _reader.skipField();
      }
    }

    GetChatroomsByUserResponse.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: GetChatroomsByUserResponse,
    _writer: BinaryWriter
  ) {
    if (_instance.success) {
      _writer.writeBool(1, _instance.success);
    }
    if (_instance.count) {
      _writer.writeInt64String(2, _instance.count);
    }
    if (_instance.chatrooms && _instance.chatrooms.length) {
      _writer.writeRepeatedMessage(
        3,
        _instance.chatrooms as any,
        Chatroom.serializeBinaryToWriter
      );
    }
  }

  private _success?: boolean;
  private _count?: string;
  private _chatrooms?: Chatroom[];

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of GetChatroomsByUserResponse to deeply clone from
   */
  constructor(_value?: RecursivePartial<GetChatroomsByUserResponse>) {
    _value = _value || {};
    this.success = _value.success;
    this.count = _value.count;
    this.chatrooms = (_value.chatrooms || []).map(m => new Chatroom(m));
    GetChatroomsByUserResponse.refineValues(this);
  }
  get success(): boolean | undefined {
    return this._success;
  }
  set success(value: boolean | undefined) {
    this._success = value;
  }
  get count(): string | undefined {
    return this._count;
  }
  set count(value: string | undefined) {
    this._count = value;
  }
  get chatrooms(): Chatroom[] | undefined {
    return this._chatrooms;
  }
  set chatrooms(value: Chatroom[] | undefined) {
    this._chatrooms = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    GetChatroomsByUserResponse.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): GetChatroomsByUserResponse.AsObject {
    return {
      success: this.success,
      count: this.count,
      chatrooms: (this.chatrooms || []).map(m => m.toObject())
    };
  }

  /**
   * JSON serializer
   * Only intended to be used by `JSON.stringify` function. If you want to cast message to standard JavaScript object, use `toObject()` instead
   */
  toJSON() {
    return this.toObject();
  }
}
export module GetChatroomsByUserResponse {
  /**
   * Standard JavaScript object representation for GetChatroomsByUserResponse
   */
  export interface AsObject {
    success?: boolean;
    count?: string;
    chatrooms?: Chatroom.AsObject[];
  }
}
