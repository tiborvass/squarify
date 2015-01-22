# Squarify

Increase the canvas size of any image to make it square

## Library
[https://godoc.org/github.com/tiborvass/squarify](https://godoc.org/github.com/tiborvass/squarify)

## Command

```
$ go get github.com/tiborvass/squarify/cmd/squarify
$ squarify --help
squarify - Increase canvas size of any image to make it square

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
The format is always preserved, no conversion is made.
```
