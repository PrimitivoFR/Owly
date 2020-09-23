import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { LocalRoomsAndMessagesStore } from 'src/_models/localRoomsAndMessagesStore';
import { ChatroomService } from 'src/_services/chatroom.service';
import { map, filter } from 'rxjs/operators';
import { MessageService } from 'src/_services/message.service';
import { StoreService } from 'src/_services/store.service';


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

    

    updateNavStore(localID: string) {
        this.messageService.getMessagesForAllChatrooms()

        var currentStoreItem: LocalRoomsAndMessagesStore;
        const currentStore = this.storeService.currentChatroomsAndMessageStore.pipe(
                map(items => items.find(e => e.localID === localID))
                );
        
        currentStore.subscribe((v: LocalRoomsAndMessagesStore) => {
            if(v == undefined) {
            currentStoreItem = new LocalRoomsAndMessagesStore()
            } else {
            currentStoreItem = v
            }
        })
        this.navStore.next(currentStoreItem)
    }
}