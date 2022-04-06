package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {
	freq := make(map[string]int)

	words := strings.Fields(text)
	wordCount := len(words)

	workers := 25
	size := wordCount / workers

	if size < 1 {
		size = 1
	}

	wg := new(sync.WaitGroup)
	freqCh := make(chan map[string]int, workers+1)

	// Iterate through the words in equal parts
	for i, j := 0, size; i < wordCount; i, j = j, j+size {
		if j > wordCount {
			j = wordCount
		}
		wg.Add(1)

		go func(i, j int) {
			partFreq := make(map[string]int)

			for _, word := range words[i:j] {
				word = strings.ReplaceAll(word, ".", "")
				word = strings.ReplaceAll(word, ",", "")
				word = strings.ToLower(word)

				partFreq[word]++
			}

			freqCh <- partFreq
			wg.Done()
		}(i, j)
	}

	wg.Wait()
	close(freqCh)

	// Collect word counts from channel
	for m := range freqCh {
		for word, count := range m {
			freq[word] += count
		}
	}

	return freq
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	data, err := os.ReadFile(DataFile)
	if err != nil {
		panic(err)
	}

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
