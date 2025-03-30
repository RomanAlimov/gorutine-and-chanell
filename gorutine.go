package main

import "fmt"

func gorutine() {
	exit := make(chan int) // передача горутины

	go func() { // анонимная функция превратилась в горутину
		for i := 0; i < 10; i++ { // цикл до 10
			fmt.Println(<-data) // 0-9 ответ and exit
		}
	}() // вызов функции горутины
}

func slo(data, ex chan int) { // функция и 2 горутины
	x := 0 // счетчик
	for {
		select { // свитчи
		case data <- x: // если кто-то читает горутину то:
			x += 1 // прибавляем 1 к счетчику
		case <-exit: // если не читают то выходим из кода
			fmt.Println("exit code")
			return
		}
	}
}
