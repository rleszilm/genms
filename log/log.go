package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

// Level identifies how verbose logging should be.
type Level int

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
	// LvlDisable is for when all logging on the channel should be disabled. This does not impact
	// logging on the fatal or panic levels.
	LvlDisable
	// LvlFatal is for a problem that is non-recoverable. This also terminates the app.
	LvlFatal
	// LvlPanic is for a problem that can be recovered from at the apps discretion. This
	// causes a pannic that may result in app termination.
	LvlPanic
)

var (
	channelMux  sync.Mutex
	globalLevel = LvlInfo
	channels    = map[string]*Channel{}
)

// Logger the interface for the logger that a channel writes to.
type Logger interface {
	Println(...interface{})
	Fatalln(...interface{})
	Panicln(...interface{})
	SetFlags(int)
	SetOutput(io.Writer)
}

// Channel identifies a logging channel.
type Channel struct {
	name        string
	level       Level
	log         Logger
	nameTrace   string
	nameDebug   string
	nameInfo    string
	nameWarning string
	nameError   string
	nameFatal   string
	namePanic   string
	namePrint   string
}

// Trace logs Trace level messages
func (c *Channel) Trace(args ...interface{}) {
	if c.skipLog(LvlTrace) {
		return
	}

	c.log.Println(append([]interface{}{c.nameTrace}, args...)...)
}

// Tracef logs Trace level messages
func (c *Channel) Tracef(msg string, args ...interface{}) {
	if c.skipLog(LvlTrace) {
		return
	}

	c.Trace(fmt.Sprintf(msg, args...))
}

// Debug logs Debug level messages
func (c *Channel) Debug(args ...interface{}) {
	if c.skipLog(LvlDebug) {
		return
	}

	c.log.Println(append([]interface{}{c.nameDebug}, args...)...)
}

// Debugf logs Debug level messages
func (c *Channel) Debugf(msg string, args ...interface{}) {
	if c.skipLog(LvlDebug) {
		return
	}

	c.Debug(fmt.Sprintf(msg, args...))
}

// Info logs Info level messages
func (c *Channel) Info(args ...interface{}) {
	if c.skipLog(LvlInfo) {
		return
	}

	c.log.Println(append([]interface{}{c.nameInfo}, args...)...)
}

// Infof logs Info level messages
func (c *Channel) Infof(msg string, args ...interface{}) {
	if c.skipLog(LvlInfo) {
		return
	}

	c.Info(fmt.Sprintf(msg, args...))
}

// Warning logs Warning level messages
func (c *Channel) Warning(args ...interface{}) {
	if c.skipLog(LvlWarning) {
		return
	}

	c.log.Println(append([]interface{}{c.nameWarning}, args...)...)
}

// Warningf logs Warning level messages
func (c *Channel) Warningf(msg string, args ...interface{}) {
	if c.skipLog(LvlWarning) {
		return
	}

	c.Warning(fmt.Sprintf(msg, args...))
}

// Error logs Error level messages
func (c *Channel) Error(args ...interface{}) {
	if c.skipLog(LvlError) {
		return
	}

	c.log.Println(append([]interface{}{c.nameError}, args...)...)
}

// Errorf logs Error level messages
func (c *Channel) Errorf(msg string, args ...interface{}) {
	if c.skipLog(LvlError) {
		return
	}

	c.Error(fmt.Sprintf(msg, args...))
}

// Fatal logs Fatal level messages
func (c *Channel) Fatal(args ...interface{}) {
	c.log.Fatalln(append([]interface{}{c.nameFatal}, args...)...)
}

// Fatalf logs Fatal level messages
func (c *Channel) Fatalf(msg string, args ...interface{}) {
	c.Fatal(fmt.Sprintf(msg, args...))
}

// Panic logs Panic level messages
func (c *Channel) Panic(args ...interface{}) {
	c.log.Panicln(append([]interface{}{c.namePanic}, args...)...)
}

// Panicf logs Panic level messages
func (c *Channel) Panicf(msg string, args ...interface{}) {
	c.Panic(fmt.Sprintf(msg, args...))
}

// Print logs messages as long as the channel is not explicitly disabled.
func (c *Channel) Print(args ...interface{}) {
	c.log.Println(append([]interface{}{c.namePanic}, args...)...)
}

// Printf logs messages as long as the channel is not explicitly disabled.
func (c *Channel) Printf(msg string, args ...interface{}) {
	c.Print(fmt.Sprintf(msg, args...))
}

// WithFlags sets the flags on the underlying logger.
func (c *Channel) WithFlags(flags int) {
	c.log.SetFlags(flags)
}

// WithLogger sets the logger to use.
func (c *Channel) WithLogger(l Logger) {
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

func (c *Channel) skipLog(msg Level) bool {
	// always print fatal and panic messages
	if msg == LvlFatal || msg == LvlPanic {
		return false
	}

	// if the channel or the global system is disabled do not log
	if c.level == LvlDisable || globalLevel == LvlDisable {
		return true
	}

	// log at the global level unless the channel level is more permissive
	lvl := globalLevel
	if c.level < globalLevel {
		lvl = c.level
	}

	if msg < lvl {
		return true
	}
	return false
}

// LevelFromString returns the level from the given string.
func LevelFromString(lvl string) (Level, bool) {
	lvl = strings.ToLower(lvl)
	switch lvl {
	case "trace":
		return LvlTrace, true
	case "debug":
		return LvlDebug, true
	case "info":
		return LvlInfo, true
	case "warning":
		return LvlWarning, true
	case "error":
		return LvlError, true
	case "disable":
		return LvlDisable, true
	case "fatal":
		return LvlFatal, true
	case "panic":
		return LvlPanic, true
	}
	return LvlInfo, false
}

// SetLevel sets the global logging level. When logging the lower of the global and channel specific
// level is used. A channel level of debug and a global level of info will result in all
func SetLevel(l Level) {
	globalLevel = l
}

// SetChannelLevel sets the logging level of a specific channel if it exists.
func SetChannelLevel(name string, l Level) {
	channelMux.Lock()
	defer channelMux.Unlock()

	if c, ok := channels[name]; ok {
		c.WithLevel(l)
	}
}

// NewChannel returns a logging channel with the given name.
func NewChannel(name string) *Channel {
	channelMux.Lock()
	defer channelMux.Unlock()

	if c, ok := channels[name]; ok {
		return c
	}

	channels[name] = &Channel{
		name:        name,
		level:       LvlInfo,
		log:         log.New(os.Stderr, "", log.LstdFlags),
		nameTrace:   fmt.Sprintf("[Trace](%s):", name),
		nameDebug:   fmt.Sprintf("[Debug](%s):", name),
		nameInfo:    fmt.Sprintf("[Info](%s):", name),
		nameWarning: fmt.Sprintf("[Warning](%s):", name),
		nameError:   fmt.Sprintf("[Error](%s):", name),
		nameFatal:   fmt.Sprintf("[Fatal](%s):", name),
		namePanic:   fmt.Sprintf("[Panic](%s):", name),
		namePrint:   fmt.Sprintf("[Print](%s):", name),
	}
	return channels[name]
}
