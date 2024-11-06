package response

// 错误码
const (
	// success
	SUCCESS int = 0

	// error -1 ~ -100
	UNKNOWN_FAIL = -1

	FRAMEWORK_FAIL = -10

	DB_FAIL                 = -20
	DB_RESULT_OVERFLOW_FAIL = -21

	AUTHORIZE_FAIL = -30

	RPC_FAIL = -40

	CODE_FAIL    = -50
	TIMEOUT_FAIL = -52

	REQUEST_INPUT_FAIL = -60

	JSON_UNMARSHAL_FAIL = -70
	JSON_MARSHAL_FAIL   = -71
)

func StatusName(code int) string {
	switch code {
	case SUCCESS:
		return "success"
	case UNKNOWN_FAIL:
		return "unknown fail"
	case DB_FAIL:
		return "db fail"
	case DB_RESULT_OVERFLOW_FAIL:
		return "db result overflow fail"
	case FRAMEWORK_FAIL:
		return "framework fail"
	case AUTHORIZE_FAIL:
		return "authorize fail"
	case RPC_FAIL:
		return "rpc fail"
	case CODE_FAIL:
		return "code fail"
	case TIMEOUT_FAIL:
		return "timeout fail"
	case REQUEST_INPUT_FAIL:
		return "request input fail"
	case JSON_UNMARSHAL_FAIL:
		return "json unmarshal fail"
	case JSON_MARSHAL_FAIL:
		return "json marshal fail"
	default:
		return "default no code status"
	}

}
