// Code generated by "stringer -type=MoodState"; DO NOT EDIT.

package socialmedia

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MoodStateNeutral-0]
	_ = x[MoodStateHappy-1]
	_ = x[MoodStateSad-2]
	_ = x[MoodStateAngry-3]
	_ = x[MoodStateHopeful-4]
	_ = x[MoodStateThrilled-5]
	_ = x[MoodStateBored-6]
	_ = x[MoodStateShy-7]
	_ = x[MoodStateComical-8]
	_ = x[MoodStateOnCloudNine-9]
}

const _MoodState_name = "MoodStateNeutralMoodStateHappyMoodStateSadMoodStateAngryMoodStateHopefulMoodStateThrilledMoodStateBoredMoodStateShyMoodStateComicalMoodStateOnCloudNine"

var _MoodState_index = [...]uint8{0, 16, 30, 42, 56, 72, 89, 103, 115, 131, 151}

func (i MoodState) String() string {
	if i < 0 || i >= MoodState(len(_MoodState_index)-1) {
		return "MoodState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MoodState_name[_MoodState_index[i]:_MoodState_index[i+1]]
}
