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
import * as thisProto from './message.pb';
import { GRPC_MESSAGE_SERVICE_CLIENT_SETTINGS } from './message.pbconf';
/**
 * Service client implementation for message.MessageService
 */
@Injectable({
  providedIn: 'root'
})
export class MessageServiceClient {
  private client: GrpcClient;

  constructor(
    @Optional()
    @Inject(GRPC_MESSAGE_SERVICE_CLIENT_SETTINGS)
    settings: GrpcClientSettings,
    @Inject(GRPC_CLIENT_FACTORY) clientFactory: GrpcClientFactory,
    private handler: GrpcHandler
  ) {
    this.client = clientFactory.createClient(
      'message.MessageService',
      settings
    );
  }

  /**
   * Server streaming RPC for /message.MessageService/StreamMessagesByChatroom
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.StreamMessagesByChatroomResponse>
   */
  streamMessagesByChatroom(
    requestData: thisProto.StreamMessagesByChatroomRequest,
    requestMetadata: Metadata = {}
  ): Observable<thisProto.StreamMessagesByChatroomResponse> {
    return this.streamMessagesByChatroom$eventStream(
      requestData,
      requestMetadata
    ).pipe(throwStatusErrors(), takeMessages());
  }

  /**
   * Server streaming RPC for /message.MessageService/StreamMessagesByChatroom
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<GrpcEvent<thisProto.StreamMessagesByChatroomResponse>>
   */
  streamMessagesByChatroom$eventStream(
    requestData: thisProto.StreamMessagesByChatroomRequest,
    requestMetadata: Metadata = {}
  ): Observable<GrpcEvent<thisProto.StreamMessagesByChatroomResponse>> {
    return this.handler.handle({
      type: GrpcCallType.serverStream,
      client: this.client,
      path: '/message.MessageService/StreamMessagesByChatroom',
      requestData,
      requestMetadata,
      requestClass: thisProto.StreamMessagesByChatroomRequest,
      responseClass: thisProto.StreamMessagesByChatroomResponse
    });
  }
}
