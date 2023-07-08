package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
)

//go:embed version.txt
var version_txt string

//go:embed Screenshot-2023-07-02-162102.jpg
var Screenshot_2023_07_02_162102_jpg []byte

//go:embed files
var path embed.FS

func main() {
	fmt.Println(version_txt)

	err := os.WriteFile("Newfile.jpg", Screenshot_2023_07_02_162102_jpg, fs.ModeType)
	if err != nil {
		panic(err)
	}

	DirEntry, _ := path.ReadDir("files") // baca folder
	for _, entry := range DirEntry {     // ini bentuk for, sebelum koma itu index
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			// sampe di atas ini bakalan ngebaca isi filenya apa aja sesuai dengan go embed yang diatas, coba aja ganti ke .txt dan * aja, which is why we do it in the first place

			isifolder, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(isifolder))
		}
	}

	// build aplikasi dengan perintah
	// go build
	// di local terminal
}
