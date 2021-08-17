package sparsebitset

import "sort"

type Run struct {
	Start  int64
	End    int64
	Values []uint64
}

type SparseBitset struct {
	runs      []Run
	bitFields map[int64]uint64
}

func NewSparseBitset() *SparseBitset {
	bitset := new(SparseBitset)
	bitset.bitFields = make(map[int64]uint64)
	return bitset
}

func (bitset *SparseBitset) Add(bit uint64) {
	key := (int64)(bit / 64)
	bitPosition := (int)(bit % 64)
	bitValue := uint64(1) << bitPosition

	if value, ok := bitset.bitFields[key]; !ok {
		bitset.bitFields[key] = bitValue
	} else {
		bitset.bitFields[key] = value | bitValue
	}
}

func (bitset *SparseBitset) Remove(bit uint64) {
	key := (int64)(bit / 64)
	bitPosition := (int)(bit % 64)
	bitValue := uint64(1) << bitPosition

	if value, ok := bitset.bitFields[key]; ok {
		updatedValue := value & ^bitValue
		if updatedValue == 0 {
			delete(bitset.bitFields, key)
		} else {
			bitset.bitFields[key] = updatedValue
		}
	}
}

// Converts runs into bitfields
func (bitset *SparseBitset) Unpack() {
	bitset.bitFields = make(map[int64]uint64)

	for r := 0; r < len(bitset.runs); r++ {
		run := bitset.runs[r]
		for i := int64(0); i < run.End-run.Start+1; i++ {
			bitset.bitFields[i] = run.Values[i]
		}
	}

	bitset.runs = nil
}

// Converts bit fields into runs
func (bitset *SparseBitset) Pack() {

	var keyvaluegrouplist []KeyValueGroup

	// Select(x => new KeyValueGroup { Key = x.Key, Value = x.Value, Group = x.Key - i } )
	for key, value := range bitset.bitFields {
		var kvg KeyValueGroup
		kvg.key = key
		kvg.value = value
		keyvaluegrouplist = append(keyvaluegrouplist, kvg)
	}

	keyFunc := func(p1, p2 *KeyValueGroup) bool {
		return p1.key < p2.key
	}

	// OrderBy(x => x.Key)
	By(keyFunc).Sort(keyvaluegrouplist)

	i := int64(0)
	for _, kvg := range keyvaluegrouplist {
		keyvaluegrouplist[i].group = kvg.key - i
		i++
	}

	// GroupBy(kvg => kvg.Group)
	groups := make(map[int64][]KeyValueGroup)

	for i := 0; i < len(keyvaluegrouplist); i++ {
		kvg := keyvaluegrouplist[i]
		groups[kvg.group] = append(groups[kvg.group], kvg)
	}

	bitset.runs = []Run{}

	keys := make([]int64, 0, len(groups))
	for k := range groups {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for _, key := range keys {
		group := groups[key]
		min, max := minmax(group)
		var run Run
		run.Start = min
		run.End = max
		run.Values = selectvalues(group)
		bitset.runs = append(bitset.runs, run)
	}

	bitset.bitFields = nil
}

func (bitset *SparseBitset) GetPopCount() int64 {
	nSize := int64(0)

	for r := 0; r < len(bitset.runs); r++ {
		run := bitset.runs[r]
		i := 0
		for i < len(run.Values) {
			if run.Values[i] == MaxUint64 {
				nSize += 64
			} else if run.Values[i] > 0 {
				nSize += countSetBitsFast(run.Values[i])
			}
			i++
		}
	}

	return nSize
}

const length = int64(64)

func (bitset *SparseBitset) GetValues() []int64 {
	values := []int64{}
	for _, run := range bitset.runs {
		//var start = run.start
		var runLength = run.End - run.Start + 1

		for j := int64(0); j < runLength; j++ {
			var value = run.Values[j]
			for i := int64(0); i < length; i++ {
				if (value & 1) == 1 {
					values = append(values, (run.Start+j)*int64(64)+i)
				}
				value >>= 1
			}
		}
	}
	return values
}

func (bitset *SparseBitset) createIterator() Int64Iterator {
	run := &bitset.runs[0]
	runLength := run.End - run.Start + 1
	value := int64(run.Values[0])
	return &BitsetIterator{
		run:       run,
		runLength: runLength,
		value:     value,
		runs:      bitset.runs,
	}
}

type Int64Iterator interface {
	HasNext() bool
	GetNext() int64
}

type BitsetIterator struct {
	runs       []Run
	run        *Run
	runIndex   int   // The current run in the bitset
	runLength  int64 // The number of values in the run
	valueIndex int64 // The current value in the run
	value      int64 // The current value
	bitIndex   int64 // The current bit in the value
}

func (b *BitsetIterator) HasNext() bool {
	return b.runIndex < len(b.runs) && b.valueIndex < b.runLength && b.bitIndex < length
}

func (b *BitsetIterator) postIncrement() {
	b.value >>= 1
	b.bitIndex++
	if b.bitIndex < length {
		return
	}

	b.bitIndex = 0
	b.valueIndex++

	if b.valueIndex < b.runLength {
		b.value = int64(b.run.Values[b.valueIndex])
	} else {
		b.valueIndex = 0
		b.runIndex++
		if b.runIndex < len(b.runs) {
			b.run = &b.runs[b.runIndex]
			b.runLength = b.run.End - b.run.Start + 1
		}
	}
}

func (b *BitsetIterator) GetNext() int64 {
	next := (b.run.Start+b.valueIndex)*int64(64) + b.bitIndex
	//b.ctr++
	b.postIncrement()
	for b.HasNext() && (b.value&1) == 0 {
		b.postIncrement()
	}
	return next
}

func (bitset *SparseBitset) GetValuesIterator() Int64Iterator {
	return bitset.createIterator()
}
