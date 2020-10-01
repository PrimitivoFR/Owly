import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginUserRequest } from 'src/proto/auth.pb';
import { AuthService } from 'src/_services/auth.service';
import { ChatroomService } from 'src/_services/chatroom.service';
import { MessageService } from 'src/_services/message.service';
import { UserService } from 'src/_services/user.service';
import { LoadingSpinnerService } from '../common/components/loading-spinner/loading-spinner.service';
import { SnackAlertService } from '../common/components/snack-alert/snack-alert.service';
import { NavigationService } from '../navigation/navigation.service';

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
    private authService: AuthService,
    private chatroomService: ChatroomService,
    private messageService: MessageService,
    private navService: NavigationService,
    private loadingSpinnerService: LoadingSpinnerService
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
    this.loadingSpinnerService.showSpinner();
    const success = await this.authService.login(request);
    this.loadingSpinnerService.hideSpinner();

    if(success) {
      this.snackService.showSnack('Welcome !');
      await this.chatroomService.getChatrooms();
      this.messageService.getMessagesForAllChatrooms();
      this.navService.updateNavStore("0");
      this.router.navigate(['chatroom']);
    }
    else {
      this.submitted = false;
      this.snackService.showSnack('An error has occured !');
    }
    
  }

}
