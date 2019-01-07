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
		want *Set
	}{
		{name: "Add new element into Set",
			s:    &Set{},
			args: args{key: 1},
			want: &Set{1: struct{}{}},
		},
		{name: "Add existing element into Set",
			s:    &Set{1: struct{}{}},
			args: args{key: 1},
			want: &Set{1: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Add(tt.args.key)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Set.Add() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Remove(t *testing.T) {
	type args struct {
		key interface{}
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want *Set
	}{
		{name: "Remove an element from Set",
			s:    &Set{1: struct{}{}},
			args: args{key: 1},
			want: &Set{},
		},
		{name: "Remove an element from empty Set",
			s:    &Set{},
			args: args{key: 1},
			want: &Set{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Remove(tt.args.key)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("Set.Remove() = %v, want %v", tt.s, tt.want)
			}
		})
	}
}

func TestSet_Contains(t *testing.T) {
	type args struct {
		key interface{}
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want bool
	}{
		{name: "Check for an existing element in Set",
			s:    &Set{1: struct{}{}},
			args: args{key: 1},
			want: true,
		},
		{name: "Check for a non-existing element in Set",
			s:    &Set{2: struct{}{}},
			args: args{key: 1},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Contains(tt.args.key); got != tt.want {
				t.Errorf("Set.Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Union(t *testing.T) {
	type args struct {
		other *Set
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want *Set
	}{
		{name: "Union of two sets",
			s:    &Set{1: struct{}{}},
			args: args{other: &Set{3: struct{}{}}},
			want: &Set{1: struct{}{}, 3: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Union(tt.args.other)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Union() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSet_Intersection(t *testing.T) {
	type args struct {
		other *Set
	}
	tests := []struct {
		name string
		s    *Set
		args args
		want *Set
	}{
		{name: "Intersection of two sets",
			s:    &Set{1: struct{}{}, 3: struct{}{}},
			args: args{other: &Set{3: struct{}{}}},
			want: &Set{3: struct{}{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.Intersection(tt.args.other)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set.Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}
