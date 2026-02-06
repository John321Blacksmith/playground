package structural

import (
	"fmt"
)

// MODEL
type Adaptee struct{}

func (a *Adaptee) ExistingMethod() {
	fmt.Println("Performing an existing method")
}

type Adaptor struct {
	adaptee *Adaptee
}

func NewAdaptor() *Adaptor {
	var adaptee *Adaptee
	return &Adaptor{adaptee: adaptee}
}

// Perform an old method
// With the incoming logic
func (a *Adaptor) ExpectedMethod() {
	a.adaptee.ExistingMethod()
}
