package utils

func StrLengthCheck(params map[string]string) (bool, string) {
	for k, v := range params {
		if len(v) == 0 {
			return false, k
		}
	}
	return true, ""
}

/*
	toCheck := map[string]string{
		"deviceId":  deviceId,
		"timestamp": timestamp,
		"token":     token,
	}
	if ok, _ := utils.StrLengthCheck(toCheck); !ok {
		return
	}
*/
