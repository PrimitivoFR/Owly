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
 * Message implementation for message.messageHistory
 */
export class messageHistory implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new messageHistory();
    messageHistory.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: messageHistory) {
    _instance.timestamp = _instance.timestamp || '0';
    _instance.content = _instance.content || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: messageHistory,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.timestamp = _reader.readInt64String();
          break;
        case 2:
          _instance.content = _reader.readString();
          break;
        default:
          _reader.skipField();
      }
    }

    messageHistory.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: messageHistory,
    _writer: BinaryWriter
  ) {
    if (_instance.timestamp) {
      _writer.writeInt64String(1, _instance.timestamp);
    }
    if (_instance.content) {
      _writer.writeString(2, _instance.content);
    }
  }

  private _timestamp?: string;
  private _content?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of messageHistory to deeply clone from
   */
  constructor(_value?: RecursivePartial<messageHistory>) {
    _value = _value || {};
    this.timestamp = _value.timestamp;
    this.content = _value.content;
    messageHistory.refineValues(this);
  }
  get timestamp(): string | undefined {
    return this._timestamp;
  }
  set timestamp(value: string | undefined) {
    this._timestamp = value;
  }
  get content(): string | undefined {
    return this._content;
  }
  set content(value: string | undefined) {
    this._content = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    messageHistory.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): messageHistory.AsObject {
    return {
      timestamp: this.timestamp,
      content: this.content
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
export module messageHistory {
  /**
   * Standard JavaScript object representation for messageHistory
   */
  export interface AsObject {
    timestamp?: string;
    content?: string;
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
    _instance.history = _instance.history || [];
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
        case 11:
          const messageInitializer11 = new messageHistory();
          _reader.readMessage(
            messageInitializer11,
            messageHistory.deserializeBinaryFromReader
          );
          (_instance.history = _instance.history || []).push(
            messageInitializer11
          );
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
    if (_instance.history && _instance.history.length) {
      _writer.writeRepeatedMessage(
        11,
        _instance.history as any,
        messageHistory.serializeBinaryToWriter
      );
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
  private _history?: messageHistory[];

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
    this.history = (_value.history || []).map(m => new messageHistory(m));
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
  get history(): messageHistory[] | undefined {
    return this._history;
  }
  set history(value: messageHistory[] | undefined) {
    this._history = value;
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
      id: this.id,
      history: (this.history || []).map(m => m.toObject())
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
    history?: messageHistory.AsObject[];
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
    _instance.operation = _instance.operation || '';
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
          _instance.operation = _reader.readString();
          break;
        case 2:
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
    if (_instance.operation) {
      _writer.writeString(1, _instance.operation);
    }
    if (_instance.message) {
      _writer.writeMessage(
        2,
        _instance.message as any,
        Message.serializeBinaryToWriter
      );
    }
  }

  private _operation?: string;
  private _message?: Message;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of StreamMessagesByChatroomResponse to deeply clone from
   */
  constructor(_value?: RecursivePartial<StreamMessagesByChatroomResponse>) {
    _value = _value || {};
    this.operation = _value.operation;
    this.message = _value.message ? new Message(_value.message) : undefined;
    StreamMessagesByChatroomResponse.refineValues(this);
  }
  get operation(): string | undefined {
    return this._operation;
  }
  set operation(value: string | undefined) {
    this._operation = value;
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
      operation: this.operation,
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
    operation?: string;
    message?: Message.AsObject;
  }
}

/**
 * Message implementation for message.SendMessageRequest
 */
export class SendMessageRequest implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new SendMessageRequest();
    SendMessageRequest.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: SendMessageRequest) {
    _instance.message = _instance.message || undefined;
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: SendMessageRequest,
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

    SendMessageRequest.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: SendMessageRequest,
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
   * @param _value initial values object or instance of SendMessageRequest to deeply clone from
   */
  constructor(_value?: RecursivePartial<SendMessageRequest>) {
    _value = _value || {};
    this.message = _value.message ? new Message(_value.message) : undefined;
    SendMessageRequest.refineValues(this);
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
    SendMessageRequest.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): SendMessageRequest.AsObject {
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
export module SendMessageRequest {
  /**
   * Standard JavaScript object representation for SendMessageRequest
   */
  export interface AsObject {
    message?: Message.AsObject;
  }
}

/**
 * Message implementation for message.SendMessageResponse
 */
export class SendMessageResponse implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new SendMessageResponse();
    SendMessageResponse.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: SendMessageResponse) {
    _instance.success = _instance.success || false;
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: SendMessageResponse,
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

    SendMessageResponse.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: SendMessageResponse,
    _writer: BinaryWriter
  ) {
    if (_instance.success) {
      _writer.writeBool(1, _instance.success);
    }
  }

  private _success?: boolean;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of SendMessageResponse to deeply clone from
   */
  constructor(_value?: RecursivePartial<SendMessageResponse>) {
    _value = _value || {};
    this.success = _value.success;
    SendMessageResponse.refineValues(this);
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
    SendMessageResponse.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): SendMessageResponse.AsObject {
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
export module SendMessageResponse {
  /**
   * Standard JavaScript object representation for SendMessageResponse
   */
  export interface AsObject {
    success?: boolean;
  }
}

/**
 * Message implementation for message.GetMessagesByChatroomRequest
 */
export class GetMessagesByChatroomRequest implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new GetMessagesByChatroomRequest();
    GetMessagesByChatroomRequest.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: GetMessagesByChatroomRequest) {
    _instance.chatroomID = _instance.chatroomID || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: GetMessagesByChatroomRequest,
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

    GetMessagesByChatroomRequest.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: GetMessagesByChatroomRequest,
    _writer: BinaryWriter
  ) {
    if (_instance.chatroomID) {
      _writer.writeString(1, _instance.chatroomID);
    }
  }

  private _chatroomID?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of GetMessagesByChatroomRequest to deeply clone from
   */
  constructor(_value?: RecursivePartial<GetMessagesByChatroomRequest>) {
    _value = _value || {};
    this.chatroomID = _value.chatroomID;
    GetMessagesByChatroomRequest.refineValues(this);
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
    GetMessagesByChatroomRequest.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): GetMessagesByChatroomRequest.AsObject {
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
export module GetMessagesByChatroomRequest {
  /**
   * Standard JavaScript object representation for GetMessagesByChatroomRequest
   */
  export interface AsObject {
    chatroomID?: string;
  }
}

/**
 * Message implementation for message.GetMessagesByChatroomResponse
 */
export class GetMessagesByChatroomResponse implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new GetMessagesByChatroomResponse();
    GetMessagesByChatroomResponse.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: GetMessagesByChatroomResponse) {
    _instance.messages = _instance.messages || [];
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: GetMessagesByChatroomResponse,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          const messageInitializer1 = new Message();
          _reader.readMessage(
            messageInitializer1,
            Message.deserializeBinaryFromReader
          );
          (_instance.messages = _instance.messages || []).push(
            messageInitializer1
          );
          break;
        default:
          _reader.skipField();
      }
    }

    GetMessagesByChatroomResponse.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: GetMessagesByChatroomResponse,
    _writer: BinaryWriter
  ) {
    if (_instance.messages && _instance.messages.length) {
      _writer.writeRepeatedMessage(
        1,
        _instance.messages as any,
        Message.serializeBinaryToWriter
      );
    }
  }

  private _messages?: Message[];

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of GetMessagesByChatroomResponse to deeply clone from
   */
  constructor(_value?: RecursivePartial<GetMessagesByChatroomResponse>) {
    _value = _value || {};
    this.messages = (_value.messages || []).map(m => new Message(m));
    GetMessagesByChatroomResponse.refineValues(this);
  }
  get messages(): Message[] | undefined {
    return this._messages;
  }
  set messages(value: Message[] | undefined) {
    this._messages = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    GetMessagesByChatroomResponse.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): GetMessagesByChatroomResponse.AsObject {
    return {
      messages: (this.messages || []).map(m => m.toObject())
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
export module GetMessagesByChatroomResponse {
  /**
   * Standard JavaScript object representation for GetMessagesByChatroomResponse
   */
  export interface AsObject {
    messages?: Message.AsObject[];
  }
}

/**
 * Message implementation for message.DeleteMessageRequest
 */
export class DeleteMessageRequest implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new DeleteMessageRequest();
    DeleteMessageRequest.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: DeleteMessageRequest) {
    _instance.messageID = _instance.messageID || '';
    _instance.chatroomID = _instance.chatroomID || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: DeleteMessageRequest,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.messageID = _reader.readString();
          break;
        case 2:
          _instance.chatroomID = _reader.readString();
          break;
        default:
          _reader.skipField();
      }
    }

    DeleteMessageRequest.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: DeleteMessageRequest,
    _writer: BinaryWriter
  ) {
    if (_instance.messageID) {
      _writer.writeString(1, _instance.messageID);
    }
    if (_instance.chatroomID) {
      _writer.writeString(2, _instance.chatroomID);
    }
  }

  private _messageID?: string;
  private _chatroomID?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of DeleteMessageRequest to deeply clone from
   */
  constructor(_value?: RecursivePartial<DeleteMessageRequest>) {
    _value = _value || {};
    this.messageID = _value.messageID;
    this.chatroomID = _value.chatroomID;
    DeleteMessageRequest.refineValues(this);
  }
  get messageID(): string | undefined {
    return this._messageID;
  }
  set messageID(value: string | undefined) {
    this._messageID = value;
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
    DeleteMessageRequest.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): DeleteMessageRequest.AsObject {
    return {
      messageID: this.messageID,
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
export module DeleteMessageRequest {
  /**
   * Standard JavaScript object representation for DeleteMessageRequest
   */
  export interface AsObject {
    messageID?: string;
    chatroomID?: string;
  }
}

/**
 * Message implementation for message.DeleteMessageResponse
 */
export class DeleteMessageResponse implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new DeleteMessageResponse();
    DeleteMessageResponse.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: DeleteMessageResponse) {
    _instance.success = _instance.success || false;
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: DeleteMessageResponse,
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

    DeleteMessageResponse.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: DeleteMessageResponse,
    _writer: BinaryWriter
  ) {
    if (_instance.success) {
      _writer.writeBool(1, _instance.success);
    }
  }

  private _success?: boolean;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of DeleteMessageResponse to deeply clone from
   */
  constructor(_value?: RecursivePartial<DeleteMessageResponse>) {
    _value = _value || {};
    this.success = _value.success;
    DeleteMessageResponse.refineValues(this);
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
    DeleteMessageResponse.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): DeleteMessageResponse.AsObject {
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
export module DeleteMessageResponse {
  /**
   * Standard JavaScript object representation for DeleteMessageResponse
   */
  export interface AsObject {
    success?: boolean;
  }
}

/**
 * Message implementation for message.UpdateMessageContentRequest
 */
export class UpdateMessageContentRequest implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new UpdateMessageContentRequest();
    UpdateMessageContentRequest.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: UpdateMessageContentRequest) {
    _instance.messageId = _instance.messageId || '';
    _instance.chatroomId = _instance.chatroomId || '';
    _instance.newContent = _instance.newContent || '';
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: UpdateMessageContentRequest,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.messageId = _reader.readString();
          break;
        case 2:
          _instance.chatroomId = _reader.readString();
          break;
        case 3:
          _instance.newContent = _reader.readString();
          break;
        default:
          _reader.skipField();
      }
    }

    UpdateMessageContentRequest.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: UpdateMessageContentRequest,
    _writer: BinaryWriter
  ) {
    if (_instance.messageId) {
      _writer.writeString(1, _instance.messageId);
    }
    if (_instance.chatroomId) {
      _writer.writeString(2, _instance.chatroomId);
    }
    if (_instance.newContent) {
      _writer.writeString(3, _instance.newContent);
    }
  }

  private _messageId?: string;
  private _chatroomId?: string;
  private _newContent?: string;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of UpdateMessageContentRequest to deeply clone from
   */
  constructor(_value?: RecursivePartial<UpdateMessageContentRequest>) {
    _value = _value || {};
    this.messageId = _value.messageId;
    this.chatroomId = _value.chatroomId;
    this.newContent = _value.newContent;
    UpdateMessageContentRequest.refineValues(this);
  }
  get messageId(): string | undefined {
    return this._messageId;
  }
  set messageId(value: string | undefined) {
    this._messageId = value;
  }
  get chatroomId(): string | undefined {
    return this._chatroomId;
  }
  set chatroomId(value: string | undefined) {
    this._chatroomId = value;
  }
  get newContent(): string | undefined {
    return this._newContent;
  }
  set newContent(value: string | undefined) {
    this._newContent = value;
  }

  /**
   * Serialize message to binary data
   * @param instance message instance
   */
  serializeBinary() {
    const writer = new BinaryWriter();
    UpdateMessageContentRequest.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): UpdateMessageContentRequest.AsObject {
    return {
      messageId: this.messageId,
      chatroomId: this.chatroomId,
      newContent: this.newContent
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
export module UpdateMessageContentRequest {
  /**
   * Standard JavaScript object representation for UpdateMessageContentRequest
   */
  export interface AsObject {
    messageId?: string;
    chatroomId?: string;
    newContent?: string;
  }
}

/**
 * Message implementation for message.UpdateMessageContentResponse
 */
export class UpdateMessageContentResponse implements GrpcMessage {
  /**
   * Deserialize binary data to message
   * @param instance message instance
   */
  static deserializeBinary(bytes: ByteSource) {
    const instance = new UpdateMessageContentResponse();
    UpdateMessageContentResponse.deserializeBinaryFromReader(
      instance,
      new BinaryReader(bytes)
    );
    return instance;
  }

  /**
   * Check all the properties and set default protobuf values if necessary
   * @param _instance message instance
   */
  static refineValues(_instance: UpdateMessageContentResponse) {
    _instance.success = _instance.success || false;
    _instance.message = _instance.message || undefined;
  }

  /**
   * Deserializes / reads binary message into message instance using provided binary reader
   * @param _instance message instance
   * @param _reader binary reader instance
   */
  static deserializeBinaryFromReader(
    _instance: UpdateMessageContentResponse,
    _reader: BinaryReader
  ) {
    while (_reader.nextField()) {
      if (_reader.isEndGroup()) break;

      switch (_reader.getFieldNumber()) {
        case 1:
          _instance.success = _reader.readBool();
          break;
        case 2:
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

    UpdateMessageContentResponse.refineValues(_instance);
  }

  /**
   * Serializes a message to binary format using provided binary reader
   * @param _instance message instance
   * @param _writer binary writer instance
   */
  static serializeBinaryToWriter(
    _instance: UpdateMessageContentResponse,
    _writer: BinaryWriter
  ) {
    if (_instance.success) {
      _writer.writeBool(1, _instance.success);
    }
    if (_instance.message) {
      _writer.writeMessage(
        2,
        _instance.message as any,
        Message.serializeBinaryToWriter
      );
    }
  }

  private _success?: boolean;
  private _message?: Message;

  /**
   * Message constructor. Initializes the properties and applies default Protobuf values if necessary
   * @param _value initial values object or instance of UpdateMessageContentResponse to deeply clone from
   */
  constructor(_value?: RecursivePartial<UpdateMessageContentResponse>) {
    _value = _value || {};
    this.success = _value.success;
    this.message = _value.message ? new Message(_value.message) : undefined;
    UpdateMessageContentResponse.refineValues(this);
  }
  get success(): boolean | undefined {
    return this._success;
  }
  set success(value: boolean | undefined) {
    this._success = value;
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
    UpdateMessageContentResponse.serializeBinaryToWriter(this, writer);
    return writer.getResultBuffer();
  }

  /**
   * Cast message to standard JavaScript object (all non-primitive values are deeply cloned)
   */
  toObject(): UpdateMessageContentResponse.AsObject {
    return {
      success: this.success,
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
export module UpdateMessageContentResponse {
  /**
   * Standard JavaScript object representation for UpdateMessageContentResponse
   */
  export interface AsObject {
    success?: boolean;
    message?: Message.AsObject;
  }
}
