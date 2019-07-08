import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpClient } from '@angular/common/http'

import { SITE_HOST_URL,GET_ALL_SENTENCE,WRITE_SENTENCE,GET_ONE_SENTENCE,
    UPDATE_SENTENCE_ESSAY,DELETE_SENTENCE_ESSAY,GET_SENTENCE_TIME,GET_USER_DATA } from '../config/api';

import { SentenceStruct } from '../data/sentenceStruct'

@Injectable()
export class SentenceService {
    constructor(private http: Http,
        private httpclient:HttpClient
    ) { }
    GetAllSentenceInfo(limit:string,offset:string,searchstring:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ALL_SENTENCE+"?limit="+limit+"&offset="+offset+"&searchstring="+searchstring);
    }
    WriteNewSentence(sentenceinfo:SentenceStruct){
        return this.httpclient.post(SITE_HOST_URL+WRITE_SENTENCE,JSON.stringify(sentenceinfo));
    }
    GetOneSentence(sentenceid:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ONE_SENTENCE+"?sentenceid="+sentenceid);
    }
    GetSentenceTime(){
        return this.httpclient.get(SITE_HOST_URL+GET_SENTENCE_TIME)
    }
    UpdateOneSentence(sentenceinfo:SentenceStruct){
        return this.httpclient.post(SITE_HOST_URL+UPDATE_SENTENCE_ESSAY,JSON.stringify(sentenceinfo))
    }
    DeleteOneSentence(essayid:string){
        const data = {
            "essayid": essayid,
          };
        return this.httpclient.post(SITE_HOST_URL+DELETE_SENTENCE_ESSAY,JSON.stringify(data))
    }
    GetUserData(){
        return this.httpclient.get(SITE_HOST_URL+GET_USER_DATA)
    }
}