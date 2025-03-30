package main

import (
	"fmt"
	"time"
)

// работа горутин и каналов наглядный пример.

func main() {
	ch := make(chan string) // передача горутины
	go sayhel(ch)           //
	s := <-ch               // передаем канал в переменную
	fmt.Println(s)          // вывод s
}

func say(word string) { // функция с печатыванием переменной
	fmt.Println(word) // 5 раз hello
}

func sayhel(ex chan string) {
	for i := 0; i < 5; i++ { // от 0 до 5
		time.Sleep(100 * time.Millisecond) // замедляет вывод текста
		say("hello")                       // передаем значение в функцию
	}
	ex <- "s" // передаем значение s в канал
}
