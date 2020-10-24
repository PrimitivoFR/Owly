import { Injectable } from '@angular/core';
import { BehaviorSubject, Observable, Subscription } from 'rxjs';
import { LocalRoomsAndMessagesStore } from 'src/_models/localRoomsAndMessagesStore';
import { ChatroomService } from 'src/_services/chatroom.service';
import { map, filter } from 'rxjs/operators';
import { MessageService } from 'src/_services/message.service';
import { StoreService } from 'src/_services/store.service';
import { Message, StreamMessagesByChatroomRequest, StreamMessagesByChatroomResponse } from 'src/proto/message.pb';


@Injectable({
    providedIn: 'root'
  })
  
export class NavigationService {

    constructor(
        private messageService: MessageService,
        private storeService: StoreService
    ) { }

    private navStore = new BehaviorSubject<LocalRoomsAndMessagesStore>(new LocalRoomsAndMessagesStore());
    currentNavStore = this.navStore.asObservable();

    private currentMessageStream = new Subscription();

    updateNavStore(localID: string): boolean {
        if(this.currentMessageStream != null) {
            // We have to do this, else the subscription will be recreated each time
            // we switch channels, which leads to receive the messages twice in the chat
            this.currentMessageStream.unsubscribe();
        }

        //this.messageService.getMessagesForAllChatrooms()

        var currentStoreItem: LocalRoomsAndMessagesStore;
        var hasChatrooms: boolean;

        const currentStore = this.storeService.currentChatroomsAndMessageStore.pipe(
                map(items => items.find(e => e.localID === localID))
                );
        
        currentStore.subscribe((v: LocalRoomsAndMessagesStore) => {
            if(v == undefined) {
                currentStoreItem = new LocalRoomsAndMessagesStore();
                hasChatrooms = false;
            } else {
                currentStoreItem = v;
                hasChatrooms = true;
            }
            this.navStore.next(currentStoreItem)
        })
        
        // update messages in real time by listening to the gRPC stream server
        
        if(currentStoreItem.chatroom) {
            this.currentMessageStream = this.messageService.streamMessagesByChatroom(new StreamMessagesByChatroomRequest({
                chatroomID: currentStoreItem.chatroom.id
            })).subscribe((res) => {
                if(res.operation == "CREATE") {
                    currentStoreItem.messages.push(res.message)
                }
                if(res.operation == "DELETE") {
                    currentStoreItem.messages = currentStoreItem.messages.filter((mess: Message) => mess.id != res.message.id) // Deletes the message: return every 
                }
                
            });
        }

        return hasChatrooms;
        
    }


    // O <- Function to add a tempo Message
    // Params: message
    // currentStoreItem.messages.push
    addTempoMsg(message: Message) {
        this.navStore.value.messages.push(message);
    }


    // O <- Function to remove the tempo message
    // Params: tempo uuid of the message
    // currentStoreItem.messages.filter or find then remove
    // Do NOT USE pop() !!

    deleteTempoMsg(id) {
        this.navStore.value.messages.splice(this.navStore.value.messages.findIndex(message => message.id == id), 1);
    }
}