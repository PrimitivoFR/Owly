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

@Injectable({
  providedIn: 'root'
})



export class ChatroomService {
  
  private chatroomsAndMessageStore = new BehaviorSubject<LocalRoomsAndMessagesStore[]>([]);
  currentChatroomsAndMessageStore = this.chatroomsAndMessageStore.asObservable();


  updateStore(val: LocalRoomsAndMessagesStore[]) {
    this.chatroomsAndMessageStore.next(val);
  }

  private chatroomsList = new BehaviorSubject<LocalChatroom[]>([]);
  chatroomsListValue = this.chatroomsList.asObservable();

  currentUser: LoggedUser;

  

  constructor(
    private chatroomClient: ChatroomServiceClient,
    private authService: AuthService
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
    var localChatrooms: LocalChatroom[] = []
    if (res.count != "0" ) {
      localChatrooms = res.chatrooms.map((chatroom: Chatroom, index) => new LocalChatroom({localID: index.toString(), chatroom})) as unknown as LocalChatroom[];
    }
    this.chatroomsList.next(localChatrooms);
  }

  getStoreItemById(id: string): LocalRoomsAndMessagesStore {
    const rooms = this.chatroomsAndMessageStore.value;
    const room = rooms.find((element: LocalRoomsAndMessagesStore) => element.localID === id);
    return room

  }
}
