/* tslint:disable */
/* eslint-disable */
//
// THIS IS A GENERATED FILE
// DO NOT MODIFY IT! YOUR CHANGES WILL BE LOST
import { GrpcMessage, RecursivePartial } from '@ngx-grpc/common';
import { BinaryReader, BinaryWriter, ByteSource } from 'google-protobuf';

/**
 * Message implementation for user.User
 */
export class User implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new User();
    User.deserializeBinaryFromReader(instance, new BinaryReader(bytes));
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: User) {
    _instance.uuid = _instance.uuid || '';
    _instance.username = _instance.username || '';
    _instance.email = _instance.email || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(_instance: User, _reader: BinaryReader) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.uuid = _reader.readString();
          break;
        case 2:
          _instance.username = _reader.readString();
          break;
        case 3:
          _instance.email = _reader.readString();
          break;
        default:
          _reader.skipField();
      }
    }

    User.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(_instance: User, _writer: BinaryWriter) {
    if (_instance.uuid) {
      _writer.writeString(1, _instance.uuid);
    }
    if (_instance.username) {
      _writer.writeString(2, _instance.username);
    }
    if (_instance.email) {
      _writer.writeString(3, _instance.email);
    }
  }

  private _uuid?: string;
  private _username?: string;
  private _email?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of User to deeply clone from
   */
  constructor(_value?: RecursivePartial<User>) {
    _value = _value || {};
    this.uuid = _value.uuid;
    this.username = _value.username;
    this.email = _value.email;
    User.refineValues(this);
  }
  get uuid(): string | undefined {
    return this._uuid;
  }
  set uuid(value: string | undefined) {
    this._uuid = value;
  }
  get username(): string | undefined {
    return this._username;
  }
  set username(value: string | undefined) {
    this._username = value;
  }
  get email(): string | undefined {
    return this._email;
  }
  set email(value: string | undefined) {
    this._email = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    User.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): User.AsObject {
    return {
      uuid: this.uuid,
      username: this.username,
      email: this.email
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
export module User {
  /**
   * Standard JavaScript object representation for User
   */
  export interface AsObject {
    uuid?: string;
    username?: string;
    email?: string;
  }
}

/**
 * Message implementation for user.CreateNewUserRequest
 */
export class CreateNewUserRequest implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new CreateNewUserRequest();
    CreateNewUserRequest.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: CreateNewUserRequest) {
    _instance.firstName = _instance.firstName || '';
    _instance.lastName = _instance.lastName || '';
    _instance.username = _instance.username || '';
    _instance.email = _instance.email || '';
    _instance.password = _instance.password || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: CreateNewUserRequest,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.firstName = _reader.readString();
          break;
        case 2:
          _instance.lastName = _reader.readString();
          break;
        case 3:
          _instance.username = _reader.readString();
          break;
        case 4:
          _instance.email = _reader.readString();
          break;
        case 5:
          _instance.password = _reader.readString();
          break;
        default:
          _reader.skipField();
      }
    }

    CreateNewUserRequest.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: CreateNewUserRequest,
    _writer: BinaryWriter
  ) {
    if (_instance.firstName) {
      _writer.writeString(1, _instance.firstName);
    }
    if (_instance.lastName) {
      _writer.writeString(2, _instance.lastName);
    }
    if (_instance.username) {
      _writer.writeString(3, _instance.username);
    }
    if (_instance.email) {
      _writer.writeString(4, _instance.email);
    }
    if (_instance.password) {
      _writer.writeString(5, _instance.password);
    }
  }

  private _firstName?: string;
  private _lastName?: string;
  private _username?: string;
  private _email?: string;
  private _password?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of CreateNewUserRequest to deeply clone from
   */
  constructor(_value?: RecursivePartial<CreateNewUserRequest>) {
    _value = _value || {};
    this.firstName = _value.firstName;
    this.lastName = _value.lastName;
    this.username = _value.username;
    this.email = _value.email;
    this.password = _value.password;
    CreateNewUserRequest.refineValues(this);
  }
  get firstName(): string | undefined {
    return this._firstName;
  }
  set firstName(value: string | undefined) {
    this._firstName = value;
  }
  get lastName(): string | undefined {
    return this._lastName;
  }
  set lastName(value: string | undefined) {
    this._lastName = value;
  }
  get username(): string | undefined {
    return this._username;
  }
  set username(value: string | undefined) {
    this._username = value;
  }
  get email(): string | undefined {
    return this._email;
  }
  set email(value: string | undefined) {
    this._email = value;
  }
  get password(): string | undefined {
    return this._password;
  }
  set password(value: string | undefined) {
    this._password = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    CreateNewUserRequest.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): CreateNewUserRequest.AsObject {
    return {
      firstName: this.firstName,
      lastName: this.lastName,
      username: this.username,
      email: this.email,
      password: this.password
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
export module CreateNewUserRequest {
  /**
   * Standard JavaScript object representation for CreateNewUserRequest
   */
  export interface AsObject {
    firstName?: string;
    lastName?: string;
    username?: string;
    email?: string;
    password?: string;
  }
}

/**
 * Message implementation for user.CreateNewUserResponse
 */
export class CreateNewUserResponse implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new CreateNewUserResponse();
    CreateNewUserResponse.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: CreateNewUserResponse) {
    _instance.success = _instance.success || false;
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: CreateNewUserResponse,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.success = _reader.readBool();
          break;
        default:
          _reader.skipField();
      }
    }

    CreateNewUserResponse.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: CreateNewUserResponse,
    _writer: BinaryWriter
  ) {
    if (_instance.success) {
      _writer.writeBool(1, _instance.success);
    }
  }

  private _success?: boolean;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of CreateNewUserResponse to deeply clone from
   */
  constructor(_value?: RecursivePartial<CreateNewUserResponse>) {
    _value = _value || {};
    this.success = _value.success;
    CreateNewUserResponse.refineValues(this);
  }
  get success(): boolean | undefined {
    return this._success;
  }
  set success(value: boolean | undefined) {
    this._success = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    CreateNewUserResponse.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): CreateNewUserResponse.AsObject {
    return {
      success: this.success
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
export module CreateNewUserResponse {
  /**
   * Standard JavaScript object representation for CreateNewUserResponse
   */
  export interface AsObject {
    success?: boolean;
  }
}

/**
 * Message implementation for user.SearchUserByUsernameRequest
 */
export class SearchUserByUsernameRequest implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new SearchUserByUsernameRequest();
    SearchUserByUsernameRequest.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: SearchUserByUsernameRequest) {
    _instance.username = _instance.username || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: SearchUserByUsernameRequest,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.username = _reader.readString();
          break;
        default:
          _reader.skipField();
      }
    }

    SearchUserByUsernameRequest.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: SearchUserByUsernameRequest,
    _writer: BinaryWriter
  ) {
    if (_instance.username) {
      _writer.writeString(1, _instance.username);
    }
  }

  private _username?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of SearchUserByUsernameRequest to deeply clone from
   */
  constructor(_value?: RecursivePartial<SearchUserByUsernameRequest>) {
    _value = _value || {};
    this.username = _value.username;
    SearchUserByUsernameRequest.refineValues(this);
  }
  get username(): string | undefined {
    return this._username;
  }
  set username(value: string | undefined) {
    this._username = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    SearchUserByUsernameRequest.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): SearchUserByUsernameRequest.AsObject {
    return {
      username: this.username
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
export module SearchUserByUsernameRequest {
  /**
   * Standard JavaScript object representation for SearchUserByUsernameRequest
   */
  export interface AsObject {
    username?: string;
  }
}

/**
 * Message implementation for user.SearchUserByUsernameResponse
 */
export class SearchUserByUsernameResponse implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new SearchUserByUsernameResponse();
    SearchUserByUsernameResponse.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: SearchUserByUsernameResponse) {
    _instance.results = _instance.results || [];
    _instance.count = _instance.count || '0';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: SearchUserByUsernameResponse,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          const messageInitializer1 = new User();
          _reader.readMessage(
            messageInitializer1,
            User.deserializeBinaryFromReader
          );
          (_instance.results = _instance.results || []).push(
            messageInitializer1
          );
          break;
        case 2:
          _instance.count = _reader.readInt64String();
          break;
        default:
          _reader.skipField();
      }
    }

    SearchUserByUsernameResponse.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: SearchUserByUsernameResponse,
    _writer: BinaryWriter
  ) {
    if (_instance.results && _instance.results.length) {
      _writer.writeRepeatedMessage(
        1,
        _instance.results as any,
        User.serializeBinaryToWriter
      );
    }
    if (_instance.count) {
      _writer.writeInt64String(2, _instance.count);
    }
  }

  private _results?: User[];
  private _count?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of SearchUserByUsernameResponse to deeply clone from
   */
  constructor(_value?: RecursivePartial<SearchUserByUsernameResponse>) {
    _value = _value || {};
    this.results = (_value.results || []).map(m => new User(m));
    this.count = _value.count;
    SearchUserByUsernameResponse.refineValues(this);
  }
  get results(): User[] | undefined {
    return this._results;
  }
  set results(value: User[] | undefined) {
    this._results = value;
  }
  get count(): string | undefined {
    return this._count;
  }
  set count(value: string | undefined) {
    this._count = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    SearchUserByUsernameResponse.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): SearchUserByUsernameResponse.AsObject {
    return {
      results: (this.results || []).map(m => m.toObject()),
      count: this.count
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
export module SearchUserByUsernameResponse {
  /**
   * Standard JavaScript object representation for SearchUserByUsernameResponse
   */
  export interface AsObject {
    results?: User.AsObject[];
    count?: string;
  }
}