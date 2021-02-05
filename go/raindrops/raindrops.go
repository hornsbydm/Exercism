package raindrops

import "fmt"

//Convert  a number into a string that contains raindrop sounds corresponding to
//certain potential factors
func Convert(r int) (retVal string) {

	if r%3 == 0 {
		retVal += "Pling"
	}

	if r%5 == 0 {
		retVal += "Plang"
	}

	if r%7 == 0 {
		retVal += "Plong"
	}

	if retVal != "" {
		return retVal
	}
	return fmt.Sprintf("%d", r)

}
