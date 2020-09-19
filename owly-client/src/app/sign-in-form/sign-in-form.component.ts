import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginUserRequest } from 'src/proto/user.pb';
import { UserService } from 'src/_services/user.service';
import { SnackAlertService } from '../common/components/snack-alert/snack-alert.service';
import { FakeLoginService, LoginResponse } from '../common/services/fake-login.service';

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
    private fakeLoginService: FakeLoginService,
    private snackService: SnackAlertService,
    private router: Router,
    private userService: UserService
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
    // let response = new LoginResponse();
    let request = new LoginUserRequest({username: this.f.username.value, password: this.f.password.value })
    const success = await this.userService.login(request);
    // response = this.fakeLoginService.login(this.f.username.value, this.f.password.value);
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
