package errorx

type BizError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *BizError) Error() string {
	return e.Message
}

func New(code int, msg string) *BizError {
	return &BizError{
		Code:    code,
		Message: msg,
	}
}

// ========== 通用模块：10 开头 ==========
var (
	ErrSuccess               = New(0, "success")
	ErrForbidden             = New(403, "权限不足")
	ErrInternal              = New(500, "服务器错误")
	ErrInvalidParam          = New(100001, "参数错误")
	ErrUnauthorized          = New(100002, "未授权")
	ErrInternalServer        = New(100003, "服务器内部错误")
	ErrNotFound              = New(100004, "资源不存在")
	ErrPermissionDeny        = New(100005, "无权限访问")
	ErrAuthHeaderMissing     = New(100006, "未提供 Authorization header")
	ErrAuthHeaderFormat      = New(100007, "Authorization 格式错误")
	ErrInvalidToken          = New(100008, "token 无效")
	ErrTokenExpired          = New(100009, "token 已过期")
	ErrTokenInvalid          = New(100010, "token 无效")
	ErrTokenMalformed        = New(100011, "token 格式错误")
	ErrTokenNotValidYet      = New(100012, "token 尚未生效")
	ErrTokenSignatureInvalid = New(100013, "token 签名无效")
	ErrTokenParseFailed      = New(100014, "token 解析失败")
	ErrPermissionDenied      = New(100015, "权限校验失败")
)

var (
	ErrUserNotFound          = New(200001, "用户不存在")
	ErrUserConflict          = New(200002, "用户已存在")
	ErrUserLoginFail         = New(200003, "用户名或密码错误")
	ErrUserPasswordNotMatch  = New(200004, "密码错误")
	ErrAuthGenerateTokenFail = New(200005, "生成 token 失败")
)

var (
	ErrRoleNotFound      = New(300001, "角色不存在")
	ErrAddPoliciesFail   = New(300002, "分配权限失败")
	ErrRoleAlreadyExists = New(300003, "角色已存在")
	ErrRoleIsSystem      = New(300004, "角色为系统内置角色")
)

var (
	ErrApiNotFound = New(400001, "接口不存在")
)
