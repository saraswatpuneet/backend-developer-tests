package env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const tag string = "env"

var pattern = regexp.MustCompile(`(export\s*)?(?P<key>\w+)\s*(=)\s*(?P<value>("|')?[\w\s-$,:/\.\+]+)("|')?`)

// findKeyValue finds '=' separated key/value pair from the given string.
func findKeyValue(line string) (key, value string) {
	// Intentionally insert a linebreak here to avoid non-stop characters on [[:graph:]].
	line = fmt.Sprintf("%s\n", line)
	match := pattern.FindStringSubmatch(line)
	if len(match) != 0 {
		result := make(map[string]string)
		for index, name := range pattern.SubexpNames() {
			result[name] = match[index]
		}
		// preserve those quoted spaces for value.
		key, value = result["key"], strings.TrimSpace(result["value"])
		value = strings.Trim(value, "'")
		value = strings.Trim(value, "\"")
	}
	return
}

// Exists check if the given path exists.
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

// Load parses & set the values in the given files into os environment.
func Load(dotenv string) {
	if Exists(dotenv) {
		dotenv, _ = filepath.Abs(dotenv)

		if file, err := os.Open(dotenv); err == nil {
			defer file.Close()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				k, v := findKeyValue(scanner.Text())
				if k == "" && v == "" {
					continue
				}

				if err := Set(k, v); err != nil {
					log.Printf("Error on Env.Load %v", err)
				}
			}
		}
	}
}

// Get retrieves the string value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
func Get(key string) (value string, exists bool) {
	if v := os.Getenv(key); v != "" {
		value, exists = v, true
	}
	return
}

// Set stores the value of the environment variable named by the key. It returns an error, if any.
func Set(key string, value interface{}) error {
	var sv string

	switch T := value.(type) {
	case bool:
		sv = strconv.FormatBool(T)

	case float32, float64:
		sv = strconv.FormatFloat(reflect.ValueOf(value).Float(), 'g', -1, 64)

	case int, int8, int16, int32, int64:
		sv = strconv.FormatInt(reflect.ValueOf(value).Int(), 10)

	case uint, uint8, uint16, uint32, uint64:
		sv = strconv.FormatUint(reflect.ValueOf(value).Uint(), 10)

	case string:
		sv = value.(string)

	case []string:
		sv = strings.Join(value.([]string), ",")

	default:
		return fmt.Errorf("Unsupported type: %v", T)
	}
	return os.Setenv(key, sv)
}

// String retrieves the string value from environment named by the key.
func String(key string, fallback ...string) (value string) {
	if str, exists := Get(key); exists {
		value = str
	} else if len(fallback) > 0 {
		value = fallback[0]
	}
	return
}

// Strings retrieves the string values separated by comma from the environment.
func Strings(key string, fallback ...[]string) (value []string) {
	if v, exists := Get(key); exists {
		for _, item := range strings.Split(strings.TrimSpace(v), ",") {
			value = append(value, strings.TrimSpace(item))
		}
	} else if len(fallback) > 0 {
		value = fallback[0]
	}
	return
}

// Int retrieves the integer values separated by comma from the environment.
func Int(key string, fallback ...int) (value int) {
	if str, exists := Get(key); exists {
		if v, e := strconv.ParseInt(str, 10, 0); e == nil {
			value = int(v)
		}
	} else if len(fallback) > 0 {
		value = fallback[0]
	}
	return
}

// Int64 retrieves the 64-bit integer values separated by comma from the environment.
func Int64(key string, fallback ...int64) (value int64) {
	if str, exists := Get(key); exists {
		if v, e := strconv.ParseInt(str, 10, 64); e == nil {
			value = v
		}
	} else if len(fallback) > 0 {
		value = fallback[0]
	}
	return
}

// Uint retrieves the unsigned integer values separated by comma from the environment.
func Uint(key string, fallback ...uint) (value uint) {
	if str, exists := Get(key); exists {
		if v, e := strconv.ParseUint(str, 10, 0); e == nil {
			value = uint(v)
		}
	} else if len(fallback) > 0 {
		value = fallback[0]
	}
	return
}

// Uint64 retrieves the 64-bit unsigned integer values separated by comma from the environment.
func Uint64(key string, fallback ...uint64) (value uint64) {
	if str, exists := Get(key); exists {
		if v, e := strconv.ParseUint(str, 10, 64); e == nil {
			value = v
		}
	} else if len(fallback) > 0 {
		value = fallback[0]
	}
	return
}

// Bool retrieves boolean value from the environment.
func Bool(key string, fallback ...bool) (value bool) {
	if str, exists := Get(key); exists {
		if v, e := strconv.ParseBool(str); e == nil {
			value = v
		}
	} else if len(fallback) > 0 {
		value = fallback[0]
	}
	return
}

// Float retrieves float (64) value from the environment.
func Float(key string, fallback ...float64) (value float64) {
	if str, exists := Get(key); exists {
		if v, e := strconv.ParseFloat(str, 64); e == nil {
			value = v
		}
	} else if len(fallback) > 0 {
		value = fallback[0]
	}
	return
}
