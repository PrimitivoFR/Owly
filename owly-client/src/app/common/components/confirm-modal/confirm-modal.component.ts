import { Component, OnInit } from '@angular/core';
import { ConfirmModalService } from './confirm-modal.service';

@Component({
  selector: 'app-confirm-modal',
  templateUrl: './confirm-modal.component.html',
  styleUrls: ['./confirm-modal.component.scss']
})
export class ConfirmModalComponent implements OnInit {

  showStatus: boolean
  message: string

  constructor(private confirmModalService: ConfirmModalService) { }

  ngOnInit(): void {
    this.confirmModalService.currentShowStatus.subscribe((status) => {
      if(status == true)
        this.openModal();
      else
        this.closeModal();
    });

    this.confirmModalService.currentMessage.subscribe((message) => this.message = message);
  }

  private openModal() {
    let modalEl = document.getElementById('confirmModal');
    modalEl.classList.remove('fadeOut');
    modalEl.classList.add('fadeIn');
    modalEl.style.display = "flex";
  }

  private closeModal() {
    let modalEl = document.getElementById('confirmModal');
    modalEl.classList.remove('fadeIn');
    modalEl.classList.add('fadeOut');
    setTimeout(() => {
      modalEl.style.display = 'none';
    }, 500);
  }

  private accept() {
    this.confirmModalService.accept();
  }

  private decline() {
    this.confirmModalService.decline();
  }

}
