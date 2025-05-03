package main

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/creack/pty"
)

func RunFZF(input string) (string, error) {
	// 创建伪终端
	ptm, pts, err := pty.Open()
	if err != nil {
		return "", err
	}
	defer ptm.Close()
	defer pts.Close()

	// 创建结果缓冲区
	// 配置fzf命令
	var buf bytes.Buffer
	cmd := exec.Command("fzf", "--ansi")
	cmd.Stdout = io.MultiWriter(pts, &buf) // 实时显示并捕获

	cmd.Stderr = os.Stderr
	cmd.Stdin = io.MultiReader(strings.NewReader(input)) // 允许接收键盘输入

	// // 单 goroutine 读取输出
	// go func() {
	// 	io.Copy(io.MultiWriter(os.Stdout, &buf), ptm)
	// }()

	// 执行命令并等待完成
	if err := cmd.Run(); err != nil {
		return "", err
	}

	// 返回清理后的结果
	return strings.TrimSpace(buf.String()), nil
}

func main() {
	// 示例数据
	data := strings.Join([]string{
		"Red",
		"Green",
		"Blue",
	}, "\n")

	// 运行并捕获
	result, err := RunFZF(data)
	if err != nil {
		if err == io.EOF {
			println("用户取消选择")
			return
		}
		panic(err)
	}

	println("最终选择:", result)
}
