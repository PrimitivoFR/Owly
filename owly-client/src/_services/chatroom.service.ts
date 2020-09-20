import { Injectable } from '@angular/core';
import { CookieService } from 'ngx-cookie';
import { BehaviorSubject } from 'rxjs';
import { CreateChatroomRequest, CreateChatroomResponse, GetChatroomsByUserRequest, GetChatroomsByUserResponse } from 'src/proto/chatroom.pb';
import { ChatroomServiceClient } from 'src/proto/chatroom.pbsc';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class ChatroomService {
  
  private chatroomsList = new BehaviorSubject<GetChatroomsByUserResponse>(new GetChatroomsByUserResponse());
  chatroomsListValue = this.chatroomsList.asObservable();

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

  async getChatroom() {
    const req = new GetChatroomsByUserRequest() ;
    const token = this.authService.currentUserValue.accessToken;
    this.chatroomsList.next(await this.chatroomClient.getChatroomsByUser(req, {"authorization": token}).toPromise());
  }
}
