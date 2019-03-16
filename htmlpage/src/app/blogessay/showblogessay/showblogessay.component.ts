import { Component,OnInit,ViewEncapsulation } from '@angular/core';
import { Router,ActivatedRoute } from '@angular/router';

import { BlogEssayStruct } from '../../data/blogessayStruct'

import { BlogEssayService } from '../../service/blogessay.service'
import { SessionService } from '../../service/session.service'

import { ROUTES } from '../../config/route-api'

@Component({
  selector: 'app-showblogessay',
  templateUrl: './showblogessay.component.html',
  styleUrls: ['./showblogessay.component.css'],
  encapsulation: ViewEncapsulation.None,
})
export class ShowBlogEssayComponent implements OnInit {
  IsEmpty:boolean
  Role:number
  CurrentPage:number
  searchstring:string
  TotalPage:string

  limit:string
  offset:string
  BlogEssayArray:BlogEssayStruct[]
  TagArray=new Array()
  TimeArray=new Array()

  constructor(
    private router:Router,
    private blogessayservice:BlogEssayService,
    private sessionservice:SessionService,
    private activatedRoute:ActivatedRoute,
  ) { }
  ngOnInit(){
    this.IsEmpty=false;
      this.CurrentPage = this.activatedRoute.snapshot.queryParams["CurrentPage"];
      if (!this.CurrentPage){
        this.CurrentPage=1
      }
    this.sessionservice.GetRole().subscribe(
      fb=>{
          if(fb["code"]!=1000){
            this.Role=0
          }else{
              this.Role=fb["data"]
          }
      },
      err=>{
          this.Role=0
      })
    this.limit="5"
    this.offset=((this.CurrentPage-1)*5).toString()
    this.searchstring=""
    this.blogessayservice.GetAllBlogEssayInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        // console.log(fb)
        this.BlogEssayArray=fb["data"]
        if(fb["code"]==1000){
        if(this.BlogEssayArray.length>0){
          for(let i=0;i<this.BlogEssayArray.length;i++){
            this.BlogEssayArray[i].time = this.BlogEssayArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
          for(let i=0;i<this.BlogEssayArray.length;i++){
            this.BlogEssayArray[i].content = this.BlogEssayArray[i].content.replace(/<img [^>]*src=['"]([^'"]+)[^>]*>/gi, '[图片]')
            if(this.BlogEssayArray[i].content.length >= 100){
              this.BlogEssayArray[i].content = this.BlogEssayArray[i].content.substring(0,100) + '...';
            }
          }
        }
        this.CreateFiling()
      }
      else{
          this.IsEmpty=true;
      }
      },
      err=>{
      }
    )
    this.limit="5"
    
  }
  ToWriteBlogEssay(){
    this.router.navigate([ROUTES.writeblogessay.route])
  }
  // }
  CurrentPageOut(CurrentPageOut) {
    this.CurrentPage=CurrentPageOut
    this.ChangePage(this.CurrentPage)//返回被选择的当前页进行处理
  }
  ChangePage(choosepage){
    this.BlogEssayArray=[]
    this.offset=String(Number(this.limit)*(choosepage-1))
    this.blogessayservice.GetAllBlogEssayInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.BlogEssayArray=fb["data"]
        if(this.BlogEssayArray.length>0){
          for(let i=0;i<this.BlogEssayArray.length;i++){
            this.BlogEssayArray[i].time = this.BlogEssayArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
        }
        for(let i=0;i<this.BlogEssayArray.length;i++){
          this.BlogEssayArray[i].content = this.BlogEssayArray[i].content.replace(/<img [^>]*src=['"]([^'"]+)[^>]*>/gi, '[图片]')
          if(this.BlogEssayArray[i].content.length >= 100){
            this.BlogEssayArray[i].content = this.BlogEssayArray[i].content.substring(0,100) + '...';
          }
        }
      },
      err=>{
      }
    )
  }
  CreateFiling(){
    this.blogessayservice.GetBlogEssayTag().subscribe(
      fb=>{
        this.TagArray=fb["data"]
        this.TagArray[this.TagArray.length]={name:'全部',number:this.TotalPage}
      },
      err=>{

      }
    )
    this.blogessayservice.GetBlogEssayTime().subscribe(
      fb=>{
        this.TimeArray=fb["data"]
      },
      err=>{

      }
    )
  }
  GetBlogEssayAboutTime(Time){
    // console.log(Time)
    this.BlogEssayArray=[]
    this.searchstring=Time
    this.blogessayservice.GetAllBlogEssayInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.BlogEssayArray=fb["data"]
        if(this.BlogEssayArray.length>0){
          for(let i=0;i<this.BlogEssayArray.length;i++){
            this.BlogEssayArray[i].time = this.BlogEssayArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
        }
        for(let i=0;i<this.BlogEssayArray.length;i++){
          this.BlogEssayArray[i].content = this.BlogEssayArray[i].content.replace(/<img [^>]*src=['"]([^'"]+)[^>]*>/gi, '[图片]')
          if(this.BlogEssayArray[i].content.length >= 100){
            this.BlogEssayArray[i].content = this.BlogEssayArray[i].content.substring(0,100) + '...';
          }
        }
      },
      err=>{
      }
    )
  }
  GetBlogEssayAboutTag(Tag){
    // console.log(Tag)
    this.BlogEssayArray=[]
    this.searchstring=Tag
    if(this.searchstring=="全部"){
      this.searchstring=""
    }
    this.blogessayservice.GetAllBlogEssayInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.BlogEssayArray=fb["data"]
        if(this.BlogEssayArray.length>0){
          for(let i=0;i<this.BlogEssayArray.length;i++){
            this.BlogEssayArray[i].time = this.BlogEssayArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
        }
        for(let i=0;i<this.BlogEssayArray.length;i++){
          this.BlogEssayArray[i].content = this.BlogEssayArray[i].content.replace(/<img [^>]*src=['"]([^'"]+)[^>]*>/gi, '[图片]')
          if(this.BlogEssayArray[i].content.length >= 100){
            this.BlogEssayArray[i].content = this.BlogEssayArray[i].content.substring(0,100) + '...';
          }
        }
      },
      err=>{
      }
    )
  }
}

