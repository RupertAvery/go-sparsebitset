package bitset_tests

import (
	bitset "sparsebitset"
	"testing"
)

func Test_And_With_Empty_Matching_Runs(t *testing.T) {
	var ____left = bitset.ToRuns(0, "----------------------------")
	var ___right = bitset.ToRuns(0, "----------------------------")
	var __result = bitset.ToRuns(0, "----------------------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Full_Matching_Runs(t *testing.T) {
	var ____left = bitset.ToRuns(0, "****************************")
	var ___right = bitset.ToRuns(0, "****************************")
	var __result = bitset.ToRuns(0, "****************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Left_Empty_Right_Full_Matching_Runs(t *testing.T) {
	var ____left = bitset.ToRuns(0, "----------------------------")
	var ___right = bitset.ToRuns(0, "****************************")
	var __result = bitset.ToRuns(0, "----------------------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Left_Full_Right_Empty_Matching_Runs(t *testing.T) {
	var ____left = bitset.ToRuns(0, "****************************")
	var ___right = bitset.ToRuns(0, "----------------------------")
	var __result = bitset.ToRuns(0, "----------------------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Left_Enclosing_Run(t *testing.T) {
	var ____left = bitset.ToRuns(0, "****************************")
	var ___right = bitset.ToRuns(0, "-------**************-------")
	var __result = bitset.ToRuns(0, "-------**************-------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Right_Enclosing_Run(t *testing.T) {
	var ____left = bitset.ToRuns(0, "-------**************-------")
	var ___right = bitset.ToRuns(0, "****************************")
	var __result = bitset.ToRuns(0, "-------**************-------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Various_Runs(t *testing.T) {
	var ____left = bitset.ToRuns(0, "-----------------------------***************************-------------********----------*************-----------------************-----***--**")
	var ___right = bitset.ToRuns(0, "---------*****************************-------------**************************------*****----*------***------***------******---------***-----*")
	var __result = bitset.ToRuns(0, "-----------------------------*********-------------*****-------------********----------*----*------*-----------------******-----------*-----*")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Non_Overlapping_Runs_Returns_Empty(t *testing.T) {
	var ____left = bitset.ToRuns(0, "---------------*************")
	var ___right = bitset.ToRuns(0, "***************-------------")
	var __result = bitset.ToRuns(0, "----------------------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Left_Bridging_Run(t *testing.T) {
	var ____left = bitset.ToRuns(0, "---------********************------------")
	var ___right = bitset.ToRuns(0, "**************--------*******************")
	var __result = bitset.ToRuns(0, "---------*****--------*******------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Right_Bridging_Run(t *testing.T) {
	var ____left = bitset.ToRuns(0, "**************--------*******************")
	var ___right = bitset.ToRuns(0, "---------********************------------")
	var __result = bitset.ToRuns(0, "---------*****--------*******------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Alternate_Bridging_Runs(t *testing.T) {
	var ____left = bitset.ToRuns(0, "****--********--------*******************----------")
	var ___right = bitset.ToRuns(0, "---------********************--------*******---****")
	var __result = bitset.ToRuns(0, "---------*****--------*******--------****----------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Left_Leading_Runs(t *testing.T) {
	var ____left = bitset.ToRuns(0, "--****---*******----*********************")
	var ___right = bitset.ToRuns(0, "--------------------------***************")
	var __result = bitset.ToRuns(0, "--------------------------***************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Right_Leading_Runs(t *testing.T) {
	var ____left = bitset.ToRuns(0, "--------------------------***************")
	var ___right = bitset.ToRuns(0, "--****---*******----*********************")
	var __result = bitset.ToRuns(0, "--------------------------***************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Left_Advanced_Overlapping_Run(t *testing.T) {
	var ____left = bitset.ToRuns(0, "--****************----------")
	var ___right = bitset.ToRuns(0, "---------*******************")
	var __result = bitset.ToRuns(0, "---------*********----------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Right_Advanced_Overlapping_Run(t *testing.T) {
	var ____left = bitset.ToRuns(0, "---------*******************")
	var ___right = bitset.ToRuns(0, "--****************----------")
	var __result = bitset.ToRuns(0, "---------*********----------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Start_Synchronized_Left_Advanced_Ending_Overlapping_Run(t *testing.T) {
	var ____left = bitset.ToRuns(0, "***************--------------")
	var ___right = bitset.ToRuns(0, "*********************--------")
	var __result = bitset.ToRuns(0, "***************--------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_Start_Synchronized_Right_Advanced_Ending_Overlapping_Run(t *testing.T) {
	var ____left = bitset.ToRuns(0, "*********************--------")
	var ___right = bitset.ToRuns(0, "***************--------------")
	var __result = bitset.ToRuns(0, "***************--------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_End_Synchronized_Left_Advanced_Overlapping_Run(t *testing.T) {
	var ____left = bitset.ToRuns(0, "--------------***************")
	var ___right = bitset.ToRuns(0, "*****************************")
	var __result = bitset.ToRuns(0, "--------------***************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_End_Synchronized_Right_Advanced_Overlapping_Run(t *testing.T) {
	var ____left = bitset.ToRuns(0, "*****************************")
	var ___right = bitset.ToRuns(0, "--------------***************")
	var __result = bitset.ToRuns(0, "--------------***************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_End_Synchronized_Left_Advanced_Overlapping_Run_With_Trailing_Runs(t *testing.T) {
	var ____left = bitset.ToRuns(0, "--------------***************----***----**-")
	var ___right = bitset.ToRuns(0, "*****************************--------------")
	var __result = bitset.ToRuns(0, "--------------***************--------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_And_With_End_Synchronized_Right_Advanced_Overlapping_Run_With_Trailing_Runs(t *testing.T) {
	var ____left = bitset.ToRuns(0, "*****************************----***----**-")
	var ___right = bitset.ToRuns(0, "--------------***************--------------")
	var __result = bitset.ToRuns(0, "--------------***************--------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.And(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}
