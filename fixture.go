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
	"strings"
)

var wordsFailed = false
var ArraySize = 5
var MaxDepth = 5
var EnableDebug = false

func SetPopulateOptions(arraySize, maxDepth int, enableDebug bool) {
	ArraySize = arraySize
	MaxDepth = maxDepth
	EnableDebug = enableDebug
}

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

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func randomString(n int) string {
	if !wordsFailed {
		if word, err := randomWord(); err == nil {
			return word
		}
		wordsFailed = true
	}
	// if we cannot read /usr/share/dict/words, generate random string
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func prefix(n int) string {
	buf := strings.Builder{}
	for i := 0; i < n; i++ {
		buf.Write([]byte("  "))
	}
	return buf.String()
}

func debug(depth int, format string, a ...interface{}) {
	if !EnableDebug {
		return
	}
	line := fmt.Sprintf(format, a...)
	fmt.Printf("%s%s", prefix(depth), line)
}

func populate(v interface{}, depth int) {
	if depth == MaxDepth {
		debug(depth, "depth limit\n")
		return
	}

	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return
	}

	val = val.Elem()
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		num := int64(randomInt(-100, 100))
		val.SetInt(num)
		debug(depth, "(int) %v\n", num)

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		num := uint64(randomInt(0, 100))
		val.SetUint(num)
		debug(depth, "(uint) %v\n", num)

	case reflect.String:
		str := randomString(10)
		val.SetString(str)
		debug(depth, "(string) %v\n", str)

	case reflect.Slice:
		if val.IsNil() || val.Len() < ArraySize {
			val.Set(reflect.MakeSlice(val.Type(), ArraySize, ArraySize))
		}
		fallthrough

	case reflect.Array:
		for i := 0; i < val.Len(); i++ {
			debug(depth, "[%d]: %s\n", i, val.Index(i).Type())
			populate(val.Index(i).Addr().Interface(), depth+1)
		}

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldType := val.Type().Field(i)
			debug(depth, "%s: %s\n", fieldType.Name, field.Type())
			if field.CanSet() {
				populate(field.Addr().Interface(), depth+1)
			}
		}

	case reflect.Ptr:
		if val.IsNil() {
			val.Set(reflect.New(val.Type().Elem()))
		}
		populate(val.Interface(), depth)
	}
}

// Populate recursively populates a struct with random values.
func Populate(v interface{}) {
	populate(v, 0)
}

func Fixture[T any](fileName string, rv T) (T, error) {
	dir := path.Dir(fileName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return rv, fmt.Errorf("failed to create %s: %w", dir, err)
		}
	}
	// if file does not exist, create it with and empty return object.
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		Populate(rv)
		fp, err := os.Create(fileName)
		if err != nil {
			return rv, fmt.Errorf("failed to create %s: %w", fileName, err)
		}
		defer fp.Close()
		enc := json.NewEncoder(fp)
		enc.SetIndent("", "  ")
		if err := enc.Encode(rv); err != nil {
			return rv, fmt.Errorf("failed to encode %s: %w", fileName, err)
		}
		return rv, nil
	}
	// if file fileName exists, read it and return its contents.
	fp, err := os.Open(fileName)
	if err != nil {
		return rv, fmt.Errorf("failed to open %s: %w", fileName, err)
	}
	err = json.NewDecoder(fp).Decode(rv)
	if err != nil {
		return rv, fmt.Errorf("failed to decode %s: %w", fileName, err)
	}
	return rv, nil
}
