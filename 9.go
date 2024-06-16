// Условие задачи
// В системе есть очередь сообщений, а также 3 события, которые записываются в нее — X, Y и Z.
// Известно, что есть внешние сервисы, которые могут сделать запись событий в следующих парах — XY, XZ и YZ.
// Назовем их правильными парами. Вас попросили собрать аналитику из очереди.
// Для этого вы дождались, пока все сервисы завершат запись, и приостановили работу очереди.
//
// Сейчас очередь состоит из 𝑛 сообщений. Поскольку запись происходит асинхронно,
// события от сервиса в очереди могут лежать не последовательно,
// то есть после записи первого события одним сервисом, могло записаться событие от другого сервиса.
//
// Ваша задача — определить, возможно ли представить текущую очередь в виде набора из 2𝑛 правильных пар событий,
// где одно событие принадлежит только одной правильной паре.
//
// Входные данные
// Каждый тест состоит из нескольких наборов входных данных.
//
// Первая строка содержит целое число 𝑡 (1≤𝑡≤104) — количество наборов входных данных.
// Далее следует описание наборов входных данных.
//
// Первая строка каждого набора входных данных содержит четное целое число 𝑛 (2≤𝑛≤2⋅105) — количество сообщений в очереди.
//
// Вторая строка каждого набора входных данных содержит строку из 𝑛 символов X, Y и Z — сообщения из очереди,
// от самого первого, к самому последнему.
//
// Гарантируется, что сумма значений 𝑛 по всем наборам входных данных не превышает 2⋅105.
//
// Выходные данные
// Для каждого набора входных данных в отдельной строке выведите Yes, если это возможно. В противном случае выведите No.

package main

import (
	"fmt"
	"strings"
)

func isPossiblePairs(n int, messages string) string {
	countX := strings.Count(messages, "X")
	countY := strings.Count(messages, "Y")
	countZ := strings.Count(messages, "Z")

	if countX+countY == n || countY+countZ == n || countZ+countX == n {
		return "Yes"
	}

	return "No"
}

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		var messages string
		fmt.Scan(&messages)

		result := isPossiblePairs(n, messages)
		fmt.Println(result)
	}
}