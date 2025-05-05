# Lesson 3: Constants and Math in Go

## What is a constant?
Anything that is constant signifies something that does not change. Similarly in Go, a constant is a value name that doesn't change once it is declared or while the program runs.Thus, new values cannot be assigned to it.
They are very useful because they are only declared and initialized once but can be referred to as many times as required through the program. A constant can be of any type; integer, bollean, float, string etc. and cannot be defined using functions or variables.

## When to Use Constants:

-A constant can be used to define fixed configuration values that don't change during the execution of the program;
-It can be used for error messages or status codes;
-It is also used for any value that should be immutable;
-Again, constants can be used as label categories or options with readable names;

## overview
This lesson demonstrates the use of constants and the `math` package in Go.

## Concepts Covered

- Defining string and numeric constants using `const`
- Basic arithmetic with large numbers
- Type conversion (`float64` to `int64`)
- Using built-in math functions (e.g., `math.Sin()`)

## What the Program Does

1.Prints a string constant.
2.Divides a large float number by a constant and prints the result.
3.Converts the result to an integer and prints it.
4.Uses math.Sin() to compute and print the sine of a constant.

## Code summary

This Go program demonstrates:

- Declaring constants (string and numeric)
- Performing arithmetic operations with constants
- Using scientific notation (e.g., 9e20)
- Type conversion from float64 to int64
- Using the math package to calculate sine

## How to Run

- Open your terminal
- Navigate to the directory where the file is located:
     cd ~/Desktop/golang/lesson3-constants
- Run using the command:
    go run lesson3-constants.go

## Output

constant
6e+11
600000000000
-0.28470407323754404
