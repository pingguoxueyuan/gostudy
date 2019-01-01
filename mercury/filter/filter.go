package filter

func Replace(text string, replace string) (result string, hit bool) {
	result, hit = trie.Check(text, replace)
	return
}
