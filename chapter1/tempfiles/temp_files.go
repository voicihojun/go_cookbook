package tempfiles

import (
	"fmt"
	"io/ioutil"
	"os"
)

func WorkWithTemp() error {
	// func TempDir(dir, pattern string) (name string, err error)
	// tmp + 임의의 수 로 구성된 이름을 가진 임시 디렉토리를 만든다.
	t, err := ioutil.TempDir("", "tmp")
	fmt.Println(t)

	if err != nil {
		return err
	}
	
	defer os.RemoveAll(t)
	// t 디렉토리에 tmp + 임의의 수 로 구성된 이름을 가진 임시파일을 만든다.
	tf, err := ioutil.TempFile(t, "tmp")
	if err != nil {
		return err
	}
	fmt.Println(tf.Name())
	return nil
}