package main

import (
	"flag"
	"strings"

	datastructure "github.com/IBM/go-web-app/datastructure"
	routers "github.com/IBM/go-web-app/routers"

	// "gowebapp/plugins" if you create your own plugins import them here

	"github.com/onrik/logrus/filename"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	var conf *datastructure.Configuration
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "15-01-2018 15:04:05.00000"
	Formatter.FullTimestamp = true
	Formatter.ForceColors = true
	log.AddHook(filename.NewHook()) // Print filename + line at every log line
	log.SetFormatter(Formatter)
	log.SetLevel(log.DebugLevel)

	conf = verifyCommandLineInput()
	log.SetLevel(conf.LogLevel)

	router := gin.Default()
	router.RedirectTrailingSlash = false

	router.LoadHTMLGlob("public/*.html")
	router.Use(static.Serve("/", static.LocalFile("./public", false)))
	router.GET("/", routers.Index)
	router.NoRoute(routers.NotFoundError)
	router.GET("/500", routers.InternalServerError)
	router.GET("/health", routers.HealthGET)

	log.Info("Starting gowebapp on port-> ", conf.Port)

	err := router.Run()
	if err != nil {
		log.Error(`Unable to spaw server | Configuration -> `, conf)
		log.Fatal(`Unable to spaw server | Error -> `, err)
	}
}

// verifyCommandLineInput is delegated to manage the inputer parameter provide with the input flag from command line
func verifyCommandLineInput() *datastructure.Configuration {
	log.Debug("VerifyCommandLineInput | START")
	// This folder contains the two folder related to the CODE and the DOC
	port := flag.String("port", "", "Port to bind the service")
	host := flag.String("host", "", "Host to bind the service")
	logLevel := flag.String("logLevel", "debug", "Logging level [Panic|Fatal|Error|Warn|Info|Debug|Trace]")
	flag.Parse()
	// Verify if the mandatory parameter are provided
	if isBlank(*logLevel) {
		//flag.PrintDefaults()
		log.Warning("Be sure to choose the right logging level! (use --logLevel info)")
		*logLevel = "debug"
	}
	log.Debug("VerifyCommandLineInput | Starting command line input validation ..")
	/* OK, now we are sure that the folder exist */
	if isBlank(*port) { // TODO: If no port selected, generate select a random one from 8080 to 8090
		*port = "8080"
		log.Warning("VerifyCommandLineInput | Use '-port 8081' to bind the service on the port 8081 | Binded @" + *port)
	}
	if isBlank(*host) {
		*host = "localhost" // If no host provided set localhost
		log.Warning("VerifyCommandLineInput | Use '-host localhost' for bind the service to 127.0.0.1 | Binded @" + *host)
	}
	log.Info("Port: ", *port, " | Host: ", *host, " | LogLevel: ", logLevel)
	log.Debug("VerifyCommandLineInput | STOP")
	return &datastructure.Configuration{Port: *port, Host: *host, LogLevel: setDebugLevel(*logLevel)}
}

// setDebugLevel return the Logrus log level object by the given string
func setDebugLevel(level string) log.Level {
	switch strings.ToLower(level) {
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	case "panic":
		return log.PanicLevel
	case "warn", "warning":
		return log.WarnLevel
	}
	return log.DebugLevel
}

// isBlank is delegated to verify that the input string does not contains only empty char
func isBlank(str string) bool {
	// Check length
	if len(str) > 0 {
		// Iterate string
		for i := range str {
			// Check about char different from whitespace
			if str[i] > 32 {
				return false
			}
		}
	}
	return true
}
