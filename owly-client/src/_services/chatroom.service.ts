import { Injectable } from '@angular/core';
import { CreateChatroomRequest, CreateChatroomResponse } from 'src/proto/chatroom.pb';
import { ChatroomServiceClient } from 'src/proto/chatroom.pbsc';

@Injectable({
  providedIn: 'root'
})
export class ChatroomService {

  constructor(private chatroomClient: ChatroomServiceClient) { }


  async createChatroom(req: CreateChatroomRequest): Promise<CreateChatroomResponse> {
    const res = await this.chatroomClient.createChatroom(req).toPromise()
    return res;
  }
}
