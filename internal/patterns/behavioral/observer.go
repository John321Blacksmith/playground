// This package provides a
// basic model of the observer
// pattern
package behavioral

// MODEL

// abstraction of an observer
type Observer interface {
	Update()
}

// a concrete subject
// to be listened
type Subject struct {
	state     string
	observers []Observer
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) GetState() string {
	return s.state
}

func (s *Subject) SetState(newState string) {
	s.state = newState
	for _, o := range s.observers {
		o.Update()
	}
}

// a concrete observer
type StreamSubscribeUseCase struct {
	state   string
	subject Subject
}

func (o *StreamSubscribeUseCase) Update() {
	o.state = o.subject.GetState()
}

func (o *StreamSubscribeUseCase) SetSubject(subject Subject) {
	o.subject = subject
}
