package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	c := []byte{0x11, 0x00}

	fmt.Println(c[0] > c[1])
	//test := n.Find(&p)

	//i := 0
	// for i < 10000 {
	// 	t := geo.NewPoint(randBits(100), randBits(100))
	// 	n.InsertPoint(&t)
	// 	fmt.Println(t)
	// 	i++
	// }

}

func randBits(n int) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		s := strconv.Itoa(rand.Intn(2))
		buffer.WriteString(s)
	}
	return buffer.String()
}
