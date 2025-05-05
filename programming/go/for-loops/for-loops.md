# Lesson 4: For Loops in Go

## Brief description on for-loops
A for loop is a repetition control structure used to repeat a block of code multiple times so far as the condition is true and terminates only when it becomes false.
It is the only looping construct in Go. Other loops such as while, do-while etc. are not used in Go like they used in other programming languages.
The for-loop has three parts; initialization(j ;= 1), condition(j <=5 )and advancement to next iteration(j++). 
e.g for j := 1; j <= 5; j++ {
    ....
}

## When to used for-loops
-For-loop is used tasks are needed to be repeated for a certain number of times.
-They are used for iterating over collections such as arrays and slices.
-They can also be used to scan through strings or files.


Note that, loops that are infinite (has no condition, initialization) excutes continuously until you break it.


## Overview

This program demonstrates the different types of for-loops available in Go. Go only has one looping construct—`for`—but it can be used in various ways: like a while loop, a traditional C-style loop, a range loop, and even infinite loops. You'll also see how `break` and `continue` control flow inside loops.

## Code Summary

This program showcases multiple ways to use for-loops in Go.

Basic While-style Loop:

- i := 1
- Runs while i <= 3
- Prints 1, 2, 3

Classic For Loop:

- for j := 0; j < 3; j++
- Prints 0, 1, 2

Range-style Loop with Integer:

- for i := range 3
- Prints "Range 0", "Range 1", "Range 2"
- Note: range over an int like this works in recent Go versions (Go 1.22+)

Infinite Loop with Break:

- Prints "loop" once, then breaks

Range Loop with Continue:

- for n := range 6
- Skips even numbers using `continue`
- Prints odd numbers: 1, 3, 5

## How to Run

- Press ctrl + s to save code
- Open terminal and navigate:
     cd ~/Desktop/golang/lesson4-loops
- Run code with command:
     go run lesson4-loops.go