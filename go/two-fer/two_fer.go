// Package twofer should have a package comment that summarizes what it's about.
package twofer

// ShareWith should have a comment documenting it.
func ShareWith(name string) string {
	return "One for " + pronoun(name) + ", one for me."
}

func pronoun(name string) string {
	if name != "" {
		return name
	}
	return "you"
}
