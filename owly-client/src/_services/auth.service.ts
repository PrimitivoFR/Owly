import { Injectable } from '@angular/core';
import { JwtHelperService } from '@auth0/angular-jwt';
import { CookieService } from 'ngx-cookie';
import { LoggedUser } from 'src/_models/loggedUser';
import { UserService } from './user.service';


@Injectable()
export class AuthService {
    constructor(public jwtHelper: JwtHelperService, private cookieService: CookieService) { }
    // ...
    public isAuthenticated(): boolean {
        var user: LoggedUser = <LoggedUser>this.cookieService.getObject("owly_user_cookies")
        console.log(user)
        if (!user) {
            return false
        }
        const token = user.accessToken
        
        return !this.jwtHelper.isTokenExpired(token);
    }


    createLoggedUserFromJWT(jwt: string): LoggedUser {
        var infos = this.jwtHelper.decodeToken(jwt);
        const { email, given_name, family_name, preferred_username } = infos;
        return new LoggedUser({ username: preferred_username, firstName: given_name, lastName: family_name, email, accessToken: jwt })
    }
}