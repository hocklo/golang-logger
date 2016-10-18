package logger

// Author @hocklo 

import (
	"bytes"
	"fmt"
	"log"
)

const (
        INFO = "INFO"
        DEBUG = "DEBUG"
        TRACE = "TRACE"
        WARN = "WARN"
        AUDIT = "AUDIT"
        ERROR = "ERROR"
)

/**
 * Print a log INFO.
 */
func Info(m string) {
        // Invoke writeLog function with INFO level and message
        writeLog(INFO, m)
}

/**
 * Print a log DEBUG.
 */
func Debug(m string) {
        // Invoke writeLog function with DEBUG level and message
        writeLog(DEBUG, m)
}

/**
 * Print a log WARN.
 */
func Warn(m string) {
        // Invoke writeLog function with WARN level and message
        writeLog(WARN, m)
}

/**
 * Print a log TRACE.
 */
func Trace(m string) {
        // Invoke writeLog function with TRACE level and message
        writeLog(TRACE, m)
}

/**
 * Print a log ERROR.
 */
func Error(m string) {
        // Invoke writeLog function with ERROR level and message
        writeLog(ERROR, m)
}

/**
 * Print a log AUDIT.
 */
func Audit(m string) {
        // Invoke writeLog function with AUDIT level and message
        writeLog(AUDIT, m)
}

/**
 * Print &buf to log.
 */
func writeLog(level string, m string) {
        var buf bytes.Buffer // Instance buffer
        var logger = log.New(&buf, "logger::", (log.Ldate | log.Ltime | log.Lmicroseconds)) // Instance log with specific prefix and flags.
        logger.Print(level+"::"+m) // Print log to buffer
        fmt.Print(&buf) // throw log from buffer to file.
}
