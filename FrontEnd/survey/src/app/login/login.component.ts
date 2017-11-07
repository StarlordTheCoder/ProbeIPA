import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { User } from '../model/user.model';
import { Const } from '../const';

@Component({
  selector: 'login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  public user: User = {Id: 0} as User;
  constructor (private http: HttpClient) {}
  public login() {
    console.warn(this.user);
    this.http.post(Const.baseUrl + 'login', this.user).subscribe(data => {
      console.warn('succes');
    }, err => {
      console.warn('error');
    });
  }
}
