//https://github.com/MaksimDzhangirov/oop-go/blob/master/docs/inheritance.md
//https://medium.com/german-gorelkin/embedding-in-go-6739e46c1be1
//Данные источники помогли мне понять более детально как реализовано наследование в golang

package main

import (
	"fmt"
	"time"
)

type Human struct {
	name         string
	stamina      uint16
	intellegence uint16
	strength     uint16
}

// наследуем все поля от класса human и добавляяем счетчик сна
type Action struct {
	Human
	sleepCount uint8
}

// интерфейс, который служит переключателем между родительской функцией sleep и дочерней
type Sleeper interface {
	Sleep()
}

func (h *Human) Sleep() {
	h.stamina = 100
	fmt.Println("Stamina is recovered: stamina = ", h.stamina)
}

func (a *Action) ReadABook() {
	if a.stamina < 5 {
		fmt.Println("Your tired, go to sleep")
		return
	}
	a.stamina -= 5
	a.intellegence += 5
	fmt.Printf("You stay more well-read: intellegence = %d, stamina = %d\n", a.intellegence, a.stamina)
}

func (a *Action) MakeAWorkout() {
	if a.stamina < 25 {
		fmt.Println("Your tired, go to sleep")
		return
	}
	a.stamina -= 25
	a.strength += 5
	a.sleepCount++
	fmt.Printf("You stay more strength: strength = %d stamina = %d\n", a.intellegence, a.stamina)
}

func (a *Action) Sleep() {
	if a.sleepCount == 0 {
		fmt.Println("Go to spend a stamina")
		return
	}
	time.Sleep(time.Duration(a.sleepCount) * time.Second)
	a.stamina += 10 * uint16(a.sleepCount)
	fmt.Printf("You sleep %d hours. Stamina: %d\n", a.sleepCount, a.stamina)
	a.sleepCount = 0
}

// создаем функцию sleep которая принимает наш родительский или дочерний класс, для того, чтобы потом вызвать функцию sleep, которая определена в классе
func Sleep(s Sleeper) {
	s.Sleep()
}

func main() {
	hum := &Human{"Richard", 50, 0, 0}
	act := &Action{*hum, 2}

	act.MakeAWorkout()
	act.ReadABook()
	act.ReadABook()
	act.MakeAWorkout()
	Sleep(act)

	act.MakeAWorkout()
	act.ReadABook()
	act.ReadABook()
	Sleep(hum)
}
