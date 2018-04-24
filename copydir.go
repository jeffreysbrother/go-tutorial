package main

 import (
   "os"
   "flag"
   "fmt"
   "io"
	 "log"
	 "path/filepath"
 )

 func CopyFile(source string, dest string) (err error) {
     sourcefile, err := os.Open(source)
     if err != nil {
         return err
     }

     defer sourcefile.Close()

     destfile, err := os.Create(dest)
     if err != nil {
         return err
     }

     defer destfile.Close()

     _, err = io.Copy(destfile, sourcefile)
     if err == nil {
         sourceinfo, err := os.Stat(source)
         if err != nil {
             err = os.Chmod(dest, sourceinfo.Mode())
         }
     }

     return
 }

 func CopyDir(source string, dest string) (err error) {

     // get properties of source dir
     sourceinfo, err := os.Stat(source)
     if err != nil {
         return err
     }

     // create dest dir
     err = os.MkdirAll(dest, sourceinfo.Mode())
     if err != nil {
         return err
     }

     directory, _ := os.Open(source)
     objects, err := directory.Readdir(-1)

     for _, obj := range objects {

         sourcefilepointer := source + "/" + obj.Name()
         destinationfilepointer := dest + "/" + obj.Name()

         if obj.IsDir() {
             // create sub-directories - recursively
             err = CopyDir(sourcefilepointer, destinationfilepointer)
             if err != nil {
                 fmt.Println(err)
             }
         } else {
             // perform copy
             err = CopyFile(sourcefilepointer, destinationfilepointer)
             if err != nil {
                 fmt.Println(err)
             }
         }

     }
     return
 }


 func main() {
   flag.Parse() // get the source and destination directory

	 // get current working directory
	 cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	  if err != nil {
	          log.Fatal(err)
	  }

		source := flag.Arg(0)
		dest := flag.Arg(1)

   source_dir := cwd + "/" + source // get the source directory from 1st argument
   dest_dir := cwd + "/" + dest // get the destination directory from the 2nd argument

	 if (source == "" || dest == "") {
		 fmt.Println("Please provide a source and destination directory. Aborting.")
		 os.Exit(1)
	 }

   // if source dir doesn't exist
   src, err := os.Stat(source_dir)
   if err != nil {
      fmt.Println("Source directory does not exist. Aborting.")
			os.Exit(1)
   }

	 // if source is not a directory
   if !src.IsDir() {
     fmt.Println("Source is not a directory.")
     os.Exit(1)
   }

	 fmt.Println("Source: " + source_dir)
		fmt.Println("Destination: " + dest_dir)

	 _, err = os.Open(dest_dir)
	 if !os.IsNotExist(err) {
		 fmt.Println("Destination directory already exists. Aborting.")
		 os.Exit(1)
	 }

   err = CopyDir(source_dir, dest_dir)
   if err != nil {
      fmt.Println(err)
   } else {
      fmt.Println("Directory copied.")
   }

 }
