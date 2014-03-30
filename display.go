
package main

import (
    "errors"
    "image"
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
func displayImage(img image.Image) error {
        // Copy img into a new gdk.Pixbuf
        pixbuf, err := imageToPixbuf(img)
        if err != nil {
            return errors.New("Failed to create a pixbuf for image")
        }
        // Initialize GtkWindow
        gtk.Init(nil)
        win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
        if err != nil {
            return errors.New("Failed to create new window")
        }
        win.Connect("destroy", func() { gtk.MainQuit() })
        // Create GtkImage widget from pixbuf
        gimg, err := gtk.ImageNewFromPixbuf(pixbuf)
        if err != nil {
            return errors.New("Failed to load image into window")
        }
        win.Add(gimg)
        win.ShowAll()
        gtk.Main()
        return nil
}
