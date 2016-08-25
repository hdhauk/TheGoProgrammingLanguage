package tempconv

// CToF converts a Celsius temperature to Fahrenheit
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// CToK converts Celsius to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// FToC converts a Fahrenheit temperature to Celsius
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// FToK converts Fahrenheit to Kelvin
func FToK(f Fahrenheit) Kelvin { return Kelvin(FToC(f) - AbsoluteZeroC) }

// KToC converts Kelvin to Celsius
func KToC(k Kelvin) Celsius { return Celsius(Celsius(k) + AbsoluteZeroC) }

// KToF converts Kelvin to Fahrenheit
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(Celsius(k) + AbsoluteZeroC)) }
