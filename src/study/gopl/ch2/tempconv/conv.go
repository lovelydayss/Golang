package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC
func KToC(k Kelvin) Celsius { return Celsius(k - Kelvin(AbsoluteZeroC)) }

// CToK ......
func CToK(c Celsius) Kelvin { return Kelvin(c + AbsoluteZeroC) }

// KToF
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

// FToK
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }
