import { User } from './../model/user.model';
import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Const } from '../const';

@Component({
  selector: 'register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  public user: User = {Id: 0} as User;
  constructor (private http: HttpClient) {}
  public register() {
    console.warn(this.user);
    this.http.post(Const.baseUrl + 'register', this.user).subscribe(data => {
      console.warn('succes');
    }, err => {
      console.warn('error');
    });
  }
}
