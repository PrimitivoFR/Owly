import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Chatroom } from 'src/proto/chatroom.pb';
import { GetMessagesByChatroomRequest, Message, SendMessageRequest } from 'src/proto/message.pb';
import { LocalChatroom } from 'src/_models/localChatroom';
import { LocalMessages } from 'src/_models/localMessages';
import { LocalRoomsAndMessagesStore } from 'src/_models/localRoomsAndMessagesStore';
import { ChatroomService } from 'src/_services/chatroom.service';
import { MessageService } from 'src/_services/message.service';
import { map, filter } from 'rxjs/operators';
import { Observable } from 'rxjs';
import { LoggedUser } from 'src/_models/loggedUser';
import { AuthService } from 'src/_services/auth.service';
import { Location } from '@angular/common';
import { NavigationService } from '../navigation/navigation.service';

@Component({
  selector: 'app-chatroom',
  templateUrl: './chatroom.component.html',
  styleUrls: ['./chatroom.component.scss']
})
export class ChatroomComponent implements OnInit {

  localIdChatroom: string;
  nameChatroom: string;
  sendMsgForm: FormGroup;

  private currentStore: Observable<LocalRoomsAndMessagesStore>
  currentStoreItem: LocalRoomsAndMessagesStore = new LocalRoomsAndMessagesStore();

  private currentUser: LoggedUser;

  constructor(
    private route: ActivatedRoute,
    private navService: NavigationService,
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private messageService: MessageService,
    private chatroomService: ChatroomService
  ) { }



  ngOnInit() {
    this.navService.currentNavStore.subscribe(v => this.currentStoreItem = v)
    
    this.authService.currentUser.subscribe(v => this.currentUser = v)
    console.log("currentStoreItem: ")
    console.log(this.currentStoreItem)

    this.sendMsgForm = this.formBuilder.group({
      message: ['', Validators.required],
    });

  }

  get f() { return this.sendMsgForm.controls; }

  async onSubmit() {
    // stop here if form is invalid
    if (this.sendMsgForm.invalid) {
      return;
    }

    const res = await this.sendMessage();

    if (res) {
      this.sendMsgForm.reset();
    }
    else {
      console.log("something went wrong");
    }
  }

  async sendMessage(): Promise<Boolean> {
    const req = new SendMessageRequest({
      message: new Message({
        authorUUID: "",
        authorNAME: "",
        chatroomID: this.currentStoreItem.chatroom.id,
        content: this.f.message.value,
        timestamp: new Date().getTime().toString(),
        hasFileAttached: false,
        isAnswer: false
      })
    });

    try {
      const res = await this.messageService.sendMessage(req);
      console.log(res)

      return res.success
    } catch (e) {

      return false
    }
  }

  timestampToReadableDate(timestamp: string) {
    const tmstp = parseInt(timestamp)
    const date = new Date(tmstp)

    return date.getHours()+":"+date.getMinutes() +" - "+ date.getDate() + "/" + (date.getMonth()+1)
  }

}
