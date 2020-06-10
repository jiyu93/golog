package golog

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// 日志文件的权限: rw-r--r--
const permission = 0644

// Rotater 日志轮转器
type Rotater struct {
	FileName string
	MaxSize  int64
	Backups  int
	Gzip     bool
	gzSuffix string
	wcnt     int64 // write count
	f        *os.File
	fLock    sync.Mutex // locker of file, for writing&rotating
	cLock    sync.Mutex // locker for compressing
}

// NewRotater 创建一个新的日志轮转器
// fileName: 文件名，可以是相对路径或绝对路径，父层文件夹需要自行创建
// maxSizeMB: 单个日志文件最大大小
// backups: 备份的文件数量
// gzip: 备份文件是否需要gzip压缩
func NewRotater(fileName string, maxSizeMB int64, backups int, gzip bool) *Rotater {
	// 尝试打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, permission)
	if err != nil {
		panic(err.Error())
	}
	// 获取首个文件的大小，用以初始化wcnt
	fi, err := f.Stat()
	if err != nil {
		panic(err.Error())
	}
	// suffix of gzip
	gz := ""
	if gzip {
		gz = ".gz"
	}
	return &Rotater{
		FileName: fileName,
		MaxSize:  maxSizeMB * 1048576,
		Backups:  backups,
		Gzip:     gzip,
		gzSuffix: gz,
		f:        f,
		wcnt:     fi.Size(),
	}
}

// Write 切分并且写日志
func (r *Rotater) Write(p []byte) (int, error) {
	r.fLock.Lock()
	defer r.fLock.Unlock()
	r.rotate()
	n, err := r.f.Write(p)
	r.wcnt += int64(n)
	return n, err
}

// rotate 轮转文件
func (r *Rotater) rotate() {
	if r.wcnt >= r.MaxSize {
		// 关闭现有文件
		r.f.Close()
		// 删除编号最大的文件, 无论文件是否存在
		os.Remove(fmt.Sprintf("%s.%d%s", r.FileName, r.Backups, r.gzSuffix))
		// 重命名剩下的编号文件，无论文件是否存在
		for i := r.Backups - 1; i >= 1; i-- {
			os.Rename(
				fmt.Sprintf("%s.%d%s", r.FileName, i, r.gzSuffix),
				fmt.Sprintf("%s.%d%s", r.FileName, i+1, r.gzSuffix),
			)
		}
		// 将刚才关闭的文件编号为1，如果有需要则进行压缩
		if r.Gzip {
			// 为防止gzip完成之前与下一个.1文件发生冲突，先将本文件rename为临时文件
			tmp := fmt.Sprintf("%s.1.tmp-%d", r.FileName, time.Now().UnixNano())
			os.Rename(r.FileName, tmp)
			// gzip不占用rotate的处理时间，以免影响日志输出性能
			go r.gzip(
				tmp,
				fmt.Sprintf("%s.1.gz", r.FileName),
			)
		} else {
			os.Rename(r.FileName, fmt.Sprintf("%s.1", r.FileName))
		}
		// 打开新的文件
		nf, err := os.OpenFile(r.FileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, permission)
		if err != nil {
			panic(err)
		}
		r.f = nf
		// 清空计数器
		r.wcnt = 0
	}
}

// gzip gzip压缩
func (r *Rotater) gzip(src, dst string) {
	r.cLock.Lock()
	defer r.cLock.Unlock()
	fsrc, err := os.Open(src)
	if err != nil {
		panic(err)
	}
	fdst, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, permission)
	if err != nil {
		panic(err)
	}
	gw := gzip.NewWriter(fdst)
	_, err = io.Copy(gw, fsrc)
	if err != nil {
		panic(err)
	}
	gw.Close()
	fsrc.Close()
	os.Remove(src)
}
