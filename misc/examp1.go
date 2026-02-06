package misc

import (
	"fmt"
	"strings"
)

func Test() {
	categories := map[string]map[string]bool{
		"sport": {
			"tire":      true,
			"фитн":      true,
			"gym":       true,
			"muscle":    true,
			"water":     true,
			"форма":     true,
			"спорт":     true,
			"вода":      true,
			"мышmuscle": true,
			"график":    true,
			"run":       true,
			"exercis":   true,
			"schedule":  true,
			"упражн":    true,
			"питат":     true,
			"nutri":     true,
			"здоров":    true,
			"sport":     true,
			"shape":     true,
			"health":    true,
			"rest":      true,
		},
		"home": {
			"famil":      true,
			"house":      true,
			"kitchen":    true,
			"квартир":    true,
			"home":       true,
			"убр":        true,
			"child":      true,
			"спать":      true,
			"care":       true,
			"children":   true,
			"dad":        true,
			"sleep":      true,
			"поряд":      true,
			"daughter":   true,
			"mom":        true,
			"restfamily": true,
			"clean":      true,
			"сад":        true,
			"pet":        true,
			"комн":       true,
			"garden":     true,
			"son":        true,
			"room":       true,
			"flat":       true,
			"party":      true,
			"отд":        true,
			"комат":      true,
			"дом":        true,
			"убир":       true,
			"cook":       true,
			"yard":       true,
		},
		"work": {
			"income":      true,
			"co-work":     true,
			"friend":      true,
			"routine":     true,
			"factor":      true,
			"рутина":      true,
			"hour":        true,
			"work":        true,
			"finish":      true,
			"tommorow":    true,
			"график":      true,
			"коллег":      true,
			"рабо":        true,
			"weekday":     true,
			"schedule":    true,
			"salary":      true,
			"босс":        true,
			"будни":       true,
			"зарплата":    true,
			"collegue":    true,
			"collegues":   true,
			"офис":        true,
			"yesterday":   true,
			"director":    true,
			"cрок":        true,
			"office":      true,
			"шеф":         true,
			"завод":       true,
			"chief":       true,
			"weekdaywork": true,
			"project":     true,
			"deadline":    true,
			"boss":        true,
			"доход":       true,
		},
		"cooking": {
			"ужин":       true,
			"вечеринк":   true,
			"boil":       true,
			"stew":       true,
			"слад":       true,
			"delicious":  true,
			"meal":       true,
			"dinner":     true,
			"праздн":     true,
			"tasty":      true,
			"kitchen":    true,
			"food":       true,
			"water":      true,
			"жар":        true,
			"raw":        true,
			"drink":      true,
			"meat":       true,
			"lunch":      true,
			"овощ":       true,
			"sauce":      true,
			"vegetables": true,
			"готов":      true,
			"eat":        true,
			"fried":      true,
			"завтрак":    true,
			"family":     true,
			"кухня":      true,
			"fry":        true,
			"мороз":      true,
			"frozen":     true,
			"sweet":      true,
			"breakfast":  true,
			"wash":       true,
			"свеж":       true,
			"party":      true,
			"сыр":        true,
			"фрукт":      true,
			" туш":       true,
			"freeze":     true,
			"cook":       true,
			"вар":        true,
			"обед":       true,
		},
		"study": {
			"уч":        true,
			"by heart":  true,
			"проект":    true,
			"книг":      true,
			"learn":     true,
			"аудитори":  true,
			"inform":    true,
			"знан":      true,
			"class":     true,
			"project":   true,
			"библиотек": true,
			"класс":     true,
			"librar":    true,
			"book":      true,
			"knowledge": true,
		},
		"family": {
			"сестра":      true,
			"мам":         true,
			"friend":      true,
			"mom":         true,
			"семь":        true,
			"mother":      true,
			"sister":      true,
			"дед":         true,
			"dad":         true,
			"father":      true,
			"brother":     true,
			"grandmother": true,
			"друг":        true,
			"сёстры":     true,
			"бабушк":      true,
			"дедушк":      true,
			"party":       true,
			"child":       true,
			"мать":        true,
			"grand":       true,
			"bro":         true,
			"друзья":      true,
			"брат":        true,
			"cousin":      true,
			"grandfather": true,
		},
	}

	sentenceData := map[string]bool{
		"форма":     true,
		"спорт":     true,
		"вода":      true,
		"мышmuscle": true,
		"график":    true,
		"run":       true,
		"квартир":   true,
		"children":  true,
		"комн":      true,
		"party":     true,
	}

	frequencies := map[string]int{}

	for category := range categories {
		for pattern := range categories[category] {
			for word := range sentenceData {
				if strings.Contains(word, pattern) {
					_, exists := frequencies[category]
					if !exists {
						frequencies[category] = 1
					}
					frequencies[category]++
				}
			}
		}
	}

	fmt.Println(frequencies)
}

type Container[T comparable] interface {
	String() string
	Values() []any
	Size() int
	Empty() bool
	Clear()
}

type Set[T comparable] interface {
	Intersection(sets ...T)
	Add(elements ...T)

	Container[T]
}

var itemExists bool

type RegularSet[T comparable] struct {
	data map[T]bool
}

func New[T comparable](elements ...T) *RegularSet[T] {
	result := &RegularSet[T]{data: make(map[T]bool)}
	if len(elements) > 0 {
		result.Add(elements...)
	}
	return result
}

func (s *RegularSet[T]) Add(elements ...T) {
	if len(elements) > 0 {
		for _, el := range elements {
			s.data[el] = itemExists
		}
	}
}

func (s *RegularSet[T]) Size() int {
	return len(s.data)
}

func (s *RegularSet[T]) Intersection(another *RegularSet[T]) *RegularSet[T] {
	result := New[T]()

	if s.Size() < another.Size() {
		for item := range s.data {
			if _, contains := another.data[item]; contains {
				result.Add(item)
			}
		}
	} else {
		for item := range another.data {
			if _, contains := s.data[item]; contains {
				result.Add(item)
			}
		}
	}

	return result
}
