import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Observable, BehaviorSubject } from "rxjs";

@Injectable({
  providedIn: 'root'
})



export class AuthService {

  private onProduction: boolean = false;

  private isAuthenticatedSubject = new BehaviorSubject<boolean>(false);
  isAuthenticated$ = this.isAuthenticatedSubject.asObservable();

  private url: string = this.onProduction ? "/login" : "http://localhost:8080/login";
  private logoutUrl: string = this.onProduction ? "/logout" : "http://localhost:8080/logout";

  private httpOptions: any = {
    headers: new HttpHeaders({
      "Content-Type": "application/json"
    }),
    withCredentials: true
  }

  constructor(private httpClient: HttpClient) { }

  sendLoginData(data: any): Observable<any> {
    return this.httpClient.post(this.url, JSON.stringify(data), this.httpOptions)
  }

  register(data: any): Observable<any> {
    const registerUrl = this.onProduction ? "/register" : "http://localhost:8080/register";
    return this.httpClient.post(registerUrl, JSON.stringify(data), this.httpOptions);
  }

  logout(): Observable<any> {
    return this.httpClient.get(this.logoutUrl, this.httpOptions);
  }

  setAuthenticated(value: boolean) {
    this.isAuthenticatedSubject.next(value);
  }

  get isAuthenticated(): boolean {
    return this.isAuthenticatedSubject.value;
  }
}
