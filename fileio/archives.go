package main

import (
	"archive/zip"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
)


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
