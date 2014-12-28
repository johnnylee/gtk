package gtk

// #include <gtk/gtk.h>
import "C"

func StringToGChar(s string) *C.gchar {
	return (*C.gchar)(C.CString(s))
}

func GCharToString(s *C.gchar) string {
	if s == nil {
		return ""
	}
	return C.GoString((*C.char)(s))
}

func BoolToGBoolean(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}

func GBooleanToBool(b C.gboolean) bool {
	return int(C.int(b)) != 0
}

func FloatToGDouble(f float64) C.gdouble {
	return C.gdouble(f)
}

func GDoubleToFloat(f C.gdouble) float64 {
	return float64(C.double(f))
}

func GFileToString(gfile *C.GFile) string {
	defer C.g_object_unref((C.gpointer)(gfile))
	cstring := C.g_file_get_path(gfile)
	defer C.g_free((C.gpointer)(cstring))
	return C.GoString(cstring)
}

func StringToGFile(s string) *C.GFile {
	return C.g_file_new_for_path(C.CString(s))
}
