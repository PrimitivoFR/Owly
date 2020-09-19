import { RecursivePartial } from '@ngx-grpc/common';

export class LoggedUser {
    firstName: string;
    lastName: string;
    username: string;
    email: string;
    accessToken: string;
    refreshToken: string;

    constructor(_value?: RecursivePartial<LoggedUser>) {
        _value = _value || {};
        this.firstName = _value.firstName;
        this.lastName = _value.lastName;
        this.username = _value.username;
        this.email = _value.email;
        this.accessToken = _value.accessToken;
        this.refreshToken = _value.refreshToken;

    }

}