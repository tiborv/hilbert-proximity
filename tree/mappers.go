package tree

var mapR, mapU, mapL, mapD map[byte]int

func init() {
	mapR = map[byte]int{
		0: 0,
		2: 1,
		3: 2,
		1: 3,
	}
	mapU = map[byte]int{
		0: 0,
		2: 3,
		3: 2,
		1: 1,
	}
	mapL = map[byte]int{
		0: 2,
		2: 3,
		3: 0,
		1: 1,
	}
	mapD = map[byte]int{
		0: 2,
		2: 1,
		3: 0,
		1: 3,
	}
}
