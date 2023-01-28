package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

func setupConfig(configPath *string) {
	// init config
	// ===========
	yamlFile, err := os.ReadFile(*configPath + "taskhandler.yml")
	if err != nil {
		log.Fatalln("critical", "No config found at "+*configPath+", will stop now.", err)
	}

	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		log.Fatalln("ERROR parsing config", fmt.Sprint(err))
	}

	// load secrets
	// ============
	yamlFile, err = os.ReadFile(*configPath + ".secrets.yml")
	if err != nil {
		log.Fatalln("critical", "No .secrets file at "+*configPath+".secrets"+" found, will stop now.")
	}
	if err = yaml.Unmarshal(yamlFile, &secrets); err != nil {
		log.Fatalln("ERROR parsing secrets", fmt.Sprint(err))
	}
}

// Logging
// =======
func SetupLogfilePath(logFilePath string) string {
	filename := "taskhandler.log"
	// check, if logfile exists (eg after crash) and move it
	// set up regular log rotation with unix's logrotate
	// (e.g. https://medium.com/rahasak/golang-logging-with-unix-logrotate-41ec2672b439)
	if FileExists(logFilePath + filename) {

		today := time.Now().Format("2006-01-02")
		newfilename := filename + "_" + today
		if FileExists(logFilePath + newfilename) {
			counter := 0
			logfiles, err := os.ReadDir(logFilePath)
			if err != nil {
				LogIt("ERROR", fmt.Sprint(err))
			}
			for _, file := range logfiles {
				if strings.HasPrefix(file.Name(), newfilename) {
					counter++
				}
			}
			newfilename = newfilename + "." + fmt.Sprint(counter)
		}
		fmt.Println("logfile " + logFilePath + filename + " exists, will move it to " + logFilePath + newfilename)
		err := os.Rename(logFilePath+filename, logFilePath+newfilename)
		if err != nil {
			log.Fatal(err)
		}

	}
	return logFilePath + filename
}

func LogIt(level string, message string) {
	logLevel := config.LogLevel
	levels := []string{"DEBUG", "INFO", "WARNING", "ERROR"}
	logLevelIndex := GetStringSliceElementIndex(levels, logLevel)
	messageLevelIndex := GetStringSliceElementIndex(levels, level)
	if messageLevelIndex >= logLevelIndex {
		log.Println(level, message)
	}
}
