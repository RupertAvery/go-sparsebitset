package sparsebitset

func (a *SparseBitset) Or(b *SparseBitset) *SparseBitset {
	var result = new(SparseBitset)
	result.runs = []Run{}

	var ptrA = int64(0)
	var ptrB = int64(0)
	var ptrC = int64(0)
	var currentA = 0
	var currentB = 0

	if len(a.runs) == 0 {
		return b
	}

	if len(b.runs) == 0 {
		return a
	}

	// We.start with the first Run. They may or may not overlap.

	var currentRunA = a.runs[currentA]
	var currentRunB = b.runs[currentB]
	var currentRunC *Run

	// Only process the first and overlapping runs. Exit early when we run out of runs in either bitset.
	for currentA < len(a.runs) && currentB < len(b.runs) {
		// Check if we're in the middle of a run. This happens when one of the staggered overlapping runs end, leaving the longer run hanging
		if ptrA > 0 {
			// Since we are ORing, continue copying elements until we reach the end of this run, or the.start of the opposite run
			for ptrA+currentRunA.Start < min(ptrB+currentRunB.Start, currentRunA.End+1) {
				currentRunC.Values[ptrC] = currentRunA.Values[ptrA]
				ptrA++
				ptrC++
			}

			// Check if we reached the end of our run
			if ptrA+currentRunA.Start > currentRunA.End {
				if currentRunC != nil {
					result.runs = append(result.runs, *currentRunC)
					currentRunC = nil
					ptrC = 0
				}
				// Advance to the next run
				currentA++
				if currentA < len(a.runs) {
					currentRunA = a.runs[currentA]
				} else {
					continue
				}
				ptrA = 0
			}
		}

		// Check if we're in the middle of a run. This happens when one of the staggered overlapping runs end, leaving the longer run hanging
		if ptrB > 0 {
			// Since we are ORing, continue copying elements until we reach the end of this run, or the.start of the opposite run
			for ptrB+currentRunB.Start < min(ptrA+currentRunA.Start, currentRunB.End+1) {
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
				if currentB < len(b.runs) {
					currentRunB = b.runs[currentB]
				} else {
					continue
				}
				ptrB = 0
			}
		}

		if ptrA == 0 && ptrB == 0 {
			if currentRunC != nil {
				result.runs = append(result.runs, *currentRunC)
				currentRunC = nil
				ptrC = 0
			}
		}

		// Check if these runs overlap
		if isOverlapping(currentRunA, currentRunB) {
			// Check if we need to create a new output run
			if currentRunC == nil {
				currentRunC = newRun()
				currentRunC.Start = min(currentRunA.Start, currentRunB.Start)
			}

			// Extend the run as needed
			currentRunC.End = max(currentRunC.End, max(currentRunA.End, currentRunB.End))

			// Copy the current run array to the new one as needed
			if currentRunC.Values == nil || len(currentRunC.Values) < (int)(currentRunC.End-currentRunC.Start+1) {
				var newArray = make([]uint64, currentRunC.End-currentRunC.Start+1)
				if currentRunC.Values != nil {
					copy(newArray, currentRunC.Values)
				}
				currentRunC.Values = newArray
			}

			// Copy elements from A to the current output run until we reach the overlapping element in B
			for ptrA+currentRunA.Start < ptrB+currentRunB.Start {
				currentRunC.Values[ptrC] = currentRunA.Values[ptrA]
				ptrA++
				ptrC++
			}

			// Copy elements from B to the current output run until we reach the overlapping element in A
			for ptrB+currentRunB.Start < ptrA+currentRunA.Start {
				currentRunC.Values[ptrC] = currentRunB.Values[ptrB]
				ptrB++
				ptrC++
			}

			// Now OR the elements in A and B and write to the output run while they overlap
			for int(ptrA) < len(currentRunA.Values) && int(ptrB) < len(currentRunB.Values) {
				currentRunC.Values[ptrC] = currentRunA.Values[ptrA] | currentRunB.Values[ptrB]
				ptrA++
				ptrB++
				ptrC++
			}

			// Check if we've reached the end of current run A
			if int(ptrA) == len(currentRunA.Values) {
				// Advance to the next run
				currentA++
				if currentA < len(a.runs) {
					currentRunA = a.runs[currentA]
				}
				// And reset the pointer
				ptrA = 0
			}

			// Check if we've reached the end of current run B
			if int(ptrB) == len(currentRunB.Values) {
				// Advance to the next run
				currentB++
				if currentB < len(b.runs) {
					currentRunB = b.runs[currentB]
				}
				// And reset the pointer
				ptrB = 0
			}

		} else if currentRunA.Start > currentRunB.End {
			if currentRunC != nil {
				result.runs = append(result.runs, *currentRunC)
				currentRunC = nil
				ptrC = 0
			}
			result.runs = append(result.runs, currentRunB)
			// catchup B
			currentB++
			if currentB < len(b.runs) {
				currentRunB = b.runs[currentB]
			}
		} else if currentRunB.Start > currentRunA.End {
			if currentRunC != nil {
				result.runs = append(result.runs, *currentRunC)
				currentRunC = nil
				ptrC = 0
			}
			// catchup A
			result.runs = append(result.runs, currentRunA)
			currentA++
			if currentA < len(a.runs) {
				currentRunA = a.runs[currentA]
			}
		}
	}

	// Check if we're in the middle of a run. This happens when one of the staggered overlapping runs end, leaving the longer run hanging
	if ptrA > 0 {
		// Since we are ORing, continue copying elements until we reach the end of this run, or the.start of the opposite run
		for ptrA+currentRunA.Start < currentRunA.End+1 {
			currentRunC.Values[ptrC] = currentRunA.Values[ptrA]
			ptrA++
			ptrC++
		}

		// Check if we reached the end of our run
		if ptrA+currentRunA.Start > currentRunA.End {
			if currentRunC != nil {
				result.runs = append(result.runs, *currentRunC)
				currentRunC = nil
				ptrC = 0
			}
			// Advance to the next run
			currentA++
			if currentA < len(a.runs) {
				currentRunA = a.runs[currentA]
			}
			ptrA = 0
		}
	}

	// Check if we're in the middle of a run. This happens when one of the staggered overlapping runs end, leaving the longer run hanging
	if ptrB > 0 {
		// Since we are ORing, continue copying elements until we reach the end of this run, or the.start of the opposite run
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
			if currentB < len(b.runs) {
				currentRunB = b.runs[currentB]
			}
			ptrB = 0
		}
	}

	if currentRunC != nil {
		result.runs = append(result.runs, *currentRunC)
		currentRunC = nil
		ptrC = 0
	}

	for currentA < len(a.runs) {
		currentRunA = a.runs[currentA]
		result.runs = append(result.runs, currentRunA)
		currentA++
	}

	for currentB < len(b.runs) {
		currentRunB = b.runs[currentB]
		result.runs = append(result.runs, currentRunB)
		currentB++
	}

	return result
}
