package Google

import "testing"

func Test_longestWord(t *testing.T) {
	type args struct {
		word  string
		words []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example test",
			args: args {
				word: "abppplee",
				words: []string{"able", "ale", "apple", "bale", "kangaroo"},
			},
			want: "apple",
		},
		{
			name: "No subwords",
			args: args {
				word: "adbc",
				words: []string{"da", "cb", "ba", "aa", "xy", "xxxxxxxx", "cy"},
			},
			want: "",
		},
		{
			name: "A suffix",
			args: args {
				word: "abcdef",
				words: []string{"def", "ef", "f", "cdef", "aba", "ccc", "dea"},
			},
			want: "cdef",
		},
		{
			name: "A prefix",
			args: args {
				word: "abcdef",
				words: []string{"abc", "ab", "a", "abcd", "aba", "ccc", "dea"},
			},
			want: "abcd",
		},
		{
			name: "Skips",
			args: args {
				word: "aaabcbea",
				words: []string{"aaaaaaaaa", "aaabb", "caa"},
			},
			want: "aaabb",
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestWord(tt.args.word, tt.args.words); got != tt.want {
				t.Errorf("longestWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
