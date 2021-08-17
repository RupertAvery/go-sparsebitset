package sparsebitset

func (a *SparseBitset) And(b *SparseBitset) *SparseBitset {
	result := new(SparseBitset)
	result.runs = []Run{}
	// Our purpose here is to merge Bitsets only where they overlap:
	// Here we have two bitsets with 1 run each.
	// Each letter here represents a Value element in a Run. A dash is an unused key.
	//
	// The figure below represents two staggered runs (they don't align) A and B, with C containing
	// the ANDed Value elements.
	//
	// -------------AAAAAAAAAAAAAAAAA
	// ----------BBBBBBBBB-----------
	//
	//--------------CCCCCC-----------

	var ptrA = int64(0)
	var ptrB = int64(0)
	var ptrC = int64(0)
	var currentA = 0
	var currentB = 0

	aLen := len(a.runs)
	bLen := len(b.runs)

	// If either of the runs are empty, return an empty bitset.
	if aLen == 0 {
		return result
	}

	if bLen == 0 {
		return result
	}

	// We.start with the first Run. They may or may not overlap.

	var currentRunA = a.runs[currentA]
	var currentRunB = b.runs[currentB]
	var currentRunC *Run

	for currentA < aLen && currentB < bLen {
		// after the end of any Run, we write the current Run to the output and.start anew
		if currentRunC != nil {
			result.runs = append(result.runs, *currentRunC)
			currentRunC = nil
			ptrC = 0
		}

		/// Check if the runs overlap
		if currentRunA.Start <= currentRunB.End && currentRunB.Start <= currentRunA.End {
			//if isOverlapping(currentRunA, currentRunB) {
			// If there is no current run, create one
			if currentRunC == nil {
				currentRunC = newRun()
				currentRunC.Start = max(currentRunA.Start, currentRunB.Start)
				currentRunC.End = min(currentRunA.End, currentRunB.End)
			} else {
				currentRunC.End = min(currentRunC.End, min(currentRunA.End, currentRunB.End))
			}

			// Check if we need to initialize, or expand the.values array
			if currentRunC.Values == nil || len(currentRunC.Values) < (int)(currentRunC.End-currentRunC.Start+1) {
				var newArray = make([]uint64, currentRunC.End-currentRunC.Start+1)
				if currentRunC.Values != nil {
					// Copy the old.values to the new array
					copy(newArray, currentRunC.Values)
				}
				currentRunC.Values = newArray
			}

			diffA := (ptrB + currentRunB.Start) - (ptrA + currentRunA.Start)

			if diffA > 0 {
				ptrA += diffA
			}

			diffB := (ptrA + currentRunA.Start) - (ptrB + currentRunB.Start)

			if diffB > 0 {
				ptrB += diffB
			}

			// Now the Run pointers are aligned,.start ANDing the elements until we reach the end of either Run
			aValueLen := len(currentRunA.Values)
			bValueLen := len(currentRunB.Values)

			for int(ptrA) < aValueLen-4 && int(ptrB) < bValueLen-4 {
				currentRunC.Values[ptrC] = currentRunA.Values[ptrA] & currentRunB.Values[ptrB]
				currentRunC.Values[ptrC+1] = currentRunA.Values[ptrA+1] & currentRunB.Values[ptrB+1]
				currentRunC.Values[ptrC+2] = currentRunA.Values[ptrA+2] & currentRunB.Values[ptrB+2]
				currentRunC.Values[ptrC+3] = currentRunA.Values[ptrA+3] & currentRunB.Values[ptrB+3]
				ptrA += 4
				ptrB += 4
				ptrC += 4
			}

			for int(ptrA) < aValueLen && int(ptrB) < bValueLen {
				currentRunC.Values[ptrC] = currentRunA.Values[ptrA] & currentRunB.Values[ptrB]
				ptrA++
				ptrB++
				ptrC++
			}

			// #if Manual1
			// for int(ptrA) < aValueLen && int(ptrB) < bValueLen {
			// 	currentRunC.Values[ptrC] = currentRunA.Values[ptrA] & currentRunB.Values[ptrB]
			// 	ptrA++
			// 	ptrB++
			// 	ptrC++
			// }
			// #endif

			// Check if we reached the end of the A run
			if int(ptrA) == aValueLen {
				// advance to the next A run
				currentA++
				if currentA < len(a.runs) {
					currentRunA = a.runs[currentA]
				}
				ptrA = 0
			}

			// Check if we reached the end of the B run
			if int(ptrB) == bValueLen {
				// advance to the next B run
				currentB++
				if currentB < len(b.runs) {
					currentRunB = b.runs[currentB]
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
			currentB++
			if currentB < len(b.runs) {
				currentRunB = b.runs[currentB]
			}

			ptrB = 0

		} else if currentRunB.Start > currentRunA.End {

			if currentRunC != nil {
				result.runs = append(result.runs, *currentRunC)
				currentRunC = nil
				ptrC = 0
			}

			// catchup A
			currentA++
			if currentA < len(a.runs) {
				currentRunA = a.runs[currentA]
			}

			ptrA = 0
		}
	}

	if currentRunC != nil {
		result.runs = append(result.runs, *currentRunC)
	}

	return result
}
