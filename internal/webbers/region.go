package webbers

import "fmt"

type Region struct {
	Start int
	End   int
	Lines []string
}

func NewOpenRegion(start int) Region {
	return Region{
		Start: start,
		End:   -1,
		Lines: make([]string, 0),
	}
}

func NewEmptyRegion() Region {
	return Region{
		Start: -1,
		End:   -1,
		Lines: make([]string, 0),
	}
}

func (r *Region) Close(end int) bool {
	if r.End != -1 {
		return false
	}
	r.End = end
	return true
}

func (r *Region) IsNul() bool {
	return r.Start == -1 && r.End == -1
}

func (r *Region) IsOpen() bool {
	return r.End == -1 && r.Start != -1
}

func (r *Region) IsClosed() bool {
	return r.End != -1 && r.Start != -1
}

func (r *Region) HasLines() bool {
	return len(r.Lines) > 0
}

func (r *Region) AddLine(line string) bool {
	if r.IsClosed() {
		return false
	}
	r.Lines = append(r.Lines, line)
	return true
}

func (r *Region) GetLine(i int) (string, error) {
	if i >= len(r.Lines) {
		return "", fmt.Errorf("index out of range")
	}
	return r.Lines[i], nil
}
