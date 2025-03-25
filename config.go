package config

import (
	"encoding/json"
	"errors"
	"github.com/goccy/go-yaml"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

// WithFile first ingests configuration values from a file into a user-defined
// structure, then applies environment variables to override them. If no file
// is provided, environment variables can still be used.
func WithFile(filename string, conf any) error {
	var err error = nil

	fn := strings.TrimSpace(filename)
	if fn != "" {
		var ext = filepath.Ext(fn)
		var data []byte
		data, err = os.ReadFile(fn)
		if err == nil {
			// Read configuration values from file.
			switch ext {
			case ".yaml", ".yml":
				err = yaml.Unmarshal(data, conf)
			case ".json":
				err = json.Unmarshal(data, conf)
			}
		}
	}

	if err != nil {
		return err
	}

	return withEnvironmentOverrides(reflect.ValueOf(conf).Elem())
}

func withEnvironmentOverrides(val reflect.Value) error {
	for i := 0; i < val.NumField(); i++ {
		var err error = nil
		field := val.Field(i)
		fieldType := val.Type().Field(i)
		if fieldType.Type.Kind() == reflect.Struct {
			// Process nested structs
			err = withEnvironmentOverrides(field)
		} else {
			// Pull values from the environment, if available
			tag := fieldType.Tag.Get("env")
			if tag != "" {
				env := os.Getenv(tag)
				if env != "" {
					switch field.Kind() {
					case reflect.String:
						field.SetString(env)
					case reflect.Bool:
						field.SetBool(env == "true")
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						var i int64
						i, err = strconv.ParseInt(env, 10, 64)
						if err == nil {
							field.SetInt(i)
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
						var i uint64
						i, err = strconv.ParseUint(env, 10, 64)
						if err == nil {
							field.SetUint(i)
						}
					case reflect.Float32, reflect.Float64:
						var f float64
						f, err = strconv.ParseFloat(env, 64)
						if err == nil {
							field.SetFloat(f)
						}
					default:
						err = errors.New("unexpected field type")
					}
				}
			}
		}

		if err != nil {
			return err
		}
	}

	return nil
}
