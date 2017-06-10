package main

import (
	"fmt"
	"runtime"
	"os/exec"
	"log"
	"unsafe"
)

func main() {
	// 操作系统
	fmt.Println("computer contral system:", runtime.GOOS)
	// 架构
	fmt.Println("framework:", runtime.GOARCH)
	// GOROOT
	fmt.Println("goroot:", runtime.GOROOT())
	// go版本
	fmt.Println("go version:", runtime.Version())
	// cpu数
	fmt.Println("CPU Number:", runtime.NumCPU())
	fmt.Println()
	switch runtime.GOOS {
	case "linux":
		out, err := exec.Command("cat", "proc/version").Output()
		if err != nil {
			log.Fatal(err)
			fmt.Printf("The version is %s\n", out)
		}
	case "windows":
		fmt.Println("Computer System Version:", GetOsVersion())
	}
}
type OSVERSIONINFO struct {
	dwOSVersionInfoSize int32
	dwMajorVersion int32
	dwMinorVersion int32
	dwBuildNumber int32
	dwPlatformId int32
	}

func GetOsVersion() string {
	var version string = "Unknown Version"
	var os OSVERSIONINFO;
	os.dwOSVersionInfoSize = int32(unsafe.Sizeof(os))
	switch {
		case os.dwMajorVersion==4 && os.dwMinorVersion==0 && os.dwPlatformId==1:
			version = "Windows 95"
			break
		case os.dwMajorVersion==4 && os.dwMinorVersion==10:
			version = "Windows 98"
			break
		case os.dwMajorVersion==4 && os.dwMinorVersion==90:
			version = "Windows Me"
			break
		case os.dwMajorVersion==4 && os.dwMinorVersion==0 && os.dwPlatformId==2:
			version = "Windows NT4"
			break
		case os.dwMajorVersion==5 && os.dwMinorVersion==0:
			version = "Windows 2000"
			break
		case os.dwMajorVersion==5 && os.dwMinorVersion==1:
			version = "Windows XP"
			break
		case os.dwMajorVersion==5 && os.dwMinorVersion==2:
			version = "Windows 2003"
			break
		case os.dwMajorVersion==6 && os.dwMinorVersion==0:
			version = "Windows Vista"
			break
		case os.dwMajorVersion==6 && os.dwMinorVersion==1:
			version = "Windows 7"
			break
		case os.dwMajorVersion==6 && os.dwMinorVersion==3:
			version = "Windows 8.1"
			break
		case os.dwMajorVersion==6 && os.dwMinorVersion==2:
			version = "Windows 8"
			break
		default:
			version = "Windows 10"
			break


	}
	return version
}