package interfaces

import (
	"fmt"
	"io"
	"os"
)

func Copy(in io.ReadSeeker, out io.Writer) error {
	// out과 os.Stdout(화면)에 값을 쓰는 것인데, 다음 라인에 나오는 
	// io.Copy(w, in)에 의해서 in 값이 w의 out / os.Stdout에 
	// 써지는 것인 듯. 책에는 MultiWriter 인터페이스가 두개 writer stream을 
	// 합치고 그것을 ReadSeeker를 이용해서 두번 쓴다(write함) 
	w := io.MultiWriter(out, os.Stdout)

	// main.go에서 호출될 때 in 값이 'example' 이었으므로, 
	// out, os.Stdout = 'example'
	if _, err := io.Copy(w, in); err != nil {
		return err
	}

	// seek(offset, whence) 파이썬 설명에 의하면
	// whence 값이 0 이면 파일의 처음부터 측정, 1 이면 현재 위치부터 측정, 2 이면 파일의 끝을 기준으로 측정하는 것
	// offset은 그 기준점으로부터의 거리
	in.Seek(0, 0)

	// 64개의 byte 슬라이스를 만드는 것
	buf := make([]byte, 64)
	
	// io.CopyBuffer
	// CopyBuffer는 임시 버퍼를 할당하는 대신 제공된 버퍼 (필요한 경우)를 통해 스테이징된다는 점을 
	// 제외하면 Copy와 동일합니다.
	// buf가 nil이면 하나가 할당됩니다. 그렇지 않고 길이가 0이면 CopyBuffer가 패닉합니다.
	// 즉, 여기선 in('example') 값을 w에 넣는데 위에서 정의한 buf를 통해서 w에 넣는다는 것 같음.
	if _, err := io.CopyBuffer(w, in, buf); err != nil {
		return err
	}

	fmt.Println()

	return nil
}