package tree

import "log"

func mapR(value string, nodes []*Node) *Node {
	switch value {
	case "00":
		return nodes[0]
	case "10":
		return nodes[1]
	case "11":
		return nodes[2]
	case "01":
		return nodes[3]
	}
	log.Fatal("MapR: Out of bounds!")
	return nodes[3]
}
func mapU(value string, nodes []*Node) *Node {
	switch value {
	case "00":
		return nodes[0]
	case "10":
		return nodes[3]
	case "11":
		return nodes[2]
	case "01":
		return nodes[1]
	}
	log.Fatal("MapU: Out of bounds!")
	return nodes[3]
}
func mapL(value string, nodes []*Node) *Node {
	switch value {
	case "11":
		return nodes[0]
	case "01":
		return nodes[1]
	case "00":
		return nodes[2]
	case "10":
		return nodes[3]
	}
	log.Fatal("MapL: Out of bounds!")
	return nodes[3]
}

func mapD(value string, nodes []*Node) *Node {
	switch value {
	case "11":
		return nodes[0]
	case "10":
		return nodes[1]
	case "00":
		return nodes[2]
	case "01":
		return nodes[3]
	}
	log.Fatal("MapD: Out of bounds!")
	return nodes[3]
}
