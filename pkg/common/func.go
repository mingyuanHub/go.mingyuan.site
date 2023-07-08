package common

//三元运算符
func If(condition bool, trueV, falseV interface{}) interface{}{
	if condition {
		return trueV
	}
	return falseV
}