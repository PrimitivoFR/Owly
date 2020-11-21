import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class ConfirmModalService { 

    private showStatus = new BehaviorSubject<boolean>(false);
    currentShowStatus = this.showStatus.asObservable();

    private message = new BehaviorSubject<string>("");
    currentMessage = this.message.asObservable();

    private decision = new BehaviorSubject<boolean>(null);
    currentDecision = this.decision.asObservable();

    async confirm(message: string): Promise<boolean> {
        this.openModal(message);

        await this.waitForDecision();

        this.closeModal();
        
        return this.decision.value;
    }

    private openModal(message: string) {
        this.decision.next(null);
        this.message.next(message);
        this.showStatus.next(true);
    }

    private closeModal() {
        this.message.next("");
        this.showStatus.next(false);
    }

    async waitForDecision(): Promise<void> {
        return new Promise(resolve => {
            this.currentDecision.subscribe((decision) => {
                console.log(decision)
                if(decision != null) 
                    resolve();
            })
        });
    }

    accept() { this.decision.next(true) }
    
    decline() { this.decision.next(false) }


}