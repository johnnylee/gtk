from __future__ import print_function


def go_name(name):
    """Create a Go function name from a Gtk function name."""
    if name[:4] == "gtk_":
        name = name[4:]
    if name[:7] == "widget_":
        name = name[7:]
    name = name.title()
    name = name.replace('_', '')
    return name


def type_cast(type_):
    if type_ == 'GtkWidget':
        return 'w.ptr'

    return '(*C.{0})(unsafe.Pointer(w.ptr))'.format(type_)


def command(type_, *names):
    for name in names:
        print('func (w Widget) {0}() {{'.format(go_name(name)))
        print('ptr := (*C.{0})(w.ptr)'.format(type_))
        print('C.{0}(ptr)'.format(name))
        print('}')


def set_bool(type_, *names):
    for name in names:
        print('func (w Widget) {0}(b bool) {{'.format(go_name(name)))
        print('ptr := (*C.{0})(w.ptr)'.format(type_))
        print('C.{0}(ptr, BoolToGBoolean(b))'.format(name))
        print('}')


def get_bool(type_, *names):
    for name in names:
        print('func (w Widget) {0}() bool {{'.format(go_name(name)))
        print('ptr := (*C.{0})(w.ptr)'.format(type_))
        print('return GBooleanToBool(C.{0}(ptr))'.format(name))
        print('}')


def set_string(type_, *names):
    for name in names:
        print('func (w Widget) {0}(s string) {{'.format(go_name(name)))
        print('ptr := (*C.{0})(w.ptr)'.format(type_))
        print('C.{0}(ptr, StringToGChar(s))'.format(name))
        print('}')


def get_string(type_, *names):
    for name in names:
        print('func (w Widget) {0}() string {{'.format(go_name(name)))
        print('ptr := (*C.{0})(w.ptr)'.format(type_))
        print('return GCharToString(C.{0}(ptr))'.format(name))
        print('}')


def set_float(type_, *names):
    for name in names:
        print('func (w Widget) {0}(f float64) {{'.format(go_name(name)))
        print('ptr := (*C.{0})(w.ptr)'.format(type_))
        print('C.{0}(ptr, FloatToGDouble(f))'.format(name))
        print('}')


def get_float(type_, *names):
    for name in names:
        print('func (w Widget) {0}() float64 {{'.format(go_name(name)))
        print('ptr := (*C.{0})(w.ptr)'.format(type_))
        print('return GDoubleToFloat(C.{0}(ptr))'.format(name))
        print('}')


def set_floatfloat(type_, *names):
    for name in names:
        print('func (w Widget) {0}(f1, f2 float64) {{'.format(go_name(name)))
        print('ptr := (*C.{0})(w.ptr)'.format(type_))
        print('C.{0}(ptr, FloatToGDouble(f1), FloatToGDouble(f2))'.format(
              name))
        print('}')


