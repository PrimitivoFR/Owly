import { Component } from '@angular/core';
import { GRPC_CLIENT_FACTORY } from '@ngx-grpc/core';
import { GrpcWebClientBase } from 'grpc-web';
import { GrpcClientSettings, GrpcStatusEvent } from '@ngx-grpc/common';
import { GRPC_USER_SERVICE_CLIENT_SETTINGS } from 'src/proto/user.pbconf';
import { CreateNewUserRequest, CreateNewUserResponse } from 'src/proto/user.pb';
import { UserServiceClient } from 'src/proto/user.pbsc';


interface CreateNewUserEvent {
  request?: CreateNewUserRequest,
  response?: CreateNewUserResponse,
  status?: GrpcStatusEvent,
  error?: GrpcStatusEvent
}

@Component({
  selector: 'app-root',
  
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'owly-client';
  constructor(private userClient: UserServiceClient) {}

  private createUserRequest() {
    return new CreateNewUserRequest(
      { firstName: "Séraphin",
        lastName: "Henry",
        username:"the_capi",
        email: "seraphin.henry@primitivo.fr",
        password:"bruh" 
      })
  }

  ngOnInit(): void {

    const request = this.createUserRequest();

    this.userClient.createNewUser(request).subscribe(
      (res) => console.log(res),
      (err) => console.log(err)
    );
    
  }
}
