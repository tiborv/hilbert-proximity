package tree

func map1(value string, nodes []*Node) *Node {
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
	return nodes[3]
}
func map2(value string, nodes []*Node) *Node {
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
	return nodes[3]
}
func map3(value string, nodes []*Node) *Node {
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
	return nodes[3]
}
func map4(value string, nodes []*Node) *Node {
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
	return nodes[3]
}
