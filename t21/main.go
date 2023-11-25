/*
* Паттерн адаптер нужен для того, чтобы конвертирвать внешниий интерфейс (допустим внешняя библиотека) в тот,
* который нужен пользователю (программисту). Подходит для того, чтобы изменять и добавлять свой функционал,
* при этом не изменяя изначальнный код

* Основная информация про паттерн была взята из статьи Шевякова Ивана (https://habr.com/ru/articles/765468/)
* и https://refactoring.guru/design-patterns/adapter/go/example
 */

package main

import (
	"fmt"
	"t21/adapter"
	"t21/materials"
)

type Human struct{}

func (h Human) HitReaction() {
	fmt.Println("Oof")
}

func main() {
	reactions := []adapter.ReactionAdapter{adapter.NewSteelAdapter(&materials.Steel{}), adapter.NewWoodAdapter(&materials.Wood{}), &Human{}}
	for _, r := range reactions {
		r.HitReaction()
	}
}
