package main
​
import (
	"fmt";
    "math"
)
​
func main() {
​
	var c Converter
​
	fmt.Println(c.CentimeterToFeet(1))
	fmt.Println(c.FeetToCentimeter(1))
	fmt.Println(c.CentimeterToFeet(2))
	fmt.Println(c.FeetToCentimeter(2))
}
​
// Feet the named types that represents feet
type Feet float64
type Minute float64
type Seconds float64
type MilliSeconds float64
type Centimeter float64
type Fahrenheit float64
type Celsius float64
type Kilogram float64
type Pounds float64
type Radian float64
type Degree float64
​
​
// Converter contains func to convert values
type Converter struct {
}
​
func (c Converter) MinutesToSeconds(m Minute) Seconds {
	return Seconds(float64(m) * 60)
}
​
func (c Converter) SecondsToMilliSeconds(s Seconds) MilliSeconds {
	return MilliSeconds(float64(s) * 1000)
}
​
func (c Converter) CentimeterToFeet(cm Centimeter) Feet {
	return Feet(float64(cm) / 30.48)
}
​
func (c Converter) FeetToCentimeter(f Feet) Centimeter {
	return Centimeter(float64(f) * 30.48)
}
​
func (c Converter) CelsiusToFahrenheit(celsius Celsius) Fahrenheit {
	return Fahrenheit( (float64(celsius) * (9/5)) + 32 )
}
​
func (c Converter) FahrenheitToCelsius(f Fahrenheit) Celsius {
	return Celsius( (float64(f) - 32) * (5/9) )
}
​
func (c Converter) DegreeToRadian(d Degree) Radian {
	return Radian(float64(d) * (math.Pi/180) )
}
​
func (c Converter) RadianToDegree(r Radian) Degree {
	return Degree(float64(r) * (180/math.Pi) )
}
​
func (c Converter) PoundsToKilogram(p Pounds) Kilogram {
	return Kilogram( float64(p)/2.205  )
}
​
func (c Converter) KilogramToPounds(k Kilogram) Pounds {
	return Pounds( float64(k)*2.205  )
}