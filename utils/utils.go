package utils

func Message(code int, message string) map[string]interface{} {
	return map[string]interface{}{"code": code, "message": message}
}
