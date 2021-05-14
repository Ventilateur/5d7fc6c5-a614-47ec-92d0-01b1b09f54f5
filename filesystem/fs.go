package filesystem

import (
	"fmt"
	"golang.org/x/sys/unix"
	"io/ioutil"
	"os"
)

type FileStorage struct {
	DirName string
}

func (fs *FileStorage) Init(dirName string) error {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		unix.Umask(0000)
		return os.Mkdir(dirName, 0777)
	}
	fs.DirName = dirName
	return err
}

func (fs *FileStorage) Save(fileName, content string) error {
	f, err := os.OpenFile(
		fmt.Sprintf("%s/%s", fs.DirName, fileName),
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755,
	)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	return err
}

func (fs *FileStorage) Get(fileName string) (string, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", fs.DirName, fileName))
	return string(content), err
}

func (fs *FileStorage) Delete(fileName string) error {
	return os.Remove(fmt.Sprintf("%s/%s", fs.DirName, fileName))
}
