package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		PrintUsage()
		os.Exit(1)
	}

	action := os.Args[1]
	filename := os.Args[2]

	switch action {
	case "encode":
		encode(filename)
	case "decode":
		decode(filename)
	default:
		PrintUsage()
		os.Exit(1)
	}
}

func PrintUsage() {
	os.Stderr.WriteString("base64img\n")
	os.Stderr.WriteString("Encodes/Decodes images to/from base64 format\n\n")
	os.Stderr.WriteString("Usage:\n")
	os.Stderr.WriteString("base64img action filename\n\n")
	os.Stderr.WriteString("action can be either encode or decode\n")
	os.Stderr.WriteString("filename, in case of encode, is a jpg/png image file\n")
	os.Stderr.WriteString("filename, in case of decode, is a text file containing \"data:image/png;base64,...\"\n\n")
	os.Stderr.WriteString("Output is written to stdout\n")
}

func encode(filename string) {
	data, err := getFileContents(filename)
	DieOnError(err)

	output := base64.StdEncoding.EncodeToString(data)
	mime := http.DetectContentType(data)

	fmt.Printf("data:%s;base64,%s", mime, output)
}

func decode(filename string) {
	rawdata, err := getFileContents(filename)
	DieOnError(err)

	encdata := string(rawdata)
	encdata = strings.Replace(encdata, "\n", "", -1)

	data, err := stripMime(encdata)
	DieOnError(err)

	output, err := base64.StdEncoding.DecodeString(data)
	DieOnError(err)

	os.Stdout.Write(output)
}

func getFileContents(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("Error opening file")
	}

	info, err := f.Stat()
	if err != nil {
		return nil, errors.New("Error getting file stats")
	}

	len := info.Size()
	data := make([]byte, len)
	n, err := f.Read(data)
	if err != nil {
		return nil, errors.New("Error reading file")
	}
	if int64(n) != len {
		return nil, errors.New("Could not read entire contents of file")
	}

	return data, nil
}

func stripMime(combined string) (string, error) {
	re := regexp.MustCompile("data:(.*);base64,(.*)")
	parts := re.FindStringSubmatch(combined)

	if len(parts) < 3 {
		return "", errors.New("Invalid base64 input")
	}

	data := parts[2]
	return data, nil
}

func DieOnError(err error) {
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
}
