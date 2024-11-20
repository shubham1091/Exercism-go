package strand

import "strings"

var dnaToRna = map[rune]rune{
	'A': 'U',
	'C': 'G',
	'G': 'C',
	'T': 'A',
}

func ToRNA(dna string) string {
	var sb strings.Builder
	sb.Grow(len(dna))
	for _, nuc := range dna {
		sb.WriteRune(dnaToRna[nuc])
	}
	return sb.String()
}
