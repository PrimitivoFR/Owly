import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class SnackAlertService { 

    private showStatus = new BehaviorSubject<Boolean>(false);
    currentShowStatus = this.showStatus.asObservable();

    private message = new BehaviorSubject<string>("");
    currentMessage = this.message.asObservable();

    showSnack(message?: string, duration?: number) {
        this.message.next(message);
        this.showStatus.next(true)
        setTimeout(()=>this.showStatus.next(false), duration??2000);
    }


}