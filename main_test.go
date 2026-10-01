package mqtt_test

import (
	"strings"
	"testing"

	"github.com/SkycareOnskyHealth/mqtt"
	"github.com/SkycareOnskyHealth/rbac/model"
)

func TestPrepareBody(t *testing.T) {
	locales := []string{
		"en-US",
		"vi-VN",
		"zh-Hans",
		"zh-Hant",
		"fr-FR",
		"de-DE",
		"id-ID",
		"ja-JP",
		"ko-KR",
		"es-ES",
		"th-TH",
		"fil-PH",
	}

	devices := []struct {
		alertType  model.NotificationType
		deviceName string
	}{
		{model.BedSensorSOS, "BedSensor-01"},
		{model.BedSensorBreathStop, "SkyPad-02"},
		{model.SkySOSButtonTriggered, "SkySOS-Pendent"},
		{model.SkySOSFallDetection, "SkySOS-Pendent"},
		{model.SkyBandSpo2Low, "SkyBand-Smart"},
		{model.SafetyBreachCO, "CO-Detector"},
	}

	for _, loc := range locales {
		for _, dev := range devices {
			// Test có fullName
			msgWithUser := mqtt.PrepareBody(
				dev.alertType,
				loc,
				"Gateway-1",
				dev.deviceName,
				"LivingRoom",
				"",
				"123 Ho Chi Minh City",
				"0901234567",
				"John Doe",
			)
			if msgWithUser == "" {
				t.Fatalf("PrepareBody returned empty string for locale %s, alert %v", loc, dev.alertType)
			}
			if !strings.Contains(msgWithUser, "John Doe") {
				t.Errorf("PrepareBody for %s should contain user name 'John Doe', got: %s", loc, msgWithUser)
			}
			if !strings.Contains(msgWithUser, "123 Ho Chi Minh City") {
				t.Errorf("PrepareBody for %s should contain address, got: %s", loc, msgWithUser)
			}
			if !strings.Contains(msgWithUser, dev.deviceName) {
				t.Errorf("PrepareBody for %s should contain device name '%s', got: %s", loc, dev.deviceName, msgWithUser)
			}

			// Test không có fullName
			msgWithoutUser := mqtt.PrepareBody(
				dev.alertType,
				loc,
				"Gateway-1",
				dev.deviceName,
				"LivingRoom",
				"",
				"123 Ho Chi Minh City",
				"0901234567",
				"",
			)
			if msgWithoutUser == "" {
				t.Fatalf("PrepareBody returned empty string for empty user in locale %s", loc)
			}
			if strings.Contains(msgWithoutUser, "John Doe") {
				t.Errorf("PrepareBody without user should not contain 'John Doe', got: %s", msgWithoutUser)
			}
		}
	}
}

func TestGetTimeZone(t *testing.T) {
	cases := []struct {
		locale   string
		expected string
	}{
		{"en-US", "America/Mazatlan"},
		{"vi-VN", "Asia/Ho_Chi_Minh"},
		{"ja-JP", "Asia/Tokyo"},
		{"ko-KR", "Asia/Seoul"},
		{"zh-Hans", "Asia/Shanghai"},
		{"fr-FR", "Europe/Paris"},
		{"de-DE", "Europe/Berlin"},
		{"", "America/Mazatlan"}, // fallback en-US timezone
	}

	for _, tc := range cases {
		tz := mqtt.GetTimeZone(tc.locale)
		if tz != tc.expected {
			t.Errorf("GetTimeZone(%q) = %q; expected %q", tc.locale, tz, tc.expected)
		}
	}
}

