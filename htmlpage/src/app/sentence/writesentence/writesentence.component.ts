import { Component,OnInit } from '@angular/core';
import { Router,ActivatedRoute } from '@angular/router';
import { DatePipe } from '@angular/common';

import { SentenceStruct } from '../../data/sentenceStruct'

import { SentenceService } from '../../service/sentence.service'
import { SessionService } from '../../service/session.service'

import { ToastrService } from 'ngx-toastr'

import { ROUTES } from '../../config/route-api'

@Component({
  selector: 'app-writesentence',
  templateUrl: './writesentence.component.html',
  styleUrls: ['./writesentence.component.css'],
  providers:[DatePipe]
})
export class WriteSentenceComponent implements OnInit {
  essayid:string
  Role:number
  editorContent:string
  NewSentence:SentenceStruct
    constructor(
      private router:Router,
      private sentenceservice:SentenceService,
      private datePipe: DatePipe,
      private toastrservice:ToastrService,
      private sessionservice:SessionService,
      private activatedroute:ActivatedRoute
    ) { }
    ngOnInit(){
      this.NewSentence=new SentenceStruct
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
        this.activatedroute.queryParams.subscribe(params => {
          this.essayid = params['essayid']; 
        });
        if(this.essayid){
          this.sentenceservice.GetOneSentence(this.essayid).subscribe(
            fb=>{
              this.NewSentence.content=fb["data"][0]["content"]
              this.NewSentence.time=fb["data"][0]["time"]
            },
            err=>{
  
            }
          )
      }
    }
    quillconfig={
      toolbar: [
      ]
    };
    NewSentencePush(sentenceinfo:SentenceStruct){
      if(!sentenceinfo.content){
        this.toastrservice.error('内容不能为空')
        return
      }(sentenceinfo.content)
      sentenceinfo.time=Date.now().toString()
      sentenceinfo.time=this.datePipe.transform(sentenceinfo.time, 'yyyy-MM-dd HH:mm:ss')
      this.sentenceservice.WriteNewSentence(sentenceinfo).subscribe(
        fb=>{
          this.toastrservice.success('上传成功')
          this.router.navigate([ROUTES.showsentence.route])
        },
        err=>{
          this.toastrservice.error('上传失败')
        }
      )
    }
    ToBackEssay(){
      this.router.navigate([ROUTES.showsentence.route])
    }
    UpdateSentence(sentenceinfo:SentenceStruct){
      sentenceinfo.id=this.essayid
      if(!sentenceinfo.content){
        this.toastrservice.error("未上传图片")
      }
      this.sentenceservice.UpdateOneSentence(sentenceinfo).subscribe(
        fb=>{
          this.toastrservice.success('修改成功')
          this.router.navigate([ROUTES.showsentence.route])
        },
        err=>{
          this.toastrservice.error('修改失败')
        }
      )
    }
}
