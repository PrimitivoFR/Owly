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
import * as thisProto from './user.pb';
import { GRPC_USER_SERVICE_CLIENT_SETTINGS } from './user.pbconf';
/**
 * Service client implementation for user.UserService
 */
@Injectable({
  providedIn: 'root'
})
export class UserServiceClient {
  private client: GrpcClient;

  constructor(
    @Optional()
    @Inject(GRPC_USER_SERVICE_CLIENT_SETTINGS)
    settings: GrpcClientSettings,
    @Inject(GRPC_CLIENT_FACTORY) clientFactory: GrpcClientFactory,
    private handler: GrpcHandler
  ) {
    this.client = clientFactory.createClient('user.UserService', settings);
  }

  /**
   * Unary RPC for /user.UserService/CreateNewUser
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
   * Unary RPC for /user.UserService/CreateNewUser
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
      path: '/user.UserService/CreateNewUser',
      requestData,
      requestMetadata,
      requestClass: thisProto.CreateNewUserRequest,
      responseClass: thisProto.CreateNewUserResponse
    });
  }

  /**
   * Unary RPC for /user.UserService/LoginUser
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
   * Unary RPC for /user.UserService/LoginUser
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
      path: '/user.UserService/LoginUser',
      requestData,
      requestMetadata,
      requestClass: thisProto.LoginUserRequest,
      responseClass: thisProto.LoginUserResponse
    });
  }

  /**
   * Unary RPC for /user.UserService/SearchUserByUsername
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<thisProto.SearchUserByUsernameResponse>
   */
  searchUserByUsername(
    requestData: thisProto.SearchUserByUsernameRequest,
    requestMetadata: Metadata = {}
  ): Observable<thisProto.SearchUserByUsernameResponse> {
    return this.searchUserByUsername$eventStream(
      requestData,
      requestMetadata
    ).pipe(throwStatusErrors(), takeMessages());
  }

  /**
   * Unary RPC for /user.UserService/SearchUserByUsername
   *
   * @param requestMessage Request message
   * @param requestMetadata Request metadata
   * @returns Observable<GrpcEvent<thisProto.SearchUserByUsernameResponse>>
   */
  searchUserByUsername$eventStream(
    requestData: thisProto.SearchUserByUsernameRequest,
    requestMetadata: Metadata = {}
  ): Observable<GrpcEvent<thisProto.SearchUserByUsernameResponse>> {
    return this.handler.handle({
      type: GrpcCallType.unary,
      client: this.client,
      path: '/user.UserService/SearchUserByUsername',
      requestData,
      requestMetadata,
      requestClass: thisProto.SearchUserByUsernameRequest,
      responseClass: thisProto.SearchUserByUsernameResponse
    });
  }
}
