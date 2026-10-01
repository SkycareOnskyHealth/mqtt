package locales

import "github.com/SkycareOnskyHealth/rbac/model"

// Labels contains fixed labels used in message templates
type Labels struct {
	PhoneLabel  string // e.g. "phone", "SDT"
	DeviceLabel string // e.g. "device", "thiet bi"
	OnDateLabel string // e.g. "on", "Vao ngay"
	AtTimeLabel string // e.g. "at", "luc"
	OfLabel     string // e.g. " of ", " cua "
	AtLabel     string // e.g. "at", "tai"
	ZoneLabel   string // e.g. "zone", "khu"
}

// Bundle contains dictionary and localization configuration for a specific language
type Bundle struct {
	Code                string                            // BCP 47 language tag (e.g. "en-US", "vi-VN")
	DefaultTimezone     string                            // Default timezone for the locale/region
	Labels              Labels                            // Common word labels
	ServiceHeaders      map[model.NotificationType]string // Service header mapping per notification type
	Alerts              map[model.NotificationType]string // Alert description strings per notification type
	Actions             map[model.NotificationType]string // Call-to-action text ("Please check" / "Check Now!")
	DefaultHeader       string                            // Fallback header (e.g. "OnSky Security & Safety service")
	DefaultAlert        string                            // Fallback alert description (e.g. "Intruder detected in")
	DefaultAction       string                            // Fallback action text (e.g. "Check Now!")
	Notifications       map[model.NotificationType]string // Short push / in-app notification messages
	DefaultNotification string                            // Fallback notification message
}
