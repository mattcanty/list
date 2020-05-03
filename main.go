package main

// Lists will be unexported
var lists = make(map[string][]string)

// CreateNewList creates a new list
func CreateNewList(listName string) {
	var emptyslice []string
	lists[listName] = emptyslice
}

// ListCount returns the number of lists which are in existence
func ListCount() int {
	return len(lists)
}

// GetList returns a list by name
func GetList(listName string) []string {
	return lists[listName]
}

// AddListItem adds a string to a named list
func AddListItem(listName string, item string) {
	lists[listName] = append(lists[listName], item)
}

// DeleteAllLists does what it sez on the tin
func DeleteAllLists() {
	for list := range lists {
		delete(lists, list)
	}
}
