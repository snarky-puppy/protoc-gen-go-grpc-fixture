package protoc_gen_go_grpc_fixture

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path"
	"reflect"
)

// open /usr/share/dict/words and pick a word at random.
// this function first seeks to a random location in the file, then backtracks to the last newline character
func randomWord() (string, error) {
	file, err := os.Open("/usr/share/dict/words")
	if err != nil {
		return "", err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return "", err
	}
	fileSize := stat.Size()

	randomPosition := rand.Int63n(fileSize)

	_, err = file.Seek(randomPosition, io.SeekStart)
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(file)

	// Discard the current partial line, unless we're at the beginning of the file
	if randomPosition > 0 {
		scanner.Scan()
	}

	if !scanner.Scan() {
		return "", fmt.Errorf("failed to read word after seeking")
	}

	return scanner.Text(), nil
}

func replace(v interface{}, visited map[uintptr]bool) error {
	value := reflect.ValueOf(v)
	valueType := reflect.TypeOf(v)

	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	ptr := value.UnsafeAddr()
	if visited[ptr] {
		return nil
	}
	visited[ptr] = true

	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		min := -10
		max := 10
		value.SetInt(int64(rand.Intn(max-min+1) + min))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		min := 0
		max := 20
		value.SetUint(uint64(rand.Intn(max-min+1) + min))
	case reflect.String:
		if value.CanSet() {
			if str, err := randomWord(); err == nil {
				value.SetString(str)
			} else {
				return fmt.Errorf("%s: %w", valueType.Name, err)
			}
		}
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			fieldType := value.Type().Field(i)
			if !fieldType.IsExported() {
				continue
			}
			iface := field.Addr().Interface()
			if err := replace(iface, visited); err != nil {
				return fmt.Errorf("%s: %w", fieldType.Name, err)
			}
			if fieldType.Anonymous {
				if err := replace(field.Addr().Interface(), visited); err != nil {
					return fmt.Errorf("%s: %w", fieldType.Name, err)
				}
			}
		}
	case reflect.Ptr:
		if !value.IsNil() {
			if err := replace(value.Interface(), visited); err != nil {
				return fmt.Errorf("%s: %w", valueType.Name, err)
			}
		}
	case reflect.Slice, reflect.Array:
		if value.Type().Elem().Kind() == reflect.Uint8 {
			for j := 0; j < value.Len(); j++ {
				// 'x' as a byte is 0x78
				value.Index(j).SetUint(0x78)
			}
		} else {
			for j := 0; j < value.Len(); j++ {
				elem := value.Index(j)
				if elem.Kind() == reflect.Struct || (elem.Kind() == reflect.Ptr && elem.Elem().Kind() == reflect.Struct) {
					if err := replace(elem.Addr().Interface(), visited); err != nil {
						return fmt.Errorf("%s: %w", valueType.Name(), err)
					}
				}
			}
		}
	}

	return nil
}

func Fixture[T any](baseDir string, name string, rv T) (T, error) {
	f := path.Join(baseDir, fmt.Sprintf("%s.json", name))
	dir := path.Dir(f)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return rv, fmt.Errorf("failed to create %s: %w", dir, err)
		}
	}
	// if file f does not exist, create it with and empty return object.
	_, err := os.Stat(f)
	if os.IsNotExist(err) {
		visited := make(map[uintptr]bool)
		if err := replace(rv, visited); err != nil {
			return rv, fmt.Errorf("failed to create fixture values %s: %w", f, err)
		}
		fp, err := os.Create(f)
		if err != nil {
			return rv, fmt.Errorf("failed to create %s: %w", f, err)
		}
		defer fp.Close()
		enc := json.NewEncoder(fp)
		enc.SetIndent("", "  ")
		if err := enc.Encode(rv); err != nil {
			return rv, fmt.Errorf("failed to encode %s: %w", f, err)
		}
		return rv, nil
	}
	// if file f exists, read it and return its contents.
	fp, err := os.Open(f)
	if err != nil {
		return rv, fmt.Errorf("failed to open %s: %w", f, err)
	}
	err = json.NewDecoder(fp).Decode(rv)
	if err != nil {
		return rv, fmt.Errorf("failed to decode %s: %w", f, err)
	}
	return rv, nil
}
