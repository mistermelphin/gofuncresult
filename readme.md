Go Function Result
==================

`gofuncresult` is a Go library that simplifies handling function results by wrapping them into structures with data and errors. It helps to reduce the boilerplate code associated with error handling and increases code readability, allowing you to work with functions more conveniently.

## Installation

Install the library using go get:

```bash
go get github.com/mistermelphin/gofuncresult
```

Then include it in your project:

```go
import (
    result "github.com/mistermelphin/gofuncresult"
)
```

## Usage

### Wrapping Results

Example using the `Wrap` function, which wraps the result of a function:

```go
func DoSomething(ok bool) (string, error) {
    if ok {
        return "example", nil
    } else {
        return "", fmt.Errorf("argument is not ok")
    }
}
func main(){
    res := result.Wrap(DoSomething(true))
    if res.IsError() {
        fmt.Println("1. Error:", res.Err())
    }
    fmt.Println("2. Value:", res.Value())
    // Output:
    // 2. Value: example
}
```

### Simple Value Handling

Simplified value handling using the `Value` structure:
```go
func DoSomething(ok bool) result.Value[string] {
    if ok {
        return result.NewValue("example", nil)
    } else {
        return result.NewValue("", fmt.Errorf("argument is not ok"))
    }
}

func main(){
    res := DoSomething(true)
    res.PanicIfError()
    if res.IsOk() {
        fmt.Println("1. OK:", res.Value())
    }

    res = DoSomething(false)
    if res.IsError() {
        fmt.Println("2. Error:", res.Err())
    }

    v, err := DoSomething(true).Unwrap()
    if err != nil {
        fmt.Println("3. Not ok")
    } else {
        fmt.Println("3. OK:", v)
    }
    // Output:
    // 1. OK: example
    // 2. Error: argument is not ok
    // 3. OK: example
}
```

### Localized Results

Using the library to create localized structures:
```go
type PackageResult result.Value[string]

func DoSomething(ok bool) PackageResult {
    if ok {
        return result.NewValue("package", nil)
    }
    return result.NewValue("", fmt.Errorf("not ok"))
}

func main(){
    fmt.Println("1.", DoSomething(true).Must())
    // Output:
    // 1. package
}
```


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.