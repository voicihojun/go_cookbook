package main

import (
	"fmt"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter1/csvformat"
)

func main() {
	if err := csvformat.AddMoviesFromText(); err != nil {
		fmt.Printf("111\n")
		panic(err)
	}

	if err := csvformat.WriteCSVOutput(); err != nil {
		fmt.Printf("222")
		panic(err)
	}

	buffer, err := csvformat.WriteCSVBuffer()
	if err != nil {
		fmt.Printf("333")
		panic(err)
	}
	fmt.Println("Buffer = ", buffer.String())
}


// parseInt 관련 코드

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	// func parseInt(s string, base int, bitSize int) (i int64, err error) 
// 	// 부호 있는 정수 형태의 문자열 => 부호 있는 정수로 변환
// 	v32 := "-354634382"
// 	if s, err := strconv.ParseInt(v32, 10, 32); err == nil {
// 		fmt.Printf("1 %T, %v\n", s, s)
		
// 	}
// 	if s, err := strconv.ParseInt(v32, 16, 32); err == nil {
// 		fmt.Printf("2 %T, %v\n", s, s)
		
// 	}

// 	v64 := "-3546343826724305832"
// 	if s, err := strconv.ParseInt(v64, 10, 64); err == nil {
// 		fmt.Printf("3 %T, %v\n", s, s)
		
// 	}
// 	if s, err := strconv.ParseInt(v64, 16, 64); err == nil {
// 		fmt.Printf("4 %T, %v\n", s, s)
		
// 	}

// }