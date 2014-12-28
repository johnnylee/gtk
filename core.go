package gtk

import (
	"log"
	"sync"
	"unsafe"
)

// #include "core.h"
// #cgo pkg-config: gtk+-3.0
import "C"

//export callbackMapper
func callbackMapper(cWidget *C.GtkWidget, slotNum C.int) {
	callback(NewWidget(unsafe.Pointer(cWidget)), int(slotNum))
}

// Our development will be focused around the widget type.
type Widget struct {
	ptr unsafe.Pointer
}

func NewWidget(ptr unsafe.Pointer) Widget {
	w := Widget{}
	w.ptr = ptr
	return w
}

// The global Gtk object.
var Gtk *gtk

// gtk is a structure for a global object that will synchronize access
// shared state.
type gtk struct {
	mutex      sync.Mutex
	callbacks  map[int]func(Widget)
	builderPtr *C.GtkBuilder
}

func callback(widget Widget, slotNum int) {
	g := Gtk
	g.mutex.Lock()
	cb, ok := g.callbacks[slotNum]
	g.mutex.Unlock()

	if !ok {
		// TODO: Log.
		return
	}

	cb(widget)
}

func SignalConnect(widget Widget, signal string, cb func(Widget)) {
	g := Gtk
	g.mutex.Lock()
	defer g.mutex.Unlock()
	slotNum := len(g.callbacks)
	g.callbacks[slotNum] = cb
	C.ConnectSignal(
		(*C.GtkWidget)(widget.ptr), C.CString(signal), C.int(slotNum))
}

func Init() {
	C.gtk_init(nil, nil)
	Gtk = new(gtk)
	Gtk.callbacks = make(map[int]func(Widget))
	Gtk.builderPtr = C.gtk_builder_new()
}

func Main() {
	C.gtk_main()
}

func MainQuit() {
	C.gtk_main_quit()
}

func AddFromFile(path string) {
	C.gtk_builder_add_from_file(Gtk.builderPtr, StringToGChar(path), nil)
}

func AddFromString(defn []byte) {
	cdef := (*C.gchar)(unsafe.Pointer(&defn[0]))
	size := len(defn)
	C.gtk_builder_add_from_string(
		Gtk.builderPtr, cdef, C.gsize(size), nil)
}

func GetWidget(name string) Widget {
	objPtr := C.gtk_builder_get_object(Gtk.builderPtr, StringToGChar(name))
	ptr := unsafe.Pointer(objPtr)
	if ptr == nil {
		log.Println("Failed to find widget:", name)
	}
	return NewWidget(ptr)
}
