# 🧪 Test-Driven Development in Go

Welcome to my journey of learning **Test-Driven Development (TDD)** in **Go (Golang)**! This project is based on the amazing [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) by Chris James.

## 📚 What’s This Repo?

This repo showcases my hands-on practice with Go fundamentals through the lens of TDD. Each directory here represents a mini-project or concept covered in the book, fully covered with tests *before* code—because writing tests first is how we roll. 😎

## Repo Structure 
<pre> ``` TDD/ Arrays/ Arrays.go Arrays_test.go Concurrency/ concurrency.go concurrency_test.go Dependency_Injection/ dependency.go dependency_test.go HelloWorld/ home.go home_test.go Integers/ integer.go integer_test.go Iteration/ Iteration.go Iteration_test.go Maps/ maps.go maps_test.go Maths/ Clockace/ clock.svg main.go Clock.go Clock_test.go Clock_acceptance_test.go Mocking/ mocking.go mocking_test.go Pointers_Errors/ pointer.go pointer_test.go Property_based_tests/ Roman_numbers.go Roman_numbers_test.go.go Reflection/ reflection.go reflection_test.go Select/ select.go select_test.go Struct/ struct.go struct_test.go Sync/ sync.go sync_test.go Dockerfile go.mod link.txt readme.txt test.go ``` </pre>

## 🚀 Topics Covered
|---------------------------------------------------------------------------------------|
| Folder                   | Description                                                |
|--------------------------|------------------------------------------------------------|
| `Arrays/`                | Exploring array operations and writing tests for them      |
| `Concurrency/`           | Using goroutines and channels to write concurrent programs |
| `Context/`               | Using Context to cancel the function un no longer needed   |
| `Dependency_Injection/`  | Understanding and applying dependency injection in Go      |
| `HelloWorld/`            | The humble beginning: Hello World test                     |
| `Integers/`              | Adding integers and writing simple tests                   |
| `Iteration/`             | Practicing loops with TDD                                  |
| `Maps/`                  | Using maps and testing key existence and values            |
| `Maths/`                 | Using Math's functions, creating Clock with SVJ/XML file   |
| `Mocking/`               | Learning how to use mocks to isolate test logic            |
| `Pointers_Errors/`       | Handling pointers and custom error types with tests        |
| `Property_based_tests/`  | Creating functions that uses the quick functionality in go |
| `Reflection/`            | Dynamically inspecting and manipulating values and types   |
| `Select/`                | Handling multiple channel operations with the `select`     |
| `Struct/`                | Defining and testing Go structs                            |
| `Sync/`                  | Using sync, waitgroup and mutex to stop race condition     |
| `Dockerfile`             | Containerizing the project (because why not?)              |
|--------------------------|------------------------------------------------------------|

> More folders and tests will be added as I progress through the chapters.

---

## 🧪 Run the Tests

To run all tests:

```bash
sudo docker build -t testing.demo .
sudo docker run --rm testing.demo
```
After function is finished run these comands to close the running image.

```bash

sudo docker rmi 'testing.demo'
