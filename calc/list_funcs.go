package calc

func SUMXMY2(x []int, y []int) int {
	xLen := len(x)
	yLen := len(y)
	if xLen != yLen {
		panic("Invalid Input")
	}
	val := 0
	for i := 0; i < xLen; i++ {
		val = val + (x[i]-y[i])*(x[i]-y[i])
	}
	return val
}
func SUMX2MY2(x []int, y []int) int {
	xLen := len(x)
	yLen := len(y)
	if xLen != yLen {
		panic("Invalid Input")
	}
	xSum := 0
	ySum := 0
	for i := 0; i < xLen; i++ {
		xSum = xSum + (x[i] * x[i])
		ySum = ySum + (y[i] * y[i])
	}
	return xSum - ySum
}
func SUMX2PY2(x []int, y []int) int {
	xLen := len(x)
	yLen := len(y)
	if xLen != yLen {
		panic("Invalid Input")
	}
	xSum := 0
	ySum := 0
	for i := 0; i < xLen; i++ {
		xSum = xSum + (x[i] * x[i])
		ySum = ySum + (y[i] * y[i])
	}
	return xSum + ySum
}
