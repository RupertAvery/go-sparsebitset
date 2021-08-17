package bitset_tests

import (
	bitset "sparsebitset"
	"testing"
)

func Test_Not_Of_Empty_Is_Overall(t *testing.T) {
	var ____left = bitset.ToRuns(0, "----------------------------")
	var ____full = bitset.ToRuns(0, "****************************")
	var __result = bitset.ToRuns(0, "****************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var fullBitset = ____full.ToOptimizedBitset()
	var actual = leftBitset.Not(fullBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Not_Of_Empty_Is_Overall_With_Gap(t *testing.T) {
	var ____left = bitset.ToRuns(0, "----------------------------")
	var ____full = bitset.ToRuns(0, "**************--************")
	var __result = bitset.ToRuns(0, "**************--************")

	var leftBitset = ____left.ToOptimizedBitset()
	var fullBitset = ____full.ToOptimizedBitset()
	var actual = leftBitset.Not(fullBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Not_Of_Overall_With_Gap_Is_Empty(t *testing.T) {
	var ____left = bitset.ToRuns(0, "****************************")
	var ____full = bitset.ToRuns(0, "**************--************")
	var __result = bitset.ToRuns(0, "----------------------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var fullBitset = ____full.ToOptimizedBitset()
	var actual = leftBitset.Not(fullBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Not(t *testing.T) {
	var ____left = bitset.ToRuns(0, "------------****------------")
	var ____full = bitset.ToRuns(0, "****************************")
	var __result = bitset.ToRuns(0, "************----************")

	var leftBitset = ____left.ToOptimizedBitset()
	var fullBitset = ____full.ToOptimizedBitset()
	var actual = leftBitset.Not(fullBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Not_With_Overall_With_Gap(t *testing.T) {
	var ____left = bitset.ToRuns(0, "------------****------------")
	var ____full = bitset.ToRuns(0, "********-*******************")
	var __result = bitset.ToRuns(0, "********-***----************")

	var leftBitset = ____left.ToOptimizedBitset()
	var fullBitset = ____full.ToOptimizedBitset()
	var actual = leftBitset.Not(fullBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Not_With_Gap_Runs(t *testing.T) {
	var ____left = bitset.ToRuns(0, "------------****------------")
	var ____full = bitset.ToRuns(0, "********-*******************")
	var __result = bitset.ToRuns(0, "********-***----************")

	var leftBitset = ____left.ToOptimizedBitset()
	var fullBitset = ____full.ToOptimizedBitset()
	var actual = leftBitset.Not(fullBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}
