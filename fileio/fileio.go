// This entire file follows this guide: https://www.devdungeon.com/content/working-files-go

package main

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func ExistsOrCreate(path string) *os.File {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		fmt.Println("File doesn't exist, creating...")
		file, _ := os.Create(path)
		fmt.Println("File Created")
		return file
	}

	return nil
}

func FatalError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func CreateFile() {
	newFile, err := os.Create("creation.txt")

	FatalError(err)

	log.Printf("%+v\n", newFile)
	newFile.Close()
}

func TruncateFile() {
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

func GetFileInfo() {
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

func RenameFile() {
	fmt.Println("File Rename")

	originalPath := "test.txt"

	ExistsOrCreate(originalPath)

	newPath := "test2.txt"

	fmt.Println("Renaming", originalPath, "to ", newPath)
	err := os.Rename(originalPath, newPath)

	FatalError(err)

	fmt.Println("Rename complete")
}

func DeleteFile() {
	fmt.Println("File Deletion")

	path := "deletion.txt"

	ExistsOrCreate(path)

	fmt.Println("Deleting File...")
	// time.Sleep(time.Minute) // Uncomment to see deletion in action after a minute
	err := os.Remove(path)

	FatalError(err)

	fmt.Println("File Deleted Successfully")
}

func SeekFile() {
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

func ReadWriteFile() {
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

func CopyFile() {
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

func SeekPositionInFile() {
	path := "seekPosition.txt"

	ExistsOrCreate(path)

	file, _ := os.Open(path)
	defer file.Close()

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

func WriteBytesToFile() {
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

func QuickWriteToFile() {
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
		[]byte{65, 66, 67},
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
	bytesAvailable := bufferedWriter.Available()

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

func ReadUpNBytes() {
	path := "readUpNBytes.txt"

	ExistsOrCreate(path)
	// Open file for reading

	file, err := os.Open(path)

	FatalError(err)

	defer file.Close()

	// Read up to len(b) bytes from the File
	// Zero bytes written means end of file
	// End of file returns error type io.EOF
	bytesSlice := make([]byte, 16) // will read up to the 16th byte within the file
	bytesRead, err := file.Read(bytesSlice)

	FatalError(err)

	log.Printf("Number of bytes read: %d\n", bytesRead)
	log.Printf("Data read: %s\n", bytesSlice)
}

func ReadExactlyNBytes() {
	path := "readExactlyNBytes.txt"

	ExistsOrCreate(path)

	//Open File
	file, err := os.Open(path)

	FatalError(err)

	defer file.Close()

	// The file.Read() function will happily read a tiny file in to a large
	// byte slice, but io.ReadFull() will return an
	// error if the file is smaller than the byte slice
	bytesSlice := make([]byte, 2) // 2 is the number of bytes that will be read
	numBytesRead, err := io.ReadFull(file, bytesSlice)

	FatalError(err)

	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", bytesSlice)
}

func ReadAtLeastNBytes() {
	path := "ReadAtLeastNBytes.txt"

	ExistsOrCreate(path)

	// Open file for reading
	file, err := os.Open(path)

	FatalError(err)

	byteSlice := make([]byte, 512)

	minBytes := 8

	// io.ReadAtLeast() will return an error if it cannot
	// find at least minBytes to read. It will read as
	// many bytes as byteSlice can hold.
	numBytesRead, err := io.ReadAtLeast(file, byteSlice, minBytes)

	FatalError(err)

	log.Printf("Number of bytes read: %d\n", numBytesRead)
	log.Printf("Data read: %s\n", byteSlice)
}

func ReadAllBytesOfFile() {
	path := "ReadAllBytesOfFile.txt"

	ExistsOrCreate(path)

	// Open file for reading
	file, err := os.Open(path)

	FatalError(err)

	defer file.Close()

	// os.File.Read(), io.ReadFull(), and
	// io.ReadAtLeast() all work with a fixed
	// byte slice that you make before you read

	// io.ReadAll() will read every byte
	// from the read (in this case a file)
	// and return a slice of unknown slice
	data, err := io.ReadAll(file)

	FatalError(err)

	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Printf("Number of bytes read: %d\n", len(data))
}

func QuickReadFileIntoMemory() {
	path := "QuickReadFileIntoMemory.txt"

	ExistsOrCreate(path)

	data, err := os.ReadFile(path)

	FatalError(err)

	log.Printf("Data read: %s\n", data)
}

func BufferedReader() {
	path := "BufferedReader.txt"

	ExistsOrCreate(path)

	// Open file and create a buffered read on top
	file, err := os.Open(path)

	FatalError(err)

	defer file.Close()

	bufferedReader := bufio.NewReader(file)

	// Get bytes without advancing pointer
	byteSlice := make([]byte, 5)
	byteSlice, err = bufferedReader.Peek(len(byteSlice))

	FatalError(err)

	fmt.Printf("Peeked at 5 bytes: %s\n", byteSlice)

	// Read and advance pointer
	numBytesRead, err := bufferedReader.Read(byteSlice)

	FatalError(err)

	fmt.Printf("Read %d bytes: %s\n", numBytesRead, byteSlice)

	// Read 1 byte. Error if no byte to read
	myByte, err := bufferedReader.ReadByte()

	FatalError(err)

	fmt.Printf("Read 1 byte: %c\n", myByte)

	// Read up to and including delimiter
	// Returns byte slice
	dataBytes, err := bufferedReader.ReadBytes('\n')

	FatalError(err)

	fmt.Printf("Read string: %s\n", dataBytes)

	// Read up to and including delimiter
	// Returns string

	dataString, err := bufferedReader.ReadString('\n')

	FatalError(err)

	fmt.Printf("Read string: %s\n", dataString)
	// This example reads a few lines so test.txt
	// should have a few lines of text to work correctly
}

func ReadWithScanner() {
	path := "ReadWithScanner.txt"

	ExistsOrCreate(path)

	file, err := os.Open(path)

	FatalError(err)

	scanner := bufio.NewScanner(file)

	// Default scanner is bufio.ScanLines. Lets use ScanWords.
	// Could also use a custom function of SplitFunc type
	scanner.Split(bufio.ScanWords)

	// Scan for next token.
	success := scanner.Scan()

	if !success {
		// False on error or EOF. Check error
		err = scanner.Err()
		if err != nil {
			log.Println("Scan Completed and Reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	// Get data from scan with Bytes() or Text()
	fmt.Println("First word found: ", scanner.Text())
	// Call scanner.Scan() again to find next token
}

func ArchiveFiles() {
	path := "ArchiveFiles.zip"

	// Create a file to write archive buffer to
	// Could also use an in memory buffer.
	outfile := ExistsOrCreate(path)

	defer outfile.Close()

	// Create a zip writer on top of the file writer
	zipWriter := zip.NewWriter(outfile)

	// Add files to archive
	// we use some hard coded data to demonstrate,
	// but you could iterate through all the files
	// in a directory and pass the name and content
	// of each file, or you can take data from your
	// program and write it in to the archive without

	var filesToArchive = []struct {
		Name, Body string
	}{
		{"ArchiveFilesT1.txt", "String contents of file"},
		{"ArchiveFilesT2.txt", "\x61\x62\x63\n"},
	}

	// Create and write files to the archive, which in turn
	// are getting written to the underlying writer to the
	// .zip file we created at the beginning
	for _, file := range filesToArchive {
		fileWriter, err := zipWriter.Create(file.Name)

		FatalError(err)

		_, err = fileWriter.Write([]byte(file.Body))
		FatalError(err)
	}

	// Clean up
	err := zipWriter.Close()

	FatalError(err)
}

func ExtractArchivedFiles() {
	// Create a reader out of the zip archive
	zipReader, err := zip.OpenReader("ArchiveFiles.zip")

	FatalError(err)

	defer zipReader.Close()

	// Iterate through each file/dir found in
	for _, file := range zipReader.Reader.File {
		// Open the file inside the zip archive
		// like a normal file
		zippedFile, err := file.Open()

		FatalError(err)

		defer zippedFile.Close()

		// Specify what the extracted file name should be.
		// You can specify a full path or a prefix
		// to move it to a different directory.
		// In this case, we will extract the file from
		// the zip to a file of the same name.
		targetDir := "./"
		extractedFilePath := filepath.Join(targetDir, file.Name)

		// Extract the item (or create directory)
		if file.FileInfo().IsDir() {
			// Create directory to recreate directory
			// structure inside the zip archive. Also
			// preserve permissions
			log.Println("Creating directory: ", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			// Extract regular file since it is not a directory
			log.Println("Extracting file: ", file.Name)

			// Open an output file for writing
			outputFile, err := os.OpenFile(extractedFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())

			FatalError(err)

			defer outputFile.Close()

			// "Extract" the file by copying zipped file
			// contents to the output file
			_, err = io.Copy(outputFile, zippedFile)

			FatalError(err)
		}
	}
}

func CompressFile() {
	path := "CompressFile.txt.gz"
	// Create .gz file to write to

	outputFile := ExistsOrCreate(path)

	// Create a gzip writer on top of the file writer
	gzipWriter := gzip.NewWriter(outputFile)
	defer gzipWriter.Close()

	// When we write to the gzip writer
	// it will in turn compress the contents
	// and then write it to the underlying
	// file writer as well
	// We don't have to worry about how all
	// the compression works since we just
	// use it as a simple writer interface
	// that we send bytes to

	_, err := gzipWriter.Write([]byte("Gophers rule!\n"))

	FatalError(err)

	log.Println("Compressed data written to file.")
}

func UncompressFile() {
	// Open gzip file we want to uncompress
	// the file is a reader, but we could use any
	// data source. It is common for web servers
	// to return gzipped contents to save bandwidth
	// and in that case the data is not in a file
	// on the file system but in a memory buffer
	gzipFile, err := os.Open("CompressFile.txt.gz")

	FatalError(err)

	// Create a gzip reader on top of the file reader
	// Again, it could be any type reader though
	gzipReader, err := gzip.NewReader(gzipFile)

	FatalError(err)

	defer gzipReader.Close()

	// Uncompress to a writer. We'll use a file writer
	outfileWriter, err := os.Create("UncompressedFile.txt")

	FatalError(err)

	defer outfileWriter.Close()

	// Copy contents of gzipped file to output file
	_, err = io.Copy(outfileWriter, gzipReader)

	FatalError(err)
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
	// BufferedWriter()
	// print("\n")
	// time.Sleep(time.Second)
	// ReadUpNBytes()
	// print("\n")
	// time.Sleep(time.Second)
	// ReadExactlyNBytes()
	// print("\n")
	// time.Sleep(time.Second)
	// ReadAtLeastNBytes()
	// print("\n")
	// time.Sleep(time.Second)
	// ReadAllBytesOfFile()
	// print("\n")
	// time.Sleep(time.Second)
	// QuickReadFileIntoMemory()
	// print("\n")
	// time.Sleep(time.Second)
	// BufferedReader()
	// print("\n")
	// time.Sleep(time.Second)
	// ReadWithScanner()
	// print("\n")
	// time.Sleep(time.Second)
	// ArchiveFiles()
	// print("\n")
	// time.Sleep(time.Second)
	// ExtractArchivedFiles()
	// print("\n")
	// time.Sleep(time.Second)
	// CompressFile()
	// print("\n")
	// time.Sleep(time.Second)
	UncompressFile()
}
