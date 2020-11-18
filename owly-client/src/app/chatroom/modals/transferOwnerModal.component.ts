import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { ChatroomUser } from 'src/proto/chatroom.pb';

@Component({
  selector: 'app-transfer-owner-modal',
  template: `
  <div class="border border-teal-500 shadow-lg modal-container bg-white w-11/12 md:max-w-md mx-auto rounded shadow-lg z-50 overflow-y-auto">
    <div class="modal-content py-4 text-left px-6">
      <!--Title-->
      <div class="flex justify-between items-center pb-3">
        <p class="text-2xl font-bold">
          Who will be the next owner?
        </p>
        <div class="modal-close cursor-pointer z-50" (click)="cancelTransfer()">
          <svg class="fill-current text-black" xmlns="http://www.w3.org/2000/svg" width="18" height="18"
            viewBox="0 0 18 18">
            <path
              d="M14.53 4.53l-1.06-1.06L9 7.94 4.53 3.47 3.47 4.53 7.94 9l-4.47 4.47 1.06 1.06L9 10.06l4.47 4.47 1.06-1.06L10.06 9z">
            </path>
          </svg>
        </div>
      </div>
      <!--Body-->
      <div class="my-5">
        <div *ngFor="let user of listUsers" class="mx-4 my-2 px-2 py-2 rounded-md text-center text-blue-600 bg-gray-200 cursor-pointer hover:text-gray-400" (click)="confirmTransferOwner(user.uuid)">
          {{user.username}}
        </div>
      </div>
      <!--Footer-->
      <div class="flex justify-end pt-2">
        <button
          class="focus:outline-none modal-close px-4 bg-gray-400 p-3 rounded-lg text-black hover:bg-gray-300"
          (click)="cancelTransfer()">
          Cancel
        </button>
      </div>
    </div>
  </div>
  `
})
export class TransferOwnerModalComponent implements OnInit {

  @Input() listUsers: ChatroomUser[];
  @Output() idUserChosen = new EventEmitter<string>();
  @Output() cancelAction = new EventEmitter<boolean>();

  constructor(
  ) { }

  ngOnInit(): void {

  }

  cancelTransfer() {
    this.cancelAction.emit(true);
  }

  confirmTransferOwner(idNewOwner: string) {
    this.idUserChosen.emit(idNewOwner);
  }


}
