import { Component,OnInit,Injectable,Pipe } from '@angular/core';
import { Router,ActivatedRoute } from '@angular/router';
import { DomSanitizer } from '@angular/platform-browser';

import { BlogEssayStruct } from '../../data/blogessayStruct'

import { BlogEssayService } from '../../service/blogessay.service'
import { SessionService } from '../../service/session.service'

import { ROUTES } from '../../config/route-api'

@Component({
  selector: 'app-oneblogessay',
  templateUrl: './oneblogessay.component.html',
  styleUrls: ['./oneblogessay.component.css'],
})
export class OneBlogEssayComponent implements OnInit {
    Role:number
    private essayid:string;
    limit:string
    offset:string
     BlogEssayInfo:BlogEssayStruct

  constructor(
    private router:Router,
    private blogessayservice:BlogEssayService,
    private activatedRoute:ActivatedRoute,
    private sessionservice:SessionService,
    private domsanitizer: DomSanitizer,
  ) { }
  ngOnInit(){
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
    this.BlogEssayInfo=new BlogEssayStruct
    this.essayid = this.activatedRoute.snapshot.queryParams["essayid"];
    this.blogessayservice.GetOneBlogEssayInfo(this.essayid).subscribe(
        fb=>{
            this.BlogEssayInfo=fb["data"][0]
            // this.BlogEssayInfo.content=this.domsanitizer.bypassSecurityTrustHtml(this.BlogEssayInfo.content)
            this.BlogEssayInfo.time=this.BlogEssayInfo.time.replace('Z','+08:00')
        },
        err=>{
        }
    )
  }
  ToBackList(){
    let CurrentPage=Number((Number(this.essayid)/5).toFixed(0))
    let Para=(Number(this.essayid)/5)
    if(CurrentPage-Para<0){
        CurrentPage+=1
      return CurrentPage
  }
}
}