package tree

var mapR, mapU, mapL, mapD = map[string]int{
	"00": 0,
	"10": 1,
	"11": 2,
	"01": 3,
}, map[string]int{
	"00": 0,
	"10": 3,
	"11": 2,
	"01": 1,
}, map[string]int{
	"00": 2,
	"10": 3,
	"11": 0,
	"01": 1,
}, map[string]int{
	"00": 2,
	"10": 1,
	"11": 0,
	"01": 3,
}
