package spell

// Words ...
type Words []string

// Flag returns if words include the flag given.
func (words Words) Flag(flag string) bool {
	for _, w := range words {
		if w == flag {
			return true
		}
	}
	return false
}

// Arg returns a value after the arg name given.
func (words Words) Arg(name string) string {
	for i, w := range words {
		if w == name {
			if i+1 < len(words) {
				return words[i+1]
			}
		}
	}
	return ""
}

// Remove ...
func (words Words) Remove(name string, count int) Words {
	for i, w := range words {
		if w == name {
			words = append(words[:i], words[i+count+1:]...)
			break
		}
	}
	return words
}
