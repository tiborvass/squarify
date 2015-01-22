package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"

	"github.com/tiborvass/squarify"
)

func square(src io.Reader, w io.Writer) error {
	srcImg, format, err := image.Decode(src)
	if err != nil {
		return err
	}
	dstImg := squarify.Image(srcImg, nil)
	switch format {
	case "jpeg":
		return jpeg.Encode(w, dstImg, nil)
	case "gif":
		return gif.Encode(w, dstImg, nil)
	case "png":
		return png.Encode(w, dstImg)
	}
	return fmt.Errorf("Could not find encoder for format %q", format)
}

var help bool

func init() {
	flag.BoolVar(&help, "-help", false, "Display help")
	flag.Usage = func() {
		fmt.Println(`squarify - Increase the canvas size of any image to make it square

Usage: squarify [--help] [filenames...]

If filenames are passed as arguments, then each corresponding image is converted
to a new square image by increasing the canvas size, without cropping or skewing
the image.
The void is filled with white. New images have '.square' in their filename.

Example:
$ squarify image.jpeg
New file: image.square.jpeg

If no arguments are provided, then STDIN is used to decode the source image, and
the result is sent to STDOUT.

Example:
$ cat image.jpeg | squarify > squareimage.jpeg

Supported image formats are: GIF, JPEG, PNG
The format is always preserved, no conversion is made.`)
		os.Exit(0)
	}
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if len(os.Args[1:]) == 0 {
		if err := square(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	} else {
		for _, file := range os.Args[1:] {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			defer f.Close()
			ext := filepath.Ext(file)
			newFile := file[:len(file)-len(ext)] + ".square" + ext
			dest, err := os.Create(newFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			defer dest.Close()
			if err := square(f, dest); err != nil {
				fmt.Fprintln(os.Stderr, err)
				return
			}
			fmt.Fprintf(os.Stderr, "New file: %s\n", newFile)
		}
	}
}
