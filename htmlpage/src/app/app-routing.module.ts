import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

import { HomeComponent } from './home/home.component';
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

import { RouteguardService } from './service/routeguard.service' 

export const AppRoute: Routes = [

    {
        path: '',  // 初始路由重定向[写在第一个]
        redirectTo: 'blogessay',
        pathMatch: 'full',
    },
    {
        path: 'building',
        component: BuildingComponent,
        // ShowUserData
    },
    {
        path:'blogessay',
        component: ShowBlogEssayComponent
    },
    {
        path:'writeblogessay',
        component: WriteBlogEssayComponent,
        canActivate: [RouteguardService]
    },
    {
        path:'oneblogessay',
        component: OneBlogEssayComponent
    },
    {
        path:'gameessay',
        component: ShowGameEssayComponent
    },
    {
        path:'writegameessay',
        component: WriteGameEssayComponent,
        canActivate: [RouteguardService]
    },
    {
        path:'onegameessay',
        component: OneGameEssayComponent
    },    {
        path:'sentence',
        component: ShowSentenceComponent
    },
    {
        path:'writesentence',
        component: WriteSentenceComponent,
        canActivate: [RouteguardService]
    },
    {
        path:"drawpicture",
        component:ShowDrawpictureComponent,
    },
    {
        path:"writedrawpicture",
        component:WriteDrawpictureComponent,
        canActivate:[RouteguardService]
    },
    {
        path: '**',// 错误路由重定向[写在最后一个]
        redirectTo: 'blogessay',
        pathMatch: 'full',
    },
];
@NgModule({
    imports: [
        RouterModule.forRoot(AppRoute),
    ],
    exports: [
        RouterModule
    ],
})
export class AppRouteModule { }
