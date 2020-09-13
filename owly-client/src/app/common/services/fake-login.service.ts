import { Injectable } from "@angular/core";
import { RecursivePartial } from '@ngx-grpc/common';
import { v4 as uuidv4 } from 'uuid';

class LoginResponse {

    accessToken: string;
    uuid: string;
    expiresIn: number;
    success: boolean;
    message: string;

    constructor(_value?: RecursivePartial<LoginResponse>) {
        _value = _value || {};
        this.accessToken = _value.accessToken;
        this.expiresIn = 3600;
        this.uuid = _value.uuid;
        this.success = _value.success;
        this.message = _value.message;
    }
}

@Injectable({ providedIn: 'root' })
export class FakeLoginService {
    // FakeLoginService to implement login logic in the client, waiting for the backend to be done

    login(username: string, password: string): LoginResponse {

        if (username == "test" && password == "test") {
            const uuid = uuidv4();
            const fakeToken = uuidv4();
            return new LoginResponse({ accessToken: fakeToken, uuid, success: true })
        }
        return new LoginResponse({success: false, message:"WRONG_CREDS"});
    }


}