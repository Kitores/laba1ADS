package pkg

func Int64ToFloat64(arr []int64) (floatArr []float64) {
	for _, val := range arr {
		floatArr = append(floatArr, float64(val))
	}
	return floatArr
}
func IntToFloat64(arr []int) (floatArr []float64) {
	for _, val := range arr {
		floatArr = append(floatArr, float64(val))
	}
	return floatArr
}
func Float64ToInt(arr []float64) (intArr []int) {
	for _, val := range arr {
		intArr = append(intArr, int(val))
	}
	return intArr
}
