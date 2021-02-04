package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Level identifies how verbose logging should be.
type Level int

var (
	globalLevel = LvlInfo
)

const (
	// LvlTrace is for highly verbose trace level logging
	LvlTrace Level = iota
	// LvlDebug is for highly verbose logging
	LvlDebug
	// LvlInfo is for informational logging. This is the default log level.
	LvlInfo
	// LvlWarning is for messages that may indicate a problem but are recoverable.
	LvlWarning
	// LvlError is for messages that indicate a problem occurred that isn't fatal.
	LvlError
	// LvlFatal is for a problem that is non-recoverable. This also terminates the app.
	LvlFatal
)

// Channel identifies a logging channel.
type Channel struct {
	name  string
	level Level
	log   *log.Logger
}

// Trace logs Trace level messages
func (c *Channel) Trace(args ...interface{}) {
	if c.level <= LvlTrace {
		return
	}

	c.log.Println(append([]interface{}{"[Trace](", c.name, ")"}, args...)...)
}

// Tracef logs Trace level messages
func (c *Channel) Tracef(msg string, args ...interface{}) {
	if c.level <= LvlTrace {
		return
	}

	c.Trace(fmt.Sprintf(msg, args...))
}

// Debug logs Debug level messages
func (c *Channel) Debug(args ...interface{}) {
	if c.level <= LvlDebug {
		return
	}

	c.log.Println(append([]interface{}{"[Debug](", c.name, ")"}, args...)...)
}

// Debugf logs Debug level messages
func (c *Channel) Debugf(msg string, args ...interface{}) {
	if c.level <= LvlDebug {
		return
	}

	c.Debug(fmt.Sprintf(msg, args...))
}

// Info logs Info level messages
func (c *Channel) Info(args ...interface{}) {
	if c.level <= LvlInfo {
		return
	}

	c.log.Println(append([]interface{}{"[Info](", c.name, ")"}, args...)...)
}

// Infof logs Info level messages
func (c *Channel) Infof(msg string, args ...interface{}) {
	if c.level <= LvlInfo {
		return
	}

	c.Info(fmt.Sprintf(msg, args...))
}

// Warning logs Warning level messages
func (c *Channel) Warning(args ...interface{}) {
	if c.level <= LvlWarning {
		return
	}

	c.log.Println(append([]interface{}{"[Warning](", c.name, ")"}, args...)...)
}

// Warningf logs Warning level messages
func (c *Channel) Warningf(msg string, args ...interface{}) {
	if c.level <= LvlWarning {
		return
	}

	c.Warning(fmt.Sprintf(msg, args...))
}

// Error logs Error level messages
func (c *Channel) Error(args ...interface{}) {
	if c.level <= LvlError {
		return
	}

	c.log.Println(append([]interface{}{"[Error](", c.name, ")"}, args...)...)
}

// Errorf logs Error level messages
func (c *Channel) Errorf(msg string, args ...interface{}) {
	if c.level <= LvlError {
		return
	}

	c.Error(fmt.Sprintf(msg, args...))
}

// Fatal logs Fatal level messages
func (c *Channel) Fatal(args ...interface{}) {
	if c.level <= LvlFatal {
		return
	}

	c.log.Fatalln(append([]interface{}{"[Fatal](", c.name, ")"}, args...)...)
}

// Fatalf logs Fatal level messages
func (c *Channel) Fatalf(msg string, args ...interface{}) {
	if c.level <= LvlFatal {
		return
	}

	c.Fatal(fmt.Sprintf(msg, args...))
}

// WithLogger sets the logger to use.
func (c *Channel) WithLogger(l *log.Logger) {
	c.log = l
}

// WithOutput sets the io.Writer to use.
func (c *Channel) WithOutput(w io.Writer) {
	c.log.SetOutput(w)
}

// WithLevel sets the logging level to use.
func (c *Channel) WithLevel(l Level) {
	c.level = l
}

// NewChannel returns a logging channel with the given name.
func NewChannel(name string) *Channel {
	return &Channel{
		name:  name,
		level: LvlInfo,
		log:   log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile),
	}
}
