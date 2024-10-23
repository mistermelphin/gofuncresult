package result_test

import (
	"fmt"

	result "github.com/mistermelphin/gofuncresult"
)

func ExampleWrap() {
	f := func(ok bool) (string, error) {
		if ok {
			return "example", nil
		} else {
			return "", fmt.Errorf("argument is not ok")
		}
	}

	res := result.Wrap(f(true))
	if res.IsError() {
		fmt.Println("1. Error:", res.Err())
	}
	fmt.Println("2. Value:", res.Value())
	//Output:
	//2. Value: example
}

func ExampleValue() {
	f := func(ok bool) result.Value[string] {
		if ok {
			return result.NewValue("example", nil)
		} else {
			return result.NewValue("", fmt.Errorf("argument is not ok"))
		}
	}
	res := f(true)
	res.PanicIfError()
	if res.IsOk() {
		fmt.Println("1. OK:", res.Value())
	}
	res = f(false)
	if res.IsError() {
		fmt.Println("2. Error:", res.Err())
	}

	v, err := f(true).Unwrap()
	if err != nil {
		fmt.Println("3. Not ok")
	} else {
		fmt.Println("3. OK:", v)
	}
	//Output:
	//1. OK: example
	//2. Error: argument is not ok
	//3. OK: example
}

func ExampleLocalizedResult() {
	type PackageResult result.Value[string]

	f := func(ok bool) PackageResult {
		if ok {
			return result.NewValue("package", nil)
		}
		return result.NewValue("", fmt.Errorf("not ok"))
	}
	fmt.Println("1.", f(true).Must())
	//Output:
	//1. package
}
