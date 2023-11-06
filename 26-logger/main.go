package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// log.Println("error")
	log.SetFlags(log.Ldate | log.Lshortfile)
	log.Println("Sup ninjas!")

	// log.Panic("Panicking...")
	// log.Fatal("Uh-oh...")

	//output a file
	file, _ := os.Create("file.log")
	// log.SetOutput(file)
	// log.Println("Hello world")

	// log.SetOutput(os.Stdout)
	// log.Println("Printing into standard out again")

	//common loggers
	flags := log.LstdFlags | log.Lshortfile
	infoLogger := log.New(os.Stdout, "INFO: ", flags)
	warnLogger := log.New(os.Stdout, "WARN: ", flags)
	errorLogger := log.New(os.Stdout, "ERROR: ", flags)
	// infoLogger.Println("This is an info log")
	// warnLogger.Println("This is an warning log")
	// errorLogger.Println("This is an error log")

	// You can aggregate all three into one
	al := aggregatedLogger{
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
	}
	al.info("msg1")
	al.warn("warn")
	al.error("error")

	file.Close()
	readFile(file.Name())
}

type aggregatedLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func (l *aggregatedLogger) info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

func (l *aggregatedLogger) warn(v ...interface{}) {
	l.warnLogger.Println(v...)
}

func (l *aggregatedLogger) error(v ...interface{}) {
	l.errorLogger.Println(v...)
}

func readFile(filepath string) {
	dataByte, _ := os.ReadFile(filepath)
	fmt.Println("Text inside file:", string(dataByte))
}
