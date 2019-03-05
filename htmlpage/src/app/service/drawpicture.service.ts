import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpClient } from '@angular/common/http'

import { SITE_HOST_URL,GET_ALL_DRAWPICTURE,WRITE_DRAWPICTURE,
    GET_ONE_DRAWPICTURE,GET_DRAWPICTURE_TAG,GET_DRAWPICTURE_TIME,
    UPDATE_ONE_DRAWPICTURE,DELETE_ONE_DRAWPICTURE,GET_USER_DATA,UPLOAD_IMAGE } from '../config/api';

import { DrawpictureStruct } from '../data/drawpictureStruct'

@Injectable()
export class DrawpictureService {
    constructor(private http: Http,
        private httpclient:HttpClient
        // public httpservice:HttpService
    ) { }
    GetAllDrawpictureInfo(limit:string,offset:string,searchstring:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ALL_DRAWPICTURE+"?limit="+limit+"&offset="+offset+"&searchstring="+searchstring);

    }
    GetOneDrawpictureInfo(pictureid:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ONE_DRAWPICTURE+"?pictureid="+pictureid);
    }
    WriteNewDrawpicture(drawpictureinfo:DrawpictureStruct){
        return this.httpclient.post(SITE_HOST_URL+WRITE_DRAWPICTURE,JSON.stringify(drawpictureinfo));
    }
    imageHandler(data) {
        return this.httpclient.post(SITE_HOST_URL+UPLOAD_IMAGE, data)
    }
    GetDrawpictureTag(){
        return this.httpclient.get(SITE_HOST_URL+GET_DRAWPICTURE_TAG)
    }
    GetDrawpictureTime(){
        return this.httpclient.get(SITE_HOST_URL+GET_DRAWPICTURE_TIME)
    }
    UpdateOneDrawpicture(drawpictureinfo:DrawpictureStruct){
        return this.httpclient.post(SITE_HOST_URL+UPDATE_ONE_DRAWPICTURE,JSON.stringify(drawpictureinfo))
    }
    DeleteOneDrawpicture(essayid:string){
        return this.httpclient.get(SITE_HOST_URL+DELETE_ONE_DRAWPICTURE+"?essayid="+essayid)
    }
    GetUserData(){
        return this.httpclient.get(SITE_HOST_URL+GET_USER_DATA)
    }
}