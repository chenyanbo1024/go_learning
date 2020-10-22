package err_test

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less than 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	} else if n > 100 {
		return nil, LargerThanHundredError
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-50); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less.")
		} else if err == LargerThanHundredError {
			fmt.Println("It is larger")
		}
	} else {
		t.Log(v)
	}
}
