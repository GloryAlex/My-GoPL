package boiling

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	absoluteZero = -273.15
	freezing     = 0.0
	boiling      = 100.0
)

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - absoluteZero)
}

func CToK(c Celsius) Kelvin {
	return Kelvin(c + absoluteZero)
}

//get methods

func AbsoluteZero() Celsius {
	return absoluteZero
}

func FreezingCelsius() Celsius {
	return freezing
}

func BoilingCelsius() Celsius {
	return boiling
}
