package main

import (
  "fmt"
 "image"
 "image/jpeg"
	"image/png"
	"log"
 "os"
)

func main() {
  jobs := len(os.Args)-1
  if jobs == 0 {
    fmt.Println("No images provided")
  }

  for i := 0; i < jobs; i++ {
    filename := os.Args[1+i]
    fmt.Printf("Starting on %s", filename)
    ScaleImage(filename);
    fmt.Println("...done")
  }
}

func ScaleImage(filename string) {
  reader, err := os.Open(filename)
	 if err != nil {
    log.Fatal(err)
	 }

	 img, format, err := image.Decode(reader)
	 if err != nil {
    log.Fatal(err)
	 }

  fitDim := 640
  yScale := img.Bounds().Max.Y / fitDim
  xScale := img.Bounds().Max.X / fitDim

  scale := xScale
  if yScale < xScale {
    scale = xScale
  }

  yMax := img.Bounds().Max.Y / scale
  xMax := img.Bounds().Max.X / scale

  output := image.NewNRGBA(image.Rect(0,0,xMax,yMax))

	 for y := 0; y < yMax; y++ {
		  for x := 0; x < xMax; x++ {
      xPos := x * scale
      yPos := y * scale
      c := img.At(xPos, yPos)
      output.Set(x, y, c)
		  }
	 }

  f, err := os.Create(filename)
	 if err != nil {
    log.Fatal(err)
	 }

  if format == "png" {
    if err := png.Encode(f, output); err != nil {
      f.Close()
		    log.Fatal(err)
    }
  }

  if format == "jpeg" {
    if err := jpeg.Encode(f, output, &jpeg.Options{100}); err != nil {
      f.Close()
		    log.Fatal(err)
    }
  }


	 if err := f.Close(); err != nil {
    log.Fatal(err)
	 }
}
