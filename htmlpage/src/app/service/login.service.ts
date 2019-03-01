import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpClient } from '@angular/common/http'

import { SITE_HOST_URL,LOGIN,LOGOUT,UPLOAD_IMAGE } from '../config/api';

@Injectable()
export class LoginService {
    constructor(private http: Http,
        private httpclient:HttpClient
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