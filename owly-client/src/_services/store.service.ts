import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { LocalRoomsAndMessagesStore } from 'src/_models/localRoomsAndMessagesStore';

@Injectable({
    providedIn: 'root'
})


export class StoreService {
    private chatroomsAndMessageStore = new BehaviorSubject<LocalRoomsAndMessagesStore[]>([]);
    currentChatroomsAndMessageStore = this.chatroomsAndMessageStore.asObservable();


    updateWholeStore(val: LocalRoomsAndMessagesStore[]) {
        this.chatroomsAndMessageStore.next(val);
    }
}