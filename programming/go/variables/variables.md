# Lesson 2 - Variables and how they are declared

## Brief description
A variable is a piece of memory that stores a value that can be changed. It is basically a container or a named storage location that holds a value. The main function of a variable is that it allows you to store and manipulate data.

Go automatically infers the type of initialized variables.
A Variable declared without a corresponding initialization are considered zero-valued. A variable can store any type; integer, string,float64, float32.
## When is a variable used
-A variable is used by the compiler to check type-correctness of function calls;
-It can be used to store user input;
-It stores value that can be re-used;
It aslo stores data to be used for calculations;
Used when there is a need to accept or return values in functions.

## Overview

This lesson is on the basics of variables, how to declare and initialize variables and some types of variables in Golang.


## Here are the topics covered

1. Types of variables
    - integer
    - boolean
    - string
    - float(float32 and float64). These numbers indicate the precision of both types. 32 has less precision witha about 7 decimal digits, while 64 has more precision with 15 decimal digits.

2. Declaring variables;
    - using the ":=" method eg. x:= (value)
    - using "var (name &type)" method eg. var a string= (value)
    - using "var (name)" method, eg. var b = (value)
    - and the "no-value" method which is zero-valued by default. eg. var x int

3. Initializing variables

    NOTE: nothing here?

4. Variable naming conventions
    - camelCase method eg.var userName string
    - avoiding single-letter variable names except in loops
    - avoiding underscores
    - being descriptive eg. userName, userId etc.

## Code summary

## How to Run
