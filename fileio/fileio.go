// This entire file follows this guide: https://www.devdungeon.com/content/working-files-go

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func ExistsOrCreate(path string) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("File doesn't exist to delete, creating...")
		os.Create(path)
		fmt.Println("File Created")
	}
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createFile() {
	newFile, err := os.Create("creation.txt")

	FatalError(err)


	log.Printf("%+v\n", newFile)
	newFile.Close()
}

func truncateFile() {
	// Truncate a file to 100 bytes. If file
	// is less than 100 bytes the original contents will remain
	// at the beginning, and the rest of the space is
	// filled will null bytes. If it is over 100 bytes,
	// Everything past 100 bytes will be lost. Either way
	// we will end up with exactly 100 bytes.
	// Pass in 0 to truncate to a completely empty file
	fmt.Println("Truncating file")
	err := os.Truncate("truncation.txt", 100)
	FatalError(err)

	fmt.Println("File truncated")
}

func getFileInfo() {
	// Stat returns file info. It will return
	// an error if there is no file
	file, err := os.Stat("truncation.txt")

	FatalError(err)


	fmt.Println("FileName: ", file.Name())
	fmt.Println("Size in bytes: ", file.Size())
	fmt.Println("Permissions: ", file.Mode())
	fmt.Println("Last Modified: ", file.ModTime())
	fmt.Println("Is Directory: ", file.IsDir())
	fmt.Printf("System interface type: %T\n", file.Sys())
	data, _ := json.MarshalIndent(file.Sys(), " ", "    ")
	fmt.Printf("System info: %+v\n\n", string(data))
}

func renameFile() {
	fmt.Println("File Rename")

	originalPath := "test.txt"
	_, err := os.Stat(originalPath)

	if err != nil && os.IsNotExist(err) {
		fmt.Println("File didn't exist, creating...")
		os.Create(originalPath)
		fmt.Println("File created")
	}

	newPath := "test2.txt"

	fmt.Println("Renaming", originalPath, "to ", newPath)
	err = os.Rename(originalPath, newPath)

	FatalError(err)


	fmt.Println("Rename complete")
}

func deleteFile() {
	fmt.Println("File Deletion")

	path := "deletion.txt"

	ExistsOrCreate(path)

	fmt.Println("Deleting File...")
	// time.Sleep(time.Minute) // Uncomment to see deletion in action after a minute
	err := os.Remove(path)

	FatalError(err)


	fmt.Println("File Deleted Successfully")

}

func seekFile() {
	fmt.Println("Seeking out file")
	path := "seek.txt"
	// Simple read only open. We will cover actually reading
	// and writing to files in examples further down the page

	ExistsOrCreate(path)

	file, err := os.Open(path)

	FatalError(err)


	fmt.Println("Close initial File Seek")
	file.Close()

	// OpenFile with more options. Last param is the permission mode
	// Second param is the attributes when opening
	fmt.Println("Secondary File seek, but with ")
	file, err = os.OpenFile("test.txt", os.O_APPEND, 0666)
	FatalError(err)

	file.Close()
	// Use these attributes individually or combined
	// with an OR for second arg of OpenFile()
	// e.g. os.O_CREATE|os.O_APPEND
	// or os.O_CREATE|os.O_TRUNC|os.O_WRONLY

	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create is none exist
	// os.O_TRUNC // Truncate file when opening
}

func readWriteFile() {
	// Test write permissions. It is possible the file
	// does not exit and that will return a different
	// error that can be checked with os.IsNotExist(err)
	path := "readWriteFile.txt"

	ExistsOrCreate(path)

	file, err := os.OpenFile(path, os.O_WRONLY, 0666)

	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Write Permission Denied.")
		}
	}

	file.Close()

	// Test read permissions
	file, err = os.OpenFile(path, os.O_RDONLY, 0666)

	if err != nil && os.IsPermission(err) {
		log.Println("Error: Read Permission Denied")
	}

	file.Close()
}

func changePermissions() {
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

func HardLinkFiles() {
	// Creating a hard link
	// You will have two file names that point to the same contents
	// changing the contents of one will change the other
	// Deleting/Renaming one will not affect the other
	path := "HardLink.txt"
	path2 := "HardLink_Other.txt"
	// Simple read only open. We will cover actually reading
	// and writing to files in examples further down the page
	ExistsOrCreate(path)

	err := os.Link(path, path2)

	FatalError(err)

}

func SymLinkFiles() {
	// Creating a symlink
	path := "SymLink.txt"
	sym := "SymLink_SYM.txt"

	ExistsOrCreate(path)

	err := os.Symlink(path, sym)

	FatalError(err)


	// LStat will return file info, but if it is actually
	// a symlink, it will return info about the SymLink
	// It will not follow the link and give information
	// about the real file
	// Symlinks do not work in Windows (Running this in WSL2 - UBUNTU Dev Contianer)
	fileInfo, err := os.Lstat(sym)

	FatalError(err)


	data, err := json.MarshalIndent(fileInfo, " ", "	")

	FatalError(err)


	fmt.Printf("Link info: %+v", data)

	// Change ownership of a symlink only
	// and not the file it points to
	err = os.Lchown(sym, os.Geteuid(), os.Getgid())
	FatalError(err)

}

func copyFile() {
	// Copy a file
	// Open original file
	path := "copy.txt"

	ExistsOrCreate(path)

	original, err := os.Open(path)

	FatalError(err)


	defer original.Close()

	// Create new copy
	newFile, err := os.Create("test_copy.txt")

	FatalError(err)


	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, original)

	FatalError(err)


	log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file content
	// Flushes Memory To Disk
	err = newFile.Sync()
	FatalError(err)

}

