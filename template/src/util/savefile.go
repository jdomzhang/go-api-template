package util

import (
	"bytes"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var (
	// errBucket       = errors.New("Invalid bucket!")
	// errSize         = errors.New("Invalid size!")
	errInvalidImage = errors.New("invalid image")
)

// SaveImageToDisk will save image to disk
func SaveImageToDisk(fileNameBase, data string) (string, error) {
	idx1 := strings.Index(data, ";base64,")
	idx2 := strings.Index(data, "data:image/")
	if idx1 < 0 || idx2 != 0 {
		return "", errInvalidImage
	}

	ext := data[11:idx1]

	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx1+8:]))
	buff := bytes.Buffer{}
	if _, err := buff.ReadFrom(reader); err != nil {
		logrus.Error("buff.ReadFrom:", err.Error())
		return "", err
	}

	// imgCfg, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	// if err != nil {
	// 	logrus.Error("image.DecodeConfig:", err.Error())
	// 	return "", err
	// }

	// fmt.Printf("file size: %d", buff.Len())

	// if imgCfg.Width > 400 || imgCfg.Height > 400 {
	// 	// TODO: resize the image
	// 	// return "", errSize
	// }

	fileName := fileNameBase + "." + ext

	safeFileName := fileName
	if safeFileName[0] == '/' {
		safeFileName = safeFileName[1:]
	}

	if err := makeSureFolders(safeFileName); err != nil {
		return "", err
	}

	if err := ioutil.WriteFile(safeFileName, buff.Bytes(), 0644); err != nil {
		return "", err
	}

	return fileName, nil
}

func makeSureFolders(fullPathWithFileName string) error {
	folderPath := fullPathWithFileName[:strings.LastIndex(fullPathWithFileName, "/")]
	if len(folderPath) > 0 {
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			return err
		}
	}

	return nil
}
