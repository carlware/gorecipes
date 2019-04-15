package main

import "fmt"

func arr() {
	var movies [4]string
	movies[0] = "Hellboy"
	movies[1] = "Shazam"
	movies[2] = "Avengers"

	fmt.Println(movies)
	for _, ele := range movies {
		fmt.Println(ele)
	}
}

func updateByRef(sl []string) {
	sl[0] = "0"
	sl[1] = "1"
	sl[3] = "3"

	sl = append(sl, "a")
	sl = append(sl, "b")
	sl = append(sl, "c")
	fmt.Println(sl)
}

func main() {
	arr()

	s := make([]string, 5)

	fmt.Println(s)
	updateByRef(s)
	fmt.Println(s)
}
