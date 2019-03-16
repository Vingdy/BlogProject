import { Component,OnInit,ViewEncapsulation } from '@angular/core';
import { Router } from '@angular/router';

import { SentenceStruct } from '../../data/sentenceStruct'

import { SentenceService } from '../../service/sentence.service'
import { SessionService } from '../../service/session.service'

import { ToastrService } from 'ngx-toastr'
import { BsModalService } from 'ngx-bootstrap/modal';
import { BsModalRef } from 'ngx-bootstrap/modal/bs-modal-ref.service';

import { ModalComponent } from '../../modal/modal-pop.component'

import { ROUTES } from '../../config/route-api'

@Component({
  selector: 'app-showsentence',
  templateUrl: './showsentence.component.html',
  styleUrls: ['./showsentence.component.css'],
  encapsulation: ViewEncapsulation.None,
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
  TimeArray=new Array()

  constructor(
    private router:Router,
    private sentenceservice:SentenceService,
    private sessionservice:SessionService,
    private modalService: BsModalService,
    private toastrservice:ToastrService,
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
    this.sentenceservice.GetAllSentenceInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.SentenceArray=fb["data"]
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
      },
      err=>{
      }
    )
    this.limit="5"
    this.CreateFiling()
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
    this.sentenceservice.GetAllSentenceInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.SentenceArray=fb["data"]
        if(this.SentenceArray.length>0){
          for(let i=0;i<this.SentenceArray.length;i++){
            this.SentenceArray[i].time = this.SentenceArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
        }
      },
      err=>{
      }
    )
}
CreateFiling(){
    this.sentenceservice.GetSentenceTime().subscribe(
      fb=>{
        this.TimeArray=fb["data"]
      },
      err=>{

      }
    )
  }
GetSentenceAboutTime(Time){
    // console.log(Time)
    this.SentenceArray=[]
    this.searchstring=Time
    this.sentenceservice.GetAllSentenceInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.SentenceArray=fb["data"]
        if(this.SentenceArray.length>0){
          for(let i=0;i<this.SentenceArray.length;i++){
            this.SentenceArray[i].time = this.SentenceArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
        }
      },
      err=>{
      }
    )
  }
  bsModalRef: BsModalRef;//弹出的子模块的引用
    openModalWithComponent(essayid) {//按钮的click执行函数，打开一个子模块
    //给子组件的成员赋值，如果子组件中含有list、title同名成员，则自动进行了赋值
     const initialState = {
     }
    this.bsModalRef = this.modalService.show(ModalComponent, { initialState });//显示子组件
    this.modalService.onHidden.subscribe((r: string) => {//子组件关闭后，触发的订阅函数
        if (this.bsModalRef.content.isCancel){}//this.bsModalRef.content代表子组件对象，isCancel是子组件中的一个成员，可以直接访问
            // console.log("取消了" + this.bsModalRef.content.value);//value是子组件的一个数据成员
        else
            // console.log("确定了" + this.bsModalRef.content.value);
           {this.ToDeleteSentence(essayid)}
    })
    }
  ToDeleteSentence(essayid){
    this.sentenceservice.DeleteOneSentence(essayid).subscribe(
      fb=>{
        if(fb["code"]==1000){
          this.toastrservice.success("删除成功")
          this.sentenceservice.GetAllSentenceInfo(this.limit,this.offset,this.searchstring).subscribe(
            fb=>{
              this.SentenceArray=fb["data"]
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
            },
            err=>{
            }
          )
        }
        else{
          this.toastrservice.error("删除失败")
        }
      },
      err=>{
        this.toastrservice.error("删除失败")
      }
    )
  }
}

