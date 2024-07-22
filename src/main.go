package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type result struct {
	min, max, sum, count float64
}

type agg struct {
	m    map[string]*result
	keys []string
}

func main() {
	file, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	a := &agg{m: make(map[string]*result), keys: make([]string, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		a.f(line)
	}

	a.print(os.Stdout)
}

func (a *agg) f(line string) {
	i := strings.IndexRune(line, ';')
	city := line[:i]
	temp := mustParseFloat64(line[i+1:])
	curr, ok := a.m[city]
	if !ok {
		a.m[city] = &result{
			min:   temp,
			max:   temp,
			sum:   temp,
			count: 1,
		}
		a.keys = append(a.keys, city)
		return
	}

	curr.min = min(curr.min, temp)
	curr.max = max(curr.max, temp)
	curr.sum += temp
	curr.count += 1
}

func (a *agg) print(out io.Writer) {
	sort.Strings(a.keys)

	var sep string = ", "
	fmt.Fprintf(out, "{")
	for i, key := range a.keys {
		if i == len(a.keys)-1 {
			sep = "}\n"
		}
		entry := a.m[key]
		fmt.Fprintf(out, "%s=%.1f/%.1f/%.1f%s", key, entry.min, mean(entry), entry.max, sep)
	}
}

func mean(entry *result) float64 {
	return round(entry.sum/entry.count, 10)
}

func round(x, unit float64) float64 {
	return math.Round(x*unit) / unit
}

func mustParseFloat64(s string) float64 {
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return num
}
