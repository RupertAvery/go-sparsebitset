package bitset_tests

import (
	bitset "sparsebitset"
	"testing"
)

func Test_Or_With_Empty_Matching_Values(t *testing.T) {
	var ____left = bitset.ToValues(0, "----------------------------")
	var ___right = bitset.ToValues(0, "----------------------------")
	var __result = bitset.ToValues(0, "----------------------------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Full_Matching_Values(t *testing.T) {
	var ____left = bitset.ToValues(0, "****************************")
	var ___right = bitset.ToValues(0, "****************************")
	var __result = bitset.ToValues(0, "****************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Left_Empty_Right_Full_Matching_Values(t *testing.T) {
	var ____left = bitset.ToValues(0, "----------------------------")
	var ___right = bitset.ToValues(0, "****************************")
	var __result = bitset.ToValues(0, "****************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Left_Full_Right_Empty_Matching_Values(t *testing.T) {
	var ____left = bitset.ToValues(0, "****************************")
	var ___right = bitset.ToValues(0, "----------------------------")
	var __result = bitset.ToValues(0, "****************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Left_Enclosing_Value(t *testing.T) {
	var ____left = bitset.ToValues(0, "****************************")
	var ___right = bitset.ToValues(0, "-------**************-------")
	var __result = bitset.ToValues(0, "****************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Right_Enclosing_Value(t *testing.T) {
	var ____left = bitset.ToValues(0, "-------**************-------")
	var ___right = bitset.ToValues(0, "****************************")
	var __result = bitset.ToValues(0, "****************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Various_Values(t *testing.T) {
	var ____left = bitset.ToValues(0, "-----------------------------***************************-------------********----------*************-----------------************-----***--**")
	var ___right = bitset.ToValues(0, "---------*****************************-------------**************************------*****----*------***------***------******---------***-----*")
	var __result = bitset.ToValues(0, "---------********************************************************************------*******************------***------************---*****--**")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Non_Overlapping_Values_Returns_Empty(t *testing.T) {
	var ____left = bitset.ToValues(0, "---------------*************")
	var ___right = bitset.ToValues(0, "***************-------------")
	var __result = bitset.ToValues(0, "****************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Left_Bridging_Value(t *testing.T) {
	var ____left = bitset.ToValues(0, "---------********************------------")
	var ___right = bitset.ToValues(0, "**************--------*******************")
	var __result = bitset.ToValues(0, "*****************************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Right_Bridging_Value(t *testing.T) {
	var ____left = bitset.ToValues(0, "**************--------*******************")
	var ___right = bitset.ToValues(0, "---------********************------------")
	var __result = bitset.ToValues(0, "*****************************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Alternate_Bridging_Values(t *testing.T) {
	var ____left = bitset.ToValues(0, "****--********--------*******************----------")
	var ___right = bitset.ToValues(0, "---------********************--------*******---****")
	var __result = bitset.ToValues(0, "****--**************************************---****")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Left_Leading_Values(t *testing.T) {
	var ____left = bitset.ToValues(0, "--****---*******----*********************")
	var ___right = bitset.ToValues(0, "--------------------------***************")
	var __result = bitset.ToValues(0, "--****---*******----*********************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Right_Leading_Values(t *testing.T) {
	var ____left = bitset.ToValues(0, "--------------------------***************")
	var ___right = bitset.ToValues(0, "--****---*******----*********************")
	var __result = bitset.ToValues(0, "--****---*******----*********************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Left_Advanced_Overlapping_Value(t *testing.T) {
	var ____left = bitset.ToValues(0, "--****************----------")
	var ___right = bitset.ToValues(0, "---------*******************")
	var __result = bitset.ToValues(0, "--**************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Right_Advanced_Overlapping_Value(t *testing.T) {
	var ____left = bitset.ToValues(0, "---------*******************")
	var ___right = bitset.ToValues(0, "--****************----------")
	var __result = bitset.ToValues(0, "--**************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Start_Synchronized_Left_Advanced_Ending_Overlapping_Value(t *testing.T) {
	var ____left = bitset.ToValues(0, "***************--------------")
	var ___right = bitset.ToValues(0, "*********************--------")
	var __result = bitset.ToValues(0, "*********************--------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_Start_Synchronized_Right_Advanced_Ending_Overlapping_Value(t *testing.T) {
	var ____left = bitset.ToValues(0, "*********************--------")
	var ___right = bitset.ToValues(0, "***************--------------")
	var __result = bitset.ToValues(0, "*********************--------")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_End_Synchronized_Left_Advanced_Overlapping_Value(t *testing.T) {
	var ____left = bitset.ToValues(0, "--------------***************")
	var ___right = bitset.ToValues(0, "*****************************")
	var __result = bitset.ToValues(0, "*****************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_End_Synchronized_Right_Advanced_Overlapping_Value(t *testing.T) {
	var ____left = bitset.ToValues(0, "*****************************")
	var ___right = bitset.ToValues(0, "--------------***************")
	var __result = bitset.ToValues(0, "*****************************")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_End_Synchronized_Left_Advanced_Overlapping_Value_With_Trailing_Values(t *testing.T) {
	var ____left = bitset.ToValues(0, "--------------***************----***----**-")
	var ___right = bitset.ToValues(0, "*****************************--------------")
	var __result = bitset.ToValues(0, "*****************************----***----**-")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}

func Test_Or_With_End_Synchronized_Right_Advanced_Overlapping_Value_With_Trailing_Values(t *testing.T) {
	var ____left = bitset.ToValues(0, "*****************************----***----**-")
	var ___right = bitset.ToValues(0, "--------------***************--------------")
	var __result = bitset.ToValues(0, "*****************************----***----**-")

	var leftBitset = ____left.ToOptimizedBitset()
	var rightBitset = ___right.ToOptimizedBitset()
	var actual = leftBitset.Or(rightBitset).GetValues()
	var expected = __result.ToOptimizedBitset().GetValues()

	AreEqual(t, actual, expected)
}
