import { Injectable } from '@angular/core';
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";

@Injectable({
  providedIn : 'root'
})



export class AuthService {

  private onProduction :boolean = false;

  private isAuthenticated :boolean = false
  private url :string = this.onProduction ? "/login" : "http://localhost:8080/login"
  private httpOptions :any = {
    header : new HttpHeaders({
      "Content-Type" : "application/json"
    }),
    withCredentials : true
  }

  constructor(private httpClient : HttpClient) { }
  sendLoginData(data :any) : Observable<any>{
    return this.httpClient.post(this.url, JSON.stringify(data), this.httpOptions)
  }

  setAuthenticated(value : boolean) {
    this.isAuthenticated = value;
    alert(this.isAuthenticated)
  }
}
