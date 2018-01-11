package main

import (
	"os"
	"flag"
	"fmt"
)

func main() {
	size := flag.Int("s", 1, "File size in MB")
	path := flag.String("f", "iamwastingdiskspace", "File name")

	flag.Parse()

	fmt.Println("deleting...")
	_, err := os.Stat(*path)
	if os.IsExist(err) {
		err := os.Remove(*path)
		check(err)
	}
	fmt.Println("deleted.")

	f, err := os.Create(*path)
	defer f.Close()
	check(err)

	fmt.Println("writing...")
	err = f.Truncate(int64(*size) * 1048576)
	check(err)

	fmt.Println("syncing...")
	f.Sync()
	fmt.Println("synced.")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
