import { Injectable } from '@angular/core';
import { GRPC_CLIENT_FACTORY } from '@ngx-grpc/core';
import { GrpcWebClientBase } from 'grpc-web';
import { GrpcClientSettings, GrpcStatusEvent } from '@ngx-grpc/common';
import { GRPC_USER_SERVICE_CLIENT_SETTINGS } from 'src/proto/user.pbconf';

import { UserServiceClient } from 'src/proto/user.pbsc';
import { LoggedUser } from 'src/_models/loggedUser';
import { BehaviorSubject } from 'rxjs';
import { CookieService } from 'ngx-cookie';
import { AuthService } from './auth.service';
import { SearchUserByUsernameRequest, User } from 'src/proto/user.pb';
import { AuthServiceClient } from 'src/proto/auth.pbsc';



@Injectable({ providedIn: 'root' })
export class UserService {

    currentUser: LoggedUser;

    constructor(
        private userClient: UserServiceClient,
        private authService: AuthService
    ) {
        this.authService.currentUser.subscribe((user) => {
            this.currentUser = user;
          });
     }


    async searchUserByUsername(req: SearchUserByUsernameRequest): Promise<User[]> {
        const token = this.currentUser.accessToken;

        const res = await this.userClient.searchUserByUsername(req, {"authorization": token}).toPromise();
        return res.results;
    }

}