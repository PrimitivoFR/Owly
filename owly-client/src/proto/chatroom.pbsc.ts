/* tslint:disable */
/* eslint-disable */
//
// THIS IS A GENERATED FILE
// DO NOT MODIFY IT! YOUR CHANGES WILL BE LOST
import { Inject, Injectable, Optional } from '@angular/core';
import {
  GrpcCallType,
  GrpcClient,
  GrpcClientFactory,
  GrpcClientSettings,
  GrpcEvent
} from '@ngx-grpc/common';
import {
  GRPC_CLIENT_FACTORY,
  GrpcHandler,
  takeMessages,
  throwStatusErrors
} from '@ngx-grpc/core';
import { Metadata } from 'grpc-web';
import { Observable } from 'rxjs';
import * as thisProto from './chatroom.pb';
import { GRPC_CHATROOM_SERVICE_CLIENT_SETTINGS } from './chatroom.pbconf';
/**
 * Service client implementation for chatroom.ChatroomService
 */
@Injectable({
  providedIn: 'root'
})
export class ChatroomServiceClient {
  private client: GrpcClient;

  constructor(
    @Optional()
    @Inject(GRPC_CHATROOM_SERVICE_CLIENT_SETTINGS)
    settings: GrpcClientSettings,
    @Inject(GRPC_CLIENT_FACTORY) clientFactory: GrpcClientFactory,
    private handler: GrpcHandler
  ) {
    this.client = clientFactory.createClient(
      'chatroom.ChatroomService',
      settings
    );
  }

  /**
   * Unary RPC for /chatroom.ChatroomService/CreateChatroom
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.CreateChatroomResponse>
   */
  createChatroom(
    requestData: thisProto.CreateChatroomRequest,
    requestMetadata: Metadata = {}
  ): Observable<thisProto.CreateChatroomResponse> {
    return this.createChatroom$eventStream(requestData, requestMetadata).pipe(
      throwStatusErrors(),
      takeMessages()
    );
  }

  /**
   * Unary RPC for /chatroom.ChatroomService/CreateChatroom
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<GrpcEvent<thisProto.CreateChatroomResponse>>
   */
  createChatroom$eventStream(
    requestData: thisProto.CreateChatroomRequest,
    requestMetadata: Metadata = {}
  ): Observable<GrpcEvent<thisProto.CreateChatroomResponse>> {
    return this.handler.handle({
      type: GrpcCallType.unary,
      client: this.client,
      path: '/chatroom.ChatroomService/CreateChatroom',
      requestData,
      requestMetadata,
      requestClass: thisProto.CreateChatroomRequest,
      responseClass: thisProto.CreateChatroomResponse
    });
  }

  /**
   * Unary RPC for /chatroom.ChatroomService/GetChatroomsByUser
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.GetChatroomsByUserResponse>
   */
  getChatroomsByUser(
    requestData: thisProto.GetChatroomsByUserRequest,
    requestMetadata: Metadata = {}
  ): Observable<thisProto.GetChatroomsByUserResponse> {
    return this.getChatroomsByUser$eventStream(
      requestData,
      requestMetadata
    ).pipe(throwStatusErrors(), takeMessages());
  }

  /**
   * Unary RPC for /chatroom.ChatroomService/GetChatroomsByUser
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<GrpcEvent<thisProto.GetChatroomsByUserResponse>>
   */
  getChatroomsByUser$eventStream(
    requestData: thisProto.GetChatroomsByUserRequest,
    requestMetadata: Metadata = {}
  ): Observable<GrpcEvent<thisProto.GetChatroomsByUserResponse>> {
    return this.handler.handle({
      type: GrpcCallType.unary,
      client: this.client,
      path: '/chatroom.ChatroomService/GetChatroomsByUser',
      requestData,
      requestMetadata,
      requestClass: thisProto.GetChatroomsByUserRequest,
      responseClass: thisProto.GetChatroomsByUserResponse
    });
  }

  /**
   * Unary RPC for /chatroom.ChatroomService/DeleteChatroom
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.DeleteChatroomResponse>
   */
  deleteChatroom(
    requestData: thisProto.DeleteChatroomRequest,
    requestMetadata: Metadata = {}
  ): Observable<thisProto.DeleteChatroomResponse> {
    return this.deleteChatroom$eventStream(requestData, requestMetadata).pipe(
      throwStatusErrors(),
      takeMessages()
    );
  }

  /**
   * Unary RPC for /chatroom.ChatroomService/DeleteChatroom
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<GrpcEvent<thisProto.DeleteChatroomResponse>>
   */
  deleteChatroom$eventStream(
    requestData: thisProto.DeleteChatroomRequest,
    requestMetadata: Metadata = {}
  ): Observable<GrpcEvent<thisProto.DeleteChatroomResponse>> {
    return this.handler.handle({
      type: GrpcCallType.unary,
      client: this.client,
      path: '/chatroom.ChatroomService/DeleteChatroom',
      requestData,
      requestMetadata,
      requestClass: thisProto.DeleteChatroomRequest,
      responseClass: thisProto.DeleteChatroomResponse
    });
  }

  /**
   * Unary RPC for /chatroom.ChatroomService/LeaveChatroom
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.LeaveChatroomResponse>
   */
  leaveChatroom(
    requestData: thisProto.LeaveChatroomRequest,
    requestMetadata: Metadata = {}
  ): Observable<thisProto.LeaveChatroomResponse> {
    return this.leaveChatroom$eventStream(requestData, requestMetadata).pipe(
      throwStatusErrors(),
      takeMessages()
    );
  }

  /**
   * Unary RPC for /chatroom.ChatroomService/LeaveChatroom
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<GrpcEvent<thisProto.LeaveChatroomResponse>>
   */
  leaveChatroom$eventStream(
    requestData: thisProto.LeaveChatroomRequest,
    requestMetadata: Metadata = {}
  ): Observable<GrpcEvent<thisProto.LeaveChatroomResponse>> {
    return this.handler.handle({
      type: GrpcCallType.unary,
      client: this.client,
      path: '/chatroom.ChatroomService/LeaveChatroom',
      requestData,
      requestMetadata,
      requestClass: thisProto.LeaveChatroomRequest,
      responseClass: thisProto.LeaveChatroomResponse
    });
  }
}
