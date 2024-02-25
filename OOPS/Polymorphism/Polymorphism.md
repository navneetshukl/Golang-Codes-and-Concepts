# Polymorphism

Polymorphism is a core concept in object-oriented programming. It allows values of different data types to be handled using a uniform interface. There are two types of polymorphism:

## 1. Compile Time Polymorphism

In this type of polymorphism, the compiler knows which exact functions will be executed for a particular call. Some examples of compile time polymorphism include:

- **Function Overloading**: More than one method/function exists with the same name but with different signatures or possibly different return types.

- **Operator Overloading**: The same operator is used for operating on different data types.

# Overloading in Go

Go does not directly support method, function, or operator overloading, which are common features in many object-oriented languages. However, Go provides other ways to achieve similar outcomes.

## Method Overloading

Method overloading is not supported in Go. If you try to declare two methods with the same name but different parameters (as in Program-1), the compiler will throw an error. Here's an example of such an error:

```plaintext
.\main.go:11:16: method math.add already declared at .\main.go:7:16
.\main.go:18:23: too many arguments in call to mt.add
        have (number, number, number)
        want (int, int)

```

# Method Overloading in Go

In Go, method overloading can be achieved using variadic functions. A variadic function is a function that can be called with varying numbers of arguments. 


The Program-2 is correct way of implementing method overloading in golang.

## 2. Run Time Polymorphism

RunTime Polymorphism means that a call is resolved at runtime. It is achieved in GO using interfaces. 