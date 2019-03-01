import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpClient } from '@angular/common/http'

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