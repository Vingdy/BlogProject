import { Component,OnInit,ViewEncapsulation } from '@angular/core';
import { Router,ActivatedRoute } from '@angular/router';

import { GameEssayStruct } from '../../data/gameessayStruct'

import { GameEssayService } from '../../service/gameessay.service'
import { SessionService } from '../../service/session.service'

import { ROUTES } from '../../config/route-api'

import { ToastrService } from 'ngx-toastr'
import { BsModalService } from 'ngx-bootstrap/modal';
import { BsModalRef } from 'ngx-bootstrap/modal/bs-modal-ref.service';

import { ModalComponent } from '../../modal/modal-pop.component'

@Component({
  selector: 'app-onegameessay',
  templateUrl: './onegameessay.component.html',
  styleUrls: ['./onegameessay.component.css'],
  encapsulation: ViewEncapsulation.None
})
export class OneGameEssayComponent implements OnInit {
    Role:number
  private essayid:string;
  limit:string
  offset:string
  GameEssayInfo:GameEssayStruct

  constructor(
    private router:Router,
    private gameessayservice:GameEssayService,
    private activatedRoute:ActivatedRoute,
    private sessionservice:SessionService,
    private modalService: BsModalService,
    private toastrservice:ToastrService,
    
  ) { }
  bsModalRef: BsModalRef;//弹出的子模块的引用
    openModalWithComponent() {//按钮的click执行函数，打开一个子模块
    //给子组件的成员赋值，如果子组件中含有list、title同名成员，则自动进行了赋值
     const initialState = {
     }
    this.bsModalRef = this.modalService.show(ModalComponent, { initialState });//显示子组件
    this.modalService.onHidden.subscribe((r: string) => {//子组件关闭后，触发的订阅函数
        if (this.bsModalRef.content.isCancel){}//this.bsModalRef.content代表子组件对象，isCancel是子组件中的一个成员，可以直接访问
            // console.log("取消了" + this.bsModalRef.content.value);//value是子组件的一个数据成员
        else
            // console.log("确定了" + this.bsModalRef.content.value);
           {this.ToDeleteGameEssay()}
    })
    }
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
    this.GameEssayInfo=new GameEssayStruct
    this.essayid = this.activatedRoute.snapshot.queryParams["essayid"];
    this.gameessayservice.GetOneGameEssayInfo(this.essayid).subscribe(
        fb=>{
            this.GameEssayInfo=fb["data"][0]
            this.GameEssayInfo.time=this.GameEssayInfo.time.replace('Z','+08:00')
        },
        err=>{
        }
    )
  }

  ToDeleteGameEssay(){
    this.gameessayservice.DeleteOneGameEssay(this.essayid).subscribe(
        fb=>{
            console.log(fb)
            if(fb["code"]==1000){
                this.toastrservice.success("删除成功")
                this.router.navigate([ROUTES.showgameessay.route])
            }else{
                this.toastrservice.error("删除失败")
            }
        },
        err=>{
            this.toastrservice.error("删除失败")
        }
    )
}
}