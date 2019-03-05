import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpClient } from '@angular/common/http'

import { SITE_HOST_URL,GET_ALL_GAME_ESSAY,WRITE_GAME_ESSAY,
    GET_ONE_GAME_ESSAY,GET_GAME_ESSAY_TAG,GET_GAME_ESSAY_TIME,
    UPDATE_ONE_GAME_ESSAY,DELETE_ONE_GAME_ESSAY,GET_USER_DATA,UPLOAD_IMAGE } from '../config/api';

import { GameEssayStruct } from '../data/gameessayStruct'

@Injectable()
export class GameEssayService {
    constructor(private http: Http,
        private httpclient:HttpClient
    ) { }
    GetAllGameEssayInfo(limit:string,offset:string,searchstring:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ALL_GAME_ESSAY+"?limit="+limit+"&offset="+offset+"&searchstring="+searchstring);

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
    GetGameEssayTag(){
        return this.httpclient.get(SITE_HOST_URL+GET_GAME_ESSAY_TAG)
    }
    GetGameEssayTime(){
        return this.httpclient.get(SITE_HOST_URL+GET_GAME_ESSAY_TIME)
    }
    UpdateOneGameEssay(gameessayinfo:GameEssayStruct){
        return this.httpclient.post(SITE_HOST_URL+UPDATE_ONE_GAME_ESSAY,JSON.stringify(gameessayinfo))
    }
    DeleteOneGameEssay(essayid:string){
        return this.httpclient.get(SITE_HOST_URL+DELETE_ONE_GAME_ESSAY+"?essayid="+essayid)
    }
    GetUserData(){
        return this.httpclient.get(SITE_HOST_URL+GET_USER_DATA)
    }
}