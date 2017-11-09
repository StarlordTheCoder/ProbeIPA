import { SurveyCreateComponent } from './survey-create/survey-create.component';
import { LoginComponent } from './login/login.component';
import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { RegisterComponent } from './register/register.component';
import { MainAdminComponent } from './main-admin/main-admin.component';
import { SurveyComponent } from './survey/survey.component';
import { ResultComponent } from './results/result.component';
import { AuthGuardService } from './guard';

const routes: Routes = [
  { path: 'login', component: LoginComponent},
  { path: 'register', component: RegisterComponent},
  { path: 'admin', component: MainAdminComponent, canActivate: [AuthGuardService]},
  { path: 'survey/:id', component: SurveyComponent},
  { path: 'createSurvey', component: SurveyCreateComponent, canActivate: [AuthGuardService]},
  { path: 'result/:id', component: ResultComponent, canActivate: [AuthGuardService]},
  { path: '', component: LoginComponent}];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
