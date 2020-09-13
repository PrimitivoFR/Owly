import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CreateChatroomComponent } from './create-chatroom/create-chatroom.component';
import { SignUpFormComponent } from './sign-up-form/sign-up-form.component';
import { SignInFormComponent } from './sign-in-form/sign-in-form.component';
import { HomeComponent } from './home/home.component';

const routes: Routes = [
  { path: 'register', component: SignUpFormComponent }, 
  { path: "login", component: SignInFormComponent },

  { path: "home", component: HomeComponent },
  { path: "create-chatroom", component: CreateChatroomComponent },

  { path: '', redirectTo: '/home', pathMatch: 'full' },
  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
