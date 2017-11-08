import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { User } from '../model/user.model';
import { Const } from '../const';
import { Router } from '@angular/router';

@Component({
  selector: 'login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  public user: User = {Id: 0} as User;
  constructor (private http: HttpClient, private router: Router) {}
  public login() {
    console.warn(this.user);
    this.http.post(Const.baseUrl + 'login', this.user).subscribe(data => {
      sessionStorage.setItem('username', this.user.Username);
      sessionStorage.setItem('password', this.user.Password);
      this.router.navigate(['/admin']);
    }, err => {
      console.warn('error');
    });
  }
}
