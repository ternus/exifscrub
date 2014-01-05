/*

EXIF Scrubber

Removes EXIF and other user-defined extensions (JPEG marker tag
0xFFE0-0xFFEF) from JPEG files, including GPS data.



*/ 

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <in.jpg> <out.jpg>\n", os.Args[0])
		return
	}

	buf, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	if string(buf[:2]) != "\xFF\xD8" {
		log.Fatal("Not a JPEG file!")
	}

	offset := int64(2) // Start from 2 bytes (after FF D8)

	for {

		if buf[offset] == 0xFF && (buf[offset+1]&0xf0 == 0xE0) {
			osize := int64(buf[offset+2])<<8 + int64(buf[offset+3])
			fmt.Printf("Removing extension (%x%x); size %d\n", buf[offset], buf[offset+1], osize)
			offset += osize + 2
		} else {
			break
		}
	}

	out, err := os.Create(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	out.Write(buf[:2])

	out.Write(buf[offset:])

	fmt.Println("Done.")
}
