## My Go learnings

#### Notes for reference

* Where was the Go programming language designed? <br> Google

* Go is a compiled language

 * Go was created to support :
	1) the ease of programming of an interpreted, dynamically typed language
	2) provide the efficiency and safety of a statically typed, compiled language
	3) aimed to be modern, with support for networked and multicore computing.

* What is a package in go? <br>
   A collection of files that live in the same directory and have the same package statement at the beginning.

* In a go program, where are packages declared? <br> At the start of the program

* what is the entry point in a go program? <br> The main function

* Which package consists of main() function? <br> package main

* Type checking occurs at compile time for statically typed languages. Type checking occures at run time for dynamically type languages. This is the main difference between statically and dynamically typed programming languages.

* In dynamic typing, a variable is allowed to change its data type.

* In static typing, a variable is not allowed to change its data type.

* Running  ``` go mod init my-go-code ``` would generate a  ```go.mod``` file, which has module name and go version details.

* In go, package is a way of organizing and reusing code.

* You can have multiple go source files with same package name in a directory, and they will be treated as a single package by the go compiler.

* There are many predefined packages and you can use them using ```package``` keyword.