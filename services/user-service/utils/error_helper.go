package utils

type ErrorAPI struct {
	Field string `json:"field"`
	Msg   string `json:"message"`
}

func MsgFromTag(tag string) string {
	switch tag {
	case "required":
		return "is required"
	case "email":
		return "is invalid"
	default:
		return ""
	}
}
