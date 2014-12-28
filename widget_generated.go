package gtk

// #include <gtk/gtk.h>
import "C"

func (w Widget) Show() {
	ptr := (*C.GtkWidget)(w.ptr)
	C.gtk_widget_show(ptr)
}
func (w Widget) Hide() {
	ptr := (*C.GtkWidget)(w.ptr)
	C.gtk_widget_hide(ptr)
}
func (w Widget) GrabFocus() {
	ptr := (*C.GtkWidget)(w.ptr)
	C.gtk_widget_grab_focus(ptr)
}
func (w Widget) GrabDefault() {
	ptr := (*C.GtkWidget)(w.ptr)
	C.gtk_widget_grab_default(ptr)
}
func (w Widget) Activate() bool {
	ptr := (*C.GtkWidget)(w.ptr)
	return GBooleanToBool(C.gtk_widget_activate(ptr))
}
func (w Widget) IsFocus() bool {
	ptr := (*C.GtkWidget)(w.ptr)
	return GBooleanToBool(C.gtk_widget_is_focus(ptr))
}
func (w Widget) GetSensitive() bool {
	ptr := (*C.GtkWidget)(w.ptr)
	return GBooleanToBool(C.gtk_widget_get_sensitive(ptr))
}
func (w Widget) IsSensitive() bool {
	ptr := (*C.GtkWidget)(w.ptr)
	return GBooleanToBool(C.gtk_widget_is_sensitive(ptr))
}
func (w Widget) GetVisible() bool {
	ptr := (*C.GtkWidget)(w.ptr)
	return GBooleanToBool(C.gtk_widget_get_visible(ptr))
}
func (w Widget) SetSensitive(b bool) {
	ptr := (*C.GtkWidget)(w.ptr)
	C.gtk_widget_set_sensitive(ptr, BoolToGBoolean(b))
}
func (w Widget) SetVisible(b bool) {
	ptr := (*C.GtkWidget)(w.ptr)
	C.gtk_widget_set_visible(ptr, BoolToGBoolean(b))
}
func (w Widget) GetTooltipText() string {
	ptr := (*C.GtkWidget)(w.ptr)
	return GCharToString(C.gtk_widget_get_tooltip_text(ptr))
}
func (w Widget) GetTooltipMarkup() string {
	ptr := (*C.GtkWidget)(w.ptr)
	return GCharToString(C.gtk_widget_get_tooltip_markup(ptr))
}
func (w Widget) SetTooltipMarkup(s string) {
	ptr := (*C.GtkWidget)(w.ptr)
	C.gtk_widget_set_tooltip_markup(ptr, StringToGChar(s))
}
func (w Widget) SetTooltipText(s string) {
	ptr := (*C.GtkWidget)(w.ptr)
	C.gtk_widget_set_tooltip_text(ptr, StringToGChar(s))
}
func (w Widget) WindowGetTitle() string {
	ptr := (*C.GtkWindow)(w.ptr)
	return GCharToString(C.gtk_window_get_title(ptr))
}
func (w Widget) WindowSetTitle(s string) {
	ptr := (*C.GtkWindow)(w.ptr)
	C.gtk_window_set_title(ptr, StringToGChar(s))
}
func (w Widget) LabelGetText() string {
	ptr := (*C.GtkLabel)(w.ptr)
	return GCharToString(C.gtk_label_get_text(ptr))
}
func (w Widget) LabelGetLabel() string {
	ptr := (*C.GtkLabel)(w.ptr)
	return GCharToString(C.gtk_label_get_label(ptr))
}
func (w Widget) LabelSetText(s string) {
	ptr := (*C.GtkLabel)(w.ptr)
	C.gtk_label_set_text(ptr, StringToGChar(s))
}
func (w Widget) LabelSetMarkup(s string) {
	ptr := (*C.GtkLabel)(w.ptr)
	C.gtk_label_set_markup(ptr, StringToGChar(s))
}
func (w Widget) LabelSetMarkupWithMnemonic(s string) {
	ptr := (*C.GtkLabel)(w.ptr)
	C.gtk_label_set_markup_with_mnemonic(ptr, StringToGChar(s))
}
func (w Widget) LabelSetPattern(s string) {
	ptr := (*C.GtkLabel)(w.ptr)
	C.gtk_label_set_pattern(ptr, StringToGChar(s))
}
func (w Widget) ProgressBarPulse() {
	ptr := (*C.GtkProgressBar)(w.ptr)
	C.gtk_progress_bar_pulse(ptr)
}
func (w Widget) ProgressBarGetFraction() float64 {
	ptr := (*C.GtkProgressBar)(w.ptr)
	return GDoubleToFloat(C.gtk_progress_bar_get_fraction(ptr))
}
func (w Widget) ProgressBarGetPulseStep() float64 {
	ptr := (*C.GtkProgressBar)(w.ptr)
	return GDoubleToFloat(C.gtk_progress_bar_get_pulse_step(ptr))
}
func (w Widget) ProgressBarSetFraction(f float64) {
	ptr := (*C.GtkProgressBar)(w.ptr)
	C.gtk_progress_bar_set_fraction(ptr, FloatToGDouble(f))
}
func (w Widget) ProgressBarSetPulseStep(f float64) {
	ptr := (*C.GtkProgressBar)(w.ptr)
	C.gtk_progress_bar_set_pulse_step(ptr, FloatToGDouble(f))
}
func (w Widget) SpinnerStart() {
	ptr := (*C.GtkSpinner)(w.ptr)
	C.gtk_spinner_start(ptr)
}
func (w Widget) SpinnerStop() {
	ptr := (*C.GtkSpinner)(w.ptr)
	C.gtk_spinner_stop(ptr)
}
func (w Widget) ButtonGetLabel() string {
	ptr := (*C.GtkButton)(w.ptr)
	return GCharToString(C.gtk_button_get_label(ptr))
}
func (w Widget) ButtonSetLabel(s string) {
	ptr := (*C.GtkButton)(w.ptr)
	C.gtk_button_set_label(ptr, StringToGChar(s))
}
func (w Widget) ToggleButtonSetActive(b bool) {
	ptr := (*C.GtkToggleButton)(w.ptr)
	C.gtk_toggle_button_set_active(ptr, BoolToGBoolean(b))
}
func (w Widget) ToggleButtonSetInconsistent(b bool) {
	ptr := (*C.GtkToggleButton)(w.ptr)
	C.gtk_toggle_button_set_inconsistent(ptr, BoolToGBoolean(b))
}
func (w Widget) ToggleButtonGetActive() bool {
	ptr := (*C.GtkToggleButton)(w.ptr)
	return GBooleanToBool(C.gtk_toggle_button_get_active(ptr))
}
func (w Widget) ToggleButtonGetInconsistent() bool {
	ptr := (*C.GtkToggleButton)(w.ptr)
	return GBooleanToBool(C.gtk_toggle_button_get_inconsistent(ptr))
}
func (w Widget) LinkButtonSetUri(s string) {
	ptr := (*C.GtkLinkButton)(w.ptr)
	C.gtk_link_button_set_uri(ptr, StringToGChar(s))
}
func (w Widget) LinkButtonGetUri() string {
	ptr := (*C.GtkLinkButton)(w.ptr)
	return GCharToString(C.gtk_link_button_get_uri(ptr))
}
func (w Widget) LinkButtonSetVisited(b bool) {
	ptr := (*C.GtkLinkButton)(w.ptr)
	C.gtk_link_button_set_visited(ptr, BoolToGBoolean(b))
}
func (w Widget) LinkButtonGetVisited() bool {
	ptr := (*C.GtkLinkButton)(w.ptr)
	return GBooleanToBool(C.gtk_link_button_get_visited(ptr))
}
func (w Widget) ScaleButtonSetValue(f float64) {
	ptr := (*C.GtkScaleButton)(w.ptr)
	C.gtk_scale_button_set_value(ptr, FloatToGDouble(f))
}
func (w Widget) ScaleButtonGetValue() float64 {
	ptr := (*C.GtkScaleButton)(w.ptr)
	return GDoubleToFloat(C.gtk_scale_button_get_value(ptr))
}
func (w Widget) SwitchGetActive() bool {
	ptr := (*C.GtkSwitch)(w.ptr)
	return GBooleanToBool(C.gtk_switch_get_active(ptr))
}
func (w Widget) SwitchSetActive(b bool) {
	ptr := (*C.GtkSwitch)(w.ptr)
	C.gtk_switch_set_active(ptr, BoolToGBoolean(b))
}
func (w Widget) EntryProgressPulse() {
	ptr := (*C.GtkEntry)(w.ptr)
	C.gtk_entry_progress_pulse(ptr)
}
func (w Widget) EntrySetText(s string) {
	ptr := (*C.GtkEntry)(w.ptr)
	C.gtk_entry_set_text(ptr, StringToGChar(s))
}
func (w Widget) EntryGetText() string {
	ptr := (*C.GtkEntry)(w.ptr)
	return GCharToString(C.gtk_entry_get_text(ptr))
}
func (w Widget) EntrySetVisibility(b bool) {
	ptr := (*C.GtkEntry)(w.ptr)
	C.gtk_entry_set_visibility(ptr, BoolToGBoolean(b))
}
func (w Widget) EntryGetVisibility() bool {
	ptr := (*C.GtkEntry)(w.ptr)
	return GBooleanToBool(C.gtk_entry_get_visibility(ptr))
}
func (w Widget) EntrySetProgressFraction(f float64) {
	ptr := (*C.GtkEntry)(w.ptr)
	C.gtk_entry_set_progress_fraction(ptr, FloatToGDouble(f))
}
func (w Widget) EntrySetProgressPulseStep(f float64) {
	ptr := (*C.GtkEntry)(w.ptr)
	C.gtk_entry_set_progress_pulse_step(ptr, FloatToGDouble(f))
}
func (w Widget) EntryGetProgressFraction() float64 {
	ptr := (*C.GtkEntry)(w.ptr)
	return GDoubleToFloat(C.gtk_entry_get_progress_fraction(ptr))
}
func (w Widget) EntryGetProgressPulseStep() float64 {
	ptr := (*C.GtkEntry)(w.ptr)
	return GDoubleToFloat(C.gtk_entry_get_progress_pulse_step(ptr))
}
func (w Widget) RangeGetValue() float64 {
	ptr := (*C.GtkRange)(w.ptr)
	return GDoubleToFloat(C.gtk_range_get_value(ptr))
}
func (w Widget) RangeSetValue(f float64) {
	ptr := (*C.GtkRange)(w.ptr)
	C.gtk_range_set_value(ptr, FloatToGDouble(f))
}
func (w Widget) RangeSetIncrements(f1, f2 float64) {
	ptr := (*C.GtkRange)(w.ptr)
	C.gtk_range_set_increments(ptr, FloatToGDouble(f1), FloatToGDouble(f2))
}
func (w Widget) RangeSetRange(f1, f2 float64) {
	ptr := (*C.GtkRange)(w.ptr)
	C.gtk_range_set_range(ptr, FloatToGDouble(f1), FloatToGDouble(f2))
}
func (w Widget) SpinButtonGetValue() float64 {
	ptr := (*C.GtkSpinButton)(w.ptr)
	return GDoubleToFloat(C.gtk_spin_button_get_value(ptr))
}
func (w Widget) SpinButtonSetValue(f float64) {
	ptr := (*C.GtkSpinButton)(w.ptr)
	C.gtk_spin_button_set_value(ptr, FloatToGDouble(f))
}
func (w Widget) SpinButtonGetWrap() bool {
	ptr := (*C.GtkSpinButton)(w.ptr)
	return GBooleanToBool(C.gtk_spin_button_get_wrap(ptr))
}
func (w Widget) SpinButtonSetWrap(b bool) {
	ptr := (*C.GtkSpinButton)(w.ptr)
	C.gtk_spin_button_set_wrap(ptr, BoolToGBoolean(b))
}
func (w Widget) SpinButtonSetIncrements(f1, f2 float64) {
	ptr := (*C.GtkSpinButton)(w.ptr)
	C.gtk_spin_button_set_increments(ptr, FloatToGDouble(f1), FloatToGDouble(f2))
}
func (w Widget) SpinButtonSetRange(f1, f2 float64) {
	ptr := (*C.GtkSpinButton)(w.ptr)
	C.gtk_spin_button_set_range(ptr, FloatToGDouble(f1), FloatToGDouble(f2))
}
func (w Widget) ComboBoxTextRemoveAll() {
	ptr := (*C.GtkComboBoxText)(w.ptr)
	C.gtk_combo_box_text_remove_all(ptr)
}
func (w Widget) ComboBoxTextAppend(s1, s2 string) {
	ptr := (*C.GtkComboBoxText)(w.ptr)
	C.gtk_combo_box_text_append(ptr, StringToGChar(s1), StringToGChar(s2))
}
func (w Widget) ComboBoxTextPrepend(s1, s2 string) {
	ptr := (*C.GtkComboBoxText)(w.ptr)
	C.gtk_combo_box_text_prepend(ptr, StringToGChar(s1), StringToGChar(s2))
}
func (w Widget) ComboBoxTextAppendText(s string) {
	ptr := (*C.GtkComboBoxText)(w.ptr)
	C.gtk_combo_box_text_append_text(ptr, StringToGChar(s))
}
func (w Widget) ComboBoxTextPrependText(s string) {
	ptr := (*C.GtkComboBoxText)(w.ptr)
	C.gtk_combo_box_text_prepend_text(ptr, StringToGChar(s))
}
func (w Widget) ComboBoxTextGetActiveText() string {
	ptr := (*C.GtkComboBoxText)(w.ptr)
	return GCharToString(C.gtk_combo_box_text_get_active_text(ptr))
}
func (w Widget) MenuItemGetLabel() string {
	ptr := (*C.GtkMenuItem)(w.ptr)
	return GCharToString(C.gtk_menu_item_get_label(ptr))
}
func (w Widget) MenuItemSetLabel(s string) {
	ptr := (*C.GtkMenuItem)(w.ptr)
	C.gtk_menu_item_set_label(ptr, StringToGChar(s))
}
func (w Widget) CheckMenuItemGetActive() bool {
	ptr := (*C.GtkCheckMenuItem)(w.ptr)
	return GBooleanToBool(C.gtk_check_menu_item_get_active(ptr))
}
func (w Widget) CheckMenuItemGetInconsistent() bool {
	ptr := (*C.GtkCheckMenuItem)(w.ptr)
	return GBooleanToBool(C.gtk_check_menu_item_get_inconsistent(ptr))
}
func (w Widget) CheckMenuItemSetActive(b bool) {
	ptr := (*C.GtkCheckMenuItem)(w.ptr)
	C.gtk_check_menu_item_set_active(ptr, BoolToGBoolean(b))
}
func (w Widget) CheckMenuItemSetInconsistent(b bool) {
	ptr := (*C.GtkCheckMenuItem)(w.ptr)
	C.gtk_check_menu_item_set_inconsistent(ptr, BoolToGBoolean(b))
}
func (w Widget) ToolButtonSetLabel(s string) {
	ptr := (*C.GtkToolButton)(w.ptr)
	C.gtk_tool_button_set_label(ptr, StringToGChar(s))
}
func (w Widget) ToolButtonGetLabel() string {
	ptr := (*C.GtkToolButton)(w.ptr)
	return GCharToString(C.gtk_tool_button_get_label(ptr))
}
func (w Widget) ToggleToolButtonGetActive() bool {
	ptr := (*C.GtkToggleToolButton)(w.ptr)
	return GBooleanToBool(C.gtk_toggle_tool_button_get_active(ptr))
}
func (w Widget) ToggleToolButtonSetActive(b bool) {
	ptr := (*C.GtkToggleToolButton)(w.ptr)
	C.gtk_toggle_tool_button_set_active(ptr, BoolToGBoolean(b))
}
