import { Component,OnInit,ViewEncapsulation,TemplateRef } from '@angular/core';
import { Router,ActivatedRoute } from '@angular/router';
import { DomSanitizer } from '@angular/platform-browser';

import { BlogEssayStruct } from '../../data/blogessayStruct'

import { BlogEssayService } from '../../service/blogessay.service'
import { SessionService } from '../../service/session.service'

import { ROUTES } from '../../config/route-api'

import { ToastrService } from 'ngx-toastr'
import { BsModalService } from 'ngx-bootstrap/modal';
import { BsModalRef } from 'ngx-bootstrap/modal/bs-modal-ref.service';

import { ModalComponent } from '../../modal/modal-pop.component'

@Component({
  selector: 'app-oneblogessay',
  templateUrl: './oneblogessay.component.html',
  styleUrls: ['./oneblogessay.component.css'],
  encapsulation: ViewEncapsulation.None,
  
})
export class OneBlogEssayComponent implements OnInit {
    Role:number
    private essayid:string;
    CurrentPage:number
    limit:string
    offset:string
    BlogEssayInfo:BlogEssayStruct
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
           {this.ToDeleteBlogEssay()}
    })
    }
  constructor(
    private router:Router,
    private blogessayservice:BlogEssayService,
    private activatedRoute:ActivatedRoute,
    private sessionservice:SessionService,
    private domsanitizer: DomSanitizer,
    private toastrservice:ToastrService,
    private modalService: BsModalService,
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
    this.CurrentPage = this.activatedRoute.snapshot.queryParams["CurrentPage"];
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
ToDeleteBlogEssay(){
    this.blogessayservice.DeleteOneBlogEssay(this.essayid).subscribe(
        fb=>{
            console.log(fb)
            if(fb["code"]==1000){
                this.router.navigate([ROUTES.showblogessay.route])
                this.toastrservice.success("删除成功")
            }else{
                this.toastrservice.error("删除失败")
            }
        },
        err=>{
            console.log(err)
            this.toastrservice.error("删除失败")
        }
    )
}
}