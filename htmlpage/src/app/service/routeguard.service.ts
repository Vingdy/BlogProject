import { Injectable } from '@angular/core';
import { CanActivate, ActivatedRouteSnapshot, RouterStateSnapshot } from '@angular/router';
import { Router } from "@angular/router";
import LoginStatus from '../data/loginStatus';

import { ROUTES } from '../config/route-api'

import { SessionService } from './session.service'

@Injectable()
export class RouteguardService implements CanActivate{

  constructor(
    private router: Router,
    private sessionservice:SessionService,
  ) { }

  canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean{
    // 返回值 true: 跳转到当前路由 false: 不跳转到当前路由
    // 当前路由名称
    var path = route.routeConfig.path; 
    // nextRoute: 设置需要路由守卫的路由集合
    const nextRoute = ['home', ROUTES.writeblogessay.route, ROUTES.writegameessay.route, ROUTES.writesentence.route,ROUTES.writedrawpicture.route];
    let isLogin = LoginStatus.isLogin;  // 是否登录
    // 当前路由是nextRoute指定页时
    if (nextRoute.indexOf(path) >= 0) {
      if (!isLogin) {
        this.router.navigate([ROUTES.showblogessay.route]);
        return false;
      }else{
        return true;
      }
    }
  }

}