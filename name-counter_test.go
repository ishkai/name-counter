package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestCountNamesFromReader(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]int
	}{
		{
			name:  "counts names",
			input: "Алёна\nМиша\nАлёна\nДима\n",
			want: map[string]int{
				"Алёна": 2,
				"Миша":  1,
				"Дима":  1,
			},
		},
		{
			name:  "skips empty lines and spaces",
			input: "  Алёна  \n\nМиша\n   \nМиша\n",
			want: map[string]int{
				"Алёна": 1,
				"Миша":  2,
			},
		},
		{
			name:  "empty input",
			input: "",
			want:  map[string]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CountNamesFromReader(strings.NewReader(tt.input))
			if err != nil {
				t.Fatalf("CountNamesFromReader() returned error: %v", err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountNamesFromReader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeNameList(t *testing.T) {
	tests := []struct {
		name        string
		counts      map[string]int
		sortByCount bool
		want        []NameCount
	}{
		{
			name: "sorts by name",
			counts: map[string]int{
				"Миша":  1,
				"Алёна": 2,
				"Дима":  1,
			},
			sortByCount: false,
			want: []NameCount{
				{Name: "Алёна", Count: 2},
				{Name: "Дима", Count: 1},
				{Name: "Миша", Count: 1},
			},
		},
		{
			name: "sorts by frequency",
			counts: map[string]int{
				"Миша":  1,
				"Алёна": 2,
				"Дима":  1,
			},
			sortByCount: true,
			want: []NameCount{
				{Name: "Алёна", Count: 2},
				{Name: "Дима", Count: 1},
				{Name: "Миша", Count: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MakeNameList(tt.counts, tt.sortByCount)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeNameList() = %v, want %v", got, tt.want)
			}
		})
	}
}
