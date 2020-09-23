import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Chatroom, GetChatroomsByUserRequest, GetChatroomsByUserResponse } from 'src/proto/chatroom.pb';
import { LocalChatroom } from 'src/_models/localChatroom';
import { LocalRoomsAndMessagesStore } from 'src/_models/localRoomsAndMessagesStore';
import { AuthService } from 'src/_services/auth.service';
import { ChatroomService } from 'src/_services/chatroom.service';
import { MessageService } from 'src/_services/message.service';
import { StoreService } from 'src/_services/store.service';
import { NavigationService } from './navigation.service';

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss']
})
export class NavigationComponent implements OnInit {

  chatroomsList: LocalChatroom[];
  roomStore: LocalRoomsAndMessagesStore[];

  constructor(
    private router: Router,
    private chatroomService: ChatroomService,
    private authService: AuthService,
    private messageService: MessageService,
    private navService: NavigationService,
    private storeService: StoreService
  ) { }

  ngOnInit(): void {
    
    if(this.authService.isAuthenticated())
      this.getChatroomsAndMessages();
      
  }

  goToChatroom(localID: string) {
    console.log("going to "+localID);
    this.navService.updateNavStore(localID)
    this.router.navigateByUrl('/chatroom');

  }

  async getChatroomsAndMessages() {
    await this.chatroomService.getChatrooms();
    await this.messageService.getMessagesForAllChatrooms();
    this.storeService.currentChatroomsAndMessageStore.subscribe((v) => {
      console.log(v)
      this.roomStore = v;      
    });
    this.navService.updateNavStore("0")
  }
}
