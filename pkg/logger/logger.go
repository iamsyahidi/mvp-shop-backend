package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var log *logrus.Entry

func setDefault() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	log = logrus.WithField("app", "notification")
}

func GetLogger(pkg, funcName string) *logrus.Entry {
	if log == nil {
		setDefault()
	}
	return log.WithFields(logrus.Fields{
		"function": funcName,
		"package":  pkg,
	})
}

func GetLoggerContext(ctx context.Context, pkg, funcName string) *logrus.Entry {
	if log == nil {
		setDefault()
	}
	return log.WithContext(ctx).WithFields(logrus.Fields{
		"function": funcName,
		"package":  pkg,
	})
}

func SingleTrace(ioFunction string, data map[string]interface{}) {
	logrus.SetFormatter(loggerFormat())
	field := logrus.Fields{
		"data":     data,
		"function": ioFunction,
	}
	log.WithFields(field).Info("traces")
}

func loggerFormat() *logrus.JSONFormatter {
	return &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "@timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	}
}

func GetFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}

func LogRequest(customer string, r *http.Request, body string) {
	logrus.SetFormatter(loggerFormat())

	field := logrus.Fields{
		"customer": customer,
		"body":     body,
		"url":      r.URL.RequestURI(),
		"device":   r.Header.Get("Grpc-metadata-device"),
	}
	log.WithFields(field).Println()
}

func LogResponse(customer string, response string) {
	logrus.SetFormatter(loggerFormat())

	field := logrus.Fields{
		"customer": customer,
		"response": response,
	}
	logrus.WithFields(field).Println()
}

// EscChar constant
var EscChar = "\x1B"

// ResetChar constant
var ResetChar = EscChar + "[0m"

// DateFormat date format
type DateFormat = string

const (
	// DefaultDateFormat for Default Date Format
	DefaultDateFormat DateFormat = "Y-m-d"

	// INDateFormat for Indonesian Date Format
	INDateFormat = "d-m-Y"

	// DefaultDateTimeFormat for Default Date Time Format
	DefaultDateTimeFormat = "Y-m-d H:i:s"

	// INDateTimeFormat for Indonesian Date Time Format
	INDateTimeFormat = "d-m-Y H:i:s"

	// DefaultDateTimeWithTimezoneFormat for Default Date Time Format With Timezone
	DefaultDateTimeWithTimezoneFormat = "Y-m-d H:i:s TZ"

	// DefaultTimeFormat for Default Time Format
	DefaultTimeFormat = "H:i:s"
)

const (
	// Year4Digits for Years in 4 digits
	Year4Digits string = "2006"

	// Month2Digits for Months in 2 digits
	Month2Digits = "01"

	// Day2Digits for Days in 2 digits
	Day2Digits = "02"

	// Hour2Digits for Hours in 2 digits
	Hour2Digits = "15"

	// Minute2Digits for Minutes in 2 digits
	Minute2Digits = "04"

	// Second2Digits for Second in 2 digits
	Second2Digits = "05"

	// Timezone for Timezone Location
	Timezone = "MST"
)

// ErrorStruct object
type ErrorStruct struct {
	_error    error
	_key      string
	_comments []string
	_code     int

	File string
	Line int
	Fn   string
}

type LogBaseStruct struct {
	FunctionName interface{}
	Request      interface{}
	Response     interface{}
}

type SError = *ErrorStruct

type Color int

type ColorType int

const (
	// FOREGROUND color type
	FOREGROUND ColorType = 1 + iota

	// BACKGROUND color type
	BACKGROUND
)

const (
	// DEFAULT color
	DEFAULT Color = 1 + iota

	// BLACK color
	BLACK

	// RED color
	RED

	// GREEN color
	GREEN

	// GREEN color
	YELLOW

	// BLUE color
	BLUE

	// MAGENTA color
	MAGENTA

	// CYAN color
	CYAN

	// LIGHT_GRAY color
	LIGHT_GRAY

	// DARK_GRAY color
	DARK_GRAY

	// LIGHT_RED color
	LIGHT_RED

	// LIGHT_GREEN color
	LIGHT_GREEN

	// LIGHT_YELLOW color
	LIGHT_YELLOW

	// LIGHT_BLUE color
	LIGHT_BLUE

	// LIGHT_MAGENTA color
	LIGHT_MAGENTA

	// LIGHT_CYAN color
	LIGHT_CYAN

	// WHITE color
	WHITE
)

// ParseToGoFormat to parsing format to go format
func ParseToGoFormat(format DateFormat) string {
	rl := map[string]string{
		"Y":  Year4Digits,
		"m":  Month2Digits,
		"d":  Day2Digits,
		"H":  Hour2Digits,
		"i":  Minute2Digits,
		"s":  Second2Digits,
		"TZ": Timezone,
	}

	for k, v := range rl {
		format = strings.ReplaceAll(format, k, v)
	}
	return format
}

// GetStandardFormat function
func GetStandardFormat() string {
	return "In %s[%s:%d] %s%s"
}

// GetStandardColorFormat function
func GetStandardColorFormat() string {
	frmt := ""
	frmt += ApplyForeColor("In", DARK_GRAY) + " "
	frmt += ApplyForeColor("%s", LIGHT_YELLOW)
	frmt += ApplyForeColor("[", DARK_GRAY)
	frmt += ApplyForeColor("%s:%d", MAGENTA)
	frmt += ApplyForeColor("]", DARK_GRAY)
	frmt += " %s%s"
	return frmt
}

