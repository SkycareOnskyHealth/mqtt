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
		{model.SkySOSButtonTriggered, "SkySOS-Pendant"},
		{model.SkySOSFallDetection, "SkySOS-Pendant"},
		{model.SkyBandSpo2Low, "SkyBand-Smart"},
		{model.SafetyBreachCO, "CO-Detector"},
	}

	for _, loc := range locales {
		for _, dev := range devices {
			// Test with fullName
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

			// Test without fullName
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

func TestPrepareNotificationMessage(t *testing.T) {
	// 1. Verify exact string matches for en-US and vi-VN
	enSOS := mqtt.PrepareNotificationMessage(model.BedSensorSOS, "en-US")
	expectedEnSOS := "Detect SOS alert from OnSky device. Please check."
	if enSOS != expectedEnSOS {
		t.Errorf("PrepareNotificationMessage(BedSensorSOS, en-US) = %q; expected %q", enSOS, expectedEnSOS)
	}

	vnSOS := mqtt.PrepareNotificationMessage(model.BedSensorSOS, "vi-VN")
	expectedVnSOS := "Phát hiện cảnh báo SOS từ thiết bị OnSky. Vui lòng kiểm tra."
	if vnSOS != expectedVnSOS {
		t.Errorf("PrepareNotificationMessage(BedSensorSOS, vi-VN) = %q; expected %q", vnSOS, expectedVnSOS)
	}

	enMoving := mqtt.PrepareNotificationMessage(model.SkySOSDeviceMoving, "en-US")
	expectedEnMoving := "Detect user moving by OnSky Alert Pendant."
	if enMoving != expectedEnMoving {
		t.Errorf("PrepareNotificationMessage(SkySOSDeviceMoving, en-US) = %q; expected %q", enMoving, expectedEnMoving)
	}

	vnMoving := mqtt.PrepareNotificationMessage(model.SkySOSDeviceMoving, "vi-VN")
	expectedVnMoving := "Phát hiện người dùng đang di chuyển từ OnSky Alert Pendant."
	if vnMoving != expectedVnMoving {
		t.Errorf("PrepareNotificationMessage(SkySOSDeviceMoving, vi-VN) = %q; expected %q", vnMoving, expectedVnMoving)
	}

	// 2. Verify all 12 languages return non-empty messages for all alert types
	locales := []string{
		"en-US", "vi-VN", "zh-Hans", "zh-Hant",
		"fr-FR", "de-DE", "id-ID", "ja-JP",
		"ko-KR", "es-ES", "th-TH", "fil-PH",
	}
	alertTypes := []model.NotificationType{
		model.BedSensorSOS,
		model.BedSensorHeartStop,
		model.BedSensorBreathStop,
		model.BedSensorTachycardia,
		model.BedSensorBradycardia,
		model.BedSensorSeizure,
		model.BedSensorBodyTempHeight,
		model.BedSensorRoomTempHeight,
		model.BedSensorHumidityHeight,
		model.BedSensorHeartRateHeight,
		model.BedSensorHeartRateLow,
		model.BedSensorBedLeaving,
		model.BedSensorCrying,
		model.SkySOSButtonTriggered,
		model.SkySOSFallDetection,
		model.SkySOSGeofenceEnter,
		model.SkySOSGeofenceExit,
		model.SkySOSDeviceMoving,
		model.SkySOSDeviceStopped,
		model.SkyBandSpo2Low,
		model.SkyBandHeartRateLow,
		model.SkyBandHeartRateHeight,
	}

	for _, loc := range locales {
		for _, at := range alertTypes {
			msg := mqtt.PrepareNotificationMessage(at, loc)
			if msg == "" {
				t.Fatalf("PrepareNotificationMessage(%v, %q) returned empty string", at, loc)
			}
		}
	}

	// 3. Verify fallback when alert type does not exist in map (returns locale's DefaultNotification)
	unknownAlertEn := mqtt.PrepareNotificationMessage(model.NotificationType(120), "en-US")
	if unknownAlertEn != "Detect default alert from OnSky device. Please check." {
		t.Errorf("Expected fallback DefaultNotification for en-US, got: %q", unknownAlertEn)
	}

	unknownAlertVn := mqtt.PrepareNotificationMessage(model.NotificationType(120), "vi-VN")
	if unknownAlertVn != "Phát hiện cảnh báo mặc định từ thiết bị OnSky. Vui lòng kiểm tra." {
		t.Errorf("Expected fallback DefaultNotification for vi-VN, got: %q", unknownAlertVn)
	}

	// 4. Verify fallback when non-existent locale is provided (falls back to en-US)
	fallbackLocaleMsg := mqtt.PrepareNotificationMessage(model.BedSensorSOS, "non-existent-locale")
	if fallbackLocaleMsg != expectedEnSOS {
		t.Errorf("Expected fallback to en-US for invalid locale, got: %q", fallbackLocaleMsg)
	}

	// 5. Verify fallback when empty locale string is provided (falls back to en-US)
	emptyLocaleMsg := mqtt.PrepareNotificationMessage(model.BedSensorSOS, "")
	if emptyLocaleMsg != expectedEnSOS {
		t.Errorf("Expected fallback to en-US for empty locale, got: %q", emptyLocaleMsg)
	}

	// 6. Verify locale normalization is case-insensitive and handles underscores (e.g. VI_vn -> vi-VN)
	normalizedMsg := mqtt.PrepareNotificationMessage(model.BedSensorSOS, "VI_vn")
	if normalizedMsg != expectedVnSOS {
		t.Errorf("Expected normalization VI_vn to return vi-VN message, got: %q", normalizedMsg)
	}
}
