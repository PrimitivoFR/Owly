import { Injectable } from '@angular/core';
import { GRPC_CLIENT_FACTORY } from '@ngx-grpc/core';
import { GrpcWebClientBase } from 'grpc-web';
import { GrpcClientSettings, GrpcStatusEvent } from '@ngx-grpc/common';
import { GRPC_USER_SERVICE_CLIENT_SETTINGS } from 'src/proto/user.pbconf';
import { CreateNewUserRequest, CreateNewUserResponse, JWT, LoginUserRequest, SearchUserByUsernameRequest, User } from 'src/proto/user.pb';
import { UserServiceClient } from 'src/proto/user.pbsc';
import { LoggedUser } from 'src/_models/loggedUser';
import { BehaviorSubject } from 'rxjs';
import { CookieService } from 'ngx-cookie';
import { AuthService } from './auth.service';

interface CreateNewUserEvent {
    request?: CreateNewUserRequest,
    response?: CreateNewUserResponse,
    status?: GrpcStatusEvent,
    error?: GrpcStatusEvent
}


@Injectable({ providedIn: 'root' })
export class UserService {

    constructor(
        private userClient: UserServiceClient,
        private cookieService: CookieService,
        private authService: AuthService
    ) { 
        
    }


    createUser(data) {
        const request = this.createUserRequest(data);
        this.userClient.createNewUser(request).subscribe(
            (res) => console.log(res),
            (err) => console.log(err)
        );
    }

    private createUserRequest(data) {
        return new CreateNewUserRequest({
            firstName: data.firstName,
            lastName: data.lastName,
            username: data.username,
            email: data.email,
            password: data.password
        })
    }

    async searchUserByUsername(req: SearchUserByUsernameRequest): Promise<User[]> {
        const res = await this.userClient.searchUserByUsername(req).toPromise();
        return res.results;
    }

}