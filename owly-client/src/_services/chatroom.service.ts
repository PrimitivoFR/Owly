import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie';
import { BehaviorSubject } from 'rxjs';
import { Chatroom, CreateChatroomRequest, CreateChatroomResponse, GetChatroomsByUserRequest, GetChatroomsByUserResponse } from 'src/proto/chatroom.pb';
import { ChatroomServiceClient } from 'src/proto/chatroom.pbsc';
import { LoggedUser } from 'src/_models/loggedUser';
import { AuthService } from './auth.service';
import { v4 as uuidv4 } from 'uuid';
import { RecursivePartial } from '@ngx-grpc/common';
import { LocalChatroom } from 'src/_models/localChatroom';
import { LocalRoomsAndMessagesStore } from 'src/_models/localRoomsAndMessagesStore';
import { StoreService } from './store.service';

@Injectable({
  providedIn: 'root'
})



export class ChatroomService {
  
  

  currentUser: LoggedUser;

  

  constructor(
    private chatroomClient: ChatroomServiceClient,
    private authService: AuthService,
    private storeService: StoreService
  ) {
    this.authService.currentUser.subscribe((user) => {
      this.currentUser = user;
    });
  }


  async createChatroom(req: CreateChatroomRequest): Promise<CreateChatroomResponse> {
    const token = this.currentUser.accessToken;
    const res = await this.chatroomClient.createChatroom(req, {"authorization": token}).toPromise();
    return res;
  }

  async getChatrooms() {
    const req = new GetChatroomsByUserRequest();
    const token = this.currentUser.accessToken;

    var res = await this.chatroomClient.getChatroomsByUser(req, {"authorization": token}).toPromise()
    console.log(res)
    var localChatrooms: LocalRoomsAndMessagesStore[] = []
    if (res.count != "0" ) {
      localChatrooms = res.chatrooms.map((chatroom: Chatroom, index) => {
        var room = new LocalRoomsAndMessagesStore();
        room.localID = index.toString();
        room.chatroom = chatroom;
        return room
      }) as unknown as LocalRoomsAndMessagesStore[];
    }
    
    this.storeService.updateWholeStore(localChatrooms);
  }

}
