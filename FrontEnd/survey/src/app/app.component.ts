import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'app';
  public get isLogedin() {
    const savedUsername = sessionStorage.getItem('username');
    return  savedUsername !== null;
  }
}
