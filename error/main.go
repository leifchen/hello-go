package main

import (
	"errors"
	"fmt"
	"strconv"
)

// CommonError 结构
type CommonError struct {
	errorCode int    //错误码
	errorMsg  string //错误信息
}

func (ce *CommonError) Error() string {
	return ce.errorMsg
}

func main() {
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	sum, err := add(-1, 2)
	var cm *CommonError
	if errors.As(err, &cm) {
		fmt.Println("错误代码为:", cm.errorCode, "，错误信息为：", cm.errorMsg)
	} else {
		fmt.Println(sum)

	}

	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e)
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w))
	fmt.Println(errors.Is(w, e))

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	defer func() {
		fmt.Println("first defer...")
	}()
	defer func() {
		fmt.Println("second defer...")
	}()
	connectMySQL("", "root", "123456")
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &CommonError{
			errorCode: 1,
			errorMsg:  "a或者b不能为负数"}
	}
	return a + b, nil
}

func connectMySQL(ip, username, password string) {
	if ip == "" {
		panic("ip不能为空")
	}
	//省略其他代码
}
