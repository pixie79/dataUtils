package utils

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func die(err error, msg string) {
	Logger.Error(fmt.Sprintf("%+v: %+v", msg, err))
	os.Exit(1)
}

func MaybeDie(err error, msg string) {
	if err != nil {
		die(err, msg)
	}
}

// GetEnv Simple helper function to read an environment or return a default value
func GetEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetEnvOrDie(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		die(Err, fmt.Sprintf("missing environment variable %s", key))
	}
	return value
}

func LinesFromReader(r io.Reader) []string {
	var lines []string

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err := scanner.Err()
	MaybeDie(err, "could not parse lines")

	return lines
}

func UrlToLines(url string, username string, password string) []string {
	client := http.Client{Timeout: 5 * time.Second}

	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	MaybeDie(err, "could not create http request")

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	res, err := client.Do(req)
	MaybeDie(err, "could not authenticate")

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		MaybeDie(err, "error closing connection")
	}(res.Body)

	if !InBetween(res.StatusCode, 200, 299) {
		die(fmt.Errorf("%d", res.StatusCode), fmt.Sprintf("url access error %s", url))
	}

	return LinesFromReader(res.Body)
}

func InBetween(i, min, max int) bool {
	if (i >= min) && (i <= max) {
		return true
	} else {
		return false
	}
}

func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func B64DecodeMsg(b64Key string, offsetF ...int) ([]byte, error) {
	offset := 7
	if len(offsetF) > 0 {
		offset = offsetF[0]
	}
	//logger.Debug(fmt.Sprintf("Base64 Encoded String: %s", b64Key))
	var key []byte
	var err error
	if len(b64Key)%4 != 0 {
		key, err = base64.RawStdEncoding.DecodeString(b64Key)
	} else {
		key, err = base64.StdEncoding.DecodeString(b64Key)
	}
	if err != nil {
		return []byte{}, err
	}
	result := key[offset:]
	//logger.Debug(fmt.Sprintf("Base64 Decoded String: %s", result))
	return result, nil
}
