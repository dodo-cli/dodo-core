package decoder

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"

	"github.com/oclaussen/dodo/pkg/types"
	"gopkg.in/yaml.v2"
)

type Group struct {
	Backdrops map[string]types.Backdrop
	Groups    map[string]Group
}

func (groups *Group) Names() []string {
	var result []string
	if groups.Backdrops != nil {
		for name := range groups.Backdrops {
			result = append(result, name)
		}
	}
	if groups.Groups != nil {
		for _, group := range groups.Groups {
			result = append(result, group.Names()...)
		}
	}
	return result
}

func (groups *Group) Strings() []string {
	var result []string
	if groups.Backdrops != nil {
		for name, _ := range groups.Backdrops {
			result = append(result, fmt.Sprintf("backdrop/%s (%s)", name, "unknown")) // TODO backdrop.filename))
		}
	}
	if groups.Groups != nil {
		for name, group := range groups.Groups {
			for _, substring := range group.Strings() {
				result = append(result, fmt.Sprintf("%s/%s", name, substring))
			}
		}
	}
	return result
}

func (target *Group) Merge(source *Group) {
	if source.Backdrops != nil {
		if target.Backdrops == nil {
			target.Backdrops = map[string]types.Backdrop{}
		}
		for name, backdrop := range source.Backdrops {
			if _, ok := target.Backdrops[name]; !ok {
				target.Backdrops[name] = backdrop
			}
		}
	}

	if source.Groups != nil {
		if target.Groups == nil {
			target.Groups = map[string]Group{}
		}
		for name, sgroup := range source.Groups {
			if tgroup, ok := target.Groups[name]; ok {
				tgroup.Merge(&sgroup)
			} else {
				target.Groups[name] = sgroup
			}
		}
	}
}

func (d *decoder) DecodeGroups(name string, config interface{}) (map[string]Group, error) {
	result := map[string]Group{}
	switch t := reflect.ValueOf(config); t.Kind() {
	case reflect.Map:
		for k, v := range t.Interface().(map[interface{}]interface{}) {
			key := k.(string)
			decoded, err := d.DecodeGroup(key, v)
			if err != nil {
				return result, err
			}
			result[key] = decoded
		}
	default:
		return result, &ConfigError{Name: name, UnsupportedType: t.Kind()}
	}
	return result, nil
}

func (d *decoder) DecodeGroup(name string, config interface{}) (Group, error) {
	result := Group{Backdrops: map[string]types.Backdrop{}, Groups: map[string]Group{}}
	switch t := reflect.ValueOf(config); t.Kind() {
	case reflect.Map:
		for k, v := range t.Interface().(map[interface{}]interface{}) {
			switch key := k.(string); key {
			case "groups":
				decoded, err := d.DecodeGroups(key, v)
				if err != nil {
					return result, err
				}
				for name, group := range decoded {
					result.Groups[name] = group
				}
			case "backdrops":
				decoded, err := d.DecodeBackdrops(key, v)
				if err != nil {
					return result, err
				}
				for name, backdrop := range decoded {
					result.Backdrops[name] = backdrop
				}
			case "include":
				decoded, err := d.DecodeIncludes(key, v)
				if err != nil {
					return result, err
				}
				for _, include := range decoded {
					result.Merge(&include)
				}
			default:
				return result, &ConfigError{Name: name, UnsupportedKey: &key}
			}
		}
	default:
		return result, &ConfigError{Name: name, UnsupportedType: t.Kind()}
	}
	return result, nil
}

func (d *decoder) DecodeIncludes(name string, config interface{}) ([]Group, error) {
	result := []Group{}
	switch t := reflect.ValueOf(config); t.Kind() {
	case reflect.Map:
		decoded, err := d.DecodeInclude(name, config)
		if err != nil {
			return result, err
		}
		result = append(result, decoded)
	case reflect.Slice:
		for _, v := range t.Interface().([]interface{}) {
			decoded, err := d.DecodeInclude(name, v)
			if err != nil {
				return result, err
			}
			result = append(result, decoded)
		}
	default:
		return result, &ConfigError{Name: name, UnsupportedType: t.Kind()}
	}
	return result, nil
}

func (d *decoder) DecodeInclude(name string, config interface{}) (Group, error) {
	var result Group
	switch t := reflect.ValueOf(config); t.Kind() {
	case reflect.Map:
		for k, v := range t.Interface().(map[interface{}]interface{}) {
			switch key := k.(string); key {
			case "file":
				decoded, err := d.DecodeString(key, v)
				if err != nil {
					return result, err
				}
				return d.includeFile(decoded)
			case "text":
				decoded, err := d.DecodeString(key, v)
				if err != nil {
					return result, err
				}
				return d.includeText(name, []byte(decoded))
			default:
				return result, &ConfigError{Name: name, UnsupportedKey: &key}
			}
		}
	default:
		return result, &ConfigError{Name: name, UnsupportedType: t.Kind()}
	}
	return result, nil
}

func (d *decoder) includeFile(filename string) (Group, error) {
	if !filepath.IsAbs(filename) {
		directory, err := os.Getwd()
		if err != nil {
			return Group{}, err
		}
		filename, err = filepath.Abs(filepath.Join(directory, filename))
		if err != nil {
			return Group{}, err
		}
	}
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Group{}, fmt.Errorf("could not read file '%s'", filename)
	}
	return d.WithFile(filename).includeText(filename, bytes)
}

func (d *decoder) includeText(name string, bytes []byte) (Group, error) {
	var mapType map[interface{}]interface{}
	err := yaml.Unmarshal(bytes, &mapType)
	if err != nil {
		return Group{}, err
	}
	config, err := d.DecodeGroup(name, mapType)
	if err != nil {
		return Group{}, err
	}
	return config, nil
}
