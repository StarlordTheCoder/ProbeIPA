import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Const } from '../const';
import { Survey } from '../model/survey.model';
import { Header } from '../header';

@Component({
  selector: 'main-admin',
  templateUrl: './main-admin.component.html',
  styleUrls: ['./main-admin.component.css']
})
export class MainAdminComponent {
  public surveys: Survey[] = [];
  private headers: Header;
  constructor (private http: HttpClient) {
    this.headers = new Header();
    this.http.get<Survey[]>(Const.baseUrl + 'getSurveyByUser', { headers: this.headers.headers }).subscribe(data => {
      this.surveys = data;
    }, err => {
      console.warn('error', err);
    });
  }
}
