package GoLangToSplitString

import "strings"

func SplitString(String string, Data string) ([]string, int) {
	Split := strings.Split(String, Data)
	return Split, len(Split)
}
