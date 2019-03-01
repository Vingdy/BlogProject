import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpClient } from '@angular/common/http'

import { SITE_HOST_URL,GET_ALL_GAME_ESSAY,WRITE_GAME_ESSAY,GET_ONE_GAME_ESSAY,UPLOAD_IMAGE } from '../config/api';

import { GameEssayStruct } from '../data/gameessayStruct'

@Injectable()
export class GameEssayService {
    constructor(private http: Http,
        private httpclient:HttpClient
    ) { }
    GetAllGameEssayInfo(limit:string,offset:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ALL_GAME_ESSAY+"?limit="+limit+"&offset="+offset);

    }
    GetOneGameEssayInfo(essayid:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ONE_GAME_ESSAY+"?essayid="+essayid);
    }
    WriteNewGameEssay(gameessayinfo:GameEssayStruct){
        return this.httpclient.post(SITE_HOST_URL+WRITE_GAME_ESSAY,JSON.stringify(gameessayinfo));
    }
    imageHandler(data) {
        return this.httpclient.post(SITE_HOST_URL+UPLOAD_IMAGE, data)
    }
}