package main

import (
  "testing"
)

func TestAdd(t *testing.T){
  angka1 := 1
  angka2 := 6
  result := Add(angka1, angka2)
  expected := angka1 + angka2

  if result != expected {
    t.Errorf("Add(%d, %d) = %d; want %d", angka1, angka2, result, expected)
  }
}

func TestAddTableDriven(t *testing.T) {
  tests := []struct {
      a, b     int
      expected int
  }{
      {2, 3, 5},
      {-1, 1, 0},
      {0, 0, 0},
  }

  for _, tt := range tests {
      result := Add(tt.a, tt.b)
      if result != tt.expected {
          t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
      }
  }
}