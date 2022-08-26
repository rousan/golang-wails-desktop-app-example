package main

import (
	"changeme/utils"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strconv"
)

type downloadProgress struct {
	contentLength *uint64
	progress      uint64
}

func downloadURL(ctx context.Context, urlStr string, cb func(dp downloadProgress)) (savedFilePath string, err error) {
	url, err := url.Parse(urlStr)
	if err != nil {
		return "", fmt.Errorf("can't parse url string: %v", err)
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return "", fmt.Errorf("can't create http request: %v", err)
	}
	defer resp.Body.Close()

	contentLen := new(uint64)
	rawContentLen, err := strconv.ParseUint(resp.Header.Get("Content-Length"), 10, 64)
	if err == nil {
		*contentLen = rawContentLen
	} else {
		contentLen = nil
	}

	downloadsPath, err := utils.UserDownloadsDir()
	if err != nil {
		return "", fmt.Errorf("can't get user downloads dir: %v", err)
	}

	fileName := filepath.Base(url.Path)
	if fileName == "" || fileName == "." || fileName == fmt.Sprint(os.PathSeparator) {
		fileName = utils.GenUUID()
	} else {
		fileName = fmt.Sprintf("%v_%v", utils.GenUUID(), fileName)
	}

	savedFilePath = path.Join(downloadsPath, fileName)
	outFile, err := os.OpenFile(savedFilePath, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		return "", fmt.Errorf("can't open a file for writing")
	}

	var downloaded uint64
	buf := make([]byte, 1024*16)

	for {
		nr, er := resp.Body.Read(buf)

		if nr > 0 {
			nw, ew := outFile.Write(buf[0:nr])
			if nw < 0 || nr < nw {
				nw = 0
				if ew == nil {
					ew = fmt.Errorf("invalid write operation")
				}
			}

			downloaded += uint64(nw)
			cb(downloadProgress{
				contentLength: contentLen,
				progress:      downloaded,
			})

			if ew != nil {
				err = ew
				break
			}

			if nr != nw {
				err = io.ErrShortWrite
				break
			}
		}

		if er != nil {
			if er != io.EOF {
				err = er
			}
			break
		}
	}

	if err != nil {
		return "", fmt.Errorf("can't download: %v", err)
	}

	return savedFilePath, nil
}
