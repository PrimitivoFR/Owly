import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { DeleteMessageRequest, DeleteMessageResponse, GetMessagesByChatroomRequest, GetMessagesByChatroomResponse, SendMessageRequest, SendMessageResponse, StreamMessagesByChatroomRequest } from 'src/proto/message.pb';
import { MessageServiceClient } from 'src/proto/message.pbsc';
import { LocalChatroom } from 'src/_models/localChatroom';
import { LocalMessages as LocalMessage } from 'src/_models/localMessages';
import { LocalRoomsAndMessagesStore } from 'src/_models/localRoomsAndMessagesStore';
import { LoggedUser } from 'src/_models/loggedUser';
import { AuthService } from './auth.service';
import { ChatroomService } from './chatroom.service';
import { StoreService } from './store.service';

@Injectable({
  providedIn: 'root'
})
export class MessageService {

  currentUser: LoggedUser;

  constructor(
    private authService: AuthService,
    private messageClient: MessageServiceClient,
    private chatroomService: ChatroomService,
    private storeService:  StoreService
  ) {
    this.authService.currentUser.subscribe((user) => {
      this.currentUser = user;
    });
  }


  async sendMessage(req: SendMessageRequest): Promise<SendMessageResponse> {
    req.message.authorNAME = this.currentUser.username;
    const res = await this.messageClient.sendMessage(req, {"authorization": this.currentUser.accessToken}).toPromise();
    return res;
  }

  async deleteMessage(req: DeleteMessageRequest): Promise<DeleteMessageResponse> {
    return await this.messageClient.deleteMessage(req, {"authorization": this.currentUser.accessToken}).toPromise();
  }

  async getMessagesForAllChatrooms() {

    var rooms: LocalRoomsAndMessagesStore[];
    this.storeService.currentChatroomsAndMessageStore.subscribe((v) => rooms = v);
  
    var roomsAndMessages: LocalRoomsAndMessagesStore[] = []
    roomsAndMessages = rooms.map(async (room: LocalRoomsAndMessagesStore) =>{

      const res = await this.getMessagesByChatroom(new GetMessagesByChatroomRequest({
        chatroomID: room.chatroom.id
      }))
      res.messages.sort((a,b) => parseInt(a.timestamp) > parseInt(b.timestamp) ? 1 : -1)
      var storeEl = new LocalRoomsAndMessagesStore();
      storeEl.chatroom = room.chatroom;
      storeEl.localID = room.localID;
      storeEl.messages = res.messages;
      return storeEl;

    }) as unknown as LocalRoomsAndMessagesStore[]
    roomsAndMessages = await Promise.all(roomsAndMessages)

    this.storeService.updateWholeStore(roomsAndMessages);

  }

  async getMessagesByChatroom(req: GetMessagesByChatroomRequest): Promise<GetMessagesByChatroomResponse> {
    return await this.messageClient.getMessagesByChatroom(req, {"authorization": this.currentUser.accessToken}).toPromise();
  }

  streamMessagesByChatroom(req: StreamMessagesByChatroomRequest) {
    return this.messageClient.streamMessagesByChatroom(req, {"authorization": this.currentUser.accessToken})
  }

}
