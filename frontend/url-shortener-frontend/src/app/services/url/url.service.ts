import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
    providedIn: 'root'
})
export class UrlService {

    private onProduction: boolean = false;
    private baseUrl: string = this.onProduction ? "" : "http://localhost:8080";

    private httpOptions: any = {
        headers: new HttpHeaders({
            "Content-Type": "application/json"
        }),
        withCredentials: true
    }

    constructor(private httpClient: HttpClient) { }

    addUrl(fullUrl: string): Observable<any> {
        const body = { fullUrl: fullUrl };
        return this.httpClient.post(`${this.baseUrl}/add-url`, JSON.stringify(body), this.httpOptions);
    }

    getLinks(): Observable<any> {
        return this.httpClient.get(`${this.baseUrl}/links`, this.httpOptions);
    }

    deleteLink(urlId: number): Observable<any> {
        return this.httpClient.delete(`${this.baseUrl}/links/${urlId}`, this.httpOptions);
    }
}
