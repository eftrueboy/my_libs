package log_task

import (
	"fmt"
	"log"
	"my_libs/file_operator"
	"my_libs/logger/log_config_manager"
	"my_libs/utils/date_util"
	"strings"
	"time"
)

type LogTask struct {
	fileOperator *file_operator.FileOperator
	directory    string
	created      int64
	fileSize     int
	filename     string
	fileFullName string
	splitNum     int
}

func NewLogTask(fileFullName string) *LogTask {
	lastIndex := strings.LastIndex(fileFullName, "/")
	directory := fileFullName[:lastIndex]
	filename := fileFullName[lastIndex+1:]
	file_operator.MkdirAllFull(directory)

	logTask := &LogTask{}
	logTask.fileFullName = fileFullName
	logTask.filename = filename
	logTask.directory = directory
	logTask.setFileOperator()
	return logTask
}

func (this *LogTask) setFileOperator() {
	var err error
	this.fileOperator, err = file_operator.NewFileAppend(this.fileFullName)
	if err != nil {
		log.Fatal(err)
	}
	this.created = time.Now().UnixMicro()
}

func (this *LogTask) Write(text string) {
	this.fileSize += len(text)
	this.splitFile()

	this.fileOperator.Write(text)
}

func (this *LogTask) filenameTimePrefix() string {
	timePrefix := strings.Builder{}
	timePrefix.WriteString(date_util.DateSecondFilename(this.created))
	timePrefix.WriteString("-")
	timePrefix.WriteString(date_util.NowDateSecondTimezoneFilename())
	return timePrefix.String()
}

func (this *LogTask) newFilename() string {
	newFilename := fmt.Sprintf("%s-%d-%s", this.filenameTimePrefix(), time.Now().UnixNano(), this.filename)
	return newFilename
}

func (this *LogTask) newFullname() string {
	return fmt.Sprintf("%s/%s", this.directory, this.newFilename())
}

func (this *LogTask) splitFile() {
	if this.fileSize < log_config_manager.CurrentConfig().MaxFileSize {
		return
	}
	this.fileOperator.Close()

	err := file_operator.Rename(this.fileFullName, this.newFullname())
	if err != nil {
		log.Fatal(err)
	}
	this.fileSize = 0
	this.splitNum = 0
	this.setFileOperator()
}
