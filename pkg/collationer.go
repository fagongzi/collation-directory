package pkg

import (
	"fmt"
	"github.com/CodisLabs/codis/pkg/utils/log"
	"os"
	"path/filepath"
	"time"
)

var counter = 1

type Collationer struct {
	source string
	dest   string
}

func NewCollationer(source string, dest string) *Collationer {
	return &Collationer{
		source: source,
		dest:   dest,
	}
}

func (self *Collationer) Start() error {
	if err := createDir(self.dest); err != nil {
		log.PanicErrorf(err, "Create dir <%s> failure", self.dest)
	}

	err := filepath.Walk(self.source, self.process)

	if err != nil {
		log.ErrorErrorf(err, "Iteration dir <%s> failure", self.source)
	}

	return nil
}

func (self *Collationer) process(path string, f os.FileInfo, err error) error {
	log.Debugf("Process <%s>.", path)

	if f == nil {
		return err
	}

	if f.IsDir() {
		log.Debugf("Is Dir <%s>.", path)
		return nil
	}

	dateDir := getDateFormatDir(self.dest, f.ModTime())
	if !isExit(dateDir) {
		e := createDir(dateDir)
		if e != nil {
			return e
		}

		log.Debugf("Date format dir <%s> created.", dateDir)
	}

	e := moveTo(path, dateDir)

	if e == nil {
		log.Debugf("<%s> moved to <%s> success.", filepath.Base(path), dateDir)
	}

	return e
}

func createDir(dir string) error {
	return os.MkdirAll(dir, os.ModeDir)
}

func isExit(dir string) bool {
	if _, err := os.Stat(dir); err == nil {
		return true
	}

	return false
}

func moveTo(path, dir string) error {
	newPath := filepath.Join(dir, filepath.Base(path))

	if isExit(newPath) {
		newPath = filepath.Join(dir, fmt.Sprintf("%d-%s", counter, filepath.Base(path)))
		counter += 1
	}

	return os.Rename(path, newPath)
}

func getDateFormatDir(dir string, t time.Time) string {
	return filepath.Join(dir, fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day()))
}
