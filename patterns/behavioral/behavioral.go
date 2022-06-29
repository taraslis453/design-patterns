package behavioral

import (
	"fmt"

	"github.com/taraslis453/design-patterns/patterns/behavioral/observer"
	"github.com/taraslis453/design-patterns/patterns/behavioral/strategy"
	templatemethod "github.com/taraslis453/design-patterns/patterns/behavioral/template-method"
)

func Behavioral() {
	fmt.Println("Behavioral ------------")
	fmt.Println("Template Method")
	templatemethod.TemplateMethod()
	fmt.Println("Strategy")
	strategy.Strategy()
	fmt.Println("Observer")
	observer.RunObserver()
}
