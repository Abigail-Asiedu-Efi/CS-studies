# Lesson 2 - Variables and how they are declared

##Overview:
This lesson is on the basics of variables, how to declare and initialize variables and some types of variables in Golang.


##Here are the topics covered:
1.Types of variables
    -integer
    -boolean
    -string
    -float(float32 and float64- these numbers indicate the precision of both types.32 has less precision witha about 7 decimal digits, while 64 has more precision with 15 decimal digits)

2.Declaring variables;
    -using the ":=" method eg. x:= (value)
    -using "var (name &type)" method eg. var a string= (value)
    -using "var (name)" method, eg. var b = (value)
    -and the "no-value" method which is zero-valued by default. eg. var x int

3.Initializing variables

4.Variable naming conventions
    -camelCase method eg.var userName string
    -avoiding single-letter variable names except in loops
    -avoiding underscores
    -being descriptive eg. userName, userId etc.


 ##code summary
 package: main
 Function: main()

constant declared using const keyword; cannot be changed. Prints salutation: Hello, learners!

string variable explicitly declared with type(string) and initialized. Prints to console: Programming

multiple variables(b,c,d) declared and initialized on a single line

short declaration made(:=);type is automatically infered in Go. 
prints: false to console

default zero-value declared on line 20; 0 is printed to console 

mixed-type explicit declaration on lines 23 and 24 
prints:Integer: 24
        Float: 2.56

Automatic inference of type(string) even when type was not explicitly declared; userName as string


##How to Run
 ''bash
    go run lesson2-variables.go


 ##Output
Hello, Learners!
Programming
2 7 6
false
0
Integer:  24
Float:  2.56
Abigail


