package space

type Planet string

const (
	earthOrbit = 31557600 //Second in 1 earth year
)

// Age (seconds) in earth years
func Age(age float64, p Planet) float64 {
	planetOrbits := map[string]float64{
		"Earth":   earthOrbit,
		"Mercury": earthOrbit * 00.24084670,
		"Venus":   earthOrbit * 00.61519726,
		"Mars":    earthOrbit * 01.88081580,
		"Jupiter": earthOrbit * 11.86261500,
		"Saturn":  earthOrbit * 29.44749800,
		"Uranus":  earthOrbit * 84.01684600,
		"Neptune": earthOrbit * 164.7913200,
	}
	return age / planetOrbits[string(p)]
}
