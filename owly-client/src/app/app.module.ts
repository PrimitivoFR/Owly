import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { GRPC_CLIENT_FACTORY } from '@ngx-grpc/core';
import { GrpcWebClientFactory } from '@ngx-grpc/grpc-web-client';
import { GRPC_USER_SERVICE_CLIENT_SETTINGS } from '../proto/user.pbconf'
import { ReactiveFormsModule } from '@angular/forms';
import { SignUpFormComponent } from './sign-up-form/sign-up-form.component';
import { NavigationComponent } from './navigation/navigation.component';
import { CreateChatroomComponent } from './create-chatroom/create-chatroom.component';
import { GRPC_CHATROOM_SERVICE_CLIENT_SETTINGS } from 'src/proto/chatroom.pbconf';
import { SnackAlertComponent } from './common/components/snack-alert/snack-alert.component';
import { LoadingSpinnerComponent } from './common/components/loading-spinner/loading-spinner.component';
import { AngularEntypoModule, AngularEntypoComponent } from 'angular-entypo';

@NgModule({
  declarations: [
    AppComponent,
    SignUpFormComponent,
    NavigationComponent,
    CreateChatroomComponent,
    SnackAlertComponent,
    LoadingSpinnerComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    AngularEntypoModule
  ],
  providers: [
    { provide: GRPC_CLIENT_FACTORY, useClass: GrpcWebClientFactory },
    { provide: GRPC_USER_SERVICE_CLIENT_SETTINGS, useValue: { host: 'http://localhost:8085' } },
    {provide: GRPC_CHATROOM_SERVICE_CLIENT_SETTINGS, useValue: {host: 'http://localhost:8085'}}
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
