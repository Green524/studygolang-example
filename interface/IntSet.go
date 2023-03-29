package count

type IntSet struct {
}

func (*IntSet) String() string {
	return "string"
}
func (IntSet) Write(p []byte) (n int, err error) {
	return 0, err
}
