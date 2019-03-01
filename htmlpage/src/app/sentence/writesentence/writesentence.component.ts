import { Component,OnInit } from '@angular/core';
import { Router } from '@angular/router';
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
  Role:number
  editorContent:string
  NewSentence:SentenceStruct
    constructor(
      private router:Router,
      private sentenceservice:SentenceService,
      private datePipe: DatePipe,
      private toastrservice:ToastrService,
      private sessionservice:SessionService,
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
    }
    quillconfig={
      toolbar: [
      ]
    };
    NewSentencePush(sentenceinfo:SentenceStruct){
      // sentenceinfo.content.replace("\n","<br>")
      if(!sentenceinfo.content){
        this.toastrservice.error('内容不能为空')
        return
      }(sentenceinfo.content)
      sentenceinfo.time=Date.now().toString()
      sentenceinfo.time=this.datePipe.transform(sentenceinfo.time, 'yyyy-MM-dd HH:mm:ss')
      this.sentenceservice.WriteNewSentence(sentenceinfo).subscribe(
        fb=>{
          this.toastrservice.success('上传成功')
        },
        err=>{
          this.toastrservice.error('上传失败')
        }
      )
    }
    ToBackEssay(){
      this.router.navigate([ROUTES.showsentence.route])
    }
}
