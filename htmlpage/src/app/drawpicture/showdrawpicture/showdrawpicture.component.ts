import { Component,OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { Observable } from 'Rxjs'

import { DrawpictureStruct } from '../../data/DrawpictureStruct'
// import { UserStruct } from '../../data/DrawpictureStruct'

import { DrawpictureService } from '../../service/drawpicture.service'
import { SessionService } from '../../service/session.service'

import { ROUTES } from '../../config/route-api'

import { ToastrService } from 'ngx-toastr'
import { BsModalService,BsModalRef } from 'ngx-bootstrap/modal';
// import { BsModalRef } from 'ngx-bootstrap/modal/bs-modal-ref.service';

import { ModalComponent } from '../../modal/modal-pop.component'

@Component({
  selector: 'app-showdrawpicture',
  templateUrl: './showdrawpicture.component.html',
  styleUrls: ['./showdrawpicture.component.css'],
})
export class ShowDrawpictureComponent implements OnInit {
    IsEmpty:boolean
    Role:number

  CurrentPage:number
  searchstring:string
  TotalPage:string

  limit:string
  offset:string
  DrawpictureArray:DrawpictureStruct[]
  TagArray=new Array()
  TimeArray=new Array()

  constructor(
    private router:Router,
    private drawpictureservice:DrawpictureService,
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
    this.drawpictureservice.GetAllDrawpictureInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.DrawpictureArray=fb["data"]
        if(fb["code"]==1000){
        if(this.DrawpictureArray.length>0){
          for(let i=0;i<this.DrawpictureArray.length;i++){
            this.DrawpictureArray[i].time = this.DrawpictureArray[i].time.replace('Z','+08:00')
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
  ToWriteDrawpicture(){
    this.router.navigate([ROUTES.writedrawpicture.route])
  }
  CurrentPageOut(CurrentPageOut) {
    this.CurrentPage=CurrentPageOut
    this.ChangePage(this.CurrentPage)//返回被选择的当前页进行处理
  }
  ChangePage(choosepage){
    this.DrawpictureArray=[]
    this.offset=String(Number(this.limit)*(choosepage-1))
    this.drawpictureservice.GetAllDrawpictureInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.DrawpictureArray=fb["data"]
        if(this.DrawpictureArray.length>0){
          for(let i=0;i<this.DrawpictureArray.length;i++){
            this.DrawpictureArray[i].time = this.DrawpictureArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
        }
      },
      err=>{
      }
    )
}
CreateFiling(){
    this.drawpictureservice.GetDrawpictureTag().subscribe(
      fb=>{
        this.TagArray=fb["data"]
      },
      err=>{

      }
    )
    this.drawpictureservice.GetDrawpictureTime().subscribe(
      fb=>{
        this.TimeArray=fb["data"]
      },
      err=>{

      }
    )
  }
GetDrawpictureAboutTime(Time){
    // console.log(Time)
    this.DrawpictureArray=[]
    this.searchstring=Time
    this.drawpictureservice.GetAllDrawpictureInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.DrawpictureArray=fb["data"]
        if(this.DrawpictureArray.length>0){
          for(let i=0;i<this.DrawpictureArray.length;i++){
            this.DrawpictureArray[i].time = this.DrawpictureArray[i].time.replace('Z','+08:00')
          }
          this.TotalPage=fb["total"]
        }
      },
      err=>{
      }
    )
  }
  GetDrawpictureAboutTag(Tag){
    this.DrawpictureArray=[]
    this.searchstring=Tag
    this.drawpictureservice.GetAllDrawpictureInfo(this.limit,this.offset,this.searchstring).subscribe(
      fb=>{
        this.DrawpictureArray=fb["data"]
        if(this.DrawpictureArray.length>0){
          for(let i=0;i<this.DrawpictureArray.length;i++){
            this.DrawpictureArray[i].time = this.DrawpictureArray[i].time.replace('Z','+08:00')
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
           {this.ToDeleteDrawPicture(essayid)}
    })
    }
  ToDeleteDrawPicture(essayid){
    this.drawpictureservice.DeleteOneDrawpicture(essayid).subscribe(
      fb=>{
        if(fb["code"]==1000){
          this.toastrservice.success("删除成功")
          this.drawpictureservice.GetAllDrawpictureInfo(this.limit,this.offset,this.searchstring).subscribe(
            fb=>{
              this.DrawpictureArray=fb["data"]
              if(fb["code"]==1000){
              if(this.DrawpictureArray.length>0){
                for(let i=0;i<this.DrawpictureArray.length;i++){
                  this.DrawpictureArray[i].time = this.DrawpictureArray[i].time.replace('Z','+08:00')
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

