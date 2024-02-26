package main

import (
	"log"
	"os"
	"time"
)

func ChangePermissions() {
	path := "changePermission.txt"

	ExistsOrCreate(path)

	// Change permissions using linux style
	err := os.Chmod(path, 0777)

	if err != nil {
		log.Println(err)
	}

	// Change Ownership
	err = os.Chown(path, os.Getuid(), os.Getgid())

	if err != nil {
		log.Println(err)
	}

	// Change timestamps
	twodaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twodaysFromNow
	lastModifyTime := twodaysFromNow

	err = os.Chtimes(path, lastAccessTime, lastModifyTime)

	if err != nil {
		log.Println(err)
	}
}
