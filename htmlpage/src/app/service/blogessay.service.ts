import { Injectable } from '@angular/core';
import { Http } from '@angular/http'
import { HttpClient } from '@angular/common/http'

import { SITE_HOST_URL,GET_ALL_BLOG_ESSAY,WRITE_BLOG_ESSAY,
    GET_ONE_BLOG_ESSAY,GET_BLOG_ESSAY_TAG,GET_BLOG_ESSAY_TIME,
    UPDATE_ONE_BLOG_ESSAY,DELETE_ONE_BLOG_ESSAY,GET_USER_DATA,UPLOAD_IMAGE } from '../config/api';

import { BlogEssayStruct } from '../data/blogessayStruct'

@Injectable()
export class BlogEssayService {
    constructor(private http: Http,
        private httpclient:HttpClient
        // public httpservice:HttpService
    ) { }
    GetAllBlogEssayInfo(limit:string,offset:string,searchstring:string){
        return this.httpclient.get(SITE_HOST_URL+GET_ALL_BLOG_ESSAY+"?limit="+limit+"&offset="+offset+"&searchstring="+searchstring);
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
    GetBlogEssayTag(){
        return this.httpclient.get(SITE_HOST_URL+GET_BLOG_ESSAY_TAG)
    }
    GetBlogEssayTime(){
        return this.httpclient.get(SITE_HOST_URL+GET_BLOG_ESSAY_TIME)
    }
    UpdateOneBlogEssay(blogessayinfo:BlogEssayStruct){
        return this.httpclient.post(SITE_HOST_URL+UPDATE_ONE_BLOG_ESSAY,JSON.stringify(blogessayinfo))
    }
    DeleteOneBlogEssay(essayid:string){
        return this.httpclient.get(SITE_HOST_URL+DELETE_ONE_BLOG_ESSAY+"?essayid="+essayid)
    }
    GetUserData(){
        return this.httpclient.get(SITE_HOST_URL+GET_USER_DATA)
    }
}