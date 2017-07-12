package main

import (
	"flag"
	"io/ioutil"
	"os"
	"runtime"

	"strings"

	"github.com/golang/glog"
)

const (
	sm = "sm"
)

var (
	smPtr = flag.Bool(sm, false, "handle //s/m style")
)

func main() {
	defer glog.Flush()
	glog.CopyStandardLogTo("info")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	err := do()
	if err != nil {
		glog.Exit(err)
	}
}

func do() error {
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	glog.V(1).Infoln("converting:", string(content))
	if *smPtr {
		str := string(content)
		str = strings.TrimSpace(str)
		str = strings.TrimSuffix(str, ".smhss.de")
		content = []byte(str)
		glog.V(1).Infoln("trimming .smhss.de from str:", str)
	}
	output := rewind(content)
	if *smPtr {
		output = append(output, []byte(".smhss.de")...)
		glog.V(1).Infoln("appending .smhss.de:", string(output))
	}
	str := string(output)
	str = strings.TrimSpace(str)
	str = strings.TrimPrefix(str, "de.smhss.")
	output = []byte(str)
	if _, err := os.Stdout.Write(output); err != nil {
		return err
	}
	return nil
}

//Rewind the given data by it's separator "."
func rewind(data []byte) []byte {
	str := string(data)
	str = strings.TrimSpace(str)
	arr := strings.Split(str, ".")
	glog.V(2).Infoln("splitting string:", arr)
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	str = strings.Join(arr, ".")
	glog.V(2).Infof("joining string(%s) from array(%v)", str, arr)
	return []byte(str)
}
