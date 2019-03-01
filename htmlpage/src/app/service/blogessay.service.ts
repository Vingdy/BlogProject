import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpClient } from '@angular/common/http'

import { SITE_HOST_URL,GET_ALL_BLOG_ESSAY,WRITE_BLOG_ESSAY,GET_ONE_BLOG_ESSAY,UPLOAD_IMAGE } from '../config/api';

import { BlogEssayStruct } from '../data/blogessayStruct'

@Injectable()
export class BlogEssayService {
    constructor(private http: Http,
        private httpclient:HttpClient
        // public httpservice:HttpService
    ) { }
    GetAllBlogEssayInfo(limit:string,offset:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ALL_BLOG_ESSAY+"?limit="+limit+"&offset="+offset);
    }
    GetOneBlogEssayInfo(essayid:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ONE_BLOG_ESSAY+"?essayid="+essayid);
    }
    WriteNewBlogEssay(blogessayinfo:BlogEssayStruct){
        return this.httpclient.post(SITE_HOST_URL+WRITE_BLOG_ESSAY,JSON.stringify(blogessayinfo));
    }
    imageHandler(data) {
        return this.httpclient.post(SITE_HOST_URL+UPLOAD_IMAGE, data)
    }
}