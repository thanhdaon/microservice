package utils

var exists = struct{}{}

type Set struct {
	m map[string]struct{}
}

func NewSet() *Set {
	s := &Set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *Set) Add(value string) {
	s.m[value] = exists
}

func (s *Set) Remove(value string) {
	delete(s.m, value)
}

func (s *Set) Contains(value string) bool {
	_, isContain := s.m[value]
	return isContain
}

func (s *Set) Iterator() map[string]struct{} {
	return s.m
}

func (s *Set) Clear() {
	s.m = make(map[string]struct{})
}
