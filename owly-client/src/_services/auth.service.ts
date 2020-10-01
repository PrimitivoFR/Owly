import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { CookieService } from 'ngx-cookie';
import { BehaviorSubject } from 'rxjs';
import { AuthServiceClient } from 'src/proto/auth.pbsc';
import { Chatroom } from 'src/proto/chatroom.pb';
import { CreateNewUserRequest, LoginUserRequest } from 'src/proto/auth.pb';
import { UserServiceClient } from 'src/proto/user.pbsc';
import { LoggedUser } from 'src/_models/loggedUser';
import { ChatroomService } from './chatroom.service';
import { MessageService } from './message.service';
import { StoreService } from './store.service';
import { UserService } from './user.service';


@Injectable()
export class AuthService {
    constructor(
        public jwtHelper: JwtHelperService,
        private cookieService: CookieService, 
        private authClient: AuthServiceClient,
        private userClient: UserServiceClient,
        private storeService: StoreService) { 
        const user:LoggedUser = <LoggedUser>this.cookieService.getObject("owly_user_cookies");
        this.user.next(user)
     }
    
    
    private user = new BehaviorSubject<LoggedUser>(new LoggedUser());
    currentUser = this.user.asObservable();


    public isAuthenticated(): boolean {
        var user: LoggedUser = <LoggedUser>this.cookieService.getObject("owly_user_cookies");

        console.log(user)
        if (!user) {
            return false
        }
        const token = user.accessToken
        
        return !this.jwtHelper.isTokenExpired(token);
    }

    createUser(data) {
        const request = this.createUserRequest(data);
        this.authClient.createNewUser(request).subscribe(
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

    createLoggedUserFromJWT(jwt: string): LoggedUser {
        var infos = this.jwtHelper.decodeToken(jwt);

        const { email, given_name, family_name, preferred_username } = infos;
        return new LoggedUser({ username: preferred_username, firstName: given_name, lastName: family_name, email, accessToken: jwt })
    }


    async login(req: LoginUserRequest): Promise<Boolean> {
        try {
            const res = await this.authClient.loginUser(req).toPromise();
            if (res.result.accessToken == "") {
                return false;
            }
            const { accessToken, refreshToken } = res.result;
            
            const loggedUser = this.createLoggedUserFromJWT(accessToken);
    
            loggedUser.refreshToken = refreshToken;
    
            this.user.next(loggedUser);
            
            this.cookieService.putObject("owly_user_cookies", loggedUser)
    
            // this.messageService.getMessagesForAllChatrooms()
            return true;
        }
        catch(error) {
            console.log(error)
            return false;
        }
        
    }

    logout() {
        this.user.next(undefined);
        this.cookieService.remove("owly_user_cookies");
        this.storeService.emptyStore()
    }
}