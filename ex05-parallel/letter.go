package letter

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
    freqs := FreqMap{}
    for _, v := range s {
	if v >= 97 && v <= 122 {
            v = v - 97
	    } else {
		if v >= 65 && v<= 90 {
                    v = v - 65
		}
	    }
	freqs[v]++
    }
    return freqs
}

func ConcurrentFrequency(texts []string) FreqMap {
    res_chan := make(chan FreqMap, len(texts))
    for _, s := range texts {
	go func(s string) {
	    res_chan <- Frequency(s)
	}(s)
    }
    res := FreqMap{}
    for range texts {
	for i, v := range <-res_chan {
	    res[i] += v
	}
    }
    return res
}
