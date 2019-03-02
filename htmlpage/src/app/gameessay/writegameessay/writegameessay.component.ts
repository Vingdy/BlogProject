import { Component,OnInit,ViewChild } from '@angular/core';
import { Router } from '@angular/router';
import { DatePipe } from '@angular/common';

import { GameEssayStruct } from '../../data/gameessayStruct'

import { GameEssayService } from '../../service/gameessay.service'
import { SessionService } from '../../service/session.service'

import { ToastrService } from 'ngx-toastr'
import { ImageCropperComponent, CropperSettings } from 'ng2-img-cropper';

import { ROUTES } from '../../config/route-api'

@Component({
  selector: 'app-writegameessay',
  templateUrl: './writegameessay.component.html',
  styleUrls: ['./writegameessay.component.css'],
  providers:[DatePipe]
})
export class WriteGameEssayComponent implements OnInit {
    data: any;
    cropperSettings: CropperSettings;
    Role:number

    OpenCover:boolean
    ChangeCover:boolean

    @ViewChild('cropper', undefined)
    cropper:ImageCropperComponent;

    CoverImage:any=new Image()

  editorContent:string
  NewGameEssay:GameEssayStruct
    constructor(
      private router:Router,
      private gameessayservice:GameEssayService,
      private datePipe: DatePipe,
      private toastrservice:ToastrService,
      private sessionservice:SessionService,
      
    ) { 
        this.cropperSettings = new CropperSettings();
        this.cropperSettings.width = 720;
        this.cropperSettings.height = 280;
        this.cropperSettings.croppedWidth = 720;
        this.cropperSettings.croppedHeight = 280;
        this.cropperSettings.canvasWidth = 720;
        this.cropperSettings.canvasHeight = 280;
        this.data = {};
    }
    ngOnInit(){
        this.OpenCover=false
        this.ChangeCover=false
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
      this.NewGameEssay=new GameEssayStruct
      this.NewGameEssay.author="左糖"
    }
    quillconfig={
        toolbar: [
          ['bold', 'italic', 'underline', 'strike'],        // toggled buttons
          ['blockquote', 'code-block'],
      
          [{ 'header': 1 }, { 'header': 2 }],               // custom button values
          [{ 'list': 'ordered'}, { 'list': 'bullet' }],
          [{ 'script': 'sub'}, { 'script': 'super' }],      // superscript/subscript
          [{ 'indent': '-1'}, { 'indent': '+1' }],          // outdent/indent
          // [{ 'direction': 'rtl' }],                         // text direction
      
          [{ 'size': ['small', false, 'large', 'huge'] }],  // custom dropdown
          [{ 'header': [1, 2, 3, 4, 5, 6, false] }],
      
          [{ 'color': [] }, { 'background': [] }],          // dropdown with defaults from theme
          [{ 'font': [] }],
          [{ 'align': [] }],
      
          // ['clean'],                                         // remove formatting button
      
          ['image']                         // link and image, video
        ]
      };
      
      public editor;
      EditorCreated(quill) {
        const toolbar = quill.getModule('toolbar');
        toolbar.addHandler('image', this.ImageHandle.bind(this));
        this.editor = quill;
      }
      ImageHandle(){
        const Imageinput = document.createElement('input');
        Imageinput.setAttribute('type', 'file');
        Imageinput.setAttribute('accept','image/png, image/gif, image/jpeg, image/bmp, image/x-icon');
        Imageinput.classList.add('ql-image');
        Imageinput.addEventListener('change', () => {
          const file = Imageinput.files[0];
          const data: FormData = new FormData();
          data.append('file', file, file.name);
          if (Imageinput.files != null && Imageinput.files[0] != null) {
             this.gameessayservice.imageHandler(data).subscribe(fb => {
              const range = this.editor.getSelection(true);
              const index = range.index + range.length;
              let img = '<img src="'+fb["link"]+'" style="max-width:100%;">';
              if (!this.NewGameEssay.content){
                this.NewGameEssay.content=img
              }
              else this.NewGameEssay.content+=img
              this.toastrservice.success('上传成功')
            },
        err=>{
            this.toastrservice.error('上传失败')
        }
    );
          }
        });
        Imageinput.click();
    }
    NewGameEssayPush(gameessayinfo:GameEssayStruct){
        if(!gameessayinfo.title){
            this.toastrservice.error('标题不能为空')
            return
          }
          if(!gameessayinfo.content){
            this.toastrservice.error('内容不能为空')
            return
          }
          if(!gameessayinfo.tag){
            this.toastrservice.error('标签不能为空')
            return
          }
          if(!gameessayinfo.cover){
              gameessayinfo.cover=""
          }
      gameessayinfo.time=Date.now().toString()
      gameessayinfo.time=this.datePipe.transform(gameessayinfo.time, 'yyyy-MM-dd HH:mm:ss')
      this.gameessayservice.WriteNewGameEssay(gameessayinfo).subscribe(
        fb=>{
            this.toastrservice.success('上传成功')
          },
          err=>{
            this.toastrservice.error('上传失败')
          }
      )
    }
    ToBackEssay(){
      this.router.navigate([ROUTES.showgameessay.route])
    }
    IsChangeCover(){
        this.ChangeCover=!this.ChangeCover
        this.OpenCover=true
        this.NewGameEssay.cover=""
    }
    ChangeCoverOK(){
        this.CoverImage.src=this.data.image
        let a=this.dataURLtoFile(this.data.image)
        this.ChangeCover=false
        this.OpenCover=true
        const data: FormData = new FormData();
        data.append('file', a );
        this.gameessayservice.imageHandler(data).subscribe(
            fb=>{
                this.NewGameEssay.cover=fb["link"]
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
        this.NewGameEssay.cover=""
    }
    /**
    * 将dataurl转为blob对象
    * @param dataurl
     * @returns {File}
    */
    dataURLtoFile(dataurl: string) {
        let arr = dataurl.split(','), mime = arr[0].match(/:(.*?);/)[1],
          bstr = atob(arr[1]), n = bstr.length, u8arr = new Uint8Array(n);
        while (n--) {
          u8arr[n] = bstr.charCodeAt(n);
        }
        return new File([u8arr],mime.replace("/","."));
      }
}
