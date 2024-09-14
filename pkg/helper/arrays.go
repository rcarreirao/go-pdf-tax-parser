package helper

/**
 */
func ArrayInsertString(array []string, element string, i int) []string {
	return append(array[:i], append([]string{element}, array[i:]...)...)
}
