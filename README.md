# 🧪 Test-Driven Development in Go

Welcome to my journey of learning **Test-Driven Development (TDD)** in **Go (Golang)**! This project is based on the amazing [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/) by Chris James.

## 📚 What’s This Repo?

This repo showcases my hands-on practice with Go fundamentals through the lens of TDD. Each directory here represents a mini-project or concept covered in the book, fully covered with tests *before* code—because writing tests first is how we roll. 😎

## Repo Structure 
<details>
  <summary><strong>📁 Arrays</strong></summary>

  - Arrays.go  
  - Arrays_test.go
</details>
<details>
  <summary><strong>📁 Blogposts</strong></summary>

  - blogposts.go  
  - blogposts_test.go
</details>

<details>
  <summary><strong>📁 Concurrency</strong></summary>

  - concurrency.go  
  - concurrency_test.go
</details>

<details>
  <summary><strong>📁 Dependency_Injection</strong></summary>

  - dependency.go  
  - dependency_test.go
</details>

<details>
  <summary><strong>📁 HelloWorld</strong></summary>

  - home.go  
  - home_test.go
</details>

<details>
  <summary><strong>📁 Integers</strong></summary>

  - integer.go  
  - integer_test.go
</details>

<details>
  <summary><strong>📁 Iteration</strong></summary>

  - Iteration.go  
  - Iteration_test.go
</details>

<details>
  <summary><strong>📁 Maps</strong></summary>

  - maps.go  
  - maps_test.go
</details>

<details>
  <summary><strong>📁 Maths</strong></summary>

  <details>
    <summary>📂 Clockace</summary>

    - clock.svg  
    - main.go
  </details>

  - Clock.go  
  - Clock_test.go  
  - Clock_acceptance_test.go
</details>

<details>
  <summary><strong>📁 Mocking</strong></summary>

  - mocking.go  
  - mocking_test.go
</details>

<details>
  <summary><strong>📁 Pointers_Errors</strong></summary>

  - pointer.go  
  - pointer_test.go
</details>

<details>
  <summary><strong>📁 Property_based_tests</strong></summary>

  - Roman_numbers.go  
  - Roman_numbers_test.go.go
</details>

<details>
  <summary><strong>📁 Reflection</strong></summary>

  - reflection.go  
  - reflection_test.go
</details>

<details>
  <summary><strong>📁 Select</strong></summary>

  - select.go  
  - select_test.go
</details>

<details>
  <summary><strong>📁 Struct</strong></summary>

  - struct.go  
  - struct_test.go
</details>

<details>
  <summary><strong>📁 Sync</strong></summary>

  - sync.go  
  - sync_test.go
</details>

<details>
  <summary><strong>📁 Root Files</strong></summary>

  - Dockerfile  
  - go.mod  
  - link.txt  
  - readme.txt  
  - test.go
</details>


## 🚀 Topics Covered

| Folder                   | Description                                                |
|--------------------------|------------------------------------------------------------|
| `Arrays/`                | Exploring array operations and writing tests for them      |
| `Blogposts/`             | Using fstest to test and read files in memory              |
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
| `docker-compose.yml`     | Can containerizing multiple containers if need be          |


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
```
To run all tests via compose:

```bash
sudo docker-compose run --rm app
```
After function is finished run these comands to close the running image.

```bash
sudo docker rmi 'tdd.testing'
```
