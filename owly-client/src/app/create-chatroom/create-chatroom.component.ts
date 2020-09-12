import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { SearchUserByUsernameRequest, User } from 'src/proto/user.pb';
import { UserService } from 'src/_services/user.service';

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

  constructor(
    private formBuilder: FormBuilder,
    private userService: UserService
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
      const req = new SearchUserByUsernameRequest({username});
      res = await this.userService.searchUserByUsername(req);
    }
    res = res.filter(user => this.addedUsers.findIndex((addedUser) => addedUser.uuid == user.uuid ) == -1);
    this.foundUsers = res;
  }

  addUser(foundUser: User) {
    this.addedUsers.push(foundUser);
    this.foundUsers = [];
    this.searchUsername = "";
  }


  onSubmit() {
    this.submitted = true;
    // stop here if form is invalid
    if (this.chatroomForm.invalid) {
        return;
    }

    console.log(this.f)
    
  }

}
