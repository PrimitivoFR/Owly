/* tslint:disable */
/* eslint-disable */
//
// THIS IS A GENERATED FILE
// DO NOT MODIFY IT! YOUR CHANGES WILL BE LOST
import { GrpcMessage, RecursivePartial } from '@ngx-grpc/common';
import { BinaryReader, BinaryWriter, ByteSource } from 'google-protobuf';

/**
 * Message implementation for message.Attachments
 */
export class Attachments implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new Attachments();
    Attachments.deserializeBinaryFromReader(instance, new BinaryReader(bytes));
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: Attachments) {
    _instance.fileName = _instance.fileName || '';
    _instance.bucket = _instance.bucket || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: Attachments,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.fileName = _reader.readString();
          break;
        case 2:
          _instance.bucket = _reader.readString();
          break;
        default:
          _reader.skipField();
      }
    }

    Attachments.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: Attachments,
    _writer: BinaryWriter
  ) {
    if (_instance.fileName) {
      _writer.writeString(1, _instance.fileName);
    }
    if (_instance.bucket) {
      _writer.writeString(2, _instance.bucket);
    }
  }

  private _fileName?: string;
  private _bucket?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of Attachments to deeply clone from
   */
  constructor(_value?: RecursivePartial<Attachments>) {
    _value = _value || {};
    this.fileName = _value.fileName;
    this.bucket = _value.bucket;
    Attachments.refineValues(this);
  }
  get fileName(): string | undefined {
    return this._fileName;
  }
  set fileName(value: string | undefined) {
    this._fileName = value;
  }
  get bucket(): string | undefined {
    return this._bucket;
  }
  set bucket(value: string | undefined) {
    this._bucket = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    Attachments.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): Attachments.AsObject {
    return {
      fileName: this.fileName,
      bucket: this.bucket
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
export module Attachments {
  /**
   * Standard JavaScript object representation for Attachments
   */
  export interface AsObject {
    fileName?: string;
    bucket?: string;
  }
}

/**
 * Message implementation for message.Message
 */
export class Message implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new Message();
    Message.deserializeBinaryFromReader(instance, new BinaryReader(bytes));
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: Message) {
    _instance.authorUUID = _instance.authorUUID || '';
    _instance.chatroomID = _instance.chatroomID || '';
    _instance.authorNAME = _instance.authorNAME || '';
    _instance.content = _instance.content || '';
    _instance.timestamp = _instance.timestamp || '0';
    _instance.hasFileAttached = _instance.hasFileAttached || false;
    _instance.attach = _instance.attach || [];
    _instance.isAnswer = _instance.isAnswer || false;
    _instance.answersTo = _instance.answersTo || '';
    _instance.id = _instance.id || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: Message,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.authorUUID = _reader.readString();
          break;
        case 2:
          _instance.chatroomID = _reader.readString();
          break;
        case 3:
          _instance.authorNAME = _reader.readString();
          break;
        case 4:
          _instance.content = _reader.readString();
          break;
        case 5:
          _instance.timestamp = _reader.readInt64String();
          break;
        case 6:
          _instance.hasFileAttached = _reader.readBool();
          break;
        case 7:
          const messageInitializer7 = new Attachments();
          _reader.readMessage(
            messageInitializer7,
            Attachments.deserializeBinaryFromReader
          );
          (_instance.attach = _instance.attach || []).push(messageInitializer7);
          break;
        case 8:
          _instance.isAnswer = _reader.readBool();
          break;
        case 9:
          _instance.answersTo = _reader.readString();
          break;
        case 10:
          _instance.id = _reader.readString();
          break;
        default:
          _reader.skipField();
      }
    }

    Message.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(_instance: Message, _writer: BinaryWriter) {
    if (_instance.authorUUID) {
      _writer.writeString(1, _instance.authorUUID);
    }
    if (_instance.chatroomID) {
      _writer.writeString(2, _instance.chatroomID);
    }
    if (_instance.authorNAME) {
      _writer.writeString(3, _instance.authorNAME);
    }
    if (_instance.content) {
      _writer.writeString(4, _instance.content);
    }
    if (_instance.timestamp) {
      _writer.writeInt64String(5, _instance.timestamp);
    }
    if (_instance.hasFileAttached) {
      _writer.writeBool(6, _instance.hasFileAttached);
    }
    if (_instance.attach && _instance.attach.length) {
      _writer.writeRepeatedMessage(
        7,
        _instance.attach as any,
        Attachments.serializeBinaryToWriter
      );
    }
    if (_instance.isAnswer) {
      _writer.writeBool(8, _instance.isAnswer);
    }
    if (_instance.answersTo) {
      _writer.writeString(9, _instance.answersTo);
    }
    if (_instance.id) {
      _writer.writeString(10, _instance.id);
    }
  }

  private _authorUUID?: string;
  private _chatroomID?: string;
  private _authorNAME?: string;
  private _content?: string;
  private _timestamp?: string;
  private _hasFileAttached?: boolean;
  private _attach?: Attachments[];
  private _isAnswer?: boolean;
  private _answersTo?: string;
  private _id?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of Message to deeply clone from
   */
  constructor(_value?: RecursivePartial<Message>) {
    _value = _value || {};
    this.authorUUID = _value.authorUUID;
    this.chatroomID = _value.chatroomID;
    this.authorNAME = _value.authorNAME;
    this.content = _value.content;
    this.timestamp = _value.timestamp;
    this.hasFileAttached = _value.hasFileAttached;
    this.attach = (_value.attach || []).map(m => new Attachments(m));
    this.isAnswer = _value.isAnswer;
    this.answersTo = _value.answersTo;
    this.id = _value.id;
    Message.refineValues(this);
  }
  get authorUUID(): string | undefined {
    return this._authorUUID;
  }
  set authorUUID(value: string | undefined) {
    this._authorUUID = value;
  }
  get chatroomID(): string | undefined {
    return this._chatroomID;
  }
  set chatroomID(value: string | undefined) {
    this._chatroomID = value;
  }
  get authorNAME(): string | undefined {
    return this._authorNAME;
  }
  set authorNAME(value: string | undefined) {
    this._authorNAME = value;
  }
  get content(): string | undefined {
    return this._content;
  }
  set content(value: string | undefined) {
    this._content = value;
  }
  get timestamp(): string | undefined {
    return this._timestamp;
  }
  set timestamp(value: string | undefined) {
    this._timestamp = value;
  }
  get hasFileAttached(): boolean | undefined {
    return this._hasFileAttached;
  }
  set hasFileAttached(value: boolean | undefined) {
    this._hasFileAttached = value;
  }
  get attach(): Attachments[] | undefined {
    return this._attach;
  }
  set attach(value: Attachments[] | undefined) {
    this._attach = value;
  }
  get isAnswer(): boolean | undefined {
    return this._isAnswer;
  }
  set isAnswer(value: boolean | undefined) {
    this._isAnswer = value;
  }
  get answersTo(): string | undefined {
    return this._answersTo;
  }
  set answersTo(value: string | undefined) {
    this._answersTo = value;
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
    Message.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): Message.AsObject {
    return {
      authorUUID: this.authorUUID,
      chatroomID: this.chatroomID,
      authorNAME: this.authorNAME,
      content: this.content,
      timestamp: this.timestamp,
      hasFileAttached: this.hasFileAttached,
      attach: (this.attach || []).map(m => m.toObject()),
      isAnswer: this.isAnswer,
      answersTo: this.answersTo,
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
export module Message {
  /**
   * Standard JavaScript object representation for Message
   */
  export interface AsObject {
    authorUUID?: string;
    chatroomID?: string;
    authorNAME?: string;
    content?: string;
    timestamp?: string;
    hasFileAttached?: boolean;
    attach?: Attachments.AsObject[];
    isAnswer?: boolean;
    answersTo?: string;
    id?: string;
  }
}

/**
 * Message implementation for message.StreamMessagesByChatroomRequest
 */
export class StreamMessagesByChatroomRequest implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new StreamMessagesByChatroomRequest();
    StreamMessagesByChatroomRequest.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: StreamMessagesByChatroomRequest) {
    _instance.chatroomID = _instance.chatroomID || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: StreamMessagesByChatroomRequest,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.chatroomID = _reader.readString();
          break;
        default:
          _reader.skipField();
      }
    }

    StreamMessagesByChatroomRequest.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: StreamMessagesByChatroomRequest,
    _writer: BinaryWriter
  ) {
    if (_instance.chatroomID) {
      _writer.writeString(1, _instance.chatroomID);
    }
  }

  private _chatroomID?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of StreamMessagesByChatroomRequest to deeply clone from
   */
  constructor(_value?: RecursivePartial<StreamMessagesByChatroomRequest>) {
    _value = _value || {};
    this.chatroomID = _value.chatroomID;
    StreamMessagesByChatroomRequest.refineValues(this);
  }
  get chatroomID(): string | undefined {
    return this._chatroomID;
  }
  set chatroomID(value: string | undefined) {
    this._chatroomID = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    StreamMessagesByChatroomRequest.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): StreamMessagesByChatroomRequest.AsObject {
    return {
      chatroomID: this.chatroomID
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
export module StreamMessagesByChatroomRequest {
  /**
   * Standard JavaScript object representation for StreamMessagesByChatroomRequest
   */
  export interface AsObject {
    chatroomID?: string;
  }
}

/**
 * Message implementation for message.StreamMessagesByChatroomResponse
 */
export class StreamMessagesByChatroomResponse implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new StreamMessagesByChatroomResponse();
    StreamMessagesByChatroomResponse.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: StreamMessagesByChatroomResponse) {
    _instance.message = _instance.message || undefined;
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: StreamMessagesByChatroomResponse,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.message = new Message();
          _reader.readMessage(
            _instance.message,
            Message.deserializeBinaryFromReader
          );
          break;
        default:
          _reader.skipField();
      }
    }

    StreamMessagesByChatroomResponse.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: StreamMessagesByChatroomResponse,
    _writer: BinaryWriter
  ) {
    if (_instance.message) {
      _writer.writeMessage(
        1,
        _instance.message as any,
        Message.serializeBinaryToWriter
      );
    }
  }

  private _message?: Message;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of StreamMessagesByChatroomResponse to deeply clone from
   */
  constructor(_value?: RecursivePartial<StreamMessagesByChatroomResponse>) {
    _value = _value || {};
    this.message = _value.message ? new Message(_value.message) : undefined;
    StreamMessagesByChatroomResponse.refineValues(this);
  }
  get message(): Message | undefined {
    return this._message;
  }
  set message(value: Message | undefined) {
    this._message = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    StreamMessagesByChatroomResponse.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): StreamMessagesByChatroomResponse.AsObject {
    return {
      message: this.message ? this.message.toObject() : undefined
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
export module StreamMessagesByChatroomResponse {
  /**
   * Standard JavaScript object representation for StreamMessagesByChatroomResponse
   */
  export interface AsObject {
    message?: Message.AsObject;
  }
}
