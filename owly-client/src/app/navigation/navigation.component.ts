import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Chatroom, GetChatroomsByUserRequest, GetChatroomsByUserResponse } from 'src/proto/chatroom.pb';
import { LocalChatroom } from 'src/_models/localChatroom';
import { AuthService } from 'src/_services/auth.service';
import { ChatroomService } from 'src/_services/chatroom.service';
import { MessageService } from 'src/_services/message.service';
import { NavigationService } from './navigation.service';

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss']
})
export class NavigationComponent implements OnInit {

  chatroomsList: LocalChatroom[];

  constructor(
    private router: Router,
    private chatroomService: ChatroomService,
    private authService: AuthService,
    private messageService: MessageService,
    private navService: NavigationService
  ) { }

  ngOnInit(): void {
    this.chatroomService.chatroomsListValue.subscribe((chatrooms) => {
      this.chatroomsList = chatrooms;
      console.log(this.chatroomsList);
    });
    if(this.authService.isAuthenticated())
      this.getChatroomsAndMessages();
      this.navService.currentNavStore.subscribe(v => console.log(v))
  }

  goToChatroom(localID: string) {
    console.log("going to "+localID);
    this.navService.updateNavStore(localID)
    this.router.navigateByUrl('/chatroom');

  }

  async getChatroomsAndMessages() {
    await this.chatroomService.getChatrooms();
    await this.messageService.getMessagesForAllChatrooms();
     
    this.navService.updateNavStore("0")
  }
}
