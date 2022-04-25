package lib

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"io"
	"os"
	"path"
	"strings"

	_ "github.com/MeshBoxTech/mesh-chain/internal/smc-devnet/statics/statik"
	"github.com/rakyll/statik/fs"
)

var (
	/*
		UrlMap = map[string]map[string]string{
			"amd64": {
				"darwin": "/devnet.tar.gz",
			},
		}*/

	dataUrl = "/devnet.tar.gz"

	ErrorGoarchNotSupport = errors.New("goarch not support")
	ErrorGoosNotSupport   = errors.New("goos not support")
	dataDir               string
	withoutNodekey        bool
)

func Restore(dir string, reset, masterNode bool) (string, error) {
	if reset {
		os.RemoveAll(dir)
	}
	os.MkdirAll(dir, 0755)
	dataDir = dir
	withoutNodekey = !masterNode
	return restore()
}

func restore() (string, error) {
	var binUrl = dataUrl

	/*	if m, ok := UrlMap[runtime.GOARCH]; !ok {
			return "", ErrorGoarchNotSupport
		} else if url, ok := m[runtime.GOOS]; !ok {
			return "", ErrorGoosNotSupport
		} else {
			binUrl = url[:]
		}*/

	if strings.LastIndex(binUrl, "tar.gz") != -1 {
		err := opentargz(dataDir, binUrl)
		if err != nil {
			panic(err)
		}
	}
	return binUrl, nil
}

func opentargz(binDir, binUrl string) error {
	os.MkdirAll(binDir, 0755)
	statikFS, err := fs.New()
	srcFile, err := statikFS.Open(binUrl)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	gr, err := gzip.NewReader(srcFile)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	for {
		hdr, err := tr.Next()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		if hdr.FileInfo().IsDir() && strings.Index(hdr.Name, "devnet") == 0 ||
			withoutNodekey && strings.LastIndex(hdr.Name, "nodekey") > 0 {
			continue
		}
		n := strings.Replace(hdr.Name, "devnet/", "", 1)
		fpath := path.Join(binDir, n)
		dir, _ := path.Split(fpath)
		os.MkdirAll(dir, 0755)
		if !hdr.FileInfo().IsDir() {
			file, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				return err
			}
			io.Copy(file, tr)
		}
	}
	return nil
}
