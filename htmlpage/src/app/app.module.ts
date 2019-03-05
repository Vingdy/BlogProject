import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NgModule } from '@angular/core';

import { AppRouteModule } from './app-routing.module'

import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
// import { NgZorroAntdModule } from 'ng-zorro-antd';
import { HttpModule } from '@angular/http'

import { AppComponent } from './app.component';

import { HomeComponent } from './home/home.component'
import { BuildingComponent } from './building/building.component'

import { ShowBlogEssayComponent } from './blogessay/showblogessay/showblogessay.component'
import { WriteBlogEssayComponent } from './blogessay/writeblogessay/writeblogessay.component'
import { OneBlogEssayComponent } from './blogessay/oneblogessay/oneblogessay.component'
// import { UpdateBlogEssayComponent } from './blogessay/updateblogessay/updateblogessay.component'

import { ShowGameEssayComponent } from './gameessay/showgameessay/showgameessay.component'
import { WriteGameEssayComponent } from './gameessay/writegameessay/writegameessay.component'
import { OneGameEssayComponent } from './gameessay/onegameessay/onegameessay.component'

import { ShowSentenceComponent } from './sentence/showsentence/showsentence.component'
import { WriteSentenceComponent } from './sentence/writesentence/writesentence.component'

import { ShowDrawpictureComponent } from './drawpicture/showdrawpicture/showdrawpicture.component'
import { WriteDrawpictureComponent } from './drawpicture/writedrawpicture/writedrawpicture.component'

import { BlogEssayService } from './service/blogessay.service'
import { GameEssayService } from './service/gameessay.service'
import { SentenceService } from './service/sentence.service'
import { LoginService } from './service/login.service'
import { SessionService } from './service/session.service'
import { RouteguardService } from './service/routeguard.service'
import { DrawpictureService } from './service/drawpicture.service'

// import { SingletonService } from './data/loginStatus'

import { QuillModule } from 'ngx-quill'
import { ToastrModule } from 'ngx-toastr';
import { ImageCropperModule } from 'ng2-img-cropper';
import { ModalModule } from 'ngx-bootstrap';
import { BsModalService,BsModalRef } from 'ngx-bootstrap/modal';
// import { BsModalRef } from 'ngx-bootstrap/modal/bs-modal-ref.service';

import { PageComponent } from './page/page'
import { ModalComponent } from './modal/modal-pop.component'

import { SafeHtmlPipe } from './pipe/html.pipe'




@NgModule({
  declarations: [
    AppComponent,
    HomeComponent,
    BuildingComponent,
    ShowBlogEssayComponent,
    WriteBlogEssayComponent,
    OneBlogEssayComponent,
    PageComponent,
    ShowGameEssayComponent,
    WriteGameEssayComponent,
    OneGameEssayComponent,
    ShowSentenceComponent,
    WriteSentenceComponent,
    SafeHtmlPipe,
    ShowDrawpictureComponent,
    WriteDrawpictureComponent,
    ModalComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,//要补上不然报没有注入错误
    HttpClientModule,
    BrowserAnimationsModule,
    AppRouteModule,
    QuillModule,
    ToastrModule.forRoot(),
    ImageCropperModule,
    ModalModule.forRoot(),
    
    // NgZorroAntdModule
  ],
  providers: [
    BlogEssayService,
    GameEssayService,
    SentenceService,
    LoginService,
    SessionService,
    RouteguardService,
    DrawpictureService,
    // BsModalRef,
    // BsModalService,
],
exports:[
  ImageCropperModule,
  ],
  entryComponents: [
    ModalComponent,
  ], 
  bootstrap: [AppComponent]
})
export class AppModule { }
