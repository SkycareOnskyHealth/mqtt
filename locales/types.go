package locales

import "github.com/SkycareOnskyHealth/rbac/model"

// Labels chứa các nhãn cố định trong message template
type Labels struct {
	PhoneLabel  string // en: "phone", vi: "SDT"
	DeviceLabel string // en: "device", vi: "thiet bi"
	OnDateLabel string // en: "on", vi: "Vao ngay"
	AtTimeLabel string // en: "at", vi: "luc"
	OfLabel     string // en: " of ", vi: " cua "
	AtLabel     string // en: "at", vi: "tai"
	ZoneLabel   string // en: "zone", vi: "khu"
}

// Bundle chứa toàn bộ từ điển và cấu hình của một ngôn ngữ
type Bundle struct {
	Code            string                            // Mã chuẩn BCP 47 (ví dụ "en-US", "vi-VN")
	DefaultTimezone string                            // Timezone mặc định của khu vực
	Labels          Labels                            // Nhãn từ ngữ
	ServiceHeaders  map[model.NotificationType]string // Header theo loại cảnh báo
	Alerts          map[model.NotificationType]string // Chuỗi mô tả cảnh báo
	Actions         map[model.NotificationType]string // Lời kêu gọi hành động ("Please check" / "Check Now!")
	DefaultHeader   string                            // Fallback header (ví dụ "OnSky Security & Safety service")
	DefaultAlert    string                            // Fallback alert (ví dụ "Intruder detected in")
	DefaultAction   string                            // Fallback action (ví dụ "Check Now!")
}

