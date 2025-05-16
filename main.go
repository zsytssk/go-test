package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type Candidate struct {
	Text   string
	Weight float64
	Score  int
}

var items = []map[string]interface{}{
	{
		"ID":       1,
		"Priority": 1,
		"Name":     "测试 1",
	},
	{
		"ID":       2,
		"Priority": 2,
		"Name":     "测试 2",
	},
	{
		"ID":       3,
		"Priority": 3,
		"Name":     "测试 3",
	},
}

func main() {
	reader, writer := io.Pipe()

	fzf := exec.Command("fzf", "--ansi")
	fzf.Stdin = reader
	fzf.Stdout = os.Stdout
	fzf.Stderr = os.Stderr

	go func() {
		defer writer.Close()

		// // 写入历史记录
		for _, item := range items {
			fmt.Fprintf(writer, "%s [%d:%d]\n", item["Name"], item["ID"], item["Priority"])
		}

		// 实时执行 find 命令并写入
		findCmd := exec.Command("bash", "-c", "find /home/zsy")
		findOut, _ := findCmd.StdoutPipe()
		findCmd.Stderr = os.Stderr
		_ = findCmd.Start()

		io.Copy(writer, findOut)
		findCmd.Wait()
	}()

	if err := fzf.Run(); err != nil {
		if IsCanceled(err) {
			log.Println("选择已取消")
			return
		}
		log.Fatal(err)
	}

}

func IsCanceled(err error) bool {
	if exitErr, ok := err.(*exec.ExitError); ok {
		return exitErr.ExitCode() == 130
	}
	return false
}
