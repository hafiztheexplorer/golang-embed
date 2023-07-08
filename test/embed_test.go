package golangembed

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"testing"
)

/*----------------------------------------------------------------
perhatikan penulisan di bawah ini harus mepet //tulisan utuk baca file / embed file
----------------------------------------------------------------*/

//go:embed version.txt
var version_txt string

func TestString(t *testing.T) {

	fmt.Println(version_txt)

}

/*----------------------------------------------------------------
untuk embed menduplikasi file gambar, atau audio, video juga seperti tersebut
----------------------------------------------------------------*/

//go:embed Screenshot-2023-07-02-162102.jpg
var Screenshot_2023_07_02_162102_jpg []byte

func TestByte(t *testing.T) {
	// kita coba buka file dijadikan string
	err := ioutil.WriteFile("Newfile.jpg", Screenshot_2023_07_02_162102_jpg, fs.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

}

/*----------------------------------------------------------------
untuk embed beberapa file
----------------------------------------------------------------*/

//go:embed files/Newfile1.jpg
//go:embed files/Newfile2.txt
//go:embed files/Newfile3.txt
var files embed.FS

func TestEmbedMultiple(t *testing.T) {
	byte1, error := files.ReadFile("files/Newfile1.jpg")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(string(byte1))

	byte2, error := files.ReadFile("files/Newfile2.txt")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(string(byte2))

	byte3, error := files.ReadFile("files/Newfile3.txt")
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(string(byte3))

}

/*----------------------------------------------------------------
untuk embed beberapa file dengan otomatis search di sat folder, dengan path matcher
----------------------------------------------------------------*/

// di bawah ini nama direktori yang akan dituju dan filenya apa
// atau bisa juga files/* berati semua file yang di dalam serta kalau ada folder lagi di dalamnya, kalau *.txt berarti txt aja, kalau file saja cuma di foder itu aja dan misal ada folder lagi di dalamnya ngga dihitung juga

//go:embed files
var path embed.FS

func TestPathMatcher(t *testing.T) {
	DirEntry, error := path.ReadDir("files") // baca folder
	if error != nil {
		log.Fatal(error)
	}
	for _, entry := range DirEntry { // ini bentuk for, sebelum koma itu index
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			// sampe di atas ini bakalan ngebaca isi filenya apa aja sesuai dengan go embed yang diatas, coba aja ganti ke .txt dan * aja, which is why we do it in the first place

			isifolder, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("content: ", isifolder)
		}
	}
}
