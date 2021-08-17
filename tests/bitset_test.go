package bitset_tests

import (
	"archive/zip"
	"io"
	"log"
	bitset "sparsebitset"
	"strconv"
	"strings"
	"testing"
)

type MicroBenchmark struct {
	bitsets []*bitset.SparseBitset
}

//const filename = "census1881.zip"

const filename = "census-income.zip"

//const filename = "wikileaks-noquotes.zip"

var benchmark = ReadData(filename)

func ReadData(filename string) MicroBenchmark {
	r, err := zip.OpenReader("data/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	var benchmark MicroBenchmark
	benchmark.bitsets = []*bitset.SparseBitset{}

	for _, f := range r.File {
		newbitset := bitset.NewSparseBitset()

		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		buf, err := io.ReadAll(rc)
		if err != nil {
			log.Fatal(err)
		}
		myString := string(buf[:])

		dp := strings.Split(myString, ",")
		for _, p := range dp {
			val, _ := strconv.Atoi(p)
			newbitset.Add(uint64(val))
			// fmt.Println(err != nil)
			// if err != nil {
			// }
		}
		newbitset.Pack()
		benchmark.bitsets = append(benchmark.bitsets, newbitset)
		rc.Close()
	}

	//fmt.Println(len(benchmark.bitsets[0].GetValues()))

	return benchmark
}

func BenchmarkAnd(b *testing.B) {
	bitsets := benchmark.bitsets

	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var total = int64(0)
		for k := 0; k < len(bitsets)-1; k++ {
			total += bitsets[k].And(bitsets[k+1]).GetPopCount()
		}
		_ = total
	}
}

func BenchmarkOr(b *testing.B) {
	bitsets := benchmark.bitsets

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var total = int64(0)
		for k := 0; k < len(bitsets)-1; k++ {
			total += bitsets[k].Or(bitsets[k+1]).GetPopCount()
		}
	}
}

func BenchmarkIterate(b *testing.B) {
	bitsets := benchmark.bitsets

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var total = int64(0)
		for k := 0; k < len(bitsets)-1; k++ {
			values := bitsets[k].GetValues()
			for j := 0; j < len(values); j++ {
				total += values[j]
			}
		}
	}

}
