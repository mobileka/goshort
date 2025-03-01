// Package storetest provides utilities for testing the internal/store package.
package storetest

type StoreMock struct {
	URL    string
	Result bool
}

func (s *StoreMock) Get(_ string) (string, bool) {
	return s.URL, s.Result
}

func (s *StoreMock) Add(_, _ string) bool {
	return s.Result
}

func NewStoreMock(url string, result bool) *StoreMock {
	return &StoreMock{url, result}
}

// NewFailingStoreMock creates an MockStore that will always fail
// when trying to Get or Add something from/to it.
func NewFailingStoreMock(url string) *StoreMock {
	return NewStoreMock(url, false)
}

// NewSucceedingStoreMock creates an MockStore that will always succeed
// when trying to Get or Add something from/to it.
func NewSucceedingStoreMock(url string) *StoreMock {
	return NewStoreMock(url, true)
}
