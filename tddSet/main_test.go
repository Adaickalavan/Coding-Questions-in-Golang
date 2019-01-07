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

func TestSet_Add(t *testing.T) {
	type args struct {
		key interface{}
	}
	tests := []struct {
		name string
		s    *Set
		args args
	}{
		{name: "Initialize new Set",
			want: &Set{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.key)
		})
	}
}
