package sparsebitset

func (a *SparseBitset) Not(full *SparseBitset) *SparseBitset {
	var result = new(SparseBitset)

	result.runs = []Run{}

	var ptrA = int64(0)
	var ptrB = int64(0)
	var ptrC = int64(0)
	var currentA = 0
	var currentB = 0

	if len(a.runs) == 0 {
		return full
	}

	if len(full.runs) == 0 {
		return result
	}

	var currentRunA = a.runs[currentA]
	var currentRunB = full.runs[currentB]
	var currentRunC *Run

	// Loop while we still have data in both bitsets
	for currentA < len(a.runs) && currentB < len(full.runs) {
		// Check if we're in the middle of a Run in A
		if ptrA > 0 {
			// Nothing to do, since the full bitset doesn't have data here, the output bitset should likewise be empty
			// So, just flush the current output Run to the result bitset
			if currentRunC != nil {
				result.runs = append(result.runs, *currentRunC)
				currentRunC = nil
				ptrC = 0
			}
			// Check if we're in the middle of a Run in the full bitset
		} else if ptrB > 0 {
			// Since this is a NOT, anywhere there is no data (0) in the source bitset means it should be
			// 1 in the output bitset, or equivalent to the full bitset (since the full bitset may have 0s where there are no respondents)
			for ptrB+currentRunB.Start < min(ptrA+currentRunA.Start, currentRunB.End+1) {
				currentRunC.Values[ptrC] = currentRunB.Values[ptrB]
				ptrB++
				ptrC++
			}

			if ptrB+currentRunB.Start > currentRunB.End {
				if currentRunC != nil {
					result.runs = append(result.runs, *currentRunC)
					currentRunC = nil
					ptrC = 0
				}
				// Advance to the next run
				currentB++
				if currentB < len(full.runs) {
					currentRunB = full.runs[currentB]
				} else {
					continue
				}
				ptrB = 0
			}
		} else if ptrA == 0 && ptrB == 0 {
			if currentRunC != nil {
				result.runs = append(result.runs, *currentRunC)
				currentRunC = nil
				ptrC = 0
			}
		}

		/// Check if the runs overlap
		if isOverlapping(currentRunA, currentRunB) {
			// If there is no current run, create one
			if currentRunC == nil {
				currentRunC = newRun()
				currentRunC.Start = min(currentRunA.Start, currentRunB.Start)
			}

			// Calculate the end of the run
			currentRunC.End = max(currentRunC.End, max(currentRunA.End, currentRunB.End))

			// Check if we need to initialize, or expand the.values array
			if currentRunC.Values == nil || len(currentRunC.Values) < int(currentRunC.End-currentRunC.Start+1) {
				var newArray = make([]uint64, currentRunC.End-currentRunC.Start+1)
				if currentRunC.Values != nil {
					copy(newArray, currentRunC.Values)
				}
				currentRunC.Values = newArray
			}

			for ptrA+currentRunA.Start < ptrB+currentRunB.Start {
				ptrA++
			}

			for ptrB+currentRunB.Start < ptrA+currentRunA.Start {
				currentRunC.Values[ptrC] = currentRunB.Values[ptrB]
				ptrB++
				ptrC++
			}

			for int(ptrA) < len(currentRunA.Values) && int(ptrB) < len(currentRunB.Values) {
				currentRunC.Values[ptrC] = ^currentRunA.Values[ptrA] & currentRunB.Values[ptrB]
				ptrA++
				ptrB++
				ptrC++
			}

			if int(ptrA) == len(currentRunA.Values) {
				currentA++
				if currentA < len(a.runs) {
					currentRunA = a.runs[currentA]
				}
				ptrA = 0
			}

			if int(ptrB) == len(currentRunB.Values) {
				currentB++
				if currentB < len(full.runs) {
					currentRunB = full.runs[currentB]
				}
				ptrB = 0
			}

		} else if currentRunA.Start > currentRunB.End {
			if currentRunC != nil {
				result.runs = append(result.runs, *currentRunC)
				currentRunC = nil
				ptrC = 0
			}
			// catchup B
			result.runs = append(result.runs, currentRunB)
			currentB++
			if currentB < len(full.runs) {
				currentRunB = full.runs[currentB]
			}
			ptrB = 0
		} else if currentRunB.Start > currentRunA.End {
			if currentRunC != nil {
				result.runs = append(result.runs, *currentRunC)
				currentRunC = nil
				ptrC = 0
			}
			// catchup A
			// Nothing to write to output,
			currentA++
			if currentA < len(a.runs) {
				currentRunA = a.runs[currentA]
			}
			ptrA = 0
		}
	}

	if ptrA > 0 {
		if currentRunC != nil {
			result.runs = append(result.runs, *currentRunC)
			currentRunC = nil
			ptrC = 0
		}
	}

	if ptrB > 0 {
		for ptrB+currentRunB.Start < currentRunB.End+1 {
			currentRunC.Values[ptrC] = currentRunB.Values[ptrB]
			ptrB++
			ptrC++
		}

		// Check if we reached the end of our run
		if ptrB+currentRunB.Start > currentRunB.End {
			if currentRunC != nil {
				result.runs = append(result.runs, *currentRunC)
				currentRunC = nil
				ptrC = 0
			}
			// Advance to the next run
			currentB++
			if currentB < len(full.runs) {
				currentRunB = full.runs[currentB]
			}
			ptrB = 0
		}
	}

	if currentRunC != nil {
		result.runs = append(result.runs, *currentRunC)
		currentRunC = nil
		ptrC = 0
	}

	//while (currentA < len(a.runs))
	//{
	//    currentA++
	//}

	for currentB < len(full.runs) {
		currentRunB = full.runs[currentB]
		result.runs = append(result.runs, currentRunB)
		currentB++
	}

	return result
}
