package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

// go_file_examples.go
// Clean, well-commented collection of common file-system operations in Go.
// Each function shows a safe, idiomatic way to perform an action and explains
// important reasons and pitfalls in comments above the code.
//
// Usage:
//  - Open this file, read the examples and uncomment **one** example inside main()
//    to try it. Don't run multiple destructive operations (like DeleteFile)
//    at the same time.
//  - Run from the directory that contains the files you want to operate on:
//      cd C:/golang/23_file
//      go run go_file_examples.go

// NOTE: This file purposely contains small, self-contained helper functions.
// They are written defensively (error checks, defer Close()) and use
// clear, explicit variable assignment when reusing variables to avoid
// "no new variables on left side of :=" compilation errors.

// ---------- Utility helpers ----------

// checkPanic is a small helper to stop execution and print the error.
// For real programs prefer returning errors and handling them higher up.
func checkPanic(err error) {
	if err != nil {
		panic(err)
	}
}

// fileExists checks whether a path exists and is not a directory.
func fileExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err == nil {
		return !info.IsDir(), nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// ---------- Reading ----------

// readSmallFile reads the entire file into memory and returns its content.
// Use this only for small files to avoid high memory usage.
func readSmallFile(path string) (string, error) {
	b, err := os.ReadFile(path) // reads whole file
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// readStream reads a file byte-by-byte using a buffered reader.
// This is suitable for streaming or large files where you don't want
// to load the entire contents into memory.
func readStream(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	var out []byte
	for {
		b, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		out = append(out, b)
	}
	return string(out), nil
}

// ---------- Writing ----------

// replaceFileContent overwrites the file completely with newContent.
// This is the simplest and recommended method when you want a full replace.
func replaceFileContent(path, newContent string) error {
	// os.WriteFile will create the file if it does not exist and
	// will overwrite it if it does exist. Permissions are used if the
	// file must be created.
	return os.WriteFile(path, []byte(newContent), 0644)
}

// appendToFile appends text to the end of the file, creating the file
// if it doesn't exist.
func appendToFile(path, text string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(text)
	return err
}

// ---------- Copying ----------

// copyFile streams data from srcPath to dstPath using io.Copy which is
// efficient and works well for large files.
func copyFile(srcPath, dstPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath) // creates or truncates
	if err != nil {
		return err
	}
	defer dst.Close()

	// io.Copy uses an internal buffer and is the standard way to copy streams.
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	// Ensure data is flushed to storage. os.File.Sync is more strict than
	// just relying on Close; use Sync when you need to guarantee write to disk.
	return dst.Sync()
}

// ---------- Directory listing ----------

// listDirectory lists files and folders at the given directory path.
func listDirectory(dirPath string) ([]os.DirEntry, error) {
	dir, err := os.Open(dirPath)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	entries, err := dir.ReadDir(-1) // -1 => read all entries
	if err != nil {
		return nil, err
	}
	return entries, nil
}

// ---------- Deleting ----------

// deleteFile removes a file. Be careful — this is destructive and irreversible.
func deleteFile(path string) error {
	// os.Remove removes the named file or (empty) directory.
	return os.Remove(path)
}

// ---------- Main: demo driver ----------

func main() {
	// For safety, the main contains commented demo calls. Uncomment one
	// action you want to try. This prevents accidental deletion or
	// overwriting while experimenting.

	// ---- 1) Read full small file (safe for small files) ----
	// content, err := readSmallFile("example.txt")
	// checkPanic(err)
	// fmt.Println("example.txt contents:\n", content)

	// ---- 2) Stream read (good for large files) ----
	contentStream, err := readStream("example.txt")
	checkPanic(err)
	fmt.Println("(stream) example.txt contents:\n", contentStream)

	// ---- 3) Replace file content (overwrites file) ----
	//err := replaceFileContent("example.txt", "Hello: replaced at "+time.Now().Format(time.RFC3339)+"\n")
	//checkPanic(err)
	//fmt.Println("Replaced content of example.txt")

	// ---- 4) Append to file ----
	//err := appendToFile("example.txt", "Appended line\n")
	//checkPanic(err)
	//fmt.Println("Appended to example.txt")

	// ---- 5) Copy file (stream copy) ----
	//err := copyFile("example.txt", "example_copy.txt")
	//checkPanic(err)
	//fmt.Println("Copied example.txt -> example_copy.txt")

	// ---- 6) List directory ----
	//entries, err := listDirectory(".")
	//checkPanic(err)
	//for _, e := range entries {
	//	fmt.Printf("%s  (dir? %v)\n", e.Name(), e.IsDir())
	//}

	// ---- 7) Delete file (DESTRUCTIVE: be careful) ----
	//err := deleteFile("example2.txt")
	//checkPanic(err)
	//fmt.Println("Deleted example2.txt")

	// ---- Helpful tip: check existence before dangerous ops ----
	//exists, err := fileExists("example2.txt")
	//checkPanic(err)
	//fmt.Println("example2.txt exists?", exists)

	// If you prefer a quick demonstration that does not modify your files,
	// uncomment the following read-only example which safely lists the current
	// directory and prints file names and sizes.

	entries, err := listDirectory(".")
	checkPanic(err)
	fmt.Println("Current directory listing:")
	for _, e := range entries {
		info, err := e.Info()
		if err != nil {
			// If Info fails for a single entry continue with next one
			fmt.Println("  (error reading info)", e.Name(), err)
			continue
		}
		fmt.Printf("  %-30s dir=%-5v size=%8d modified=%s\n",
			filepath.Clean(e.Name()), e.IsDir(), info.Size(), info.ModTime().Format(time.RFC3339))
	}

	// End of demo. Uncomment other examples above to experiment.

	// Final notes (important reasons / tips):
	//  - Use os.WriteFile for simple full replacements (atomicity: NOT atomic).
	//    If you require atomic replace consider writing to a temp file then
	//    renaming it to the final name.
	//  - Use io.Copy for efficient streaming/copying of file contents.
	//  - Always defer Close() on opened files to avoid resource leaks.
	//  - Be explicit with = vs := to avoid "no new variables on left side" errors.
}
