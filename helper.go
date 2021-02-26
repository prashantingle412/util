package Helper

import "strconv"

// conversion from string to integer and float
func StringToFloat32(value string) float32 {
	valueConv, err := strconv.ParseFloat(value, 32)
	if err != nil {
		return 0
	}
	return float32(valueConv)
}
func StringToFloat64(value string) float64 {
	valueConv, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return valueConv
}
func StringToInt(value string) int {
	valueConv, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return valueConv
}

// conversion from integer to string
func IntToString(value int) string {
	valueConv := strconv.Itoa(value)
	return valueConv
}
