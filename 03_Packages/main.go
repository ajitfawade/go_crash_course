package main

import (
	"fmt"
	"math"

	"github.com/ajitfawade/go_crash_course/03_Packages/strutil"
)

func main() {
	fmt.Println(math.Floor(3.2))
	fmt.Println(math.Ceil(3.2))
	fmt.Println(math.Sqrt(16))
	fmt.Println(strutil.Reverse("olleh"))
}
