package bytestrings

import (
	"bufio"
	"bytes"
	"fmt"
)

func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array"
	defer fmt.Println("\n===========")
	b := Buffer(rawString)
	fmt.Println(b.String()) //rawString 을 그대로 출력. 1회

	s, err := toString(b) // buffer에 있는 메소드로 에러가 없는 한 b에 있는 것을 그대로 출력
	if err != nil {
		return err
	}

	fmt.Println(s) // rawString 을 그대로 출력. 2회째

	reader := bytes.NewReader([]byte(rawString)) // rawString 에서 읽은 것을 바이트?로 출력??

	// func NewScanner(r io.Reader) *Scanner
	// NewScanner returns a new Scanner to read from r. The split function defaults to ScanLines. 
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	// ScanWords is a split function for a Scanner that returns each 
	// space-separated word of text, with surrounding spaces deleted. 
	// It will never return an empty string. The definition of space is 
	// set by unicode.IsSpace.
	
	// func (s *Scanner) Split(split SplitFunc)
	// Split sets the split function for the Scanner. The default split function is ScanLines.
	// Split panics if it is called after scanning has started.

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}
