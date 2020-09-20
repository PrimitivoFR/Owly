import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { CookieService } from 'ngx-cookie';
import { BehaviorSubject } from 'rxjs';
import { LoginUserRequest } from 'src/proto/user.pb';
import { UserServiceClient } from 'src/proto/user.pbsc';
import { LoggedUser } from 'src/_models/loggedUser';
import { UserService } from './user.service';


@Injectable()
export class AuthService {
    constructor(public jwtHelper: JwtHelperService,private cookieService: CookieService, private userClient: UserServiceClient) { 
        const user:LoggedUser = <LoggedUser>this.cookieService.getObject("owly_user_cookies");
        this.user.next(user)
     }
    
    
    private user = new BehaviorSubject<LoggedUser>(new LoggedUser());
    currentUser = this.user.asObservable();

    public get currentUserValue() {
        return this.user.value;
    }

    public isAuthenticated(): boolean {
        var user: LoggedUser = <LoggedUser>this.cookieService.getObject("owly_user_cookies");

        console.log(user)
        if (!user) {
            return false
        }
        const token = user.accessToken
        
        return !this.jwtHelper.isTokenExpired(token);
    }


    createLoggedUserFromJWT(jwt: string): LoggedUser {
        var infos = this.jwtHelper.decodeToken(jwt);
        console.log(infos);
        const { email, given_name, family_name, preferred_username } = infos;
        return new LoggedUser({ username: preferred_username, firstName: given_name, lastName: family_name, email, accessToken: jwt })
    }


    async login(req: LoginUserRequest): Promise<Boolean> {
        const res = await this.userClient.loginUser(req).toPromise();
        if (res.result.accessToken == "") {
            return false;
        }
        const { accessToken, refreshToken } = res.result;
        
        const loggedUser = this.createLoggedUserFromJWT(accessToken);

        loggedUser.refreshToken = refreshToken;

        this.user.next(loggedUser);
        
        this.cookieService.putObject("owly_user_cookies", loggedUser)
        return true
    }

    logout() {
        this.user.next(new LoggedUser());
        this.cookieService.remove("owly_user_cookies");
    }
}