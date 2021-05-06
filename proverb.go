// Package proverb is Exercism.io side exercise
package proverb

// Proverb — generate the relevant proverb, based on a list of inputs.
func Proverb(rhyme []string) []string {
	result := []string{}
	if len(rhyme) > 0 {
		for i := 0; i < len(rhyme)-1; i++ {
			result = append(result, "For want of a "+rhyme[i]+" the "+rhyme[i+1]+" was lost.")
		}
		result = append(result, "And all for the want of a "+rhyme[0]+".")
	}
	return result
}
