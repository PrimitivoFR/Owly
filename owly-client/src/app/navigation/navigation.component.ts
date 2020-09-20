import { Component, OnInit } from '@angular/core';
import { Chatroom, GetChatroomsByUserRequest, GetChatroomsByUserResponse } from 'src/proto/chatroom.pb';
import { AuthService } from 'src/_services/auth.service';
import { ChatroomService } from 'src/_services/chatroom.service';

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss']
})
export class NavigationComponent implements OnInit {

  chatroomsList: GetChatroomsByUserResponse;

  constructor(
    private chatroomService: ChatroomService,
    private authService: AuthService,
  ) { }

  ngOnInit(): void {
    this.chatroomService.chatroomsListValue.subscribe((chatrooms) => {
      this.chatroomsList = chatrooms;
      console.log(this.chatroomsList);
    });
    if(this.authService.isAuthenticated())
      this.getChatrooms();
  }

  getChatrooms() {
    this.chatroomService.getChatroom();
  }
}
