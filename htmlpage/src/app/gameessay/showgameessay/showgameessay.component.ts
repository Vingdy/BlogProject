import { Component,OnInit,ViewEncapsulation } from '@angular/core';
import { Router,ActivatedRoute } from '@angular/router';

import { GameEssayStruct } from '../../data/gameessayStruct'

import { GameEssayService } from '../../service/gameessay.service'
import { SessionService } from '../../service/session.service'

import { ROUTES } from '../../config/route-api'

@Component({
  selector: 'app-showgameessay',
  templateUrl: './showgameessay.component.html',
  styleUrls: ['./showgameessay.component.css'],
  encapsulation: ViewEncapsulation.None,
})
export class ShowGameEssayComponent implements OnInit {
    Role:number

    IsEmpty:boolean

  CurrentPage:number
  searchstring:string
  TotalPage:string

  limit:string
  offset:string
  GameEssayArray:GameEssayStruct[]
  TagArray=new Array()
  TimeArray=new Array()

  constructor(
    private router:Router,
    private gameessayservice:GameEssayService,
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
    this.offset="0"
    this.searchstring=""
    this.gameessayservice.GetAllGameEssayInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.GameEssayArray=fb["data"]
        if(fb["code"]==1000)
        {
        if(this.GameEssayArray.length>0){
            for(let i=0;i<this.GameEssayArray.length;i++){
              this.GameEssayArray[i].time = this.GameEssayArray[i].time.replace('Z','+08:00')
            }
            this.TotalPage=fb["total"]
            for(let i=0;i<this.GameEssayArray.length;i++){
                this.GameEssayArray[i].content = this.GameEssayArray[i].content.replace(/<img [^>]*src=['"]([^'"]+)[^>]*>/gi, '[图片]')
                if(this.GameEssayArray[i].content.length >= 100){
                  this.GameEssayArray[i].content = this.GameEssayArray[i].content.substring(0,100) + '...';
                }
          }
          this.CreateFiling()
        }
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
  ToWriteGameEssay(){
    this.router.navigate([ROUTES.writegameessay.route])
  }
  CurrentPageOut(CurrentPageOut) {
    this.CurrentPage=CurrentPageOut
    this.ChangePage(this.CurrentPage)//返回被选择的当前页进行处理
  }
  ChangePage(choosepage){
    this.GameEssayArray=[]
    this.offset=String(Number(this.limit)*(choosepage-1))
    this.gameessayservice.GetAllGameEssayInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.GameEssayArray=fb["data"]
        if(this.GameEssayArray.length>0){
            for(let i=0;i<this.GameEssayArray.length;i++){
              this.GameEssayArray[i].time = this.GameEssayArray[i].time.replace('Z','+08:00')
            }
            this.TotalPage=fb["total"]
          }
          for(let i=0;i<this.GameEssayArray.length;i++){
            this.GameEssayArray[i].content = this.GameEssayArray[i].content.replace(/<img [^>]*src=['"]([^'"]+)[^>]*>/gi, '[图片]')
            if(this.GameEssayArray[i].content.length >= 100){
              this.GameEssayArray[i].content = this.GameEssayArray[i].content.substring(0,100) + '...';
            }
          }
      },
      err=>{
      }
    )
}
CreateFiling(){
    this.gameessayservice.GetGameEssayTag().subscribe(
      fb=>{
        this.TagArray=fb["data"]
        this.TagArray[this.TagArray.length]={name:'全部',number:this.TotalPage}
      },
      err=>{

      }
    )
    this.gameessayservice.GetGameEssayTime().subscribe(
      fb=>{
        this.TimeArray=fb["data"]
      },
      err=>{

      }
    )
  }
GetGameEssayAboutTime(Time){
    // console.log(Time)
    this.GameEssayArray=[]
    this.searchstring=Time
    this.gameessayservice.GetAllGameEssayInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.GameEssayArray=fb["data"]
        if(this.GameEssayArray.length>0){
          for(let i=0;i<this.GameEssayArray.length;i++){
            this.GameEssayArray[i].time = this.GameEssayArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
        }
        for(let i=0;i<this.GameEssayArray.length;i++){
          this.GameEssayArray[i].content = this.GameEssayArray[i].content.replace(/<img [^>]*src=['"]([^'"]+)[^>]*>/gi, '[图片]')
          if(this.GameEssayArray[i].content.length >= 100){
            this.GameEssayArray[i].content = this.GameEssayArray[i].content.substring(0,100) + '...';
          }
        }
      },
      err=>{
      }
    )
  }
  GetGameEssayAboutTag(Tag){
    // console.log(Tag)
    this.GameEssayArray=[]
    this.searchstring=Tag
    if(this.searchstring=="全部"){
      this.searchstring=""
    }
    this.gameessayservice.GetAllGameEssayInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.GameEssayArray=fb["data"]
        if(this.GameEssayArray.length>0){
          for(let i=0;i<this.GameEssayArray.length;i++){
            this.GameEssayArray[i].time = this.GameEssayArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
        }
        for(let i=0;i<this.GameEssayArray.length;i++){
          this.GameEssayArray[i].content = this.GameEssayArray[i].content.replace(/<img [^>]*src=['"]([^'"]+)[^>]*>/gi, '[图片]')
          if(this.GameEssayArray[i].content.length >= 100){
            this.GameEssayArray[i].content = this.GameEssayArray[i].content.substring(0,100) + '...';
          }
        }
      },
      err=>{
      }
    )
  }
}

