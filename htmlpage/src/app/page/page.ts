import { Component,OnInit,Input,Output,OnChanges,ChangeDetectorRef,EventEmitter } from '@angular/core';
import { Router } from '@angular/router';
import { Time } from '@angular/common';
import { FormGroup } from '@angular/forms';
declare let layui
@Component({
  selector: 'page',
  templateUrl: './page.html',
  styleUrls: ['./page.css'],
})

export class PageComponent implements OnInit {
@Input('pageParams') pageParams;// 父组件向子组件传值
@Output() change = new EventEmitter()
TotalPage:number
PageCount:number
PageList:any[]=new Array()
CurrentPage:any
limit:any
    constructor(
      private router:Router,
      private changeDetectorRef:ChangeDetectorRef
    ) {
    }
    ngOnInit(){
  }
  ngOnChanges(){
    this.TotalPage=this.pageParams.TotalPage
    this.CurrentPage=this.pageParams.CurrentPage
    this.limit=this.pageParams.limit
    if(!this.TotalPage){
        this.TotalPage=0
    }
    this.InitPageList(this.CurrentPage,this.TotalPage,this.limit)
  }
InitPageList(currentpage,totalnumber:number,limit){
    var totalpage
    this.PageCount=(totalnumber/limit)
    totalpage=this.PageCount.toFixed(0)
    if(Number(totalpage)<this.PageCount){
        totalpage=(Number(totalpage)+1).toFixed(0)
    }
    this.PageCount=Number(totalpage)+1
    this.PageList=[]

    if(this.PageCount<5){
        for (let i = 1; i < this.PageCount; i++) {
            this.PageList.push(i);
        }
    }else if(currentpage<3 ){
        for (let i = 1; i < 6; i++) {
            this.PageList.push(i);
        }
    }else if((this.PageCount-currentpage)<5){
        for (let i = this.PageCount-5; i < this.PageCount; i++) {
            this.PageList.push(i);
        }
    }else {
        for (let i = currentpage-2; i < currentpage+3; i++) {
            this.PageList.push(i);
        }
    }
    this.changeDetectorRef.markForCheck()
    this.changeDetectorRef.detectChanges()
}
ChangePage(choosepage){
    if(choosepage<=0){
        return
    }
    if(choosepage>=this.PageCount){
        return
    }
            this.InitPageList(this.CurrentPage,this.TotalPage,this.limit)
    this.CurrentPage=choosepage
    this.change.emit(this.CurrentPage)
}
}
