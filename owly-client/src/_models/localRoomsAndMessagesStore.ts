import { RecursivePartial } from '@ngx-grpc/common';
import { Message } from 'src/proto/message.pb';
import { LocalChatroom } from './localChatroom';

export class LocalRoomsAndMessagesStore extends LocalChatroom {
    messages: Message[];

    constructor() {
        super();      
      }
}