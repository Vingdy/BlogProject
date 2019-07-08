// export var BUILD_TYPE = 'LOCAL';  
// export var BUILD_TYPE = 'DEV';  
export var BUILD_TYPE = 'DEPLOY';  //部署模式

export var SITE_HOST_URL;
if (BUILD_TYPE == 'LOCAL'){
    SITE_HOST_URL = 'http://localhost:80';
    
}
if (BUILD_TYPE == 'DEV'){
    SITE_HOST_URL = 'https://111.230.186.233:443';
}
if (BUILD_TYPE == 'DEPLOY'){
    SITE_HOST_URL = 'https://vingdream.cn';
}

export const GET_ALL_BLOG_ESSAY="/api/getallblogessay"
export const WRITE_BLOG_ESSAY="/api/writeblogessay"
export const GET_ONE_BLOG_ESSAY="/api/getoneblogessay"
export const GET_BLOG_ESSAY_TAG="/api/getblogessaythetag"
export const GET_BLOG_ESSAY_TIME="/api/getblogessaythetime"
export const UPDATE_ONE_BLOG_ESSAY="/api/updateoneblogessay"
export const DELETE_ONE_BLOG_ESSAY="/api/deleteoneblogessay"

export const GET_ALL_GAME_ESSAY="/api/getallgameessay"
export const WRITE_GAME_ESSAY="/api/writegameessay"
export const GET_ONE_GAME_ESSAY="/api/getonegameessay"
export const GET_GAME_ESSAY_TAG="/api/getgameessaythetag"
export const GET_GAME_ESSAY_TIME="/api/getgameessaythetime"
export const UPDATE_ONE_GAME_ESSAY="/api/updateonegameessay"
export const DELETE_ONE_GAME_ESSAY="/api/deleteonegameessay"

export const GET_ALL_SENTENCE="/api/getallsentence"
export const WRITE_SENTENCE="/api/writesentence"
export const GET_ONE_SENTENCE="/api/getonesentence"
export const GET_SENTENCE_TIME="/api/getsentencethetime"
export const UPDATE_SENTENCE_ESSAY="/api/updateonesentence"
export const DELETE_SENTENCE_ESSAY="/api/deleteonesentence"

export const GET_ALL_DRAWPICTURE="/api/getalldrawpicture"
export const WRITE_DRAWPICTURE="/api/writedrawpicture"
export const GET_ONE_DRAWPICTURE="/api/getonedrawpicture"
export const GET_DRAWPICTURE_TAG="/api/getdrawpicturethetag"
export const GET_DRAWPICTURE_TIME="/api/getdrawpicturethetime"
export const UPDATE_ONE_DRAWPICTURE="/api/updateonedrawpicture"
export const DELETE_ONE_DRAWPICTURE="/api/deleteonedrawpicture"

export const LOGIN="/api/login"
export const LOGOUT="/api/logout"
export const UPLOAD_IMAGE="/api/uploadimage"

export const GET_ROLE="/api/getrole"

export const GET_USER_DATA="/api/getuserdata"
export const UPDATE_USER_DATA="/api/updateuserdata"