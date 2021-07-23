package utils

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// GetCurrentAbPath 最终方案-全兼容
func GetCurrentAbPath() string {
	dir := GetCurrentAbPathByExecutable()
	tmpDir, _ := filepath.EvalSymlinks(os.TempDir())
	if strings.Contains(dir, tmpDir) {
		return GetCurrentAbPathByCaller()
	}
	return dir
}

// GetCurrentAbPathByExecutable 获取当前执行文件绝对路径
func GetCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}

// GetCurrentAbPathByCaller 获取当前执行文件绝对路径（go run）
func GetCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// GetParentDirectory 获取上级方法
func GetParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

// GetCurrentDirectory 获取项目当前根路径。不能用go run运行。go run可能会获取到临时文件路径
func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

// GetAppDataPath 获取appdata路径
func GetAppDataPath(orgName string, appName string) string {
	// appPath := fmt.Sprintf("C:/Users/%s/AppData/Roaming/%s/%s", GetUserName(), GetInstance().OrgName, GetInstance().AppName)

	appPath := fmt.Sprintf("%s/%s/%s", os.Getenv("APPDATA"), orgName, appName)

	os.MkdirAll(appPath, os.ModeDir)

	return appPath
}

// GetExeFileName 获取当前正在执行的文件的名称
func GetExeFileName() string {
	filePath := os.Args[0]
	filePath = strings.Replace(filePath, "\\", "/", -1)
	return path.Base(filePath)
}

// 子路径
func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}