if __name__ == "__main__":
    print('package gtk')
    print('// #include <gtk/gtk.h>')
    print('import "C"')

    ##################################################
    # Widget
    command('GtkWidget',
            'gtk_widget_show',
            'gtk_widget_hide',
            'gtk_widget_grab_focus',
            'gtk_widget_grab_default')

    get_bool('GtkWidget',
             'gtk_widget_activate',
             'gtk_widget_is_focus',
             'gtk_widget_get_sensitive',
             'gtk_widget_is_sensitive',
             'gtk_widget_get_visible')

    set_bool('GtkWidget',
             'gtk_widget_set_sensitive',
             'gtk_widget_set_visible')

    get_string('GtkWidget',
               'gtk_widget_get_tooltip_text',
               'gtk_widget_get_tooltip_markup')

    set_string('GtkWidget',
               'gtk_widget_set_tooltip_markup',
               'gtk_widget_set_tooltip_text')

    ##################################################
    # Window
    get_string('GtkWindow',
               'gtk_window_get_title')

    set_string('GtkWindow',
               'gtk_window_set_title')

    ##################################################
    # Label
    get_string('GtkLabel',
               'gtk_label_get_text',
               'gtk_label_get_label')

    set_string('GtkLabel',
               'gtk_label_set_text',
               'gtk_label_set_markup',
               'gtk_label_set_markup_with_mnemonic',
               'gtk_label_set_pattern')

    ##################################################
    # ProgressBar
    command('GtkProgressBar',
            'gtk_progress_bar_pulse')

    get_float('GtkProgressBar',
              'gtk_progress_bar_get_fraction',
              'gtk_progress_bar_get_pulse_step')

    set_float('GtkProgressBar',
              'gtk_progress_bar_set_fraction',
              'gtk_progress_bar_set_pulse_step')

    ##################################################
    # Spinner
    command('GtkSpinner',
            'gtk_spinner_start',
            'gtk_spinner_stop')

    ##################################################
    # Button
    get_string('GtkButton',
               'gtk_button_get_label')

    set_string('GtkButton',
               'gtk_button_set_label')

    ##################################################
    # ToggleButton
    set_bool('GtkToggleButton',
             'gtk_toggle_button_set_active',
             'gtk_toggle_button_set_inconsistent')

    get_bool('GtkToggleButton',
             'gtk_toggle_button_get_active',
             'gtk_toggle_button_get_inconsistent')

    ##################################################
    # LinkButton
    set_string('GtkLinkButton', 'gtk_link_button_set_uri')
    get_string('GtkLinkButton', 'gtk_link_button_get_uri')
    set_bool('GtkLinkButton', 'gtk_link_button_set_visited')
    get_bool('GtkLinkButton', 'gtk_link_button_get_visited')

    ##################################################
    # ScaleButton
    set_float('GtkScaleButton', 'gtk_scale_button_set_value')
    get_float('GtkScaleButton', 'gtk_scale_button_get_value')

    ##################################################
    # Switch
    get_bool('GtkSwitch', 'gtk_switch_get_active')
    set_bool('GtkSwitch', 'gtk_switch_set_active')

    ##################################################
    # Entry
    command('GtkEntry', 'gtk_entry_progress_pulse')
    set_string('GtkEntry', 'gtk_entry_set_text')
    get_string('GtkEntry', 'gtk_entry_get_text')
    set_bool('GtkEntry', 'gtk_entry_set_visibility')
    get_bool('GtkEntry', 'gtk_entry_get_visibility')
    set_float('GtkEntry',
              'gtk_entry_set_progress_fraction',
              'gtk_entry_set_progress_pulse_step')
    get_float('GtkEntry',
              'gtk_entry_get_progress_fraction',
              'gtk_entry_get_progress_pulse_step')

    ##################################################
    # Range
    get_float('GtkRange', 'gtk_range_get_value')
    set_float('GtkRange', 'gtk_range_set_value')
    set_floatfloat('GtkRange',
                   'gtk_range_set_increments',
                   'gtk_range_set_range')

    ##################################################
    # SpinButton
    w = 'GtkSpinButton'
    get_float(w, 'gtk_spin_button_get_value')
    set_float(w, 'gtk_spin_button_set_value')
    get_bool(w, 'gtk_spin_button_get_wrap')
    set_bool(w, 'gtk_spin_button_set_wrap')
    set_floatfloat(w,
                   'gtk_spin_button_set_increments',
                   'gtk_spin_button_set_range')

    ##################################################
    # ComboBoxText
    w = 'GtkComboBoxText'
    command(w, 'gtk_combo_box_text_remove_all')
    set_string(w,
               'gtk_combo_box_text_append_text',
               'gtk_combo_box_text_prepend_text')
    get_string(w, 'gtk_combo_box_text_get_active_text')

    ##################################################
    # MenuItem
    w = 'GtkMenuItem'
    get_string(w, 'gtk_menu_item_get_label')
    set_string(w, 'gtk_menu_item_set_label')

    ##################################################
    # CheckMenuItem
    w = 'GtkCheckMenuItem'
    get_bool(w,
             'gtk_check_menu_item_get_active',
             'gtk_check_menu_item_get_inconsistent')
    set_bool(w,
             'gtk_check_menu_item_set_active',
             'gtk_check_menu_item_set_inconsistent')

    ##################################################
    # ToolButton
    w = 'GtkToolButton'
    set_string(w, 'gtk_tool_button_set_label')
    get_string(w, 'gtk_tool_button_get_label')

    ##################################################
    # ToggleToolButton
    w = 'GtkToggleToolButton'
    get_bool(w, 'gtk_toggle_tool_button_get_active')
    set_bool(w, 'gtk_toggle_tool_button_set_active')
