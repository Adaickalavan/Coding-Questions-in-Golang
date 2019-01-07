package main

import (
	"reflect"
	"testing"
)

func TestNewSet(t *testing.T) {
	tests := []struct {
		name string
		want *Set
	}{
		{name: "Initialize new Set",
			want: &Set{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
