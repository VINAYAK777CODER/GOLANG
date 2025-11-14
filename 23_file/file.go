package main

import (
	// "bufio"
	// "fmt"
	"os"
)

func main() {

	// 🔹 Try to open the file named "example.txt"
	//    If the file does NOT exist → err will NOT be nil.
	//----------------------------------------------------------//
	// file, err := os.Open("example.txt")

	// // 🔹 If error occurs (file not found or can't open)
	// //    panic() stops the program and prints the error.
	// if err != nil {
	// 	panic(err)
	// }

	// defer file.Close()
	//---------------------------------------------------------------------//
	// buff:=make([]byte,10)
	// length,err:=file.Read(buff)
	// if err!=nil {
	// 	panic(err)

	// }
	// fmt.Println("data",length,buff)
	// for i := 0; i < len(buff); i++ {
	// 	fmt.Println("data",length,string(buff[i]))
	// }

	// ---------------------------
	//  this is onnly useful for small files
	// data,err:=os.ReadFile("example.txt")
	// if(err!=nil){
	// 	panic(err)
	// }
	// fmt.Println(string(data))
	//------------------------------------//

	// read folders

	// dir,err:=os.Open("../") // . means current directry // ../ means one outer back directry
	// if err!=nil{
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfor,err:=dir.ReadDir(-1)

	// for _,fi:=range fileInfor{
	// 	fmt.Println(fi.Name(),fi.IsDir())
	// }
	//-----------------------------------------------------//

	// how i can write , first we will creat ethe file
	// f, err := os.Create("example2.txt") // if already present then  will it not create the new one??
	// if err != nil {
	// 	panic(err)
	// }
	// defer func() {
	// 	f.Close()
	// 	fmt.Println("file is closed now")
	// }()
	// writing the string
	// f.WriteString("hi jatin you are awesome")
	// f.WriteString(" ji kahiye aap") // it will append
	// fmt.Println("file wrting is done")

	// bytes:=[]byte("hello ji to kaise hai aap")
	// f.Write(bytes)

	//---------------------------------------------------------------//
	// to replace the content with new content
	// newContent := "Hello, this is the NEW content!"

	// // 📌 WriteFile overwrites the file if it already exists
	// err2 := os.WriteFile("example.txt", []byte(newContent), 0644)

	// if err1 != nil {
	// 	panic(err1)
	// }
	//---------------------------------------------------------------//
	//read and write to another file (streaming fasion) source file -> destination file
	// sourceFile,err:=os.Open("example.txt")
	// if err!=nil {
	// 	panic(err)

	// }
	// defer sourceFile.Close()
	// destFile,err:=os.Create("example2.txt")
	// if err!=nil {
	// 	panic(err)

	// }
	// defer destFile.Close()
	// reader:=bufio.NewReader(sourceFile)
	// writer:=bufio.NewWriter(destFile)
	// //to read
	// for {
	// 	b,err:=reader.ReadByte()
	// 	if err!=nil{
	// 		if err.Error()!="EOF"{ // reason of this condition -> file is not ended yet
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	err2:=writer.WriteByte(b)
	// 	if err2!=nil {
	// 		panic(err2)
	// 	}
	// }

	// writer.Flush()// reason??
	// fmt.Println("written to new file succes fully")
	//----------------------------------------------------------------//

	// how to delete a file
	err:=os.Remove("example2.txt")
	if err!=nil {
		panic(err)
		
	}

	//----------------------------------------------------------------//

	// // 🔹 file.Stat() returns metadata/info about the file
	// //    like name, size, permissions, modified time etc.
	// fileInfo, err := file.Stat()

	// // 🔹 Again, if error occurs while reading file info → stop
	// if err != nil {
	// 	panic(err)
	// }

	// // 🔹 Print file name (example: example.txt)
	// fmt.Println("file name :", fileInfo.Name())

	// // 🔹 Check if file is a directory (true/false)
	// fmt.Println("file or folder :", fileInfo.IsDir())

	// // 🔹 Print file size in bytes
	// fmt.Println("file size :", fileInfo.Size())

	// // 🔹 Print file permission bits (like -rw-r--r-- or 0644)
	// fmt.Println("file Permission :", fileInfo.Mode())

	// // 🔹 Print last modified time of the file
	// fmt.Println("file modfies at :", fileInfo.ModTime())
}

/*
⚠️ If you get this error:

panic: open example.txt: The system cannot find the file specified.

➡ It means Go cannot find the file in your current working directory.

✅ Solution:
Go to the folder where your Go file and example.txt are located
and run the program from there:

    cd C:/golang/23_file/
    go run file.go

Make sure "example.txt" exists in the same folder.
*/
