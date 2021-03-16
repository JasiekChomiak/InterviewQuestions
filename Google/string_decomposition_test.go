package Google

import "testing"

func Test_decompress(t *testing.T) {
	tests := []struct {
		name string
		arg string
		want string
	}{
		{
			name: "Basic test, no numbers",
			arg: "abcdef",
			want: "abcdef",
		},
		{
			name: "Numbers, no numbers inside brackets",
			arg: "3[abc]4[d]ef",
			want: "abcabcabcddddef",
		},
		{
			name: "Numbers inside brackets",
			arg: "2[ab3[d]ef]g",
			want: "abdddefabdddefg",
		},
		{
			name: "Everything",
			arg: "10[a2[b]c]24[f]ghi",
			want: "abbcabbcabbcabbcabbcabbcabbcabbcabbcabbcffffffffffffffffffffffffghi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decompress(tt.arg); got != tt.want {
				t.Errorf("decompress() = %v, want %v", got, tt.want)
			}
		})
	}
}