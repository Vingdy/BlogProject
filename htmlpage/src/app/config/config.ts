export var BUILD_TYPE = 'LOCAL';  
// export var BUILD_TYPE = 'DEV';  
//  export var BUILD_TYPE = 'DEPLOY';  //部署模式


export var SITE_HOST_URL;
if (BUILD_TYPE == 'LOCAL'){
    SITE_HOST_URL = 'http://localhost:8080';
    
}
if (BUILD_TYPE == 'DEV'){
    SITE_HOST_URL = 'http://111.230.186.233:8080';
}
if (BUILD_TYPE == 'DEPLOY'){
    SITE_HOST_URL = '';
}