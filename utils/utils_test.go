package utils

import (
	"testing"
	"time"
)

func TestBase62Encode(t *testing.T) {
	type args struct {
		num int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				num: 3,
			},
			want: "3",
		},
		{
			name: "2",
			args: args{
				num: -1,
			},
			want: "",
		},
		{
			name: "3",
			args: args{
				num: 62,
			},
			want: "10",
		},
		{
			name: "4",
			args: args{
				num: 68,
			},
			want: "16",
		},
		{
			name: "5",
			args: args{
				num: 23498716264,
			},
			want: "pEidyE",
		},
		{
			name: "6",
			args: args{
				num: 1019591930676,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Base62Encode(tt.args.num); got != tt.want {
				t.Errorf("Base62Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsCurrentTimeBeforeThan(t *testing.T) {
	type args struct {
		target time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsCurrentTimeBeforeThan(tt.args.target); got != tt.want {
				t.Errorf("IsCurrentTimeBeforeThan() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUrlValid(t *testing.T) {
	type args struct {
		originUrl string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUrlValid(tt.args.originUrl); got != tt.want {
				t.Errorf("IsUrlValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
