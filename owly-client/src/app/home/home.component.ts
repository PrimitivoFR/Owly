import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { AuthService } from 'src/_services/auth.service';
import { SnackAlertService } from '../common/components/snack-alert/snack-alert.service';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.scss']
})
export class HomeComponent implements OnInit {

  constructor(
    private route: ActivatedRoute,
    private snackService: SnackAlertService,
    private authService: AuthService,
  ) { }

  ngOnInit(): void {
    if (this.route.snapshot.queryParams['login']) {
      this.snackService.showSnack('Connection successfully');
    }
  }

  logout() {
    this.authService.logout();
  }

}