func GetColorCode(c Color, t ColorType) (string, bool) {
	fcs := map[Color][]string{
		DEFAULT:       []string{"39", "49"},
		BLACK:         []string{"30", "40"},
		RED:           []string{"31", "41"},
		GREEN:         []string{"32", "42"},
		YELLOW:        []string{"33", "43"},
		BLUE:          []string{"34", "44"},
		MAGENTA:       []string{"35", "45"},
		CYAN:          []string{"36", "46"},
		LIGHT_GRAY:    []string{"37", "47"},
		DARK_GRAY:     []string{"90", "100"},
		LIGHT_RED:     []string{"91", "101"},
		LIGHT_GREEN:   []string{"92", "102"},
		LIGHT_YELLOW:  []string{"93", "103"},
		LIGHT_BLUE:    []string{"94", "104"},
		LIGHT_MAGENTA: []string{"95", "105"},
		LIGHT_CYAN:    []string{"96", "106"},
		WHITE:         []string{"97", "107"},
	}

	cc, ok := fcs[c]
	if ok {
		if t == FOREGROUND {
			return cc[0], true
		} else {
			return cc[1], true
		}
	}
	return "", false
}

func ApplyForeColor(s string, c Color) string {
	col, ok := GetColorCode(c, FOREGROUND)
	if ok {
		return EscChar + "[" + col + "m" + s + ResetChar
	}
	return s
}

func getMessage(v interface{}) (string, string) {
	m := fmt.Sprintf("%v", v)
	m2 := m

	switch v.(type) {
	case error:
		pc, fn, line, _ := runtime.Caller(2)
		m = fmt.Sprintf(GetStandardFormat(), runtime.FuncForPC(pc).Name(), fn, line, m)
		m2 = fmt.Sprintf(GetStandardColorFormat(), runtime.FuncForPC(pc).Name(), fn, line, m)
		break

	case SError:
		v2 := v.(*ErrorStruct)
		m = v2.ToFormatedString()
		m2 = v2.ToFormatedColorString()
		break
	}

	if !isLocal() {
		m2 = m
	}

	return m, m2
}

// ToFormatedColorString function
func (ox ErrorStruct) ToFormatedColorString() string {
	return fmt.Sprintf(GetStandardColorFormat(), ox.fParams()...)
}

func (ox ErrorStruct) fParams() []interface{} {
	pars := []interface{}{
		ox.Fn,
		ox.File,
		ox.Line,
		"",
		ox.String(),
	}
	if ox._comments != nil && len(ox._comments) > 0 {
		pars[3] = fmt.Sprintf("%s, details: ", ox.Comments())
	}
	return pars
}

// Comments to get error comments
func (ox ErrorStruct) Comments() string {
	return strings.Join(ox._comments, " < ")
}

// String function
func (ox ErrorStruct) String() string {
	return fmt.Sprintf("%v", ox._error)
}

// ToFormatedString function
func (ox ErrorStruct) ToFormatedString() string {
	return fmt.Sprintf(GetStandardFormat(), ox.fParams()...)
}

func print(msg string) {
	fmt.Fprintf(os.Stdout, "[%s] %s\n", ToString(DefaultDateTimeWithTimezoneFormat, time.Now()), msg)
}

// ToString to convert time to string
func ToString(format DateFormat, value time.Time) string {
	return value.Format(ParseToGoFormat(format))
}

func printErr(msg string) {
	fmt.Fprintf(os.Stderr, "[%s] %s\n", ToString(DefaultDateTimeWithTimezoneFormat, time.Now()), msg)
}

func isLocal() bool {
	return true
}

func applyForeColor(txt string, clr Color) string {
	if isLocal() {
		return ApplyForeColor(txt, clr)
	}
	return txt
}

// Info to logging info level
func Info(msg interface{}) {
	m := fmt.Sprintf("%v", msg)
	print(applyForeColor("INFO", LIGHT_BLUE) + ": " + m)
}

// Infof to logging info level with function
func Infof(msg string, args ...interface{}) {
	Info(fmt.Sprintf(msg, args...))
}

// Log to logging log level
func Log(name string, msg interface{}) {
	m := fmt.Sprintf("%v", msg)
	print(applyForeColor(name, LIGHT_CYAN) + ": " + m)
}

// Logf to logging log level with function
func Logf(msg string, args ...interface{}) {
	Log("", fmt.Sprintf(msg, args...))
}

// Warn to logging warning level
func Warn(msg interface{}) {
	_, m2 := getMessage(msg)
	print(applyForeColor("WARN", LIGHT_YELLOW) + ": " + m2)
}

// Warnf to logging warning level with function
func Warnf(msg string, args ...interface{}) {
	Warn(fmt.Sprintf(msg, args...))
}

// Err to logging error level
func Err(msg interface{}) {
	_, m2 := getMessage(msg)
	printErr(applyForeColor(" ERR", RED) + ": " + m2)
}

// Errf to logging error level with function
func Errf(msg string, args ...interface{}) {
	Err(fmt.Sprintf(msg, args...))
}

func LogBase(in LogBaseStruct) interface{} {
	getLogFunctionName, _ := json.Marshal(in.FunctionName)
	getLogRequest, _ := json.Marshal(in.Request)
	getLogResponse, _ := json.Marshal(in.Response)

	Log("Function Name", string(getLogFunctionName))
	Log("Request", string(getLogRequest))
	Log("Response", string(getLogResponse))
	return in
}

func ErrorFormat(message string, err error) error {
	return fmt.Errorf("[ asliri usecase error | %s ] %s", message, err.Error())
}
