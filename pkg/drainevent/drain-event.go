// Copyright 2016-2017 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package drainevent

import (
	"time"
)

// DrainEvent gives more context of the drainable event
type DrainEvent struct {
	EventID     string
	Kind        string
	Description string
	State       string
	StartTime   time.Time
	EndTime     time.Time
	Drained     bool
}

// TimeUntilEvent returns the duration until the event start time
func (e *DrainEvent) TimeUntilEvent() time.Duration {
	return e.StartTime.Sub(time.Now())
}
