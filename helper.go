package main

var (
	validColumns = []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	validRows    = []string{"1", "2", "3", "4", "5", "6", "7", "8"}

	columnToIndex = map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
		"g": 6,
		"h": 7,
	}

	rowToIndex = map[string]int{
		"8": 0,
		"7": 1,
		"6": 2,
		"5": 3,
		"4": 4,
		"3": 5,
		"2": 6,
		"1": 7,
	}

	indexToColumn = map[int]string{
		0: "a",
		1: "b",
		2: "c",
		3: "d",
		4: "e",
		5: "f",
		6: "g",
		7: "h",
	}

	indexToRow = map[int]string{
		0: "8",
		1: "7",
		2: "6",
		3: "5",
		4: "4",
		5: "3",
		6: "2",
		7: "1",
	}
)

func arrayContains(array []string, item string) bool {
	for _, s := range array {
		if s == item {
			return true
		}
	}

	return false
}
