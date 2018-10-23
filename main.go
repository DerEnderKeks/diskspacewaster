package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	size := flag.Int("s", 1, "File size in MB")
	path := flag.String("f", "iamwastingdiskspace", "File name")

	flag.Parse()

	_, err := os.Stat(*path)
	if os.IsExist(err) {
		fmt.Println("Deleting old file...")
		err := os.Remove(*path)
		check(err)
		fmt.Println("Deleted.")
	}

	f, err := os.Create(*path)
	defer f.Close()
	check(err)

	fmt.Printf("Creating new file at '%s' (Size: %d MB)...\n", *path, *size)
	err = f.Truncate(int64(*size) * 1048576)
	check(err)

	fmt.Println("Syncing...")
	f.Sync()
	fmt.Println("Synced.")
	fmt.Printf("Done! Have fun with your new %d MB big file.\n", *size)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
