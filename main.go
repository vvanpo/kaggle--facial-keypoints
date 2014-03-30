
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

// Copies image to a new GdkPixbuf
func imageToPixbuf(img image.Image) (*gdk.Pixbuf, error) {
    width := img.Bounds().Max.X - img.Bounds().Min.X
    height := img.Bounds().Max.Y - img.Bounds().Min.Y
    pixbuf, err := gdk.PixbufNew(gdk.COLORSPACE_RGB, true, 8, width, height)
    if err != nil {
        return nil, err
    }
    pix := pixbuf.GetPixels()
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            i := ((width * y) + x) * 4
            r, g, b, a := img.At(x, y).RGBA()
            pix[i] = byte(r)
            pix[i + 1] = byte(g)
            pix[i + 2] = byte(b)
            pix[i + 3] = byte(a)
        }
    }
    return pixbuf, nil
}

// Displays image in a new Gtk window
// This function blocks until the window is closed
func displayImage(img image.Image) {
        // Copy img into a new gdk.Pixbuf
        pixbuf, err := imageToPixbuf(img)
        if err != nil {
            log.Print("Failed to create a pixbuf for image")
        }
        // Initialize GtkWindow
        gtk.Init(nil)
        win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
        if err != nil {
            log.Print("Failed to create new window")
        }
        win.Connect("destroy", func() { gtk.MainQuit() })
        // Create GtkImage widget from pixbuf
        gimg, err := gtk.ImageNewFromPixbuf(pixbuf)
        if err != nil {
            log.Print("Failed to load image into window")
        }
        win.Add(gimg)
        win.ShowAll()
        gtk.Main()
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
        displayImage(img)
    }
}
