import { Component, Input, OnInit } from '@angular/core';
import { SnackAlertService } from './snack-alert.service';

@Component({
  selector: 'app-snack-alert',
  templateUrl: './snack-alert.component.html',
  styleUrls: ['./snack-alert.component.scss']
})
export class SnackAlertComponent implements OnInit {
  message: string;
  showStatus: Boolean;

  constructor(private snackService: SnackAlertService) { }

  ngOnInit(): void {
    this.snackService.currentShowStatus.subscribe((status) => this.showStatus = status)
    this.snackService.currentMessage.subscribe((message) => this.message = message);
  }

}
