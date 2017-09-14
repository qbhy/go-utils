package utils

import (
	"net/url"
	"log"
	"strings"
	"net/http"
	"io/ioutil"
	"os"
	"path/filepath"
)



/*
获取程序运行路径
*/
func CurrentPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return strings.Replace(dir, "\\", "/", -1) + "/"
}

/*
检测文件或者文件夹是否已经存在
 */
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

/**
下载文件
 */
func DownloadFile(file_url string, path string) {
	u, err := url.Parse(file_url)
	if err != nil {
		log.Println("parse url failed:", file_url, err)
		return
	}

	tmp := strings.TrimLeft(u.Path, "/")

	filename := path + strings.ToLower(strings.Replace(tmp, "/", "-", -1))
	exists, _ := PathExists(filename)

	if exists {
		return
	}

	response, err := http.Get(file_url)
	if err != nil {
		log.Println("get file_url failed:", err)
		return
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("read data failed:", file_url, err)
		return
	}
	FilePutContents(filename, data)
}

/*
写文件
 */
func FilePutContents(filename string, contents []byte) {
	image, err := os.Create(filename)
	if err != nil {
		log.Println("create file failed:", filename, err)
		return
	}
	image.Write(contents)
	defer image.Close()
}
