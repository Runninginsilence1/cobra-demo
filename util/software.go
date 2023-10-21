package util

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"os/exec"
	"regexp"
	"runtime"
	"sync"
)

// IsSoftwareInstalled validates whether a linux software or a windows software is installed.
// In Windows, it will check the Windows Registry.
// Please note that this method is only effective if the software installation modified the Windows Registry.
// It cannot detect the presence of portable or standalone versions of the software that do not modify the Registry.
func IsSoftwareInstalled(software string) bool {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("which", software)
		err := cmd.Run()
		if err != nil {
			return false
		}
		return true
	case "windows":
		appPaths := hasApp(software)
		if len(appPaths) != 0 {
			return true
		}
		return false
	default:
		return false
	}

}

// hasApp checks Windows Registry to find software path.
func hasApp(appName string) []string {
	queryKey := func(w *sync.WaitGroup, startKey registry.Key, res *[]string) {
		defer w.Done()
		queryPath := "Software\\Microsoft\\Windows\\CurrentVersion\\App Paths\\"
		k, err := registry.OpenKey(startKey, queryPath, registry.READ)
		if err != nil {
			return
		}
		// 读取所有子项
		keyNames, err := k.ReadSubKeyNames(0)
		if err != nil {
			return
		}
		for _, v := range keyNames {
			matched, err := regexp.MatchString(appName, v)
			if err != nil {
				fmt.Println("regexp error:", err)
			} else {
				if matched {
					tmpRegPath := queryPath + "\\" + v
					appKey, _ := registry.OpenKey(startKey, tmpRegPath, registry.READ)
					s, _, err := appKey.GetStringValue("")
					if err != nil {
						fmt.Println(err)
					} else {
						*res = append(*res, s)
					}
				}
			}
		}
	}
	res := []string{}

	waitGroup := new(sync.WaitGroup)
	waitGroup.Add(2)

	go queryKey(waitGroup, registry.LOCAL_MACHINE, &res)
	go queryKey(waitGroup, registry.CURRENT_USER, &res)
	waitGroup.Wait()

	return res
}
