package datastructure

import log "github.com/sirupsen/logrus"

// Configuration is a struct for manage the configuration related to the webserver
type Configuration struct {
	Port     string    // Port to bind the service
	Host     string    // Hostname to bind the service
	LogLevel log.Level // Level of logging
}
