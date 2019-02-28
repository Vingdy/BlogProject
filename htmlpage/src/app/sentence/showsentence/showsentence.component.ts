import { Component,OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { SentenceStruct } from '../../data/sentenceStruct'

import { SentenceService } from '../../service/sentence.service'
import { SessionService } from '../../service/session.service'

import { ROUTES } from '../../config/route-api'

@Component({
  selector: 'app-showsentence',
  templateUrl: './showsentence.component.html',
  styleUrls: ['./showsentence.component.css'],
})
export class ShowSentenceComponent implements OnInit {
    IsEmpty:boolean
    Role:number

  CurrentPage:number
  searchstring:string
  TotalPage:string

  limit:string
  offset:string
  SentenceArray:SentenceStruct[]

  constructor(
    private router:Router,
    private sentenceservice:SentenceService,
    private sessionservice:SessionService,
  ) { }
  ngOnInit(){
    this.IsEmpty=false;
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
    this.CurrentPage=1
    this.searchstring=""
    this.sentenceservice.GetAllSentenceInfo(this.limit,this.offset).subscribe(
      fb=>{
        console.log(fb)
        this.SentenceArray=fb["data"]
        // console.log(array)
        if(fb["code"]==1000){
        if(this.SentenceArray.length>0){
          for(let i=0;i<this.SentenceArray.length;i++){
            this.SentenceArray[i].time = this.SentenceArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
        }
    }       
     else{
        this.IsEmpty=true;
    }
        console.log(this.SentenceArray)
      },
      err=>{
        console.log(err)
      }
    )
    this.limit="5"
  }
  ToWriteSentence(){
    this.router.navigate([ROUTES.writesentence.route])
  }
  CurrentPageOut(CurrentPageOut) {
    this.CurrentPage=CurrentPageOut
    this.ChangePage(this.CurrentPage)//返回被选择的当前页进行处理
  }
  ChangePage(choosepage){
    this.SentenceArray=[]
    this.offset=String(Number(this.limit)*(choosepage-1))
    this.sentenceservice.GetAllSentenceInfo(this.limit,this.offset).subscribe(
      fb=>{
        console.log(fb)
        this.SentenceArray=fb["data"]
        if(this.SentenceArray.length>0){
          for(let i=0;i<this.SentenceArray.length;i++){
            this.SentenceArray[i].time = this.SentenceArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
        }
        console.log(this.SentenceArray)
      },
      err=>{
        console.log(err)
      }
    )
}
}

