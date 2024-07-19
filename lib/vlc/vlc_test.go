package vlc

import (
	"reflect"
	"testing"
)

func Test_Encode(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"My name is Ted", "20 30 3C 18 77 4A E4 4D 28"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := Encode(tt.input); got != tt.output {
				t.Errorf("Encode() = %v, want %v", got, tt.output)
			}
		})
	}
}

func Test_prepateText(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"Hello World", "!hello !world"},
		{"test", "test"},
		{"TeSt", "!te!st"},
		{"TEST", "!t!e!s!t"},
		{"12345", "12345"},
		{"", ""},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := prepateText(test.input)
			if result != test.output {
				t.Errorf("Expected %s, got %s", test.output, result)
			}
		})
	}
}

func Test_encodeBin(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"!ted", "001000100110100101"},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := encodeBin(test.input)
			if result != test.output {
				t.Errorf("Expected %s, got %s", test.output, result)
			}
		})
	}
}

func Test_splitByChunks(t *testing.T) {
	type args struct {
		bStr      string
		chunkSize int
	}

	tests := []struct {
		name  string
		input args
		want  BinaryChunks
	}{
		{
			name: "base test",
			input: args{
				bStr:      "001000100110100101",
				chunkSize: 8,
			},
			want: BinaryChunks{"00100010", "01101001", "01000000"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := splitByChunks(test.input.bStr, test.input.chunkSize)
			if !reflect.DeepEqual(result, test.want) {
				t.Errorf("Expected %s, got %s", test.want, result)
			}
		})
	}
}

func TestBinaryChunks_ToHex(t *testing.T) {
	tests := []struct {
		name string
		bcs  BinaryChunks
		want HexChunks
	}{
		{
			name: "base test",
			bcs:  BinaryChunks{"0101111", "10000000"},
			want: HexChunks{"2F", "80"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bcs.ToHex(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
