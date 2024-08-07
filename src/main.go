package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type result struct {
	min, max, sum, count float64
}

type agg struct {
	m    map[[100]byte]*result
	keys [][100]byte
}

func main() {
	file, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	a := &agg{m: make(map[[100]byte]*result), keys: make([][100]byte, 0)}
	for scanner.Scan() {
		line := scanner.Text()
		a.f(line)
	}

	a.print(os.Stdout)
}

func (a *agg) f(line string) {
	i := strings.IndexRune(line, ';')
	var city [100]byte
	copy(city[:], line[:i])
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
	sort.Slice(a.keys, func(i, j int) bool {
		return string(a.keys[i][:]) < string(a.keys[j][:])
	})

	var sep string = ", "
	fmt.Fprintf(out, "{")
	for i, key := range a.keys {
		if i == len(a.keys)-1 {
			sep = "}\n"
		}
		entry := a.m[key]
		fmt.Fprintf(out, "%s=%.1f/%.1f/%.1f%s", decodeName(key), entry.min, mean(entry), entry.max, sep)
	}
}

func decodeName(bs [100]byte) string {
	runes := make([]rune, 0, 100)
	for i := 0; i < 100; {
		r, w := utf8.DecodeRune(bs[i:])
		if !unicode.IsGraphic(r) {
			break
		}
		i += w
		runes = append(runes, r)
	}
	return string(runes)
}

func mean(entry *result) float64 {
	return roundup(roundup(entry.sum) / entry.count)
}

func roundup(x float64) float64 {
	return math.Floor(x*10+0.5) / 10
}

func round(x, unit float64) float64 {
	return math.Round(x*unit) / unit
}

func mustParseFloat64(s string) float64 {
	c := 0
	var i int = 0
	var negative bool
	if s[i] == '-' {
		negative = true
		i++
	}
	for s[i] != '.' {
		c = c*10 + (int(s[i]) - '0')
		i++
	}
	i++
	c = c*10 + (int(s[i]) - '0')
	if negative {
		return -float64(c) / 10
	}
	return float64(c) / 10
}
