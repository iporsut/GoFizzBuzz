package main

import (
    "testing"
)

func TestInput1Output1(t *testing.T) {
    output := Solve(1)
    if output != "1" {
        t.Error("Expected 1 Actual ", output)
    }
}

func TestInput2Output2(t *testing.T) {
    output := Solve(2)
    if output != "2" {
        t.Error("Expected 2 Actual ", output)
    }
}

func TestInput3OutputFizz(t *testing.T) {
    output := Solve(3)
    if output != "Fizz" {
        t.Error("Expected Fizz Actual ", output)
    }
}

func TestInput4Output4(t *testing.T) {
    output := Solve(4)
    if output != "4" {
        t.Error("Expected 4 Actual ", output)
    }
}

func TestInput5OutputBuzz(t *testing.T) {
    output := Solve(5)
    if output != "Buzz" {
        t.Error("Expected Buzz Actual ", output)
    }
}

func TestInput6OutputFizz(t *testing.T) {
    output := Solve(6)
    if output != "Fizz" {
        t.Error("Expected Fizz Actual ", output)
    }
}

func TestInput10OutputBuzz(t *testing.T) {
    output := Solve(10)
    if output != "Buzz" {
        t.Error("Expected Buzz Actual ", output)
    }
}

func TestInput15OutputFizzBuzz(t *testing.T) {
    output := Solve(15)
    if output != "FizzBuzz" {
        t.Error("Expected FizzBuzz Actual ", output)
    }
}
