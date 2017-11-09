import { Component, OnInit, OnDestroy } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Const } from '../const';
import { Survey, Choice } from '../model/survey.model';
import { Header } from '../header';
import { Router, ActivatedRoute } from '@angular/router';
import { NewAnswers, NewAnswer } from '../model/new-answers';

@Component({
  selector: 'survey',
  templateUrl: './survey-create.component.html',
  styleUrls: ['./survey-create.component.css']
})
export class SurveyCreateComponent {
  public question: string;
  public choices: Choice[] = [{Id: 0, Choice: ''}];
  private headers: Header;
  constructor (private http: HttpClient, private router: Router) {
    this.headers = new Header();
  }
  createSurvey() {
    const survey = { Id: 0, Question: this.question, Choices: this.choices.filter(c => c.Choice !== '')} as Survey;
    this.http.post<number>(Const.baseUrl + 'createSurvey', survey, { headers: this.headers.headers }).subscribe(data => {
      const newId = data;
      this.router.navigate(['/survey', newId]);
    }, err => {
      console.warn('error', err);
    });
  }
  checkList(touched: boolean) {
    if (!touched && this.choices.length <= 10) {
      this.choices.push({Id: 0, Choice: ''});
    }
  }
}

