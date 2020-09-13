import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
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
  ) { }

  ngOnInit(): void {
    this.loginForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', [Validators.required, Validators.minLength(4)]]
    });
  }

  get f() { return this.loginForm.controls; }

  onSubmit() {
    this.submitted = true;
    // stop here if form is invalid
    if (this.loginForm.invalid) {
        return;
    }
    let response = new LoginResponse();
    response = this.fakeLoginService.login(this.f.username.value, this.f.password.value);

    if(response.success) {
      this.snackService.showSnack('Connection successfull');
      this.router.navigate(['home'], { queryParams: { login: true }});
    }
    else
      this.snackService.showSnack(response.message);

    this.submitted = false;
    
  }

}
