import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginUserRequest } from 'src/proto/user.pb';
import { AuthService } from 'src/_services/auth.service';
import { UserService } from 'src/_services/user.service';
import { SnackAlertService } from '../common/components/snack-alert/snack-alert.service';

@Component({
  selector: 'app-sign-in-form',
  templateUrl: './sign-in-form.component.html',
  styleUrls: ['./sign-in-form.component.scss']
})
export class SignInFormComponent implements OnInit {
  loginForm: FormGroup;
  submitted: boolean = false;
  
  constructor(
    private formBuilder: FormBuilder,
    private snackService: SnackAlertService,
    private router: Router,
    private authService: AuthService
  ) { }

  ngOnInit(): void {
    this.loginForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', [Validators.required, Validators.minLength(4)]]
    });
  }

  get f() { return this.loginForm.controls; }

  async onSubmit() {
    this.submitted = true;
    // stop here if form is invalid
    if (this.loginForm.invalid) {
        return;
    }

    let request = new LoginUserRequest({username: this.f.username.value, password: this.f.password.value })
    const success = await this.authService.login(request);

    console.log(success);
    if(success) {
      this.snackService.showSnack('Welcome !');
      this.router.navigate(['home'], { queryParams: { login: true }});
    }
    else {
      this.submitted = false;
      this.snackService.showSnack('An error has occured !');
    }
    
  }

}
