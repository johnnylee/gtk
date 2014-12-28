
#ifndef GO_GTK_H
#define GO_GTK_H

#include <gtk/gtk.h>

extern void callbackMapper(GtkWidget*, int);


static void CallbackWrapper(GtkWidget *object, void *slotNum) {
  callbackMapper(object, (int)slotNum);
}

static inline void ConnectSignal(GtkWidget *object, 
                                 char *signal,
                                 int slotNum) {
  g_signal_connect(object, 
                   signal, 
                   (GCallback)CallbackWrapper,
                   (gpointer)slotNum);
}

#endif
