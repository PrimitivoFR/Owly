import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { CreateChatroomRequest } from 'src/proto/chatroom.pb';
import { SearchUserByUsernameRequest, User } from 'src/proto/user.pb';
import { ChatroomService } from 'src/_services/chatroom.service';
import { UserService } from 'src/_services/user.service';
import { LoadingSpinnerService } from '../common/components/loading-spinner/loading-spinner.service';
import { SnackAlertService } from '../common/components/snack-alert/snack-alert.service';

@Component({
  selector: 'app-create-chatroom',
  templateUrl: './create-chatroom.component.html',
  styleUrls: ['./create-chatroom.component.scss']
})
export class CreateChatroomComponent implements OnInit {

  chatroomForm: FormGroup;
  submitted: Boolean;

  addedUsers: User[] = [];
  foundUsers: User[] = [];

  searchUsername: string = "";

  chatroomCreationSucess: Boolean = false;

  constructor(
    private formBuilder: FormBuilder,
    private userService: UserService,
    private chatroomService: ChatroomService,
    private snackService: SnackAlertService,
    private spinnerService: LoadingSpinnerService
  ) { }

  ngOnInit(): void {
    this.chatroomForm = this.formBuilder.group({
      room_name: ['', Validators.required],
      room_usersList: [],
    });
  }
  get f() { return this.chatroomForm.controls; }

  room_usersListChange() {
    this.searchUserByUsername(this.searchUsername);
  }

  async searchUserByUsername(username: string) {
    var res: User[] = [];
    if (username != "") {
      const req = new SearchUserByUsernameRequest({ username });
      res = await this.userService.searchUserByUsername(req);
    }
    res = res.filter(user => this.addedUsers.findIndex((addedUser) => addedUser.uuid == user.uuid) == -1);
    this.foundUsers = res;
  }

  addUser(foundUser: User) {
    this.addedUsers.push(foundUser);
    this.foundUsers = [];
    this.searchUsername = "";
  }

  removeUser(user: User) {
    this.addedUsers = this.addedUsers.filter((u: User) => u.uuid != user.uuid);
  }

  async createChatroom(name: string): Promise<Boolean> {
    var usersUUID = this.addedUsers.map((user: User) => user.uuid);
    const req = new CreateChatroomRequest({
      name, users: usersUUID
    })
    
    try {
      const res = await this.chatroomService.createChatroom(req);
      console.log(res)
    
      return res.success
    } catch (e) {
      
      return false
    }


  }

  async onSubmit() {
    this.submitted = true;
    // stop here if form is invalid
    if (this.chatroomForm.invalid) {
      return;
    }

    this.spinnerService.showSpinner()
    this.chatroomCreationSucess = await this.createChatroom(this.f.room_name.value);
    this.spinnerService.hideSpinner()
    
    if(this.chatroomCreationSucess) {
      this.snackService.showSnack("Chatroom successfully created !")
    } else {
      this.snackService.showSnack("An error has occured")
    }

  }

}
