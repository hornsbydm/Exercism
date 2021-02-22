Exercism Mentoring Comments
===========================

## 1. Hello World

## 2. Two-fer

Nicely done. Good use of Sprintf(). Try refactoring this code block to have a `func pronoun(s string) string` that does the empty string test for you.

---

Good work! It passes testing. I noticed a lot of spacing and commenting abnormalities. Check out go fmt and golint These tools can really help point out these issues.

A good editor uses these tools automatically for you. I personally use VSCode.

---

In Go parameters are passed by value and therefore it is idiomatic to modify a parameter that has been passed into the func. However, do be sure to realize that once you modify the parameter you no longer have access to the original value. In a more complex function this could cause unexpected errors and bugs. Instead consider temporary variables, named return parameters, factoring the logic out into a separate func, or early returns to name a few!

## 3. Hamming

Try running `golint hamming.go`. Getting into the habit of including comments is very important. Comments are just code for programmers.

---

Try refactoring your code to use Go's [named return values](https://tour.golang.org/basics/7). This simplifies your return statements and cuts down on temporary variables.

---

Go provides a more concise way of iterating over items through its [`range`](https://gobyexample.com/range) function. This can help clean up your loop and prevent off-by-one bugs.

> ***Note:*** The range function is particularly useful as it can return both the index of an item and its value. In the form of `for index, value := range a {`.

## 4. Raindrops

This is a very nice response! Good use of maps and slices. However, programming is often about finding the simplest working solution. 

Try refactoring your code using a switch/case statement and Go's named return values.

## 5. Scrabble Score

This is actually an impressive response! It works but it is tremendously complex and inefficient.

- [] Run your code with go test -bench ..

- [] Try refactoring your code using a switch case statement to simplify your code.

- [] Run your benchmark again and see how your changes affected benchmarking.

Submit your code and benchmarking results.



## Leap (Side)

Using nested ifs is highly discouraged as it creates unnecessary whitespace by indenting. It also makes your code very difficult to read. 

Perhaps you could refactor your the body of your function to a single line e.g. `return ...`.

>***Hint:*** Look at using Go's [boolean operators](https://golang.org/ref/spec#Operators_and_punctuation) to link the conditionals together in one statement.

***

Very Nicely Done 👍

***

Given that any logical expression returns a bool. Would it be possible to refactor your IsLeapYear() to be a single line of code consisting of return ...?

---

Do you think you could write the body of this function as a single line of code?

return ...

Does this make it easier to read?

---

## Space Age
Good work! Selecting a switch/case structure makes a lot of sense here. Since most of the data here is constant and we will only ever need one at a time it makes much more sense to use a switch over a map here.

---

I do like that you made the math easy to read. There will be no runtime penalty for this as the compiler will simplify this for us.

---

Be sure to realize that you are not actually using type Planet. Go uses a concept called "Duck Typing". Essentially if it acts like a duck and quacks like a duck it must be a duck. That is why you can use string literals in place of type Planet.

---

This is a nice textbook example of a clean solution to this exercise.

Would it make more sense to use a switch/case structure here instead of a map?

---

This is a good solution! However, do be sure to run `go test -bench . -benchmem` having the map in the func scope means that each iteration you will building the map in memory and tearing it down. This can be expensive to do. 

Since the data is constant and you will only ever need to access 1 planet at a time consider rewriting this using a switch/case structure.