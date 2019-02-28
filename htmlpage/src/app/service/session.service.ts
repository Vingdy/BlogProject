import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpHeaders,HttpClient,HttpResponse,HttpErrorResponse,HttpRequest } from '@angular/common/http'

// import { HttpService } from '../public-service/http-service/Http.Service.service';

import { SITE_HOST_URL,GET_ROLE } from '../config/api';

@Injectable()
export class SessionService {
    constructor(private http: Http,
        private httpclient:HttpClient
    ) { }
    GetRole(){
        return this.httpclient.get(SITE_HOST_URL+GET_ROLE)
    }
}