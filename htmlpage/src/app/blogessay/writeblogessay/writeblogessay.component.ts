import { Component,OnInit,ChangeDetectorRef } from '@angular/core';
import { Router } from '@angular/router';
import { DatePipe } from '@angular/common';

import { BlogEssayStruct } from '../../data/blogessayStruct'

import { BlogEssayService } from '../../service/blogessay.service'
import { SessionService } from '../../service/session.service'
import { RouteguardService } from '../../service/routeguard.service'

import { ToastrService } from 'ngx-toastr'

import { ROUTES } from '../../config/route-api'

@Component({
  selector: 'app-writeblogessay',
  templateUrl: './writeblogessay.component.html',
  styleUrls: ['./writeblogessay.component.css'],
  providers:[DatePipe,RouteguardService]
})
export class WriteBlogEssayComponent implements OnInit {
  Role:number
  NewBlogEssay:BlogEssayStruct
    constructor(
      private router:Router,
      private blogessayservice:BlogEssayService,
      private datePipe: DatePipe,
      private toastrservice:ToastrService,
      private sessionservice:SessionService,
    ) {}
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
      this.NewBlogEssay=new BlogEssayStruct
      this.NewBlogEssay.author="左糖"
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
             this.blogessayservice.imageHandler(data).subscribe(fb => {
              const range = this.editor.getSelection(true);
              const index = range.index + range.length;
              let img = '<img src="'+fb["link"]+'">';
              if (!this.NewBlogEssay.content){
                this.NewBlogEssay.content=img
              }
              else this.NewBlogEssay.content+=img
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
    NewBlogEssayPush(blogessayinfo:BlogEssayStruct){
      if(!blogessayinfo.title){
        this.toastrservice.error('标题不能为空')
        return
      }
      if(!blogessayinfo.content){
        this.toastrservice.error('内容不能为空')
        return
      }
      if(!blogessayinfo.tag){
        this.toastrservice.error('标签不能为空')
        return
      }
      blogessayinfo.time=Date.now().toString()
      blogessayinfo.time=this.datePipe.transform(blogessayinfo.time, 'yyyy-MM-dd HH:mm:ss')
      this.blogessayservice.WriteNewBlogEssay(blogessayinfo).subscribe(
        fb=>{
          this.toastrservice.success('上传成功')
        },
        err=>{
          this.toastrservice.error('上传失败')
        }
      )
    }
    ToBackEssay(){
      this.router.navigate([ROUTES.showblogessay.route])
    }
}
