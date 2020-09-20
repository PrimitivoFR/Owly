import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-chatroom',
  templateUrl: './chatroom.component.html',
  styleUrls: ['./chatroom.component.scss']
})
export class ChatroomComponent implements OnInit {

  idChatroom: string;
  nameChatroom: string;

  constructor(
    private route: ActivatedRoute,
  ) { }

  ngOnInit(): void {
    this.idChatroom = this.route.snapshot.paramMap.get('id');
    this.nameChatroom = this.route.snapshot.paramMap.get('name');
  }

}
