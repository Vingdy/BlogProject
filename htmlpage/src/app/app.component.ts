import { Component,OnInit,ViewChild } from '@angular/core';
// import { element } from 'protractor';
import { Location } from '@angular/common';
import { Router,ActivatedRoute,Params,ActivatedRouteSnapshot } from '@angular/router'
import { Title } from '@angular/platform-browser';

import { LoginStruct } from './data/loginStruct'
import { UserStruct } from './data/userStruct'

import { LoginService } from './service/login.service'
import { SessionService } from './service/session.service'

import { trigger,state,style,animate,transition } from '@angular/animations';

import { ToastrService } from 'ngx-toastr'

import { ROUTES } from './config/route-api'

import { ImageCropperComponent, CropperSettings } from 'ng2-img-cropper';

import LoginStatus from './data/loginStatus';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  animations:[
    trigger('LoginEnterLeave', [
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
  trigger('UserEnterLeave', [
    state('open', style({
      width:'550px',
      height:'500px',
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
  IsShowUserBox:boolean
  Role:number
  UserData:UserStruct
  UserName:string
  UserHeadpicture:string
  UserInfo:string
  UserInfoArray=new Array()

  data: any;
  cropperSettings: CropperSettings;
  OpenCover:boolean
  ChangeCover:boolean

  @ViewChild('cropper', undefined)
  cropper:ImageCropperComponent;

  UserImage:any=new Image()

  editorContent:string

  constructor(
    public router: Router,
   private loginservice:LoginService,
   private toastrservice:ToastrService,
   private sessionservice:SessionService,
   private title: Title,
   private acivatedroute:ActivatedRoute
  ) { 
    this.UserData=new UserStruct,
    this.cropperSettings = new CropperSettings();
    this.cropperSettings.width = 100;
    this.cropperSettings.height = 100;
    this.cropperSettings.croppedWidth = 100;
    this.cropperSettings.croppedHeight = 100;
    this.cropperSettings.canvasWidth = 150;
    this.cropperSettings.canvasHeight = 150;
    this.data = {};
    
  }
  ngOnInit(){
    this.title.setTitle('左糖的日记本')
    this.dropmeun_on= false;
    document.body.style.margin="0";
    this.IsShowLoginBox= false;
    this.IsShowUserBox= false;
    // this.IsShowUserData=true;
    this.acivatedroute.params.subscribe((params: Params) => {
      // params
      console.log(params)
      });
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
      this.loginservice.GetUserData().subscribe(
        fb=>{
          this.UserData.name=fb["data"][0]["name"]
          this.UserData.headpicture=fb["data"][0]["headpicture"]
          this.UserData.info=fb["data"][0]["info"]

          this.UserName=fb["data"][0]["name"]
          this.UserHeadpicture=fb["data"][0]["headpicture"]
          this.UserInfo=fb["data"][0]["info"]
          this.UserInfoArray=this.UserInfo.split(",")

          this.data.image=fb["data"][0]["headpicture"]
        }
      )
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
        // console.log(fb)
        if(fb["code"]==1000){
          console.log(fb["code"])
          this.sessionservice.GetRole().subscribe(
            
            fb=>{
              console.log(fb["code"])
                if(fb["code"]!=1000){
                  this.Role=0
                }else{

                    this.Role=fb["data"]
                    LoginStatus.isLogin=true;
                    window.location.reload();
                    this.toastrservice.success('登陆成功')
                }
            },
            err=>{
              console.log("test")
                this.Role=0
            })
        }else{
          this.toastrservice.error('登陆失败')
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
  ChangeUserBox(){
    this.IsShowUserBox = !this.IsShowUserBox
  }
  HideLoginBox(){
    this.IsShowLoginBox = false
  }
  HideUserBox(){
    this.IsShowUserBox = false
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
  IsChangeCover(){
    this.ChangeCover=!this.ChangeCover
    this.OpenCover=true
    this.UserData.headpicture=""
}
ChangeCoverOK(){
    this.UserImage.src=this.data.image
    let a=this.dataURLtoFile(this.data.image)
    this.ChangeCover=false
    this.OpenCover=true
    const data: FormData = new FormData();
    data.append('file', a );
    this.loginservice.imageHandler(data).subscribe(
        fb=>{
            this.UserData.headpicture=fb["link"]
            this.toastrservice.success('上传成功')
        },
        err=>{
            this.toastrservice.error('上传失败')
        }
    )
}
ChangeCoverCancel(){
    this.ChangeCover=false
    this.OpenCover=false
    this.UserData.headpicture=""
}
dataURLtoFile(dataurl: string) {
  let arr = dataurl.split(','), mime = arr[0].match(/:(.*?);/)[1],
    bstr = atob(arr[1]), n = bstr.length, u8arr = new Uint8Array(n);
  while (n--) {
    u8arr[n] = bstr.charCodeAt(n);
  }
  return new File([u8arr],mime.replace("/","."));
}
UpdateUserData(){
  this.loginservice.UpdateUserData(this.UserData).subscribe(
    fb=>{
      if(fb["code"]==1000){
        this.toastrservice.success('修改成功')
        this.loginservice.GetUserData().subscribe(
          fb=>{
            this.UserData.name=fb["data"][0]["name"]
            this.UserData.headpicture=fb["data"][0]["headpicture"]
            this.UserData.info=fb["data"][0]["info"]
  
            this.UserName=fb["data"][0]["name"]
            this.UserHeadpicture=fb["data"][0]["headpicture"]
            this.UserInfo=fb["data"][0]["info"]
            this.UserInfoArray=this.UserInfo.split(",")
  
            this.data.image=fb["data"][0]["headpicture"]
          }
        )
      }else{
        this.toastrservice.error('修改失败')
        }
      } ,
      err=>{
          this.toastrservice.error('修改失败')
      }
  )
}
ShowUserData(router: Router): boolean{
  var path = router.url; 

  const nextRoute = ['home', '/'+ROUTES.showblogessay.route, '/'+ROUTES.showgameessay.route, '/'+ROUTES.showsentence.route,'/'+ROUTES.showdrawpicture.route];

  if (nextRoute.indexOf(path) >= 0) {
      return true;
    }else{
      return false;
    }
  }
}

//ng build --prod --deploy-url /static