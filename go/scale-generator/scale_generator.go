package scale

import (
	"errors"
	"strings"
)

var (
	sharpScale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flatScale  = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	intervals  = map[rune]int{'m': 1, 'M': 2, 'A': 3}
)

//Scale generates the musical scale starting with the tonic and following the
//specified interval pattern.
func Scale(tonic, interval string) []string {
	if interval == "" {
		interval = "mmmmmmmmmmmm"
	}
	return buildScale(tonic, interval, whichScale(tonic))
}

//buildScale constructs a scales for the given tonic, interval, and chromatic scale.
func buildScale(tonic string, interval string, sc []string) (retVal []string) {
	i, err := indexPitch(tonic, sc) //determine the starting offset
	if err != nil {
		return []string{}
	}

	for _, v := range interval {
		retVal = append(retVal, sc[i])
		i += intervals[v] //increment by the interval
		i %= len(sc)      //wrap around the end of the slice
	}
	return retVal
}

//indexPitch provides the index of the given tonic in the given scale.
func indexPitch(tonic string, sc []string) (int, error) {
	u := strings.ToUpper(tonic[:1]) + tonic[1:] //deal with minors
	for i, v := range sc {
		if v == u {
			return i, nil
		}
	}
	return 0, errors.New("tonic not found in scale")
}

// whichScale provides the chromatic scale for use with the tonic.
func whichScale(tonic string) []string {
	if strings.Index("Fdgcf", tonic) != -1 {
		return flatScale
	}
	if len(tonic) == 1 {
		return sharpScale
	}
	if tonic[1] == 'b' {
		return flatScale
	}
	return sharpScale
}
