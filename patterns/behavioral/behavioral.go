package behavioral

import (
	"fmt"

	"github.com/taraslis453/design-patterns/patterns/behavioral/strategy"
	templatemethod "github.com/taraslis453/design-patterns/patterns/behavioral/template-method"
)

func Behavioral() {
	fmt.Println("Behavioral")
	templatemethod.TemplateMethod()
	strategy.Strategy()
}
