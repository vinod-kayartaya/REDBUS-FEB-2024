# Unit Testing

- [Basics of testing in Go](#Basics_of_testing_in_Go)
- [Writing unit tests using the `testing` package](#Writing_unit_tests_using_the_testing_package)
- [Running tests using `go test`](#Running_tests_using_go_test)
- [Table-driven tests for multiple input scenarios](#Table-driven_tests_for_multiple_input_scenarios)
- [Generating HTML report for coverage](#Generating_HTML_report_for_coverage)
- [Mocking in Go using interfaces](#Mocking_in_Go_using_interfaces)
- [Code coverage and profiling with `go test`](#Code_coverage_and_profiling_with_go_test)
- [Best practices for writing effective tests](#Best_practices_for_writing_effective_tests)

<div id="Basics_of_testing_in_Go"></div>

## Basics of testing in Go

Testing in Go is a fundamental aspect of writing robust and reliable code. The Go programming language comes with a built-in testing framework, making it easy to write tests for your code. Here are the basics of testing in Go:

1. **Testing Package (`testing`)**:
   Go has a built-in package called `testing` specifically designed for writing tests. This package provides functions and utilities for writing and running tests.

2. **Test Functions**:
   In Go, test functions are regular functions whose names start with `Test`. These functions reside in `_test.go` files within the same package as the code being tested.

   Example of a test function:

   ```go
   func TestAdd(t *testing.T) {
       result := add(2, 3)
       if result != 5 {
           t.Errorf("Expected 5, got %d", result)
       }
   }
   ```

3. **Test Helpers**:
   Test helpers are auxiliary functions or utilities that aid in testing. They are often used to set up test data or perform common assertions.

4. **Test Main Function**:
   If you need to perform setup or teardown tasks before or after running tests, you can define a special function called `TestMain` in your test file.

   Example:

   ```go
   func TestMain(m *testing.M) {
       // Perform setup tasks here

       // Run tests
       exitCode := m.Run()

       // Perform teardown tasks here

       // Exit with the appropriate exit code
       os.Exit(exitCode)
   }
   ```

5. **Test Assertions**:
   Go provides various assertion functions like `t.Errorf`, `t.Fatalf`, and `t.Fail` to report test failures. Additionally, the `testing` package provides functions such as `t.Error`, `t.FailNow`, `t.Fatal`, and `t.Log` for different types of test outcomes.

6. **Running Tests**:
   To run tests, you can use the `go test` command followed by the package path containing your test files. By default, `go test` runs all the tests in the current package.

   Example:

   ```
   go test
   ```

7. **Benchmark Tests**:
   Apart from regular tests, Go also supports benchmark tests. Benchmark tests measure the performance of functions or code snippets.

   Example:

   ```go
   func BenchmarkAdd(b *testing.B) {
       for i := 0; i < b.N; i++ {
           add(2, 3)
       }
   }
   ```

   To run benchmark tests:

   ```
   go test -bench=.
   ```

8. **Test Coverage**:
   Go also provides tools for measuring test coverage, which helps ensure that your tests are thorough and cover most of your codebase.

   To generate a coverage report:

   ```
   go test -cover
   ```

These are the basics of testing in Go. Writing tests alongside your code helps ensure its correctness, maintainability, and reliability over time.

<div id="Writing_unit_tests_using_the_testing_package"></div>

## Writing unit tests using the `testing` package

Let's go through an example of writing unit tests for a `factorial` function in Go using the `testing` package. The `factorial` function calculates the factorial of a non-negative integer. Here's the implementation of the `factorial` function:

```go
package mathutil

func Factorial(n int) int {
    if n < 0 {
        return -1 // Factorial is undefined for negative numbers
    }
    if n == 0 {
        return 1 // Factorial of 0 is 1
    }
    result := 1
    for i := 1; i <= n; i++ {
        result *= i
    }
    return result
}
```

Now, let's write unit tests for this function to cover various scenarios using the `testing` package:

```go
package mathutil

import "testing"

func TestFactorial(t *testing.T) {
    // Test case: Factorial of 0
    result := Factorial(0)
    if result != 1 {
        t.Errorf("Factorial(0) expected 1, got %d", result)
    }

    // Test case: Factorial of a positive number
    result = Factorial(5)
    if result != 120 {
        t.Errorf("Factorial(5) expected 120, got %d", result)
    }

    // Test case: Factorial of a negative number
    result = Factorial(-1)
    if result != -1 {
        t.Errorf("Factorial(-1) expected -1, got %d", result)
    }
}
```

In this example, we have covered the following scenarios:

1. **Factorial of 0**:
   We expect the factorial of 0 to be 1.

2. **Factorial of a positive number (5)**:
   We expect the factorial of 5 to be 120.

3. **Factorial of a negative number**:
   Factorial is undefined for negative numbers, so we expect the function to return -1.

To run these tests, you would execute `go test` in the directory containing the test file. Go will run all the test functions whose names begin with `Test`.

This is a basic example of unit testing in Go using the `testing` package. It's essential to cover various scenarios to ensure that the function behaves correctly under different conditions.

<div id="Running_tests_using_go_test"></div>

## Running tests using `go test`

To run the tests using the `go test` command, you need to navigate to the directory containing your Go package and execute the command. Here's how you would do it:

1. Navigate to the directory containing your Go package, which includes the source code file (`mathutil.go`) and the test file (`mathutil_test.go`).

2. Open a terminal or command prompt and navigate to that directory.

3. Run the `go test` command:

   ```
   go test
   ```

   This command will automatically search for test files (`*_test.go`) in the current directory and run all the test functions within those files.

4. After running the tests, Go will output the results. If all tests pass, you will see a message indicating success. If any tests fail, Go will provide information about which tests failed and why.

5. Additionally, Go provides options to customize test runs. For example, you can run tests with verbose output (`-v`), specify particular tests or packages to run, and generate coverage reports.

Here's the command to run tests with verbose output:

```
go test -v
```

And to generate coverage reports:

```
go test -cover
```

These commands provide more detailed information about the test execution and coverage statistics, respectively.

Remember to ensure that your code and test files are properly organized and named according to Go conventions for `go test` to work correctly.

<div id="Table-driven_tests_for_multiple_input_scenarios"></div>

## Table-driven tests for multiple input scenarios

Table-driven tests are a powerful technique in Go for testing functions with multiple input scenarios. Instead of writing separate test cases for each input, you can use a table of inputs and expected outputs to test the function against multiple cases in a concise and organized manner.

Here's how you can implement table-driven tests for the `factorial` function:

```go
package mathutil

import "testing"

func TestFactorial(t *testing.T) {
    // Define test cases as a table
    testCases := []struct {
        input    int
        expected int
    }{
        {0, 1},     // Factorial of 0
        {1, 1},     // Factorial of 1
        {5, 120},   // Factorial of 5
        {10, 3628800}, // Factorial of 10
        {-1, -1},   // Factorial of a negative number
    }

    // Iterate over test cases
    for _, tc := range testCases {
        // Run the test for each case
        t.Run(fmt.Sprintf("Factorial(%d)", tc.input), func(t *testing.T) {
            result := Factorial(tc.input)
            if result != tc.expected {
                t.Errorf("Factorial(%d) expected %d, got %d", tc.input, tc.expected, result)
            }
        })
    }
}
```

In this example:

- We define a slice of structs called `testCases`, where each struct contains an input value and the expected output.
- We iterate over each test case using a `for` loop.
- For each test case, we run a sub-test using `t.Run()`, which allows us to provide a descriptive name for each test case.
- Within the sub-test, we call the `Factorial` function with the input value and compare the result with the expected output.
- If the result doesn't match the expected output, we report an error using `t.Errorf()`.

This approach makes it easy to add new test cases and keeps the test code organized. When you run the tests using `go test`, each test case will be executed separately, and you'll get granular results for each case.

<div id="Generating_HTML_report_for_coverage"></div>

## Generating HTML report for coverage

In Go, you can generate an HTML page for code coverage using the built-in `go test` command with the `-coverprofile` flag followed by the `go tool cover` command. Here's how you can do it:

1. Run your tests with coverage and generate a coverage profile:

   ```
   go test -coverprofile=coverage.out
   ```

2. Once you have the coverage profile generated, you can use the `go tool cover` command to generate an HTML report:
   ```
   go tool cover -html=coverage.out -o coverage.html
   ```

This will generate a file named `coverage.html`, which you can open in your web browser to view the HTML page displaying the code coverage for your Go package.

<div id="Mocking_in_Go_using_interfaces"></div>

## Mocking in Go using interfaces

Mocking in Go using interfaces is a common technique for testing code that relies on external dependencies or collaborator objects. By defining interfaces for these dependencies, you can create mock implementations during testing to simulate different behaviors or responses without interacting with the real external systems. This approach allows you to isolate the unit of code being tested and make tests more deterministic and independent of external factors.

Here's a basic example to illustrate how mocking works using interfaces in Go:

Suppose you have a `EmailSender` interface and a function `SendWelcomeEmail` that depends on it:

```go
package myapp

type EmailSender interface {
    SendEmail(to, subject, body string) error
}

func SendWelcomeEmail(sender EmailSender, email string) error {
    subject := "Welcome to MyApp"
    body := "Welcome to MyApp! We're excited to have you on board."

    // Use the EmailSender to send the welcome email
    err := sender.SendEmail(email, subject, body)
    if err != nil {
        return err
    }

    return nil
}
```

Now, let's say you want to test the `SendWelcomeEmail` function without actually sending emails. You can create a mock implementation of the `EmailSender` interface for testing:

```go
package myapp

// MockEmailSender is a mock implementation of the EmailSender interface
type MockEmailSender struct {
    SentEmails []struct {
        To      string
        Subject string
        Body    string
    }
    ErrorToReturn error
}

// SendEmail is the implementation of the EmailSender interface for the mock
func (m *MockEmailSender) SendEmail(to, subject, body string) error {
    m.SentEmails = append(m.SentEmails, struct {
        To      string
        Subject string
        Body    string
    }{to, subject, body})

    return m.ErrorToReturn
}
```

With the `MockEmailSender` implementation, you can create instances of it in your test cases and pass them to the `SendWelcomeEmail` function instead of a real `EmailSender` implementation. This allows you to control the behavior of the email sending process during testing.

Here's an example test using the mock:

```go
package myapp_test

import (
    "testing"

    "myapp"
)

func TestSendWelcomeEmail(t *testing.T) {
    mockSender := &myapp.MockEmailSender{}

    email := "test@example.com"
    err := myapp.SendWelcomeEmail(mockSender, email)
    if err != nil {
        t.Errorf("SendWelcomeEmail returned an error: %v", err)
    }

    // Assert that the email was sent with the correct details
    if len(mockSender.SentEmails) != 1 {
        t.Errorf("Expected 1 email to be sent, got %d", len(mockSender.SentEmails))
    }
    sentEmail := mockSender.SentEmails[0]
    if sentEmail.To != email {
        t.Errorf("Expected email to be sent to %s, got %s", email, sentEmail.To)
    }
    if sentEmail.Subject != "Welcome to MyApp" {
        t.Errorf("Expected email subject to be 'Welcome to MyApp', got %s", sentEmail.Subject)
    }
    if sentEmail.Body != "Welcome to MyApp! We're excited to have you on board." {
        t.Errorf("Expected email body to be correct, got %s", sentEmail.Body)
    }
}
```

In this test, we're using the `MockEmailSender` to simulate sending a welcome email without actually sending it. We can then inspect the mock to verify that the `SendWelcomeEmail` function interacted with it correctly.

This example demonstrates how you can use interfaces and mocking in Go to write effective unit tests for code with external dependencies.

<div id="Code_coverage_and_profiling_with_go_test"></div>

## Code coverage and profiling with `go test`

Go's testing framework provides built-in support for code coverage and profiling. These features help you assess the effectiveness of your tests and identify performance bottlenecks in your code. Here's how you can use them with `go test`:

### Code Coverage

To measure code coverage, you can use the `-cover` flag with the `go test` command. This flag instructs Go to analyze your tests and report the percentage of code covered by them.

```bash
go test -cover ./...
```

- The `./...` argument tells `go test` to recursively test all packages in the current directory and its subdirectories.

When you run this command, Go will execute your tests and display a summary of code coverage at the end, indicating the percentage of code covered by your tests.

### Profiling

Go also supports profiling your tests to identify performance bottlenecks. You can profile CPU usage, memory allocation, and more using various profiling options.

#### CPU Profiling

To profile CPU usage, you can use the `-cpuprofile` flag with `go test`. This flag specifies the filename where the CPU profiling data should be written.

```bash
go test -cpuprofile cpu.prof ./...
```

After running this command, Go will execute your tests and generate a CPU profiling file named `cpu.prof`.

#### Memory Profiling

To profile memory allocation, you can use the `-memprofile` flag with `go test`. Similar to CPU profiling, this flag specifies the filename for the memory profiling data.

```bash
go test -memprofile mem.prof ./...
```

After running this command, Go will execute your tests and generate a memory profiling file named `mem.prof`.

#### Other Profiling Options

Go provides additional profiling options for specific use cases, such as block profiling (`-blockprofile`), mutex profiling (`-mutexprofile`), and tracing (`-trace`). You can find more details about these options in the `go test` documentation.

### Analyzing Profiles

Once you've generated profiling files, you can analyze them using various tools provided by the Go toolchain. For example:

- To visualize CPU or memory profiles, you can use `go tool pprof`:

  ```bash
  go tool pprof cpu.prof
  ```

- To generate a PDF or PNG visualization of a profile, you can run:

  ```bash
  go tool pprof -pdf cpu.prof > cpu.pdf
  ```

- You can also use third-party tools like `pprof`, `Graphviz`, or various profiling visualization web interfaces.

By using these profiling tools, you can gain insights into the performance characteristics of your code and optimize it accordingly.

These features make Go's testing framework robust and comprehensive, allowing you to ensure both the correctness and efficiency of your code.

<div id="Best_practices_for_writing_effective_tests"></div>

## Best practices for writing effective tests

Writing effective tests is crucial for ensuring the reliability, maintainability, and scalability of your codebase. Here are some best practices to consider when writing tests in Go:

### 1. Test Function Naming:

- Use descriptive names for test functions that clearly indicate what aspect of the code they are testing.
- Follow the convention of prefixing test function names with `Test`.

### 2. Table-Driven Tests:

- Use table-driven tests to cover multiple scenarios with different inputs and expected outputs.
- Organize test cases into tables for better readability and maintainability.

### 3. Test Coverage:

- Aim for high test coverage to ensure that most parts of your codebase are exercised by tests.
- Use the `-cover` flag with `go test` to generate coverage reports and identify areas with low coverage.

### 4. Isolation:

- Write tests that are isolated from each other and from external dependencies.
- Use interfaces and dependency injection to mock external dependencies and control test environments.

### 5. Test Organization:

- Organize tests into separate files or packages based on the code they are testing.
- Follow the same directory structure as your source code to make it easy to locate and manage tests.

### 6. Test Readability:

- Write clear and concise test cases with descriptive variable names and comments.
- Avoid complex assertions and use helper functions or assertion libraries when necessary.

### 7. Test Failures:

- Ensure that failing tests provide meaningful error messages that help diagnose the issue.
- Use `t.Errorf` or `t.Fatalf` to report failures with informative messages.

### 8. Test Coverage Analysis:

- Regularly analyze test coverage reports to identify areas of your codebase that lack sufficient test coverage.
- Prioritize adding tests for critical or frequently used code paths.

### 9. Benchmarking:

- Use benchmark tests (`Benchmark*` functions) to measure the performance of critical code paths.
- Benchmark against different input sizes and configurations to understand performance characteristics.

### 10. Continuous Integration:

- Integrate testing into your continuous integration (CI) pipeline to automatically run tests on every code change.
- Ensure that tests are fast and reliable to facilitate quick feedback during development.

### 11. Documentation:

- Document test cases using code comments to explain the purpose of each test and any special considerations.
- Use test coverage badges in your project's README to highlight the test coverage percentage.

### 12. Refactoring Tests:

- Refactor tests along with production code to maintain consistency and readability.
- Avoid duplicating test code by extracting common setup and teardown logic into helper functions.

By following these best practices, you can create effective tests that improve the quality and maintainability of your Go codebase while fostering a culture of confidence and reliability in your development process.
