import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie';
import { CreateChatroomRequest, CreateChatroomResponse, GetChatroomsByUserRequest, GetChatroomsByUserResponse } from 'src/proto/chatroom.pb';
import { ChatroomServiceClient } from 'src/proto/chatroom.pbsc';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class ChatroomService {


  constructor(
    private chatroomClient: ChatroomServiceClient,
    private authService: AuthService
  ) {

  }


  async createChatroom(req: CreateChatroomRequest): Promise<CreateChatroomResponse> {
    const token = this.authService.currentUserValue.accessToken;
    const res = await this.chatroomClient.createChatroom(req, {"authorization": token}).toPromise();
    return res;
  }

  async getChatroom(req: GetChatroomsByUserRequest): Promise<GetChatroomsByUserResponse> {
    const token = this.authService.currentUserValue.accessToken;
    const res = await this.chatroomClient.getChatroomsByUser(req, {"authorization": token}).toPromise();
    return res;
  }
}
