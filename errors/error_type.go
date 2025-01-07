package errors

// common Error type
const None Type = 0
const InternalServerError Type = 500
const Success Type = 200
const InvalidParams Type = 400
const (
	InvalidToken   Type = iota + 1000 // 无效的令牌。
	TokenExpired                      // 令牌过期。
	UploadError                       //上传失败
	NotPermission                     // 没有权限
	EmptyParameter                    //参数为空
	BadRequest                        // 请求失败
	Forbidden
	NotFound
	NotImplemented
	ServiceUnavailable
	ParameterLengthExceedsLimit //参数长度超过限制
	NoData                      //未找到有效数据
	ParameterErr                //参数错误
)

func (t Type) String() string {
	switch t {
	case None:
		return "none"
	case BadRequest:
		return "bad request"
	case InvalidToken:
		return "InvalidToken"
	case Forbidden:
		return "forbidden"
	case NotFound:
		return "not found"
	case InternalServerError:
		return "internal server errors"
	case NotImplemented:
		return "not implemented"
	case ServiceUnavailable:
		return "service unavailable"
	case NotPermission:
		return "not permission"
	case EmptyParameter:
		return "empty parameter"
	default:
		return ""
	}
}
