import { Injectable } from "@angular/core";
import { BehaviorSubject } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class LoadingSpinnerService{

    private showStatus = new BehaviorSubject<Boolean>(false);
    currentShowStatus = this.showStatus.asObservable();


    showSpinner() {
        this.showStatus.next(true)
    }
    hideSpinner() {
        this.showStatus.next(false)
    }
    
}