import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { GetMessagesByChatroomResponse, SendMessageRequest, SendMessageResponse } from 'src/proto/message.pb';
import { MessageServiceClient } from 'src/proto/message.pbsc';
import { LoggedUser } from 'src/_models/loggedUser';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root'
})
export class MessageService {

  currentUser: LoggedUser;

  constructor(
    private authService: AuthService,
    private messageClient: MessageServiceClient,
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

}
