import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { fromEvent, Observable } from 'rxjs';
import { messageHistory } from 'src/proto/message.pb';

@Component({
  selector: 'app-message-history-modal',
  template: `
  <div class="border border-teal-500 shadow-lg modal-container bg-white w-11/12 md:max-w-md mx-auto rounded shadow-lg z-50 overflow-y-auto">
    <div class="modal-content py-4 text-left px-6">
        <!--Title-->
        <div class="flex justify-between items-center pb-3">
            <p class="text-2xl font-bold">
              Message history
            </p>
            <div class="modal-close cursor-pointer z-50" (click)="close()">
                <svg class="fill-current text-black" xmlns="http://www.w3.org/2000/svg" width="18" height="18"
                    viewBox="0 0 18 18">
                    <path
                        d="M14.53 4.53l-1.06-1.06L9 7.94 4.53 3.47 3.47 4.53 7.94 9l-4.47 4.47 1.06 1.06L9 10.06l4.47 4.47 1.06-1.06L10.06 9z">
                    </path>
                </svg>
            </div>
        </div>
        <!--Body-->
        <div class="my-5 divide-y">
          <div *ngFor="let message of messageHistory" class="my-2 text-gray-700">
            <p><span class="mr-2 font-bold">{{timestampToReadableDate(message.timestamp)}} : </span>{{message.content}}</p>
          </div>
        </div>
    </div>
  </div>
  `
})
export class MessageHistoryComponent implements OnInit {

  @Input() messageHistory: messageHistory[];
  @Output() closeModal = new EventEmitter();
  clickObservable: Observable<Event> = fromEvent(document,'click');

  constructor(
  ) { }

  ngOnInit(): void {
    this.clickObservable.subscribe(event => {
      let modalEl = document.getElementById('messageHistoryModal');
      if(event.target == modalEl) {
        this.close();
      }
    })
  }

  timestampToReadableDate(timestamp: string) {
    const tmstp = parseInt(timestamp)
    const date = new Date(tmstp);
    return date.getHours()+":"+date.getMinutes() +" - "+ date.getDate() + "/" + (date.getMonth()+1)
  }

  close() {
    this.closeModal.emit();
  }
}