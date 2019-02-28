import { Component,OnInit } from '@angular/core';
import { Router,ActivatedRoute } from '@angular/router';

import { GameEssayStruct } from '../../data/gameessayStruct'

import { GameEssayService } from '../../service/gameessay.service'
import { SessionService } from '../../service/session.service'

import { ROUTES } from '../../config/route-api'

@Component({
  selector: 'app-onegameessay',
  templateUrl: './onegameessay.component.html',
  styleUrls: ['./onegameessay.component.css'],
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
  ToBackList(){
    let CurrentPage=Number((Number(this.essayid)/5).toFixed(0))
    let Para=(Number(this.essayid)/5)
    if(CurrentPage-Para<0){
        CurrentPage+=1
    }
      return CurrentPage
  }
}