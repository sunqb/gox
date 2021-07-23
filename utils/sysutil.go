package utils

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

// KillEXE  根据进程名字kill进程：
// kill调进程  参数---taskkill /im notepad.exe /T /F
// 参数说明：strGameName为需要kill的进程的名字
func KillEXE(strGameName string) bool {
	log.Println("kill exe:", strGameName)
	strGameName = strGameName + ".exe"
	cmd := exec.Command("PowerShell.exe", "taskkill /im ", strGameName, " -t -f")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.CombinedOutput()
	if nil != err {
		log.Println(fmt.Sprintf("KillEXE error:%v,%v", string(output), err))
		return false
	}
	return true
}

// IsProcessExist 检查进程是否存在
func IsProcessExist(appName string) bool {
	strAppName := appName + ".exe"
	cmd := exec.Command("PowerShell.exe", "tasklist | findstr", strAppName)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := cmd.Output()

	if nil != err {
		log.Println(fmt.Sprintf("check process exist error:%v,%v", string(output), err))
		return false
	}

	fields := strings.Fields(string(output))

	for _, v := range fields {
		if v == strAppName {
			return true
		}
	}
	return false
}

// RunEXE  参数---reg add HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /v AUTORUN /t REG_SZ /d C:\autorun.exe /f
//假如你要运行的程序名字为:"autorun.exe"使用命令为
//"reg add HKLM\SOFTWARE\Microsoft\Windows\CurrentVersion\Run /v AUTORUN /t REG_SZ /d C:\autorun.exe /f"
//(不包括引号)其中"C:\autorun.exe"为目标程序的路径.按着这样的命令就可以将你的程序添加到启动项中了
func RunEXE(strEXEName string) {
	arg := []string{"reg", " add", " HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "auto", "/t", "REG_SZ", "/d", strEXEName, "/f"}
	cmd := exec.Command("PowerShell.exe", arg...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	d, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(ConvertToByte(string(d), "gbk", "utf8")))
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(d))
	return
}

// WriteREG 写注册表
func WriteREG(strEXEName string) {
	strEXEName = strings.ReplaceAll(strEXEName, `/`, `\`)
	k, err := registry.OpenKey(registry.CURRENT_USER, `Software\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		log.Fatal(err)
	}
	if err := k.SetStringValue("auto", strEXEName); err != nil {
		log.Fatal(err)
	}
	if err := k.Close(); err != nil {
		log.Fatal(err)
	}
}
