package main

import "fmt"
import "io/ioutil"
import "os"
import "strconv"
import "github.com/atotto/clipboard"

func main() {
	files, err := ioutil.ReadDir("/Users/thomashorrobin/Projects")
	if err != nil {
		fmt.Println(err)
	}

	if len(os.Args) > 1 {
		arg := os.Args[1]
		i, _ := strconv.ParseInt(arg, 10, 64)
		dir := "/Users/thomashorrobin/Projects/" + files[i].Name()
		clipboard.WriteAll("cd " + dir);
		fmt.Println(dir)
	} else {
		for i, file := range files {
			if file.IsDir() {
				fmt.Printf("[%d] %s\n", i, file.Name())
			}
		}
	}
}