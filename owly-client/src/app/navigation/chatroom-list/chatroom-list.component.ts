import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LocalRoomsAndMessagesStore } from 'src/_models/localRoomsAndMessagesStore';
import { LoggedUser } from 'src/_models/loggedUser';
import { AuthService } from 'src/_services/auth.service';
import { ChatroomService } from 'src/_services/chatroom.service';
import { MessageService } from 'src/_services/message.service';
import { StoreService } from 'src/_services/store.service';
import { NavigationService } from '../navigation.service';

@Component({
  selector: 'app-chatroom-list',
  templateUrl: './chatroom-list.component.html',
  styleUrls: ['./chatroom-list.component.scss']
})
export class ChatroomListComponent implements OnInit {
  roomStore: LocalRoomsAndMessagesStore[];
  currentUser: LoggedUser;

  constructor(
    private router: Router,
    private chatroomService: ChatroomService,
    private authService: AuthService,
    private messageService: MessageService,
    private navService: NavigationService,
    private storeService: StoreService
  ) { }

  ngOnInit(): void {
    this.storeService.currentChatroomsAndMessageStore.subscribe(v => this.roomStore = v);
    this.authService.currentUser.subscribe(u => {
      if(u != undefined) {
        if(this.authService.isAuthenticated()) {
          this.getChatroomsAndMessages()
        }
      }
    })

  }

  goToChatroom(localID: string) {

    this.navService.updateNavStore(localID)
    this.router.navigateByUrl('/chatroom');

  }

  async getChatroomsAndMessages() {
    await this.chatroomService.getChatrooms();
    await this.messageService.getMessagesForAllChatrooms();
    this.storeService.currentChatroomsAndMessageStore.subscribe((v) => {
      this.roomStore = v;      
    });
    this.navService.updateNavStore("0")
  }

}
