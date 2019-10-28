//Package letter includes a solution for the "Parallel Letter Frequency" problem in the Go track on https://exercism.io.
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

// ConcurrentFrequency counts the frequency of each rune in a given list of texts concurrently and returns this
// data as a FreqMap.
func ConcurrentFrequency(ss []string) FreqMap {
	ch := make(chan FreqMap)
	fm := FreqMap{}

	for _, v := range ss {
		go func(s string) {
			ch <- Frequency(s)
		}(v)
	}

	for range ss {
		for k, v := range <-ch {
			fm[k] += v
		}
	}

	return fm
}
