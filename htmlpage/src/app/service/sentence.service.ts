import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpHeaders,HttpClient,HttpResponse,HttpErrorResponse,HttpRequest } from '@angular/common/http'

// import { HttpService } from '../public-service/http-service/Http.Service.service';

import { SITE_HOST_URL,GET_ALL_SENTENCE,WRITE_SENTENCE } from '../config/api';

import { SentenceStruct } from '../data/sentenceStruct'

@Injectable()
export class SentenceService {
    constructor(private http: Http,
        private httpclient:HttpClient
    ) { }
    GetAllSentenceInfo(limit:string,offset:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ALL_SENTENCE+"?limit="+limit+"&offset="+offset);
    }
    WriteNewSentence(sentenceinfo:SentenceStruct){
        return this.httpclient.post(SITE_HOST_URL+WRITE_SENTENCE,JSON.stringify(sentenceinfo));
    }
}