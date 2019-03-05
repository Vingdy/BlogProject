import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpClient } from '@angular/common/http'

import { UserStruct } from '../data/UserStruct'

import { SITE_HOST_URL,LOGIN,LOGOUT,UPLOAD_IMAGE,GET_USER_DATA,UPDATE_USER_DATA } from '../config/api';

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
    GetUserData(){
        return this.httpclient.get(SITE_HOST_URL+GET_USER_DATA)
    }
    UpdateUserData(userinfo:UserStruct){
        console.log(userinfo)
        return this.httpclient.post(SITE_HOST_URL+UPDATE_USER_DATA,JSON.stringify(userinfo))
    }
}