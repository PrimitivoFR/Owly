import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie';
import { BehaviorSubject } from 'rxjs';
import { CreateChatroomRequest, CreateChatroomResponse, GetChatroomsByUserRequest, GetChatroomsByUserResponse } from 'src/proto/chatroom.pb';
import { ChatroomServiceClient } from 'src/proto/chatroom.pbsc';
import { LoggedUser } from 'src/_models/loggedUser';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class ChatroomService {
  
  private chatroomsList = new BehaviorSubject<GetChatroomsByUserResponse>(new GetChatroomsByUserResponse());
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

  async getChatroom() {
    const req = new GetChatroomsByUserRequest();
    const token = this.currentUser.accessToken;
    this.chatroomsList.next(await this.chatroomClient.getChatroomsByUser(req, {"authorization": token}).toPromise());
  }
}
