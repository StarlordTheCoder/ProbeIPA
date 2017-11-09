import { GivenAnswers } from './../model/new-answers';
import { Component, OnInit, OnDestroy } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Const } from '../const';
import { Survey, Choice } from '../model/survey.model';
import { Header } from '../header';
import { Router, ActivatedRoute } from '@angular/router';
import { NewAnswers, NewAnswer } from '../model/new-answers';

@Component({
  selector: 'result',
  templateUrl: './result.component.html',
  styleUrls: ['./result.component.css']
})
export class ResultComponent implements OnInit, OnDestroy {
  private id: number;
  private sub: any;
  private headers: Header;
  public chartData:
    {
      name: string
      value: number
    }[];
  constructor (private http: HttpClient, private route: ActivatedRoute) {
    this.headers = new Header();
  }
  public ngOnInit() {
    this.sub = this.route.params.subscribe(params => {
      this.id = +params['id'];

      this.http.get<GivenAnswers>(Const.baseUrl + 'getAnswers/' + this.id, { headers: this.headers.headers } ).subscribe(data => {
         this.chartData = data.GivenAnswers.map((ga) =>  ({ name: ga.Choice, value: ga.Amount }));
      }, err => {
        console.warn('error', err);
      });
    });
  }

  public ngOnDestroy() {
  this.sub.unsubscribe();
  }
}

