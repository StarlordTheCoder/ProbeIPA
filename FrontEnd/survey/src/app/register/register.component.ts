import { User } from './../model/user.model';
import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Const } from '../const';
import { Router } from '@angular/router';

@Component({
  selector: 'register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  public user: User = {Id: 0} as User;
  constructor (private http: HttpClient, private router: Router) {
    if (sessionStorage.getItem('username') !== null) {
      this.router.navigate(['/admin']);
    }
  }
  public register() {
    console.warn(this.user);
    //Register user on Server
    this.http.post(Const.baseUrl + 'register', this.user).subscribe(data => {
      sessionStorage.setItem('username', this.user.Username);
      sessionStorage.setItem('password', this.user.Password);
      this.router.navigate(['/admin']);
    }, err => {
      console.warn('error');
    });
  }
}