func seekPositionInFile() {
	path := "seekPosition.txt"

	ExistsOrCreate(path)

	file, _ := os.Open(path)
	defer file.Close();


	// Offset is how many bytes to move
	// Offset can be positive or negative
	var offset int64 = 5

	// Whence is the point of reference for offset
	// 0 = Beginning of the file
	// 1 = current position
	// 2 = End of File
	var whence int = 0

	newPosition, err := file.Seek(offset, whence)

	FatalError(err)


	fmt.Println("Just moved to 5: ", newPosition)

	// Go back 2 bytes from current possition
	newPosition, err = file.Seek(-2, 1)

	FatalError(err)


	fmt.Println("Just moved back two: ", newPosition)

	// Find the current position by getting the
	// return value from seek after moving 0 bytes
	newPosition, err = file.Seek(0, 1)

	FatalError(err)


	fmt.Println("Current Position: ", newPosition)

	// Go To Beginning of file
	newPosition, err = file.Seek(0, 0)
	FatalError(err)


	fmt.Println("Position after seeking 0,0: ", newPosition)
}

func writeBytesToFile() {
	path := "writeBytes.txt"

	ExistsOrCreate(path)

	// Open a new file for writing only
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)

	FatalError(err)


	defer file.Close()

	//Write bytes to file
	bytesSlice := []byte("Bytes!\n")

	bytesWritten, err := file.Write(bytesSlice)

	FatalError(err)


	log.Printf("Wrote %d bytes. \n", bytesWritten)
}

func quickWriteToFile() {
	path := "quickwrite.txt"

	ExistsOrCreate(path)

	err := os.WriteFile(path, []byte("Hi\n"), 0666)

	FatalError(err)

}

func BufferedWriter() {
	path := "bufferedWriter.txt"

	ExistsOrCreate(path)

	// Open file for writing
	file, err := os.OpenFile(path, os.O_WRONLY, 0666)

	FatalError(err)


	defer file.Close()

	// Create a buffered writer from the file
	bufferedWriter := bufio.NewWriter(file)

	bytesWritten, err := bufferedWriter.Write(
		[]byte{65,66,67},
	)

	FatalError(err)


	log.Printf("bytes written: %d\n", bytesWritten)

	// Write string to buffer
	// Also available are WriteRune() and WriteByte()
	bytesWritten, err = bufferedWriter.WriteString("Buffered string\n")

	FatalError(err)

	log.Printf("Bytes written: %d\n", bytesWritten)

	// Check how much is stored in buffer waiting
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes Buffered: %d\n", unflushedBufferSize)

	// See how much buffer is available
	bytesAvailable:= bufferedWriter.Available()

	log.Printf("Available buffer: %d\n", bytesAvailable)

	// Write memory buffer to disk
	err = bufferedWriter.Flush()

	FatalError(err)

	// Revert any changes done to buffer that have
	// not yet been written to file with Flush()
	// We just flushed, so there are no changes to revert
	// The writer that you pass as an argument
	// is where the buffer will output to, if you want
	// to change to a new writer

	bufferedWriter.Reset(bufferedWriter)

	// See how much buffer is available
	 bytesAvailable = bufferedWriter.Available()

	log.Printf("Available buffer: %d\n", bytesAvailable)


	// Resize buffer. The firs argument is a writer
	// where the buffer should output to. In this case
	// we are using the same buffer. If we chose a number
	// that was smaller than the existing buffer, like 10
	// we would not get back a buffer of size 10, we will
	// get back a buffer the size of the original since
	// it was already large enough (default 4096)
	bufferedWriter = bufio.NewWriterSize(bufferedWriter, 8000)

	// check available buffer size after resizing
	bytesAvailable = bufferedWriter.Available()

	log.Printf("Available buffer: %d\n", bytesAvailable)
}

func main() {
	//? Timeouts are for me to be able to read the out put of one method before the other,
	//? tho it isn't really necessary

	// createFile()
	// print("\n")
	// time.Sleep(time.Second)
	// truncateFile()
	// print("\n")
	// time.Sleep(time.Second)
	// getFileInfo()
	// print("\n")
	// time.Sleep(time.Second)
	// renameFile()
	// print("\n")
	// time.Sleep(time.Second)
	// deleteFile()
	// print("\n")
	// time.Sleep(time.Second)
	// readWriteFile()
	// print("\n")
	// time.Sleep(time.Second)
	// changePermissions()
	// print("\n")
	// time.Sleep(time.Second)
	// HardLinkFiles()
	// print("\n")
	// time.Sleep(time.Second)
	// SymLinkFiles()
	// print("\n")
	// time.Sleep(time.Second)
	// copyFile();
	// print("\n")
	// time.Sleep(time.Second)
	// seekPositionInFile()
	// print("\n")
	// time.Sleep(time.Second)
	// writeBytesToFile()
	// print("\n")
	// time.Sleep(time.Second)
	// quickWriteToFile()
	// print("\n")
	// time.Sleep(time.Second)
	BufferedWriter()
}
