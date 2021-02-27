// Package bob is a teenager.
package bob

import "strings"

// Hey responds to remarks as a teenager named Bob.
func Hey(remark string) string {

	remark = strings.Trim(remark, " \r\n\t") //not worried about overwriting our input with a cleaner version.

	//He says 'Fine. Be that way!' if you address him without actually saying anything.
	if remark == "" {
		return "Fine. Be that way!"
	}

	//He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
	if isExclamation(remark) && isInterogative(remark) {
		return "Calm down, I know what I'm doing!"
	}

	//Bob answers 'Sure.' if you ask him a question, such as "How are you?".
	if isInterogative(remark) {
		return "Sure."
	}

	//He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
	if isExclamation(remark) {
		return "Whoa, chill out!"
	}

	//He answers 'Whatever.' to anything else.
	return "Whatever."
}

// isExlamation determines if a remark is an exclamatory remark.
func isExclamation(a string) bool {
	return strings.ToUpper(a) == a && strings.ToUpper(a) != strings.ToLower(a)
}

//isInterogative determines if a remark is a question.
func isInterogative(a string) bool {
	return a[len(a)-1] == '?'
}
