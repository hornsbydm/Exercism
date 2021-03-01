package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequncy of each rune in multiple given texts and returns
// this as a FreqMap.
func ConcurrentFrequency(s []string) (retVal FreqMap) {
	retVal = FreqMap{}
	ch := make(chan FreqMap)
	for _, v := range s {
		go func(v string) { ch <- Frequency(v) }(v)
	}

	for i := 0; i < len(s); i++ {
		for k, v := range <-ch {
			retVal[k] += v
		}
	}

	return
}
