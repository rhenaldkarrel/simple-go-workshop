package main

import ("testing")

func TestAddAndMultiplyIntegration(t *testing.T){
  sum := Add(3, 4)

  result := Multiply(sum, 2)

  expected := 14

  if result != expected {
    t.Errorf("Integration test failed. Got %d; expected %d", result, expected)
  }
}