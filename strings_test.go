package shelf

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	type args struct {
		s   string
		sep rune
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-no-sep",
			args{
				s:   "123",
				sep: ';',
			},
			[]string{"123"},
		},
		{
			"test-sep-at-first",
			args{
				s:   ";123",
				sep: ';',
			},
			[]string{"123"},
		},
		{
			"test-sep-at-end",
			args{
				s:   "123;",
				sep: ';',
			},
			[]string{"123"},
		},
		{
			"test-one-sep-at-mid",
			args{
				s:   "123;456",
				sep: ';',
			},
			[]string{"123", "456"},
		},
		{
			"test-two-sep-at-mid",
			args{
				s:   "123;456;789",
				sep: ';',
			},
			[]string{"123", "456", "789"},
		},
		{
			"test-two-sep-at-mid-and-end",
			args{
				s:   "123;456;789;",
				sep: ';',
			},
			[]string{"123", "456", "789"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitWithEscape(t *testing.T) {
	type args struct {
		s   string
		sep rune
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-escape-sep",
			args{
				s:   `123\;456`,
				sep: ';',
			},
			[]string{"123;456"},
		},
		{
			"test-sep-and-escape-sep",
			args{
				s:   `123;\;456`,
				sep: ';',
			},
			[]string{"123", ";456"},
		},
		{
			"test-back-slash",
			args{
				s:   `123;\\456`,
				sep: ';',
			},
			[]string{"123", `\456`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestSplitWithSpace(t *testing.T) {
	type args struct {
		s   string
		sep rune
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test-space-at-first",
			args{
				s:   ` 123\;456`,
				sep: ';',
			},
			[]string{"123;456"},
		},
		{
			"test-space-at-mid",
			args{
				s:   `123;\;4 56`,
				sep: ';',
			},
			[]string{"123", ";4 56"},
		},
		{
			"test-space-at-end",
			args{
				s:   `123;\\456 `,
				sep: ';',
			},
			[]string{"123", `\456`},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Split(tt.args.s, tt.args.sep); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Split() = %v, want %v", got, tt.want)
			}
		})
	}
}
