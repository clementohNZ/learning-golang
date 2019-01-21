# Testing Intro
No assertions in Golang testing. If no error is thrown, pass is assumed. Quite cool!

```go
func TestMySum(t *testing.T) {
	sum := mySum(2,3)
	if sum != 5 {
		t.Errorf("Expected 5, got %d", sum)
	}
}
```

## Things to keep in mind
Tests must:
- be in a file that ends with `_test.go`
- put the file in the same package as the one being tested
- be in a function with the same signature `func TestXxx(t *testing.T)`

## Running tests
`go test`

## Golang Book
```go
package math

import "testing"

func TestAverage(t *testing.T) {
  var v float64
  v = Average([]float64{1,2})
  if v != 1.5 {
    t.Error("Expected 1.5, got ", v)
  }
}
```

The go test command will look for any tests in any of the files in the current folder and run them. 
Tests are identified by starting a function with the word Test and taking one argument of type `*testing.T`. 
In our case since we're testing the Average function we name the test function TestAverage.

In this case we know the average of `[1,2]` should be `1.5` so that's what we check. It's probably a good idea 
to test many different combinations of numbers so let's change our test program a little:

```go
package math

import "testing"

type testpair struct {
  values []float64
  average float64
}

var tests = []testpair{
  { []float64{1,2}, 1.5 },
  { []float64{1,1,1,1,1,1}, 1 },
  { []float64{-1,1}, 0 },
}

func TestAverage(t *testing.T) {
  for _, pair := range tests {
    v := Average(pair.values)
    if v != pair.average {
      t.Error(
        "For", pair.values,
        "expected", pair.average,
        "got", v,
      )
    }
  }
}
```

This is a very common way to setup tests (abundant examples can be found in the source code for the packages 
included with Go). We create a struct to represent the inputs and outputs for the function. Then we create a 
list of these structs (pairs). Then we loop through each one and run the function.

### Problems
- Writing a good suite of tests is not always easy, but the process of writings tests often reveals more about 
a problem then you may at first realize. For example, with our Average function what happens if you pass in 
an empty list ([]float64{})? How could we modify the function to return 0 in this case?
- Write a series of tests for the Min and Max functions you wrote in the previous chapter.