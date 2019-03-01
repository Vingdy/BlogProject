import { Component,OnInit } from '@angular/core';
import { element } from 'protractor';
import { Router } from '@angular/router'

import { LoginStruct } from './data/loginStruct'

import { LoginService } from './service/login.service'
import { SessionService } from './service/session.service'

import { trigger,state,style,animate,transition } from '@angular/animations';

import { ToastrService } from 'ngx-toastr'

import { ROUTES } from './config/route-api'

import LoginStatus from './data/loginStatus';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  animations:[
    trigger('EnterLeave', [
      state('open', style({
        width:'300px',
        height:'400px',
        })),
      state('closed', style({
        width: '0px',
        // opacity: 1,
      })),
      transition('open => closed', [
        animate('0.5s')
      ]),
      transition('closed => open', [
        animate('0.5s')
      ]),
  ]),
  ]
})
export class AppComponent implements OnInit {
  NewLogin:LoginStruct
  dropmeun_on:boolean;
  LoginAccount:string
  LoginPassword:string
  IsShowLoginBox:boolean
  Role:number
  constructor(
   private router: Router,
   private loginservice:LoginService,
   private toastrservice:ToastrService,
   private sessionservice:SessionService,
  ) { }
  ngOnInit(){
    this.dropmeun_on= false;
    document.body.style.margin="0";
    this.IsShowLoginBox= false;
    this.sessionservice.GetRole().subscribe(
      fb=>{
          if(fb["code"]!=1000){
            this.Role=0
          }else{
              this.Role=fb["data"]
              LoginStatus.isLogin=true;
          }
      },
      err=>{
          this.Role=0
      })
    }
  ToMain(){
  }
  ToBlogEssay(){
    this.router.navigate([ROUTES.showblogessay.route])
  }
  ToGameEssay(){
    this.router.navigate([ROUTES.showgameessay.route])
  }
  ToDrawing(){        
    this.router.navigate([ROUTES.showdrawpicture.route])
  }
  ToDance(){
  }
  ToSentence(){
    this.router.navigate([ROUTES.showsentence.route])
  }
  ToTest(){
    console.log(this.Role)
    console.log(LoginStatus.isLogin)
    if(LoginStatus.isLogin)
    {
      this.toastrservice.success('已登陆')
    }
    else{
      this.toastrservice.warning('未登陆')
    }
  }
  Login(){
    this.NewLogin={loginaccount:this.LoginAccount,loginpassword:this.LoginPassword}
    this.loginservice.AdminLogin(this.NewLogin).subscribe(
      fb=>{
        console.log(fb)
        if(fb["code"]=2000){
          this.toastrservice.success('登陆成功')
          this.sessionservice.GetRole().subscribe(
            fb=>{
                if(fb["code"]!=1000){
                  this.Role=0
                }else{
                    this.Role=fb["data"]

                    LoginStatus.isLogin=true;
          window.location.reload();
                }
            },
            err=>{
                this.Role=0
            })
        }
      },
      err=>{
        this.toastrservice.error('登陆失败')
      }
    )
  }
  ChangeLoginBox(){
    this.IsShowLoginBox = !this.IsShowLoginBox
  }
  HideLoginBox(){
    this.IsShowLoginBox = false
  }
  LogOut(){
    this.loginservice.AdminLogout().subscribe(
      fb=>{
          if(fb["code"]!=1000){
            this.Role=0
          }else{
              this.Role=fb["data"]
              LoginStatus.isLogin=false;
              window.location.reload();
          }
      },
      err=>{
        this.toastrservice.error('退出失败')
      }
    )
  }
}