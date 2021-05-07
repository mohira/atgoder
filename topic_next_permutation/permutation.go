package topic_next_permutation

type IntSlice []int

func (p IntSlice) Len() int      { return len(p) }
func (p IntSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

type Permutator struct {
	first bool // It is first state of elements
	v     Interface
	b     []int
}

func NewPermutator(v Interface) *Permutator {
	n := v.Len()
	if n > 0 {
		n--
	}
	return &Permutator{
		first: true,
		v:     v,
		b:     make([]int, n),
	}
}

func (p *Permutator) Next() bool {
	if p.first {
		p.first = false
		return true
	}
	n := calcFlipSize(p.b)
	if n == -1 {
		// It is last flip. It helps to return in begin state of the elements.
		flip(p.v, p.v.Len())
		p.first = true
		return false // End of permutations
	}
	flip(p.v, n) // It is the main flip.
	return true
}

func calcFlipSize(b []int) int {
	for i := range b {
		b[i]++
		if k := i + 2; b[i] < k {
			return k
		}
		b[i] = 0
	}
	return -1
}

// flip is a function for make flip first n elements in slice (v)
func flip(v Interface, n int) {
	i, j := 0, n-1
	for i < j {
		v.Swap(i, j)
		i, j = i+1, j-1
	}
}
