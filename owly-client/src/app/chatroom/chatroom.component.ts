import { AfterViewInit, Component, ElementRef, OnInit, QueryList, ViewChild, ViewChildren } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Chatroom, ChatroomUser, DeleteChatroomRequest, LeaveChatroomRequest, TranferOwnershipRequest } from 'src/proto/chatroom.pb';
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
import { Router } from '@angular/router';

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
  private isNearBottom: boolean = true;
  private dropdownOpen: boolean = false;
  private currentIndex: number = 0;
  public isAnswer: boolean = false;
  public answersTo: string = "";
  public messageToReply: Message;
  public panelOpened: boolean = false;
  public eligibleOwnerUser: ChatroomUser[] = [];
  public wantToLeave: boolean = false;

  @ViewChild('scrollframe', {static: true}) scrollFrame: ElementRef;
  @ViewChildren('item') itemElements: QueryList<any>;

  constructor(
    private navService: NavigationService,
    private formBuilder: FormBuilder,
    private authService: AuthService,
    private messageService: MessageService,
    private storeService: StoreService,
    private snackAlertService: SnackAlertService,
    private chatroomService: ChatroomService,
    private router: Router,
  ) { }

  ngOnInit() {
    this.navService.currentNavStore.subscribe(v => {
      this.currentStoreItem = v;
      this.cancelReplyTo();
      this.closeModal();
      this.panelOpened = false;
    })

    this.authService.currentUser.subscribe(v => {
      this.currentUser = v;
    })

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
    var message = new Message({
      authorNAME: this.currentUser.username,
      chatroomID: this.currentStoreItem.chatroom.id,
      content: this.f.message.value,
      hasFileAttached: false,
      isAnswer: false
    });

    if(this.isAnswer) {
      message.isAnswer = true;
      message.answersTo = this.answersTo;
    }

    const req = new SendMessageRequest({
      message: message
    });

    // Create tempo message
    var tempoMess = message;
    tempoMess.id = "TEMPO_" + uuidv4();

    try {

      // O <- Add the tempo message to the list, using a function of navigation service
      this.navService.addTempoMsg(tempoMess);
    
      const res = await this.messageService.sendMessage(req);

      // ! <- Here the message has been sent correctly

      // X <- Delete the tempo message
      this.navService.deleteTempoMsg(tempoMess.id);

      //this.messageService.getMessagesForAllChatrooms();
      console.log(res)
      this.cancelReplyTo();
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

  replyTo(messageID) {
    this.isAnswer = true;
    this.answersTo = messageID;
    this.messageToReply = this.currentStoreItem.messages.find(message => message.id == messageID);
  }

  cancelReplyTo() {
    this.isAnswer = false;
    this.answersTo = "";
    this.messageToReply = null;
  }

  messageAnswered(messageID) {
    let message = this.currentStoreItem.messages.find(message => message.id == messageID);
    if(message) {
      return `
        <div class="font-bold text-base">
          ${message.authorNAME}
        </div>
        <div class="w-full text-sm">
          ${message.content}
        </div>
      `;
    }
    else {
      return `
        <div class="font-bold text-red-600 text-base">
          Message unavailable
        </div>
      `;
    }
    
  }


  async leaveChatroom() {
    if(this.isChatroomOwner()) {
      this.wantToLeave = true;
      this.openModal();
    }
    else {
      if(confirm("Are you sure you want to continue ?")) {
        const req = new LeaveChatroomRequest({
          chatroomId: this.currentStoreItem.chatroom.id
        })
        const chatroomName = this.currentStoreItem.chatroom.name;
        
        const res = await this.chatroomService.leaveChatroom(req);
    
        if(res.success) {
          this.snackAlertService.showSnack("You have left the chatroom : " + chatroomName);
          if(!this.navService.updateNavStore("0")) {
            this.router.navigate(['home']);
          }
          this.panelOpened = false;
        }
        else {
          this.snackAlertService.showSnack("You can't leave the chatroom when you are the owner");
        }
      }
    }
  }

  async deleteChatroom() {
    if(confirm("Are you sure you want to continue ?")) {
      const req = new DeleteChatroomRequest({
        chatroomId: this.currentStoreItem.chatroom.id
      })
      const chatroomName = this.currentStoreItem.chatroom.name;
      const res = await this.chatroomService.deleteChatroom(req);

      if(res.success) {
        this.snackAlertService.showSnack("You have deleted the chatroom : " + chatroomName);
        if(!this.navService.updateNavStore("0")) {
          this.router.navigate(['home']);
        }
        this.panelOpened = false;
      }
      else {
        this.snackAlertService.showSnack("Something went wrong, you can't leave the chatroom");
      }
    }
  }

  async transferOwnership(idUserChosen: string) {
    if(confirm("Are you sure you want to continue ?")) {
      if(!idUserChosen || idUserChosen == "") {
        return;
      }
      
      const req = new TranferOwnershipRequest({
        chatroomId: this.currentStoreItem.chatroom.id,
        newOwnerId: idUserChosen
      })

      const res = await this.chatroomService.transfertOwnershipChatroom(req);

      if(res.success) {
        this.closeModal();
        if(this.wantToLeave) {
          this.leaveChatroom();
        }
      }
      else {
        this.closeModal();
        this.snackAlertService.showSnack("Something went wrong.");
      }
    }
  }

  
  openModal() {
    this.eligibleOwnerUser = this.currentStoreItem.chatroom.users.slice()
    this.eligibleOwnerUser.splice(this.eligibleOwnerUser.findIndex(user => user.uuid == this.currentStoreItem.chatroom.owner),1);

    let modal = document.getElementById('transfermodal');
    modal.classList.remove('fadeOut');
    modal.classList.add('fadeIn');
    modal.style.display = "flex";
  }

  closeModal() {
    this.wantToLeave = false;
    let modal = document.getElementById('transfermodal');
    modal.classList.remove('fadeIn');
    modal.classList.add('fadeOut');
    setTimeout(() => {
      modal.style.display = 'none';
      this.eligibleOwnerUser = [];
    }, 500);
  }

  isChatroomOwner(): boolean {
    return this.authService.matchUUIDvsTOKEN(this.currentUser.accessToken, this.currentStoreItem.chatroom.owner);
  }

  timestampToReadableDate(timestamp: string) {
    const tmstp = parseInt(timestamp)
    const date = new Date(tmstp);
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
