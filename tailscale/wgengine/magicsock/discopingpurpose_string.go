// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

// Code generated by "stringer -type=discoPingPurpose -trimprefix=ping"; DO NOT EDIT.

package magicsock

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[pingDiscovery-0]
	_ = x[pingHeartbeat-1]
	_ = x[pingCLI-2]
}

const _discoPingPurpose_name = "DiscoveryHeartbeatCLI"

var _discoPingPurpose_index = [...]uint8{0, 9, 18, 21}

func (i discoPingPurpose) String() string {
	if i < 0 || i >= discoPingPurpose(len(_discoPingPurpose_index)-1) {
		return "discoPingPurpose(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _discoPingPurpose_name[_discoPingPurpose_index[i]:_discoPingPurpose_index[i+1]]
}