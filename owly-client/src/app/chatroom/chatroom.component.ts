import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute } from '@angular/router';
import { Message, SendMessageRequest } from 'src/proto/message.pb';
import { MessageService } from 'src/_services/message.service';

@Component({
  selector: 'app-chatroom',
  templateUrl: './chatroom.component.html',
  styleUrls: ['./chatroom.component.scss']
})
export class ChatroomComponent implements OnInit {

  idChatroom: string;
  nameChatroom: string;
  sendMsgForm: FormGroup;

  constructor(
    private route: ActivatedRoute,
    private formBuilder: FormBuilder,
    private messageService: MessageService,
  ) { }

  ngOnInit(): void {
    this.idChatroom = this.route.snapshot.paramMap.get('id');
    this.nameChatroom = this.route.snapshot.paramMap.get('name');
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

    if(res) {
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
        chatroomID: this.idChatroom,
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

}
