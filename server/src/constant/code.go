package constant

//错误码保存文件

const (
	RE_PROXY_SCHEME = "/xxxxx/"
)

//公用错误码
const (
	SUCCESS         = 1000 //成功
	SYS_ERR         = 1001 //系统出错
	PARA_ERR        = 1002 //请求参数错误
	SESSION_EXPIRED = 1003 //session过期
	NO_AUTH         = 1004 //没有权限
	DB_ERR          = 1005 //数据库操作失败
)

const (
	ADMIN_PWD_WRONG     = 2001 //管理员登陆账号密码不正确
	ADMIN_FILE_WRONG    = 2002 //文件内容不正确
	ADMIN_NOT_EXIST     = 2003 //管理员不存在
)

const (
	EVENT_NOT_FOUND = 4001 //活动不存在
	EVENT_NOT_GOING = 4002 //活动不在进行中
)

const (
	FILE_SUFFIX_NOT_MATCH = 5001 //文件后缀不支持
	FILE_SIZE_TOO_LARGE   = 5002 //文件过大
	FILE_HAS_EXISTED      = 5003 //文件已存在
	FILE_HAS_NOT_EXISTED  = 5004 //文件已存在
)

const (
	RECORD_SCORE_LESS_THAN_BEFORE = 7001
)

/**/
//构建类型
const (
	BUILD_TYPE_DEV  = "DEV"
	BUILD_TYPE_PROD = "PROD"
)

const (
	USER_ALL_PERMIT             = 511
	USER_SUPER_ADMIN            = 1
	USER_EDUCATION_BUREAU_ADMIN = 2
	USER_SCHOOL_ADMIN           = 4
	USER_PLAYER                 = 8
	USER_VISITOR                = 16
)

//活动状态
const (
	Eve_Published = 0
	Eve_Ongoing   = 1
	Eve_Over      = 2
)

var EventStatus = []int{
	Eve_Published, Eve_Ongoing, Eve_Over,
}

//账号状态
const (
	USER_STATUS_DEL           = -1
	USER_STATUS_NORMAL        = 0
	USER_STATUS_LOCK          = 1
	USER_STATUS_BAN           = 2
	USER_STATUS_SENDING_SMS   = 3
	USER_STATUS_SEND_SMS_FAIL = 4
)

func ChkEventStatus(status int) (statusValid bool) {
	statusValid = false
	for _, item := range EventStatus {
		if item == status {
			statusValid = true
		}
	}
	return
}

//操作类型
const OPER_NUM = 10

const (
	OPER_LOGIN      = 0
	OPER_CHK_EVENT  = 2
	OPER_SET_NOTIFY = 3
	OPER_CHK_VIDEO  = 9
	OPER_CHK_PLAYER = 7
)

//图表类型
const (
	CHART_RANK_CHANGE = 5
)
