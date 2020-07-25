package util

import (
	"io"
	"log"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Copy(src, dst string) error {
	in, err := os.Open(src)
	CheckErr(err)
	defer in.Close()

	out, err := os.Create(dst)
	CheckErr(err)
	defer out.Close()

	_, err = io.Copy(out, in)
	CheckErr(err)
	return out.Close()
}
