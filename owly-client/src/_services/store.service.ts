import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { Message } from 'src/proto/message.pb';
import { LocalRoomsAndMessagesStore } from 'src/_models/localRoomsAndMessagesStore';
import { ChatroomService } from './chatroom.service';
import { MessageService } from './message.service';

@Injectable({
    providedIn: 'root'
})


export class StoreService {

    constructor(
    ){}

    private chatroomsAndMessageStore = new BehaviorSubject<LocalRoomsAndMessagesStore[]>([]);
    currentChatroomsAndMessageStore = this.chatroomsAndMessageStore.asObservable();


    updateWholeStore(val: LocalRoomsAndMessagesStore[]) {
        this.chatroomsAndMessageStore.next(val);
    }

    emptyStore(){
        this.chatroomsAndMessageStore.next([]);
    }

    
    addTempoMessage(localId: string, message: Message) {
        const currentStore = this.chatroomsAndMessageStore.value;
        const nextStore = currentStore.map((e: LocalRoomsAndMessagesStore) => {
            if(e.localID === localId) {
                e.messages.push(message)
            }
            return e
        }) as unknown as LocalRoomsAndMessagesStore[];
        this.chatroomsAndMessageStore.next(nextStore);
    }


}