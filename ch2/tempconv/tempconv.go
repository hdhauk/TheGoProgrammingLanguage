// Package tempconv performs Celsius and Fahrenheit conversions
package tempconv

import "fmt"

// Celsius - Temperature in degrees Celsius
type Celsius float64

// Fahrenheit - Temperature in degrees Fahrenheit
type Fahrenheit float64

// Kelvin - Temperature in degrees Kelvin
type Kelvin float64

const (
	// AbsoluteZeroC --> Absolute zero
	AbsoluteZeroC Celsius = -273.15
	// FreezingC --> Freezing point of water at 1 Atm
	FreezingC Celsius = 0
	// BoilingC --> Boiling point of water at 1 Atm
	BoilingC Celsius = 100
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
