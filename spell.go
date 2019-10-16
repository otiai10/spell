package spell

const (
	space       = ' '
	singlequote = '\''
	doublequote = '"'
)

// Parse ...
func Parse(s string) []string {
	r := []string{}
	l := len(s)
	t := new(token)
	for i := 0; i < l; i++ {
		t.Push(s[i])
		if t.Closed {
			r = append(r, t.Flush())
		}
	}
	return r
}

// token
type token struct {
	delim  byte
	pool   []byte
	Closed bool
}

func (t *token) Flush() string {
	s := string(t.pool)
	t.delim = 0
	t.pool = nil
	t.Closed = false
	return s
}

func (t *token) Push(c byte) {
	switch c {
	case space:
		if t.delim != 0 {
			t.pool = append(t.pool, c)
		} else {
			if len(t.pool) != 0 {
				t.Closed = true
			}
		}
		return
	case singlequote:
		t.pool = append(t.pool, c)
		if t.delim == singlequote {
			t.Closed = true
		}
		if len(t.pool) == 1 {
			t.delim = singlequote
		}
	case doublequote:
		t.pool = append(t.pool, c)
		if t.delim == doublequote {
			t.Closed = true
		}
		if len(t.pool) == 1 {
			t.delim = doublequote
		}
	default:
		t.pool = append(t.pool, c)
	}
}
