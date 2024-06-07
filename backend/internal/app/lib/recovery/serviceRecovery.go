package recovery

import "fmt"

func RecoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("в ходе выполнения функции была поймана паника:", r)
	}
}
