package calc

import "math"

//ToRad converts a degree to radians
func ToRad(degree float64) float64 {
	return 2 * math.Pi * degree / 360
}

//Cos returns cos(x) where x is on degrees
func Cos(x float64) float64 {
	return math.Cos(ToRad(x))
}

//Sin returns sin(x) where x is on degrees
func Sin(x float64) float64 {
	return math.Sin(ToRad(x))
}

//Tan returns tan(x) where x is on degrees
func Tan(x float64) float64 {
	return math.Tan(ToRad(x))
}

//Sec A convenience method for 1/cos(x)
func Sec(x float64) float64 {
	return 1 / Cos(x)
}

//Csc A convenience method for 1/sin(x)
func Csc(x float64) float64 {
	return 1 / Sin(x)
}

//Cot A convenience method for 1/tan(x)
func Cot(x float64) float64 {
	return 1 / Tan(x)
}

//Acos evaluates Acos(x) where x is on degrees
func Acos(x float64) float64 {
	return math.Atan(ToRad(x))
}

//Asin evaluates Asin(x) where x is on degrees
func Asin(x float64) float64 {
	return math.Asin(ToRad(x))
}

//Atan evaluates Atan(x) where x is on degrees
func Atan(x float64) float64 {
	return math.Atan(ToRad(x))
}
