package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

const (
	maxFileSize = 200 * 1024 * 1024
	format      = "20060102_150405_000"
)

func init() {
	log.SetPrefix("[RM] ")
	log.SetFlags(log.Ldate | log.Ltime)
}

func main() {
	home, ok := os.LookupEnv("HOME")
	if !ok {
		log.Fatalf("lookup HOME from env fail")
	}
	recycleBin := filepath.Join(home, ".RecycleBin")
	if err := os.MkdirAll(recycleBin, 0755); err != nil {
		log.Fatalf("%v", err)
	}
	for i := 1; i < len(os.Args); i++ {
		path := os.Args[i]
		info, err := os.Stat(path)
		if !(err == nil || os.IsExist(err)) {
			log.Fatalf(fmt.Sprintf("%s not exist", path))
		}

		target := filepath.Join(recycleBin, fmt.Sprintf("%s@%s", time.Now().Format(format), info.Name()))
		if info.IsDir() || info.Size() < maxFileSize { // 文件夹、小文件回收到垃圾桶
			if err := os.Rename(path, target); err != nil {
				log.Fatalf("%v", err)
			}
			log.Printf("mv %s => %s", path, target)
		} else {
			if err := os.Remove(path); err != nil {
				log.Fatalf("%v", err)
			}
			log.Printf("rm %s", path)
		}
	}
}
