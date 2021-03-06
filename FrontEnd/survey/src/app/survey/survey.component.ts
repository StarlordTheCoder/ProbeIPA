import { Component, OnInit, OnDestroy } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Const } from '../const';
import { Survey } from '../model/survey.model';
import { Header } from '../header';
import { Router, ActivatedRoute } from '@angular/router';
import { NewAnswers, NewAnswer } from '../model/new-answers';

@Component({
  selector: 'survey',
  templateUrl: './survey.component.html',
  styleUrls: ['./survey.component.css']
})
export class SurveyComponent implements OnInit, OnDestroy {
  private id: number;
  private sub: any;
  public surveys: Survey = {} as Survey;
  private newAnswers: NewAnswer[] = [];
  constructor (private http: HttpClient, private route: ActivatedRoute, private router: Router) {}

  //Send Answer to Server
  public answerSurvey() {
    const answers: NewAnswers = {NewAnswers: this.newAnswers};
    this.http.post(Const.baseUrl + 'answerSurvey', answers).subscribe(data => {
      this.router.navigate(['/login']);
      console.warn(answers);
    }, err => {
      console.warn('error', err);
    });
  }
  //If a Choice gets selected or deselcted it gets stored or deleted
  public checkBox(id: number) {
    const answer: NewAnswer = {IdChoice: id};
    const index = this.newAnswers.indexOf(answer);
    if (index === -1) {
      this.newAnswers.push(answer);
    } else {
      this.newAnswers.splice(index, 1);
    }
  }

  public ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
      this.id = +params['id'];

      this.http.get<Survey>(Const.baseUrl + 'getSurvey/' + this.id).subscribe(data => {
        this.surveys = data;
      }, err => {
        console.warn('error', err);
      });
    });
  }

  public ngOnDestroy() {
  this.sub.unsubscribe();
  }
}

