
package main

import (
    "os"
    "log"
    "image"
    _ "image/png"
    _ "image/jpeg"
    "github.com/conformal/gotk3/gdk"
    "github.com/conformal/gotk3/gtk"
)

func ImageToPixbuf(img image.Image) (*gdk.Pixbuf, error) {
    width := img.Bounds().Max.X - img.Bounds().Min.X
    height := img.Bounds().Max.Y - img.Bounds().Min.Y
    pixbuf, err := gdk.PixbufNew(gdk.COLORSPACE_RGB, true, 8, width, height)
    if err != nil {
        return nil, err
    }
    pix := make([]byte, width * height * 4)
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            i := ((x * y) + x) * 4
            r, g, b, a := img.At(x, y).RGBA()
            pix[i] = byte(r / 256)
            pix[i + 1] = byte(g / 256)
            pix[i + 2] = byte(b / 256)
            pix[i + 3] = byte(a / 256)
        }
    }
    err = pixbuf.SetPixels(pix)
    if err != nil {
        return nil, err
    }
    return pixbuf, nil
}

func main() {

    for _, file := range os.Args[1:] {
        f, err := os.Open(file)
        if err != nil {
            log.Print("Failed to open file ", file)
            continue
        }
        img, _, err := image.Decode(f)
        if err != nil {
            log.Print("Failed to decode image ", file)
            continue
        }

        // Copy img into a new gdk.Pixbuf
        pb, err := ImageToPixbuf(img)
        if err != nil {
            log.Print("Failed to create a pixbuf for image ", file)
            continue
        }

        // Initialize GTK window
        gtk.Init(nil)
        win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
        win.Connect("destroy", func() { gtk.MainQuit() })
        i, err := gtk.ImageNewFromPixbuf(pb)
        win.Add(i)
        win.ShowAll()
        gtk.Main()
    }

}
