import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { GRPC_CLIENT_FACTORY } from '@ngx-grpc/core';
import { GrpcWebClientFactory } from '@ngx-grpc/grpc-web-client';
import {GRPC_USER_SERVICE_CLIENT_SETTINGS} from '../proto/user.pbconf'

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule
  ],
  providers: [
    { provide: GRPC_CLIENT_FACTORY, useClass: GrpcWebClientFactory },
    { provide: GRPC_USER_SERVICE_CLIENT_SETTINGS, useValue: { host: 'http://0.0.0.0:50051' } },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
