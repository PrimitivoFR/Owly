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
import * as thisProto from './auth.pb';
import { GRPC_AUTH_SERVICE_CLIENT_SETTINGS } from './auth.pbconf';
/**
 * Service client implementation for auth.AuthService
 */
@Injectable({
  providedIn: 'root'
})
export class AuthServiceClient {
  private client: GrpcClient;

  constructor(
    @Optional()
    @Inject(GRPC_AUTH_SERVICE_CLIENT_SETTINGS)
    settings: GrpcClientSettings,
    @Inject(GRPC_CLIENT_FACTORY) clientFactory: GrpcClientFactory,
    private handler: GrpcHandler
  ) {
    this.client = clientFactory.createClient('auth.AuthService', settings);
  }

  /**
   * Unary RPC for /auth.AuthService/CreateNewUser
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.CreateNewUserResponse>
   */
  createNewUser(
    requestData: thisProto.CreateNewUserRequest,
    requestMetadata: Metadata = {}
  ): Observable<thisProto.CreateNewUserResponse> {
    return this.createNewUser$eventStream(requestData, requestMetadata).pipe(
      throwStatusErrors(),
      takeMessages()
    );
  }

  /**
   * Unary RPC for /auth.AuthService/CreateNewUser
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<GrpcEvent<thisProto.CreateNewUserResponse>>
   */
  createNewUser$eventStream(
    requestData: thisProto.CreateNewUserRequest,
    requestMetadata: Metadata = {}
  ): Observable<GrpcEvent<thisProto.CreateNewUserResponse>> {
    return this.handler.handle({
      type: GrpcCallType.unary,
      client: this.client,
      path: '/auth.AuthService/CreateNewUser',
      requestData,
      requestMetadata,
      requestClass: thisProto.CreateNewUserRequest,
      responseClass: thisProto.CreateNewUserResponse
    });
  }

  /**
   * Unary RPC for /auth.AuthService/LoginUser
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.LoginUserResponse>
   */
  loginUser(
    requestData: thisProto.LoginUserRequest,
    requestMetadata: Metadata = {}
  ): Observable<thisProto.LoginUserResponse> {
    return this.loginUser$eventStream(requestData, requestMetadata).pipe(
      throwStatusErrors(),
      takeMessages()
    );
  }

  /**
   * Unary RPC for /auth.AuthService/LoginUser
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<GrpcEvent<thisProto.LoginUserResponse>>
   */
  loginUser$eventStream(
    requestData: thisProto.LoginUserRequest,
    requestMetadata: Metadata = {}
  ): Observable<GrpcEvent<thisProto.LoginUserResponse>> {
    return this.handler.handle({
      type: GrpcCallType.unary,
      client: this.client,
      path: '/auth.AuthService/LoginUser',
      requestData,
      requestMetadata,
      requestClass: thisProto.LoginUserRequest,
      responseClass: thisProto.LoginUserResponse
    });
  }
}
