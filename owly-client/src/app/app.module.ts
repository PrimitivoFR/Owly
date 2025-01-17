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
import { SignInFormComponent } from './sign-in-form/sign-in-form.component';
import { TopBarComponent } from './common/components/top-bar/top-bar.component';
import { HomeComponent } from './home/home.component';
import { ChatroomComponent } from './chatroom/chatroom.component';
import { UserInfoComponent } from './chatroom/modals/user-info-modal.component';
import { TransferOwnerModalComponent } from './chatroom/modals/transferOwnerModal.component'
import { MessageHistoryComponent } from './chatroom/modals/message-history-modal.component';

import { JwtModule } from '@auth0/angular-jwt';
import { CookieModule } from 'ngx-cookie';
import { TextareaAutosizeModule } from 'ngx-textarea-autosize';
import { GRPC_MESSAGE_SERVICE_CLIENT_SETTINGS } from 'src/proto/message.pbconf';
import { ChatroomListComponent } from './navigation/chatroom-list/chatroom-list.component';
import { GRPC_AUTH_SERVICE_CLIENT_SETTINGS } from 'src/proto/auth.pbconf';
import { environment } from './../environments/environment';
import { ConfirmModalComponent } from './common/components/confirm-modal/confirm-modal.component';


@NgModule({
  declarations: [
    AppComponent,
    SignUpFormComponent,
    NavigationComponent,
    CreateChatroomComponent,
    SnackAlertComponent,
    LoadingSpinnerComponent,
    SignInFormComponent,
    TopBarComponent,
    HomeComponent,
    ChatroomComponent,
    ChatroomListComponent,
    UserInfoComponent,
    TransferOwnerModalComponent,
    ConfirmModalComponent,
    MessageHistoryComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
    AngularEntypoModule,
    JwtModule.forRoot({}),
    CookieModule.forRoot(),
    TextareaAutosizeModule
  ],
  providers: [
    { provide: GRPC_CLIENT_FACTORY, useClass: GrpcWebClientFactory },
    { provide: GRPC_USER_SERVICE_CLIENT_SETTINGS, useValue: { host: environment.backendUrl } },
    {provide: GRPC_CHATROOM_SERVICE_CLIENT_SETTINGS, useValue: {host: environment.backendUrl}},
    { provide: GRPC_MESSAGE_SERVICE_CLIENT_SETTINGS, useValue: {host: environment.backendUrl}},
    { provide: GRPC_AUTH_SERVICE_CLIENT_SETTINGS, useValue: {host: environment.backendUrl}},
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
