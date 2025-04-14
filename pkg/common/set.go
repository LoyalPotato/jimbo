package common

// Ref: https://www.bytesizego.com/blog/set-in-golang

type Set map[string]struct{}

func NewSet() *Set {
	s := make(Set)
	return &s
}

// NOTE: As maps are already refernce values, this won't create another copy
// It can't be *Set because then
func (s Set) Add(elem string) {
	s[elem] = struct{}{}
}

func (s Set) Remove(elem string) {
	delete(s, elem)
}

func (s Set) Contains(elem string) bool {
	_, has := s[elem]
	return has
}

// List returns all elements in the set as a slice
func (s Set) List() []string {
	keys := make([]string, 0, len(s))
	for key := range s {
		keys = append(keys, key)
	}
	return keys
}
