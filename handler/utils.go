package handler

func jsonMsg(msg string) map[string]string {
	return map[string]string{
		"message": msg,
	}
}
