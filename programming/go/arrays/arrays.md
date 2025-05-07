
# Lesson 7: Arrays in Go

## Brief Description of arrays
An array is simply a data structure that consists of a collection of elements of the same kind. It is a numbered sequence of elements of a specific length which is fixed and cannot grow or sink when declared.
In Go, arrays start from a zero-based index, must be of the same kind and have fixed length. Example; t [4]string contains 4 string vales.

## When to Use Arrays

-Arrays can be used when the exact number of elements is known ahead of time.
-They can be used when performance with minimal memory overhead is needed.
-They can also be used when working with embedded systems where the allocation of dynamic memory is not preferred.

Arrays are considered as building blocks hence slices are used more often in Go.


## Overview

This lesson demonstrates how to work with arrays in Go, including:

- Declaring and initializing arrays
- Setting and getting array values
- Using multi-dimensional arrays

## Code Summary

1. Declare an array of 5 integers:
   var m [5]int → initialized with zero values
    Output: emp: [0 0 0 0 0]

2. Assign a value:
   m[4] = 100 → sets the last index to 100
   Output: set: [0 0 0 0 100], get: 100

3. Check array length:
     Output: length: 5

4. Declare and initialize in one line:
   n := [5]int{1, 2, 3, 4, 5} -> full initialization
   Output: dcl: [1 2 3 4 5]

5. Shorthand with inferred length:
   n = [...]int{1, 2, 3, 4, 5}
   Output: dcl: [1 2 3 4 5]

6. Indexed initialization:
   n = [...]int{100, 3: 400, 500}
   Output: idx: [100 0 0 400 500]

7. 2D Array (Manual Fill):
   var twoD [2][3]int -> filled using nested loops
   Output: 2d: [[0 1 2] [1 2 3]]

8. 2D Array (Literal Initialization):
   twoD = [2][3]int{{1, 2, 3}, {1, 2, 3}}
   Output: 2d: [[1 2 3] [1 2 3]]

## How to Run:

  1. Save the code as: lesson7-arrays.go
  2. Open terminal and navigate to the folder:
     cd ~/Desktop/golang/Arrays
  3. Run the program:
     go run lesson7-arrays.go

## Output

  emp:  [0 0 0 0 0]
  set:  [0 0 0 0 100]
  get:  100
  length:  5
  dcl:  [1 2 3 4 5]
  dcl:  [1 2 3 4 5]
  idx:  [100 0 0 400 500]
  2d:  [[0 1 2] [1 2 3]]
  2d:  [[1 2 3] [1 2 3]]
