// Это иметация внешнего пакета со своим функционалом
package materials

import "fmt"

type Steel struct{}

type Wood struct{}

func (s *Steel) HitMetallicReaction() {
	fmt.Println("Metallic remble and steel is deformed")
}

func (w *Wood) HitWoodRaction() {
	fmt.Println("Dull sound")
}
