package main

import (
  "bufio"
  "os"
  "math/rand"
  "fmt"
)

type FileIO struct {
  AllLines []string
}

func NewFileIO() *FileIO {
  return &FileIO{}
}

func (fio *FileIO) ReadLines(path string) {
  file, err := os.Open(path)
  if err != nil {
    return
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fio.AllLines = append(fio.AllLines, scanner.Text())
  }
}

func (fio *FileIO) AssembleRandomString() string{
  rs := ""
  r := rand.Intn(1000)
  for i := r; i < r+10; i++ {
    rs = fmt.Sprintf("%s%s", rs, fio.AllLines[i])
  }
  return rs
}