import { Component,OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { DrawpictureStruct } from '../../data/DrawpictureStruct'

import { DrawpictureService } from '../../service/drawpicture.service'
import { SessionService } from '../../service/session.service'

import { ROUTES } from '../../config/route-api'

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

  constructor(
    private router:Router,
    private drawpictureservice:DrawpictureService,
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
    this.drawpictureservice.GetAllDrawpictureInfo(this.limit,this.offset).subscribe(
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
    this.drawpictureservice.GetAllDrawpictureInfo(this.limit,this.offset).subscribe(
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
}

