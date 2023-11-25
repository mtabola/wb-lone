/*
* Инкапслуляция в golang достигается за счет того, что мы разбиваем структуры и методы к ним на отдельные пакеты
* Также в самих структурах мы указываем у свойств(полей), какие у них модификаторы доступа за счет их наиминования
* Идентификаторы с прописной буквы экспортируются. Прописная буква означает, что это экспортируемый идентификатор.
* Идентификаторы со строчной буквы не экспортируются. Строчные буквы указывают на то, что идентификатор не экспортируется и будет доступен только из того же пакета.
* Также все это относится не только к свойства структуры, но и к ее методам и ей самой
 */
package point

import "math"

type Point struct {
	x float64
	y float64
}

func New(x float64, y float64) Point {
	return Point{x, y}
}

func (p1 Point) DistanceCalculate(p2 Point) float64 {
	return math.Sqrt(math.Pow(p2.x-p1.x, 2) - math.Pow(p2.y-p1.y, 2))
}
