import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { fromEvent, Observable } from 'rxjs';

@Component({
  selector: 'app-user-info-modal',
  template: `
  <div class="border border-teal-500 shadow-lg modal-container bg-white w-11/12 md:max-w-md mx-auto rounded shadow-lg z-50 overflow-y-auto">
    <div class="modal-content py-4 text-left px-6">
        <!--Title-->
        <div class="flex justify-between items-center pb-3">
            <p class="text-2xl font-bold">
              Information about {{userInfo.username}}
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
        <div class="flex my-5 divide">
            <div class="px-2">
                <entypo class="text-5xl" name="user"></entypo>
            </div>
            <div class="flex-col px-2">
                <div><span class="font-bold">Username:</span> {{userInfo.username}}
                <div><span class="font-bold">Name:</span> {{userInfo.firstname}} {{userInfo.lastname}}</div>
                <div><span class="font-bold">Email:</span> {{userInfo.email}}</div>
            </div>
        </div>
    </div>
  </div>
  `
})
export class UserInfoComponent implements OnInit {

  @Input() userInfo: any/*GetUserInfosResponse*/;
  @Output() closeModal = new EventEmitter();
  clickObservable: Observable<Event> = fromEvent(document,'click');
  
  constructor(
  ) { }

  ngOnInit(): void {
    this.clickObservable.subscribe(event => {
        let modalEl = document.getElementById('userInfoModal');
        if(event.target == modalEl) {
          this.close();
        }
      })
  }

  close() {
    this.closeModal.emit();
  }
}