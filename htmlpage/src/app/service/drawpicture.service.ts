import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpClient } from '@angular/common/http'

import { SITE_HOST_URL,GET_ALL_DRAWPICTURE,WRITE_DRAWPICTURE,GET_ONE_DRAWPICTURE,UPLOAD_IMAGE } from '../config/api';

import { DrawpictureStruct } from '../data/drawpictureStruct'

@Injectable()
export class DrawpictureService {
    constructor(private http: Http,
        private httpclient:HttpClient
        // public httpservice:HttpService
    ) { }
    GetAllDrawpictureInfo(limit:string,offset:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ALL_DRAWPICTURE+"?limit="+limit+"&offset="+offset);

    }
    GetOneDrawpictureInfo(id:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ONE_DRAWPICTURE+"?id="+id);
    }
    WriteNewDrawpicture(drawpictureinfo:DrawpictureStruct){
        return this.httpclient.post(SITE_HOST_URL+WRITE_DRAWPICTURE,JSON.stringify(drawpictureinfo));
    }
    imageHandler(data) {
        return this.httpclient.post(SITE_HOST_URL+UPLOAD_IMAGE, data)
    }
}