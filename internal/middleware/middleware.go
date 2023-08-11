package middleware

import models "inter/internal"

func Validation(body models.Task) bool {
	if body.ElementCount == 0 || body.Delta == 0 || body.L == 0 || body.TTL == 0 {
		return false
	}
	return true
}
