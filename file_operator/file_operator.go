package file_operator

import (
	"fmt"
	"os"
)

type FileOperator struct {
	file *os.File
}

//------Private Methods------
func newFile(fullname string, flag int, perm os.FileMode) (fileOperator *FileOperator, err error) {
	fileOperator = &FileOperator{}
	fileOperator.file, err = os.OpenFile(fullname, flag, perm)
	if err != nil {
		return nil, err
	}
	return fileOperator, nil
}

//---------Instance Methods---------
func (this *FileOperator) Write(text string) {
	fmt.Fprintln(this.file, text)
}

func (this *FileOperator) Close() {
	this.file.Close()
}

//---------CreateInstance Methods-----------------

//创建新文件: 如果文件存在就不创建。 权限: 可写，可添加内容。
func NewFileAppend(filepath string) (file *FileOperator, err error) {
	return newFile(filepath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
}

func NewFileRead(fullname string) (file *FileOperator, err error) {
	return newFile(fullname, os.O_RDONLY|os.O_CREATE, os.ModePerm)
}

func NewFileWrite(fullname string) (file *FileOperator, err error) {
	return newFile(fullname, os.O_WRONLY|os.O_CREATE, os.ModePerm)
}

func NewFileWriteOverwrite(fullname string) (file *FileOperator, err error) {
	return newFile(fullname, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, os.ModePerm)
}

func NewFileReadWrite(fullname string) (file *FileOperator, err error) {
	return newFile(fullname, os.O_RDWR|os.O_CREATE, os.ModePerm)
}

//---------Static Methods----------

func FileExist(fullname string) bool {
	_, err := os.Lstat(fullname)
	return !os.IsNotExist(err)
}

func Rename(oldpath string, newpath string) (err error) {
	return os.Rename(oldpath, newpath)
}

//--------------目录操作-------------

//根据名称: 在当前目录创建一个文件夹, 并给予最高权限: 0777.
func MkdirFull(name string) (err error) {
	return os.Mkdir(name, os.ModePerm)
}

//根据路径: 创建一个文件夹, 并给予最高权限: 0777.
func MkdirAllFull(path string) (err error) {
	return os.MkdirAll(path, os.ModePerm)
}

//删除的一个目录.
func Remove(name string) (err error) {
	return os.Remove(name)
}

//删除多级子目录, 如果path是单个名称，那么该目录下的子目录全部删除.
func RemoveAll(path string) (err error) {
	return os.RemoveAll(path)
}
