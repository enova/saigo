package visitor

var visitorRecords = make(map[string]int)

// Record the visitor's name and visit count
func Record(name string) int {
	visitorRecords[name]++
	return visitorRecords[name]
}

// List all past visitors
func List() map[string]int {
	return visitorRecords
}
