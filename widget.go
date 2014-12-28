package gtk

// #include <gtk/gtk.h>
import "C"

// Widget
func (w Widget) SignalConnect(signal string, cb func(Widget)) {
	SignalConnect(w, signal, cb)
}

// TextView
func (w Widget) TextViewGetText() string {
	ptr := (*C.GtkTextView)(w.ptr)
	buffer := C.gtk_text_view_get_buffer(ptr)
	var start, end C.GtkTextIter
	C.gtk_text_buffer_get_bounds(buffer, &start, &end)

	return GCharToString(C.gtk_text_buffer_get_text(
		buffer, &start, &end, BoolToGBoolean(true)))
}

func (w Widget) TextViewSetText(s string) {
	ptr := (*C.GtkTextView)(w.ptr)
	buffer := C.gtk_text_view_get_buffer(ptr)
	C.gtk_text_buffer_set_text(buffer, StringToGChar(s), C.gint(-1))
}

// ComboBoxText
func (w Widget) ComboBoxTextInsertText(pos int, s string) {
	ptr := (*C.GtkComboBoxText)(w.ptr)
	C.gtk_combo_box_text_insert_text(ptr, C.gint(pos), StringToGChar(s))
}

func (w Widget) ComboBoxTextRemove(pos int) {
	ptr := (*C.GtkComboBoxText)(w.ptr)
	C.gtk_combo_box_text_remove(ptr, C.gint(pos))
}

// ComboBox
func (w Widget) ComboBoxGetActive() int {
	ptr := (*C.GtkComboBox)(w.ptr)
	return int(C.gtk_combo_box_get_active(ptr))
}

func (w Widget) ComboBoxSetActive(idx int) {
	ptr := (*C.GtkComboBox)(w.ptr)
	C.gtk_combo_box_set_active(ptr, C.gint(idx))
}

// ColorChooser
func (w Widget) ColorChooserGetRgba() string {
	ptr := (*C.GtkColorChooser)(w.ptr)
	color := C.GdkRGBA{}
	C.gtk_color_chooser_get_rgba(ptr, &color)
	return GCharToString(C.gdk_rgba_to_string(&color))
}

func (w Widget) ColorChooserSetRgba(s string) {
	ptr := (*C.GtkColorChooser)(w.ptr)
	color := C.GdkRGBA{}
	C.gdk_rgba_parse(&color, StringToGChar(s))
	C.gtk_color_chooser_set_rgba(ptr, &color)
}

// ColorButton
func (w Widget) ColorButtonGetRgba() string {
	ptr := (*C.GtkColorButton)(w.ptr)
	color := C.GdkRGBA{}
	C.gtk_color_button_get_rgba(ptr, &color)
	return GCharToString(C.gdk_rgba_to_string(&color))
}

func (w Widget) ColorButtonSetRgba(s string) {
	ptr := (*C.GtkColorButton)(w.ptr)
	color := C.GdkRGBA{}
	C.gdk_rgba_parse(&color, StringToGChar(s))
	C.gtk_color_button_set_rgba(ptr, &color)
}

// FileChooser
func (w Widget) FileChooserGetFile() string {
	ptr := (*C.GtkFileChooser)(w.ptr)
	return GFileToString(C.gtk_file_chooser_get_file(ptr))
}

func (w Widget) FileChooserSetFile(path string) {
	ptr := (*C.GtkFileChooser)(w.ptr)
	gfile := StringToGFile(path)
	defer C.g_object_unref((C.gpointer)(gfile))
	C.gtk_file_chooser_set_file(ptr, gfile, nil)
}

/* TODO
func (w Widget) FileChooserGetFiles() []string {
	paths := make([]string, 0)

}
*/
