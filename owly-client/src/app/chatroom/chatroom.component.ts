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
import { StoreService } from 'src/_services/store.service';
// import {v4 as uuidv4} from 'uuid'; <- Generate uuid

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
    private navService: NavigationService,
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private messageService: MessageService,
    private storeService: StoreService,
  ) { }

  ngOnInit() {
    this.navService.currentNavStore.subscribe(v => this.currentStoreItem = v)
    
    this.authService.currentUser.subscribe(v => this.currentUser = v)

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
      console.error(this.sendMsgForm.errors)
    }
  }

  async sendMessage(): Promise<Boolean> {

    // True message which goes to the db
    const message = new Message({
      authorNAME: this.currentUser.username,
      chatroomID: this.currentStoreItem.chatroom.id,
      content: this.f.message.value,
      timestamp: new Date().getTime().toString(),
      hasFileAttached: false,
      isAnswer: false
    });
    // Create tempo message
    // var tempoMess = message;
    // tempoMess.id = "TEMPO_+uuid"

    const req = new SendMessageRequest({
      message: message
    });

    try {

      // O <- Add the tempo message to the list, using a function of navigation service
    
      const res = await this.messageService.sendMessage(req);

      // ! <- Here the message has been sent correctly

      // X <- Delete the tempo message

      this.messageService.getMessagesForAllChatrooms();
      console.log(res)
      return res.success
    } catch (e) {

      // ! <- Here the message has NOT been sent correctly
      // O <- Display a snackabar "Message not sent..."
      // X <- Delete the tempo message

      return false
    }
  }

  timestampToReadableDate(timestamp: string) {
    const tmstp = parseInt(timestamp)
    const date = new Date(tmstp)

    return date.getHours()+":"+date.getMinutes() +" - "+ date.getDate() + "/" + (date.getMonth()+1)
  }

  // O <- Function that checks if the id field for the message contains "TEMPO_"

}
