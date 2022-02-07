package main

import (
	"os"

	"github.com/qiniu/x/log"
)

func f() {
	log.Println(os.Getwd())
}
