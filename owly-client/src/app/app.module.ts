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

@NgModule({
  declarations: [
    AppComponent,
    SignUpFormComponent,
    NavigationComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    ReactiveFormsModule,
  ],
  providers: [
    { provide: GRPC_CLIENT_FACTORY, useClass: GrpcWebClientFactory },
    { provide: GRPC_USER_SERVICE_CLIENT_SETTINGS, useValue: { host: 'http://localhost:8085' } },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
