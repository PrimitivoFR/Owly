import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { CreateChatroomComponent } from './create-chatroom/create-chatroom.component';
import { SignUpFormComponent } from './sign-up-form/sign-up-form.component';

const routes: Routes = [
  { path: 'register', component: SignUpFormComponent },
  { path: '', redirectTo: '/register', pathMatch: 'full' },
  {path: "create-chatroom", component: CreateChatroomComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
