package filter

func Replace(text string, replace string) (result string, isReplace bool) {
	isReplace, result = trie.Check(text, replace)
	return
}
