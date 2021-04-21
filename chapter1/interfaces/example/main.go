package main

import (
	"bytes"
	"fmt"

	"github.com/PacktPublishing/Go-Programming-Cookbook-Second-Edition/chapter1/interfaces"
)

func main() {
	//func NewReader(b []byte) *Reader 형식
	// b로부터 읽은 것을 새로운 reader로 반환하는 함수임
	// 즉 여기서는 example을 반환해서 in 에 넣는 역할
	in := bytes.NewReader([]byte("example"))
	fmt.Printf("%q\n", in)

	//버퍼는 읽기와 쓰기 메소드를 가진 가변크기 버퍼임. Buffer 0값은 사용할 준비가 된 빈 버퍼임.
	out := &bytes.Buffer{}
	fmt.Println("out", out)

	
	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())

	fmt.Print("stdout on PipeExample = ")
	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}
