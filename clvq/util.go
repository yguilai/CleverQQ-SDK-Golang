package clvq

import "C"
import (
	sc "golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"path/filepath"
	"unsafe"
)

func cInt(i int) C.int {
	return C.int(i)
}

func goInt(ci C.int) int {
	return int(ci)
}

func cBool(b bool) C.int {
	if b {
		return 1
	}
	return 0
}

func cString(str string) *C.char {
	gbstr, _ := sc.GB18030.NewEncoder().String(str)
	return C.CString(gbstr)
}

func goString(str *C.char) string {
	utf8str, _ := sc.GB18030.NewDecoder().String(C.GoString(str))
	return utf8str
}

func str2ptr(s string) uintptr {
	return uintptr(unsafe.Pointer(cString(s)))
}

func int2ptr(i int) uintptr {
	return uintptr(i)
}

func byte2ptr(b []byte) uintptr {
	//return uintptr(*((*int64)(unsafe.Pointer(&b))))
	return uintptr(unsafe.Pointer(&b))
}

func bool2ptr(b bool) uintptr {
	if b == true {
		return uintptr(1)
	}
	return uintptr(0)
}

func ptr2str(ptr uintptr) string {
	return goString((*C.char)(unsafe.Pointer(ptr)))
}

func ptr2bool(ptr uintptr) bool {
	i := int(ptr)
	if i == 1 {
		return true
	}
	return false
}

func ptr2int(ptr uintptr) int {
	return int(ptr)
}

func _ptr2str(ptr uintptr) string {
	var b []byte
	for {
		sbyte := *((*byte)(unsafe.Pointer(ptr)))
		if sbyte == 0 {
			break
		}
		b = append(b, sbyte)
		ptr += 1
	}
	return string(b)
}

// GetCurrPath 获取当前程序路径
// 假设框架程序CleverQQ Air.exe运行在D:\CleverQQ, 则返回D:\CleverQQ
func GetCurrPath() (string, error) {
	path, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}
	return path, nil
}

/**
 * 读取文件, 返回[]byte
 */
func file2Bytes(filename string) ([]byte, error) {
	// File
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// FileInfo:
	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// []byte
	data := make([]byte, stats.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}