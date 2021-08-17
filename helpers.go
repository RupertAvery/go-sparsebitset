package sparsebitset

import (
	"fmt"
	"strings"
)

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func newRun() *Run {
	var a = new(Run)
	a.Start = 0
	a.End = 0
	a.Values = nil
	return a
}

func isOverlapping(a Run, b Run) bool {
	return (a.Start <= b.End && b.Start <= a.End)
}

const MaxUint64 = ^uint64(0)

func countSetBitsFast(x uint64) int64 {
	x -= (x >> 1) & 0x5555555555555555                             //put count of each 2 bits into those 2 bits
	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333) //put count of each 4 bits into those 4 bits
	x = (x + (x >> 4)) & 0x0F0F0F0F0F0F0F0F                        //put count of each 8 bits into those 8 bits
	return (int64)((x * 0x0101010101010101) >> 56)
}

type KeyValueGroup struct {
	key   int64
	value uint64
	group int64
}

func minmax(array []KeyValueGroup) (int64, int64) {
	var max int64 = array[0].key
	var min int64 = array[0].key
	for _, value := range array {
		if max < value.key {
			max = value.key
		}
		if min > value.key {
			min = value.key
		}
	}
	return min, max
}

func selectvalues(array []KeyValueGroup) []uint64 {
	var out []uint64
	for _, value := range array {
		out = append(out, value.value)
	}
	return out
}

type RunCollection struct {
	runs []Run
}

type ValueCollection struct {
	values []uint64
}

func createRunCollection(runs []Run) RunCollection {
	return RunCollection{runs: runs}
}

func createValueCollection(values []uint64) ValueCollection {
	return ValueCollection{values: values}
}

func (runs RunCollection) ToOptimizedBitset() *SparseBitset {
	bitset := new(SparseBitset)
	bitset.runs = runs.runs
	return bitset
}

func (values ValueCollection) ToOptimizedBitset() *SparseBitset {
	bitset := NewSparseBitset()

	for _, bitValue := range values.values {
		bitset.Add(bitValue)
	}

	bitset.Pack()

	return bitset
}

func ToValues(start uint64, expression string) ValueCollection {
	var values []uint64
	expression = strings.Replace(expression, " ", "", -1)
	ptr := 0
	for ptr < len(expression) {
		if expression[ptr] == '*' {
			values = append(values, start+uint64(ptr))
		}
		ptr++
	}
	return createValueCollection(values)
}

func ToRuns(startKey int64, expression string) RunCollection {
	runs := make([]Run, 0)
	expression = strings.Replace(expression, " ", "", -1)
	ptr := 0
	var currentRun *Run
	buffer := make([]uint64, 256)
	bufferPtr := 0
	for ptr < len(expression) {
		if expression[ptr] != '-' {

			if currentRun == nil {
				currentRun = new(Run)
				currentRun.Start = startKey + int64(ptr)
			}

			var value uint64
			if expression[ptr] == '*' {
				value = MaxUint64
				//} else {
				// value = valueLookup[expression[ptr]]
			}

			buffer[bufferPtr] = value
			bufferPtr++
		} else if expression[ptr] == '-' {
			if currentRun != nil {
				currentRun.Values = make([]uint64, bufferPtr)
				copy(currentRun.Values, buffer[:bufferPtr])
				currentRun.End = startKey + int64(ptr) - 1
				runs = append(runs, *currentRun)
				currentRun = nil
				bufferPtr = 0
			}
		}
		ptr++
	}

	if currentRun != nil {
		currentRun.Values = make([]uint64, bufferPtr)
		copy(currentRun.Values, buffer[:bufferPtr])
		currentRun.End = startKey + int64(ptr) - 1
		runs = append(runs, *currentRun)
	}

	return createRunCollection(runs)
}

func (bitset *SparseBitset) Print() {
	values := bitset.GetValues()
	for _, value := range values {
		fmt.Printf("%d,", value)
	}
	fmt.Printf("\n")
}

func (bitset *SparseBitset) PrintIterator() {
	iterator := bitset.GetValuesIterator()
	for iterator.HasNext() {
		fmt.Printf("%d,", iterator.GetNext())
	}
	fmt.Printf("\n")
}
