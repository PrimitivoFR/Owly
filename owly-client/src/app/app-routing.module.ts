import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CreateChatroomComponent } from './create-chatroom/create-chatroom.component';
import { SignUpFormComponent } from './sign-up-form/sign-up-form.component';
import { SignInFormComponent } from './sign-in-form/sign-in-form.component';
import { HomeComponent } from './home/home.component';
import { AuthGuardGuard as AuthGuard } from './auth-guard.guard';
import { AuthService } from 'src/_services/auth.service';
import { JwtHelperService } from '@auth0/angular-jwt';
import { ChatroomComponent } from './chatroom/chatroom.component';

const routes: Routes = [
  { path: 'register', component: SignUpFormComponent }, 
  { path: 'login', component: SignInFormComponent},

  { path: 'home', component: HomeComponent,  canActivate: [AuthGuard]},
  { path: 'create-chatroom', component: CreateChatroomComponent, canActivate: [AuthGuard] },
  { path: 'chatroom', component: ChatroomComponent, canActivate: [AuthGuard],  },

  { path: '', redirectTo: '/home', pathMatch: 'full' },
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
  providers:[AuthService, JwtHelperService]
})
export class AppRoutingModule { }
