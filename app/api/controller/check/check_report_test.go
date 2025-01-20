// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package check

import (
	"testing"

	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

func Test_getStartedTime(t *testing.T) {
	type args struct {
		in    *ReportInput
		check types.Check
		now   int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "nothing to pending",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusPending},
				check: types.Check{},
				now:   1234,
			},
			want: 0,
		},

		{
			name: "nothing to running",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusRunning},
				check: types.Check{},
				now:   1234,
			},
			want: 1234,
		},
		{
			name: "nothing to completed",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusSuccess},
				check: types.Check{},
				now:   1234,
			},
			want: 1234,
		},
		{
			name: "nothing to completed1",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusSuccess, Started: 1},
				check: types.Check{},
				now:   1234,
			},
			want: 1,
		},
		{
			name: "pending to pending",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusPending, Started: 1},
				check: types.Check{Status: enum.CheckStatusPending, Started: 0},
				now:   1234,
			},
			want: 1,
		},
		{
			name: "pending to pending1",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusPending},
				check: types.Check{Status: enum.CheckStatusPending, Started: 0},
				now:   1234,
			},
			want: 0,
		},
		{
			name: "pending to running",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusRunning},
				check: types.Check{Status: enum.CheckStatusPending, Started: 0},
				now:   1234,
			},
			want: 1234,
		},
		{
			name: "pending to running1",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusRunning, Started: 1},
				check: types.Check{Status: enum.CheckStatusPending, Started: 0},
				now:   1234,
			},
			want: 1,
		},
		{
			name: "pending to completed",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusSuccess, Started: 1},
				check: types.Check{Status: enum.CheckStatusPending, Started: 0},
				now:   1234,
			},
			want: 1,
		},
		{
			name: "pending to completed1",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusSuccess},
				check: types.Check{Status: enum.CheckStatusPending, Started: 0},
				now:   1234,
			},
			want: 1234,
		},
		{
			name: "running to pending",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusPending, Started: 1},
				check: types.Check{Status: enum.CheckStatusRunning, Started: 9876},
				now:   1234,
			},
			want: 1,
		},
		{
			name: "running to pending1",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusPending},
				check: types.Check{Status: enum.CheckStatusRunning, Started: 9876},
				now:   1234,
			},
			want: 0,
		},
		{
			name: "running to running",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusRunning},
				check: types.Check{Status: enum.CheckStatusRunning, Started: 9876},
				now:   1234,
			},
			want: 9876,
		},
		{
			name: "running to running1",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusRunning, Started: 1},
				check: types.Check{Status: enum.CheckStatusRunning, Started: 9876},
				now:   1234,
			},
			want: 1,
		},
		{
			name: "running to completed",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusSuccess},
				check: types.Check{Status: enum.CheckStatusRunning, Started: 9876},
				now:   1234,
			},
			want: 9876,
		},
		{
			name: "running to completed",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusSuccess, Started: 1},
				check: types.Check{Status: enum.CheckStatusRunning, Started: 9876},
				now:   1234,
			},
			want: 1,
		},
		{
			name: "completed to pending",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusPending},
				check: types.Check{Status: enum.CheckStatusSuccess, Started: 9876},
				now:   1234,
			},
			want: 0,
		},
		{
			name: "completed to running",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusRunning},
				check: types.Check{Status: enum.CheckStatusSuccess, Started: 9876},
				now:   1234,
			},
			want: 1234,
		},
		{
			name: "completed to completed",
			args: args{
				in:    &ReportInput{Status: enum.CheckStatusSuccess},
				check: types.Check{Status: enum.CheckStatusSuccess, Started: 9876},
				now:   1234,
			},
			want: 1234,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStartTime(tt.args.in, tt.args.check, tt.args.now); got != tt.want {
				t.Errorf("getStartTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEndTime(t *testing.T) {
	type args struct {
		in  *ReportInput
		now int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "pending",
			args: args{
				in:  &ReportInput{Status: enum.CheckStatusPending},
				now: 1234,
			},
			want: 0,
		},
		{
			name: "running",
			args: args{
				in:  &ReportInput{Status: enum.CheckStatusRunning},
				now: 1234,
			},
			want: 0,
		},
		{
			name: "success",
			args: args{
				in:  &ReportInput{Status: enum.CheckStatusSuccess},
				now: 1234,
			},
			want: 1234,
		},
		{
			name: "failure",
			args: args{
				in:  &ReportInput{Status: enum.CheckStatusFailure},
				now: 1234,
			},
			want: 1234,
		},
		{
			name: "error",
			args: args{
				in:  &ReportInput{Status: enum.CheckStatusError},
				now: 1234,
			},
			want: 1234,
		},
		{
			name: "ended",
			args: args{
				in:  &ReportInput{Status: enum.CheckStatusSuccess, Ended: 4321},
				now: 1234,
			},
			want: 4321,
		},
		{
			name: "ended zero",
			args: args{
				in:  &ReportInput{Status: enum.CheckStatusSuccess, Ended: 0},
				now: 1234,
			},
			want: 1234,
		},
		{
			name: "ended negative",
			args: args{
				in:  &ReportInput{Status: enum.CheckStatusSuccess, Ended: -1},
				now: 1234,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getEndTime(tt.args.in, tt.args.now); got != tt.want {
				t.Errorf("getEndTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
