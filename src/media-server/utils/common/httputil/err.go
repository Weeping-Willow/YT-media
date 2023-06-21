package httputil

func ErrorMsgToMap(err string) map[string]string {
	return map[string]string{"error": err}
}
