import { Component, OnInit } from '@angular/core';
import { Chatroom, GetChatroomsByUserRequest, GetChatroomsByUserResponse } from 'src/proto/chatroom.pb';
import { ChatroomService } from 'src/_services/chatroom.service';

@Component({
  selector: 'app-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.scss']
})
export class NavigationComponent implements OnInit {

  chatroomsList: Chatroom[];

  constructor(
    private chatroomService: ChatroomService,
  ) { }

  ngOnInit(): void {
    this.load();
  }

  async load() {
    await this.getChatrooms();
  }

  async getChatrooms() {
    const req = new GetChatroomsByUserRequest() ;

    try {
      this.chatroomsList = await (await this.chatroomService.getChatroom(req)).chatrooms;
      console.log(this.chatroomsList);
    } catch (e) {
      
    }
  }
}
