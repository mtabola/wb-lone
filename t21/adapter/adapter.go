package adapter

import "t21/materials"

// Создаем целевую структуру метода, чтобы адаптировать сторонние классы
type ReactionAdapter interface {
	HitReaction()
}

// Создаем структуры-адаптер для структур из внешнего пакета
type SteelAdapter struct {
	*materials.Steel
}

type WoodAdapter struct {
	*materials.Wood
}

// Конструкторы адаптеров
func NewSteelAdapter(s *materials.Steel) *SteelAdapter {
	return &SteelAdapter{s}
}

func NewWoodAdapter(w *materials.Wood) *WoodAdapter {
	return &WoodAdapter{w}
}

// Создаем методы на каждую структуру-адаптер, чтобы имя метода соответствовало имени, которое мы указали в интерфейсе адаптера
func (sa *SteelAdapter) HitReaction() {
	sa.Steel.HitMetallicReaction()
}

func (wa *WoodAdapter) HitReaction() {
	wa.Wood.HitWoodRaction()
}
