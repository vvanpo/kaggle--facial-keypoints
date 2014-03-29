
package main

import (
    "fmt"
    "os"
    "log"
    "image"
    _ "image/png"
    _ "image/jpeg"
    "github.com/conformal/gotk3/gdk"
    "github.com/conformal/gotk3/gtk"
)

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
        width := img.Bounds().Max.X - img.Bounds().Min.X
        height := img.Bounds().Max.Y - img.Bounds().Min.Y
        fmt.Printf("(%d, %d)\n", width, height)
        // Initialize a Pixbuf to copy img into
        pb, err := gdk.PixbufNew(gdk.COLORSPACE_RGB, false, 8, width, height)
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
