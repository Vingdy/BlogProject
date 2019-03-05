import { Component,OnInit } from '@angular/core';
import { Router,ActivatedRoute } from '@angular/router';
import { DatePipe } from '@angular/common';
    
import { DrawpictureStruct } from '../../data/drawpictureStruct'
    
import { DrawpictureService } from '../../service/drawpicture.service'
import { SessionService } from '../../service/session.service'
    
import { ToastrService } from 'ngx-toastr'
    
import { ROUTES } from '../../config/route-api'
    
@Component({
    selector: 'app-writedrawpicture',
    templateUrl: './writedrawpicture.component.html',
    styleUrls: ['./writedrawpicture.component.css'],
    providers:[DatePipe]
  })
  export class WriteDrawpictureComponent implements OnInit {
    essayid:string
    IsChooseImage:boolean

    Image:FormData
    ImageData:any
    Role:number

    NewDrawpicture:DrawpictureStruct
      constructor(
        private router:Router,
        private drawpictureservice:DrawpictureService,
        private datePipe: DatePipe,
        private toastrservice:ToastrService,
        private sessionservice:SessionService,
        private activatedroute:ActivatedRoute,
      ) { }
      ngOnInit(){
        this.IsChooseImage=false
        this.NewDrawpicture=new DrawpictureStruct
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
          this.activatedroute.queryParams.subscribe(params => {
            this.essayid = params['essayid']; 
          });
          if(this.essayid){
            this.drawpictureservice.GetOneDrawpictureInfo(this.essayid).subscribe(
              fb=>{
                this.NewDrawpicture.title=fb["data"][0]["title"]
                this.NewDrawpicture.src=fb["data"][0]["src"]
                this.NewDrawpicture.time=fb["data"][0]["time"]
                this.NewDrawpicture.tag=fb["data"][0]["tag"]
                this.ImageData=fb["data"][0]["src"]
              },
              err=>{
    
              }
            )
        }
     }
      NewDrawpicturePush(drawpictureinfo:DrawpictureStruct){
        if(!drawpictureinfo.src){
          this.toastrservice.error("未上传图片")
          return
        }
        if(!drawpictureinfo.title){
          this.toastrservice.error("标题为空")
          return
        }
        drawpictureinfo.time=Date.now().toString()
        drawpictureinfo.time=this.datePipe.transform(drawpictureinfo.time, 'yyyy-MM-dd HH:mm:ss')
        this.drawpictureservice.WriteNewDrawpicture(drawpictureinfo).subscribe(
          fb=>{
            this.toastrservice.success('写入成功')
          },
          err=>{
            this.toastrservice.error('写入失败')
          }
        )
      }
      ToBackDrawpicture(){
        this.router.navigate([ROUTES.showdrawpicture.route])
      }
      fileChangeListener($event) {
        // var image:any = new Image();
        var file:File = $event.target.files[0];
        const data: FormData = new FormData();
        data.append('file', file, file.name);
        this.Image=data
        var myReader:FileReader = new FileReader();
        myReader.readAsDataURL(file);//读取图像文件 result 为 DataURL, DataURL 可直接 赋值给 img.src
        myReader.onloadend = (e) => {
          this.ImageData = myReader.result;
        }
        this.IsChooseImage=true
    }
    UpLoadDrawpicture(){
      this.drawpictureservice.imageHandler(this.Image).subscribe(
        fb=>{
          this.NewDrawpicture.src=fb["link"]
          this.toastrservice.success('上传成功')
        },
        err=>{
          this.toastrservice.error('上传失败')
        }
      )
      this.IsChooseImage=false
    }
    UpLoadDrawpictureCancel(){
      this.Image=null
      this.ImageData=null
      this.IsChooseImage=false
    }
    UpdateDrawpicture(drawpictureinfo:DrawpictureStruct){
      if(!drawpictureinfo.src){
        this.toastrservice.error("未上传图片")
      }
      if(!drawpictureinfo.title){
        this.toastrservice.error("标题为空")
      }
      this.drawpictureservice.UpdateOneDrawpicture(drawpictureinfo).subscribe(
        fb=>{
          this.toastrservice.success('修改成功')
          this.router.navigate([ROUTES.showdrawpicture.route])
        },
        err=>{
          this.toastrservice.error('修改失败')
        }
      )
    }
  }
  