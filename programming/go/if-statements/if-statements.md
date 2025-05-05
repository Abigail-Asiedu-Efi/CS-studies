# Lesson 5: If-Else Conditions in Go

## Brief description
An if statement is used to execute a specified block of code that meets a specific condition.Thus, the code is excuted only when a given condition is true.The else stament is used as a form of default code that should run when the condition given appears to be false.

## When to use if-staments
-If statement is used when checking input values and for data validation;
-Making decisions in flow;
-Error handling;
-As well as controlling logic in loops or functions


## Overview

This program explores conditional logic using if, else if, and else blocks in Go. It demonstrates checking for even/odd numbers, divisibility, logical OR operations, and scoped variable declaration in conditionals.

## Code Summary

  1. Even or Odd:
     - Checks if 9 is divisible by 2.
     - Output: "9 is an odd number"

  2. Divisibility:
     - Checks if 6 is divisible by 3.
     - Output: "6 is divisible by 3"

  3. Logical OR:
     - Checks if either 8 or 7 is even using `||` (OR).
     - Output: "Either 8 or 7 is even"

  4. Scoped variable in if:
     - Declares `num := 9` within the if statement scope.
     - Then checks:
       - If num < 0 → Negative
       - Else if num < 10 → Single digit
       - Else → Multiple digit
     - Output: "Num is a single digit"

## How to Run:
  1. Save the code as: lesson5-ifelse.go
  2. Open terminal and move into the folder:
     $ cd ~/Desktop/golang/lesson5-ifelse
  3. Run the code:
     $ go run lesson5-ifelse.go

## Output:
  9 is an odd number  
  6 is divisible by 3  
  Either 8 or 7 is even  
  Num is a single digit
