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

var (
	instances     = make(map[string]*locale)
	defaultLocale = "zh-CN"
)

// NewLocale will create or return existing localized instances.
func NewLocale(language string, path string) (l *locale, e error) {
	l = &locale{
		language: language,
		path:     path,
	}
	e = l.newInstance()
	return
}

// Create a localized instance.
// Returns if an instance of the language already exists.
func (l *locale) newInstance() error {
	if len(l.language) == 0 {
		l.language = defaultLocale
	}

	if len(l.path) == 0 {
		l.path = "locales/" + l.language + ".yaml"
	}

	// Returns an existing instance.
	if v, ok := instances[l.language]; ok {
		l.message = v.message
	} else {
		// Read yaml and create a new instance.
		bytes, e := ioutil.ReadFile(l.path)
		if e != nil {
			return errors.New("The locale file was not found: " + l.path)
		}
		l.message = &message{}
		e = yaml.Unmarshal(bytes, l.message)
		if e != nil {
			return e
		}
		instances[l.language] = l
	}
	return nil
}

// Get will return localized messages.
// Return a message instance if it is not a specific message, otherwise return the formatted string of the message.
func (l *locale) Get(flag string, args ...interface{}) interface{} {
	flags := strings.Split(flag, ".")
	r := (*l.message)[flags[0]]
	for i := 1; i < len(flags); i++ {
		r = r.(message)[flags[i]]
	}
	if reflect.TypeOf(r).Name() != "string" {
		return r
	}
	return parse(r, args)
}

// Group will return a message grouping for fast and efficient access to intra-group messages.
func (l *locale) Group(flag string) group {
	msg := l.Get(flag).(message)
	return group{message: &msg}
}

// Get will return a normal or formatted message string.
// Returns an empty string if the specified message does not exist.
func (g *group) Get(flag string, args ...interface{}) string {
	if v, ok := (*g.message)[flag]; ok {
		return parse(v, args)
	}
	return ""
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
