package structural

import (
	"fmt"

	"github.com/taraslis453/design-patterns/patterns/structural/adapter"
	"github.com/taraslis453/design-patterns/patterns/structural/proxy"
)

func Structural() {
	fmt.Println("Structural ----------")
	fmt.Println("Adapter")
	adapter.Adapter()
	fmt.Println("Proxy")
	proxy.RunProxy()
}
