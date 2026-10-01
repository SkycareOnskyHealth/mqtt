package locales_test

import (
	"testing"

	"github.com/SkycareOnskyHealth/mqtt/locales"
	"github.com/SkycareOnskyHealth/rbac/model"
)

var expectedLocales = []string{
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

// requiredAlerts kiểm tra tất cả các cảnh báo được định nghĩa trong Alerts map
var requiredAlerts = []model.NotificationType{
	model.BedSensorSOS,
	model.BedSensorHeartStop,
	model.BedSensorBreathStop,
	model.BedSensorTachycardia,
	model.BedSensorBradycardia,
	model.BedSensorSeizure,
	model.BedSensorAbnormalVitalSigns,
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
	model.SkyBandSpo2Low,
	model.SkyBandHeartRateLow,
	model.SkyBandHeartRateHeight,
	model.SafetyBreachCO,
	model.SafetyBreachSmoke,
	model.SafetyBreachSOS,
	model.SafetyBreachTempHumd,
	model.OSLocusSOS,
	model.OSLocusTemp,
	model.LowBattery,
}

func TestAllLocalesRegistered(t *testing.T) {
	for _, code := range expectedLocales {
		b := locales.GetBundle(code)
		if b == nil {
			t.Fatalf("Locale %s not registered", code)
		}
		if b.Code != code {
			t.Errorf("Expected bundle code %s, got %s", code, b.Code)
		}
		if b.DefaultTimezone == "" {
			t.Errorf("Locale %s has empty DefaultTimezone", code)
		}
		if b.DefaultHeader == "" {
			t.Errorf("Locale %s has empty DefaultHeader", code)
		}
		if b.DefaultAlert == "" {
			t.Errorf("Locale %s has empty DefaultAlert", code)
		}
		if b.DefaultAction == "" {
			t.Errorf("Locale %s has empty DefaultAction", code)
		}
		if b.Labels.PhoneLabel == "" || b.Labels.DeviceLabel == "" || b.Labels.OnDateLabel == "" || b.Labels.AtTimeLabel == "" {
			t.Errorf("Locale %s has empty labels", code)
		}

		// Kiểm tra toàn bộ các alert bắt buộc
		for _, alertType := range requiredAlerts {
			if _, ok := b.Alerts[alertType]; !ok {
				t.Errorf("Locale %s is missing alert translation for %v", code, alertType)
			}
		}

		// Kiểm tra toàn bộ các alert của BedSensor, SkySOS, SkyBand phải có Action và ServiceHeader riêng
		medicalAndAlertTypes := []model.NotificationType{
			model.BedSensorSOS,
			model.BedSensorHeartStop,
			model.BedSensorBreathStop,
			model.BedSensorTachycardia,
			model.BedSensorBradycardia,
			model.BedSensorSeizure,
			model.BedSensorAbnormalVitalSigns,
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
			model.SkyBandSpo2Low,
			model.SkyBandHeartRateLow,
			model.SkyBandHeartRateHeight,
		}

		for _, alertType := range medicalAndAlertTypes {
			if _, ok := b.Actions[alertType]; !ok {
				t.Errorf("Locale %s is missing action translation for %v", code, alertType)
			}
			if _, ok := b.ServiceHeaders[alertType]; !ok {
				t.Errorf("Locale %s is missing service header for %v", code, alertType)
			}
		}
	}
}

func TestGetBundleFallback(t *testing.T) {
	// Locale không tồn tại phải fallback về en-US
	fallback := locales.GetBundle("non-existent-locale")
	if fallback == nil || fallback.Code != "en-US" {
		t.Fatalf("Expected fallback to en-US, got %v", fallback)
	}

	// Chuỗi rỗng phải fallback về en-US
	emptyFallback := locales.GetBundle("")
	if emptyFallback == nil || emptyFallback.Code != "en-US" {
		t.Fatalf("Expected fallback to en-US for empty locale, got %v", emptyFallback)
	}

	// Case insensitive và thay thế _ bằng -
	jp := locales.GetBundle("JA_jp")
	if jp == nil || jp.Code != "ja-JP" {
		t.Errorf("Expected ja-JP, got %v", jp)
	}

	// Prefix fallback (vd "fr" -> "fr-FR")
	fr := locales.GetBundle("fr")
	if fr == nil || fr.Code != "fr-FR" {
		t.Errorf("Expected fr-FR, got %v", fr)
	}
}

