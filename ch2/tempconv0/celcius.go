// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

// Celcius is a float64 alias for temp
type Celcius float64

// Fahrenheit is a float64 alis for temp
type Fahrenheit float64

const (
	// AbsoluteZeroC --> Minimum acheiveable temperature
	AbsoluteZeroC Celcius = -273.15
	// FreezingC  --> Freezing point of water at 1 bar
	FreezingC Celcius = 0
	// BoilingC   --> Boiling point of water at 1 bar
	BoilingC Celcius = 100
)

// CToF convert celcius-to-fahernheit
func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts Fahrenheit-to-Celcius
func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}
