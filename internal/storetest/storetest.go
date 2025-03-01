package storetest

type Mock struct {
	URL    string
	Result bool
}

func (s *Mock) Get(_ string) (string, bool) {
	return s.URL, s.Result
}

func (s *Mock) Set(_, _ string) bool {
	return s.Result
}
