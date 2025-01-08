package errors

// common Error type
const None Type = 0
const InternalServerError Type = 500
const Success Type = 200
const InvalidParams Type = 400
const (
	InvalidToken Type = iota + 1000
	TokenExpired
	UploadError
	NotPermission
	EmptyParameter
	BadRequest
	Forbidden
	NotFound
	NotImplemented
	ServiceUnavailable
	ParameterLengthExceedsLimit
	NoData
	ParameterErr
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
