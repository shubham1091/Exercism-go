package bob

// Hey should have a comment documenting it.
func Hey(remark string) string {
	r := newRemark(remark)

	switch {
	case r.IsEmpty():
		return "Fine. Be that way!"
	case r.IsExasperated():
		return "Calm down, I know what I'm doing!"
	case r.IsYelling():
		return "Whoa, chill out!"
	case r.IsQuestion():
		return "Sure."
	default:
		return "Whatever."
	}
}
