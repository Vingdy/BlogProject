export var BUILD_TYPE = 'LOCAL';  
// export var BUILD_TYPE = 'DEV';  
//  export var BUILD_TYPE = 'DEPLOY';  //部署模式

export var SITE_HOST_URL;
if (BUILD_TYPE == 'LOCAL'){
    SITE_HOST_URL = 'http://localhost:80';
    
}
if (BUILD_TYPE == 'DEV'){
    SITE_HOST_URL = 'http://111.230.186.233:80';
}
if (BUILD_TYPE == 'DEPLOY'){
    SITE_HOST_URL = '';
}

export const GET_ALL_BLOG_ESSAY="/api/getallblogessay"
export const WRITE_BLOG_ESSAY="/api/writeblogessay"
export const GET_ONE_BLOG_ESSAY="/api/getoneblogessay"

export const GET_ALL_GAME_ESSAY="/api/getallgameessay"
export const WRITE_GAME_ESSAY="/api/writegameessay"
export const GET_ONE_GAME_ESSAY="/api/getonegameessay"

export const GET_ALL_SENTENCE="/api/getallsentence"
export const WRITE_SENTENCE="/api/writesentence"

export const GET_ALL_DRAWPICTURE="/api/getalldrawpicture"
export const WRITE_DRAWPICTURE="/api/writedrawpicture"
export const GET_ONE_DRAWPICTURE="/api/getonedrawpicture"

export const LOGIN="/api/login"
export const LOGOUT="/api/logout"
export const UPLOAD_IMAGE="/api/uploadimage"

export const GET_ROLE="/api/getrole"