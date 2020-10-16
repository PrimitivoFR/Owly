import { AfterViewInit, Component, ElementRef, OnInit, QueryList, ViewChild, ViewChildren } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Chatroom } from 'src/proto/chatroom.pb';
import { DeleteMessageRequest, GetMessagesByChatroomRequest, Message, SendMessageRequest } from 'src/proto/message.pb';
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
import {v4 as uuidv4} from 'uuid';
import { SnackAlertService } from '../common/components/snack-alert/snack-alert.service';

@Component({
  selector: 'app-chatroom',
  templateUrl: './chatroom.component.html',
  styleUrls: ['./chatroom.component.scss']
})
export class ChatroomComponent implements OnInit, AfterViewInit {

  localIdChatroom: string;
  nameChatroom: string;
  sendMsgForm: FormGroup;

  private currentStore: Observable<LocalRoomsAndMessagesStore>
  currentStoreItem: LocalRoomsAndMessagesStore = new LocalRoomsAndMessagesStore();

  private currentUser: LoggedUser;
  private scrollContainer: any;
  private isNearBottom = true;
  private dropdownOpen = false;
  private currentIndex = 0;

  @ViewChild('scrollframe', {static: true}) scrollFrame: ElementRef;
  @ViewChildren('item') itemElements: QueryList<any>;

  constructor(
    private navService: NavigationService,
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private messageService: MessageService,
    private storeService: StoreService,
    private snackAlertService: SnackAlertService,
  ) { }

  ngOnInit() {
    this.navService.currentNavStore.subscribe(v => this.currentStoreItem = v)
    
    this.authService.currentUser.subscribe(v => this.currentUser = v)

    this.sendMsgForm = this.formBuilder.group({
      message: ['', Validators.required],
    });
  }
  
  ngAfterViewInit() {
    this.scrollContainer = this.scrollFrame.nativeElement;
    this.itemElements.changes.subscribe(res  => {
      console.log(res)
      this.onItemElementsChanged();
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
    var tempoMess = message;
    tempoMess.id = "TEMPO_" + uuidv4();

    const req = new SendMessageRequest({
      message: message
    });

    try {

      // O <- Add the tempo message to the list, using a function of navigation service
      this.navService.addTempoMsg(tempoMess);
    
      const res = await this.messageService.sendMessage(req);

      // ! <- Here the message has been sent correctly

      // X <- Delete the tempo message
      this.navService.deleteTempoMsg(tempoMess.id);

      //this.messageService.getMessagesForAllChatrooms();
      console.log(res)
      return res.success
    } catch (e) {

      // ! <- Here the message has NOT been sent correctly
      // O <- Display a snackabar "Message not sent..."
      this.snackAlertService.showSnack("Something went wrong, the message was not sent");

      return false
    }
  }

  async deleteMessage(messageID) {
    if(confirm("Are you sure you want to continue ?")) {
      this.dropdownOpen = false;
      const req = new DeleteMessageRequest({
        chatroomID: this.currentStoreItem.chatroom.id,
        messageID: messageID
      })

      const res = await this.messageService.deleteMessage(req);

      if(!res.success) {
        this.snackAlertService.showSnack("Something went wrong, the message was not deleted");
      }
    }
    else {
      this.dropdownOpen = false;
    }


  }

  timestampToReadableDate(timestamp: string) {
    const tmstp = parseInt(timestamp)
    const date = new Date(tmstp)

    return date.getHours()+":"+date.getMinutes() +" - "+ date.getDate() + "/" + (date.getMonth()+1)
  }

  EnterSubmit(event) {
    if(event.which == 13) {
      this.onSubmit();
    }
  }

  // O <- Function that checks if the id field for the message contains "TEMPO_"
  checkTempoMsg(id): boolean {
    if(id.includes('TEMPO_'))
      return true;
    return false;
  }

  private onItemElementsChanged(): void {
    if (this.isNearBottom) {
      this.scrollToBottom();
    }
  }

  private scrollToBottom(): void {
    console.log(this.scrollContainer);
    this.scrollContainer.scroll({
      top: this.scrollContainer.scrollHeight,
      behavior: 'smooth'
    });
  }

  private isUserNearBottom(): boolean {
    const threshold = 150;
    const position = this.scrollContainer.scrollTop + this.scrollContainer.offsetHeight;
    const height = this.scrollContainer.scrollHeight;
    return position > height - threshold;
  }

  scrolled(event: any): void {
    this.isNearBottom = this.isUserNearBottom();
  }

}
