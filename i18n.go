package i18n

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"reflect"
	"strings"
)

type locale struct {
	language string
	path     string
	message  *message
}

type group struct {
	message *message
}

type message map[interface{}]interface{}

// Shrink default threshold size.
const threshold = 3

var (
	// cached instances.
	instances     [threshold]*locale
	DefaultLocale = "zh-CN"
	lastUsed      = -1
	errNoInstance = errors.New("must create an instance at least once")
)

// New func will create or return existing localize instances.
func New() (l *locale, e error) {
	l, e = NewLocale(DefaultLocale, "")
	return
}

// NewLocale will create or return existing localize instances.
func NewLocale(language string, path string) (l *locale, e error) {
	if len(language) == 0 {
		language = DefaultLocale
	}
	if len(path) == 0 {
		path = "locales/" + language + ".yaml"
	}
	if instance, index := findInstance(language); instance != nil && index != -1 {
		l = instance
		lastUsed = index
	} else {
		// read yaml and create a new instance.
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Print(err)
			e = errors.New("The locale file was not found: " + path)
		} else {
			l = &locale{
				language: language,
				path:     path,
				message:  &message{},
			}
			e = yaml.Unmarshal(bytes, l.message)
			if e == nil {
				cache(l)
			}
		}
	}
	return
}

// Get will return localized messages.
// Return a message instance if it is not a specific message, otherwise return the formatted string of the message.
func (l *locale) Get(flag string, args ...interface{}) interface{} {
	flags := strings.Split(flag, ".")
	val := (*l.message)[flags[0]]
	for i := 1; i < len(flags); i++ {
		val = val.(message)[flags[i]]
	}
	if val == nil {
		return ""
	}
	if reflect.TypeOf(val).Name() != "string" {
		return val
	}

	return parse(val, args)
}

// Group will return a message grouping for fast and efficient access to intra-group messages.
func (l *locale) Group(flag string) *group {
	msg := l.Get(flag).(message)
	return &group{message: &msg}
}

// Get will return a normal or formatted message string.
// Returns an empty string if the specified message does not exist.
func (g *group) Get(flag string, args ...interface{}) string {
	if v, ok := (*g.message)[flag]; ok {
		return parse(v, args)
	}
	return ""
}

// Get will return a message based on the last used instance.
func Get(flag string, args ...interface{}) (interface{}, error) {
	if lastUsed != -1 {
		return instances[lastUsed].Get(flag, args...), nil
	}
	return nil, errNoInstance
}

// Group returns a group based on the last used instance.
func Group(flag string) (*group, error) {
	if lastUsed != -1 {
		msg := instances[lastUsed].Get(flag).(message)
		return &group{message: &msg}, nil
	}
	return nil, errNoInstance
}

// Format the message or return the message text directly.
func parse(val interface{}, args []interface{}) string {
	if val == nil {
		return ""
	}
	if len(args) > 0 {
		return fmt.Sprintf(val.(string), args...)
	}
	return val.(string)
}

func findInstance(language string) (*locale, int) {
	for i, v := range instances {
		if v != nil && v.language == language {
			return v, i
		}
	}
	return nil, -1
}

func cache(l *locale) {
	cached := false
	for i := range instances {
		if instances[i] == nil {
			instances[i] = l
			cached = true
			lastUsed = i
		}
	}
	if !cached {
		end := threshold - 1
		for i := 0; i < end; i++ {
			instances[i] = instances[i+1]
		}
		instances[end] = l
		lastUsed = end
	}
}
