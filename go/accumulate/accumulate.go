package accumulate

//Accumulate performs the specified operation on the collection
func Accumulate(in []string, conv func(string) string) []string {
	retVal := make([]string, len(in))

	for i, v := range in {
		retVal[i] = conv(v)
	}
	return retVal
}
