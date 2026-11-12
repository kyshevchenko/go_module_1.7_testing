package stringutil_test

// Примеры в стиле godoc: выполняются как тесты (go test),
// вывод сверяется с блоком "// Output:".

import (
	"fmt"

	"stringutil"
)

// ExampleReverse демонстрирует разворот строк на латинице и кириллице.
func ExampleReverse() {
	fmt.Println(stringutil.Reverse("hello"))
	fmt.Println(stringutil.Reverse("привет"))
	fmt.Println(stringutil.Reverse(""))

	// Output:
	// olleh
	// тевирп
	//
}

// ExampleDivide показывает успешное деление и обработку ошибки
// при делении на ноль.
func ExampleDivide() {
	result, err := stringutil.Divide(6, 3, nil)
	fmt.Println(result, err)

	_, err = stringutil.Divide(1, 0, nil)
	fmt.Println(err)

	// Output:
	// 2 <nil>
	// division by zero
}
