package filedirs

import (
	"errors"
	"io"
	"os"
)

func Operate() error {
	// func Mkdir(name string, perm FileMode) error
	// os.Mkdir은 디렉터리 명 example_dir을 생성한다. 그리고 permission 값은 0755임
	// 읽기:4 / 쓰기:2 / 실행: 1
	// 0755는 user권한:7(rwx) / group:5(r-x) / other:5(r-x) 임. 
	if err := os.Mkdir("example_dir", os.FileMode(0755));
	err != nil {
		return err
	}

	// func Chdir(dir string) error
	// Chdir은 현재 디렉토리를 example_dir 로 바꿈. 
	if err := os.Chdir("example_dir"); err != nil {
		return err
	}

	// func Create(name string) (*File, error)
	// 이 명령어는 text.txt 파일을 생성하거나 truncate 한다. 파일이 이미 존재할 경우, 
	// truncated 된다. 파일이 존재하지 않을 경우 mode 0666으로 생성된다. 
	f, err := os.Create("test.txt")
	if err != nil {
		return err
	}

	// hellon을 value 변수에 넣어서 f변수에 쓴다. 
	// count 변수에 바이트 갯수가 들어가고, 아래 if 문에서 개수를 확인한다. 
	value := []byte("hellon")
	// func (f *File) Write(b []byte) (n int, err error)
	// file에 사용하는 Write 함수는 글자를 파일에 적고 byte 개수를 리턴한다. 
	count, err := f.Write(value)
	if err != nil {
		return err
	}
	if count != len(value) {
		return errors.New("incorrect length returned from write")
	}

	if err := f.Close(); err != nil {
		return err
	}

	f, err = os.Open("test.txt")
	if err != nil {
		return err
	}

	
	io.Copy(os.Stdout, f)

	if err := f.Close(); err != nil {
		return err
	}

	if err := os.Chdir(".."); err != nil {
		return err
	}

	if err := os.RemoveAll("example_dir"); err != nil {
		return err
	}
	return nil
}