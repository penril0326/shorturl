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

func TestIsUrlValid(t *testing.T) {
	type args struct {
		originUrl string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				originUrl: "asdasd://123.45.com",
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				originUrl: "http://localhost:8080/123",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				originUrl: "http://www.google.com/",
			},
			want: true,
		},
		{
			name: "4",
			args: args{
				originUrl: "https://www.google.com",
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				originUrl: " http://123.45.com",
			},
			want: false,
		},
		{
			name: "6",
			args: args{
				originUrl: "https://123.45.com",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUrlValid(tt.args.originUrl); got != tt.want {
				t.Errorf("IsUrlValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsT1AfterT2(t *testing.T) {
	type args struct {
		t1 string
		t2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				t1: "2021-05-23T12:00:00Z",
				t2: "2021-05-23T12:10:00Z",
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				t1: "2021-05-24T13:00:00Z",
				t2: "2021-05-23T12:10:00Z",
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				t1: "2021-05-24T12:00:00Z",
				t2: "2021-05-24T12:00:00Z",
			},
			want: false,
		},
	}

	layout := "2006-01-02T15:04:05Z"
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t1, _ := time.Parse(layout, tt.args.t1)
			t2, _ := time.Parse(layout, tt.args.t2)
			t.Log(t1, t2)
			if got := IsT1AfterT2(t1, t2); got != tt.want {
				t.Errorf("IsT1AfterT2() = %v, want %v", got, tt.want)
			}
		})
	}
}
