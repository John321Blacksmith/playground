// This package contains
// a model of the factory

package creational

import "fmt"

// MODEL

// Table of actors for
// factoring based on the key
var Options map[string]Doer = map[string]Doer{
	"firstDoer":  FirstDoer{},
	"secondDoer": SecondDoer{},
	"thirdDoer":  ThirdDoer{},
}

// Declared basic behavior
// for all mentioned actors
type Doer interface {
	DoSomething()
}

// Some different actors
// with similar behavior
type FirstDoer struct{}

func (doer FirstDoer) DoSomething() {
	fmt.Println("First doer does its job")
}

type SecondDoer struct{}

func (doer SecondDoer) DoSomething() {
	fmt.Println("Second doer does its job")
}

type ThirdDoer struct{}

func (doer ThirdDoer) DoSomething() {
	fmt.Println("Third doer does its job")
}

// An entity which creates
// a Doer based on condition
type DoerFactory struct {
	opts map[string]Doer
}

func NewDoerFactory(opts map[string]Doer) *DoerFactory {
	return &DoerFactory{
		opts: opts,
	}
}

// This method takes the
// option provided by the
// client and returns a
// corresponding instance
func (f *DoerFactory) CreateDoer(opt string) Doer {
	return f.opts[opt]
}

func main() {
	factory := NewDoerFactory(Options)
	thirdDoer := factory.CreateDoer("firstDoer")
	thirdDoer.DoSomething()
}
