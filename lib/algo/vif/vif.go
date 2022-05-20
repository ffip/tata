package vif

func Set(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	} else {
		return falseVal
	}
}
