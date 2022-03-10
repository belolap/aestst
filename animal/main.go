package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var asyncFlag = flag.Bool("a", false, "включает асинхронный режим работы")

// AnimalNumber добавляет поддержку номера животного.
type AnimalNumber struct {
	Num int
}

// Number возвращает номер животного.
func (an AnimalNumber) Number() int {
	return an.Num
}

// Cat описывает кошку.
type Cat struct {
	AnimalNumber // Встраивание номера животного.
}

// NewCat создаёт кошку.
func NewCat(num int) Animal {
	return Cat{AnimalNumber: AnimalNumber{Num: num}}
}

// Sound мяукает.
func (с Cat) Sound() string {
	return "Мяу!"
}

// Dog описывает собаку.
type Dog struct {
	AnimalNumber // Встраивание номера животного.
}

// NewDog создаёт собаку.
func NewDog(num int) Animal {
	return Dog{AnimalNumber: AnimalNumber{Num: num}}
}

// Sound гавкает.
func (d Dog) Sound() string {
	return "Гав!"
}

// Animal описывает поведение животного.
// В Go, в отличие от C++, Java и т.д., роль интерфейса другая, хотя именно
// в данном примере оба подхода внешне выглядят одинаково. Структуры
// (в данном случае можно сказать классы) Cat, не реализуют интерфейс Animal.
// Интерфейс Animal напротив, описывает то, чему должна соответствовать структура,
// чтобы она могла быть передана в методы noise*. В традиционых Java/C++ подходах
// мы бы сказали, что Dog и Cat реализуют интерфейс Animal.
type Animal interface {
	Number() int
	Sound() string
}

// noise заставляет всех животных подать голос по очереди.
func noise(animals []Animal) {
	for _, a := range animals {
		fmt.Printf("%3d: %s\n", a.Number(), a.Sound())
	}
}

// noiseAsync заставляет всех животных шуметь одновременно.
func noiseAsync(animals []Animal) {
	// Мьютекс для синхронизации вывода в stdout.
	var mu sync.Mutex

	// WorkGroup нужен для ожидания завершения всех горутин.
	var wg sync.WaitGroup

	// Инициализируем счётчик горутин.
	wg.Add(len(animals))

	for _, a := range animals {
		// Запускаем горутину, которая просит животное подать голос.
		go func(a Animal) {
			// У нас вывод производится в stdout, на всякий случай блокируем
			// мьютекс, чтобы выводы двух животных не пересеклись.
			mu.Lock()
			defer mu.Unlock()

			// Выводим номер животного и его голос.
			fmt.Printf("%3d: %s\n", a.Number(), a.Sound())

			// Помечаем горутину, как завершившуюся, уменьшая счётчик wg.
			wg.Done()
		}(a)
	}

	// Дожидаемся завершения всех горутин.
	wg.Wait()
}

func main() {
	// Обрабатываем параметры командной строки.
	flag.Parse()

	// Инициализируем генератор псевдослучайных чисел.
	rand.Seed(time.Now().UnixNano())

	// Заполняем наш зоопарк случайным образом.
	var animals []Animal
	for i := 1; i < 11; i++ {
		// Для этого генерируем псевдослучайное число, проверяем его чётность
		// по последнему биту. В зависимости от чётности добавляем либо кошку,
		// либо собаку.
		if rand.Int()&1 == 0 {
			animals = append(animals, NewCat(i))
		} else {
			animals = append(animals, NewDog(i))
		}
	}

	// Ну, и немного пошумим.
	if *asyncFlag {
		// Шумят все сразу.
		noiseAsync(animals)
	} else {
		// Культурно подают голос по-очереди.
		noise(animals)
	}
}
