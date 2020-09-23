import { RecursivePartial } from '@ngx-grpc/common';
import { Chatroom } from 'src/proto/chatroom.pb';
import { Message } from 'src/proto/message.pb';

export class LocalChatroom {
    localID: string;
    chatroom: Chatroom;
    
  
    constructor(_value?: RecursivePartial<LocalChatroom>) {
      _value = _value || {}; 
    
        this.localID = _value.localID;
        this.chatroom = _value.chatroom;
    }
  }