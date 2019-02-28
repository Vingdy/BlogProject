import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpHeaders,HttpClient,HttpResponse,HttpErrorResponse,HttpRequest } from '@angular/common/http'

// import { HttpService } from '../public-service/http-service/Http.Service.service';

import { SITE_HOST_URL,LOGIN,LOGOUT,UPLOAD_IMAGE } from '../config/api';

import { GameEssayStruct } from '../data/gameessayStruct'

@Injectable()
export class LoginService {
    constructor(private http: Http,
        private httpclient:HttpClient
        // public httpservice:HttpService
    ) { }
    AdminLogin(data){
        return this.httpclient.post(SITE_HOST_URL+LOGIN,JSON.stringify(data));
    }
    AdminLogout(){
        return this.httpclient.get(SITE_HOST_URL+LOGOUT);
    }
    imageHandler(data) {
        return this.httpclient.post(SITE_HOST_URL+UPLOAD_IMAGE, data)
    }
}