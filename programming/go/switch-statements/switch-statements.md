# Lesson 6: Switch Statements in Go

# Brief description
A switch statement is a type of control flow that allows the execution of blocks of codes depending on the value of a variable.

## When is switch-statement used
- A switch-statement is used when comparing a single value against several possible values;
- It is used when there are multiple conditions and they need to be organised clearly;
-It is also used to simplify if-else statements.

## Overview

This lesson explores Go’s switch statements with various use cases:

- Basic value matching
- Multiple conditions in a case
- Switch without an expression (if-else style)
- Type switching using interfaces

## Code Summary

1. Basic switch:

   - Variable i := 2
   - Switch checks i:
      - case 1: prints "one"
      - case 2: prints "two"
      - case 3: prints "three"
   Output: "Write 2 as two"

2. Switch on time

   - Checks the current day:
   - If Saturday or Sunday -> "It's the weekend"
   - Else -> "It's a weekday"

3. Switch without condition:
   - Checks current hour:
     - If before 12 -> "It's before noon"
     - Else -> "It's after noon"

4. Type switch using interface{}:
   - whatAmI(true)   -> "I'm a bool"
   - whatAmI(1)      -> "I'm an int"
   - whatAmI("hey")  -> "Don't know type string"

## How to Run

  1. ctrl + s to save the code as: lesson6-switch.go
  2. In terminal, navigate to your folder:
     cd ~/Desktop/golang/Switch
  3. Run the code:
     go run lesson6-switch.go

## Output (varies slightly with time/day)

  Write 2 as two  
  It's a weekday  ("It's the weekend")  
  It's before noon  ("It's after noon")  
  I'm a bool  
  I'm an int  
  Don't know type string
