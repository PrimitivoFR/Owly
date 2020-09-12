import { Component, OnInit } from '@angular/core';
import { LoadingSpinnerService } from './loading-spinner.service';

@Component({
  selector: 'app-loading-spinner',
  templateUrl: './loading-spinner.component.html',
  styleUrls: ['./loading-spinner.component.scss']
})
export class LoadingSpinnerComponent implements OnInit {

  showStatus: Boolean

  constructor(private spinnerService: LoadingSpinnerService) { }

  ngOnInit(): void {
    this.spinnerService.currentShowStatus.subscribe((status) => this.showStatus = status);
  }

}
