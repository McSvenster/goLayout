package main

import (
	"flag"
	"log"
	"os"
)

var (
	configPath = flag.String("c", "/etc/MyApp/conf.d/", "use -c to privde a custom path to the config directory (default:/etc/MyApp/conf.d/)")
	config     *ApplicationConfig
	secrets    *[]Secret
)

type Secret struct {
	Name     string `yaml:"Name"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

type ApplicationConfig struct {
	API struct {
		URL string `yaml:"URL"`
	} `yaml:"API"`
	Mail struct {
		Host    string `yaml:"Host"`
		Address string `yaml:"Address"`
	} `yaml:"Mail"`
	TempDir   string `yaml:"TempDir"`
	LogFolder string `yaml:"LogFolder"`
	LogLevel  string `yaml:"LogLevel"`
}

func main() {
	flag.Parse()
	setupConfig(configPath)

	// now setup logging
	logfilewithpath := SetupLogfilePath(config.LogFolder)
	logfile, err := os.OpenFile(logfilewithpath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("application is set up, will now enter the loop and log to logfile: " + logfilewithpath)
	log.SetOutput(logfile)
	LogIt("INFO", "TaskHandler started")
}
