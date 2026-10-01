package mqtt

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/SkycareOnskyHealth/mqtt/locales"
	"github.com/SkycareOnskyHealth/rbac/model"
	proto "github.com/SkycareOnskyHealth/rbac/proto/calling"
	notifyProto "github.com/SkycareOnskyHealth/rbac/proto/simple-notification"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// TimeZone default time zone
const TimeZone = "Asia/Ho_Chi_Minh"

// FindTimeZone find timezone of a thing
func FindTimeZone(properties []*proto.Property, defaultTimezone string) string {
	if defaultTimezone != "" {
		defaultTimezone = TimeZone
	}
	for _, n := range properties {
		if n.Name == "timezone" && n.Value != "" {
			return n.Value
		}
	}
	return defaultTimezone
}

// GetMessageKey get message key for notification simple
func GetMessageKey(notificationType model.NotificationType, value int) string {
	switch notificationType {
	case model.DoorSensor:
		if value == 1 {
			return model.DoorOpen
		}
		return model.DoorClose
	case model.SecurityBreach:
		return model.SecurityBreachMessage
	case model.SafetyBreachCO:
		return model.SafetyBreachMessageCO
	case model.SafetyBreachSOS:
		return model.SafetyBreachMessageSOS
	case model.SafetyBreachSmoke:
		return model.SafetyBreachMessageSmoke
	case model.SafetyBreachTempHumd:
		return model.SafetyBreachMessageTempHumd
	case model.SecurityAlarm:
		if value == 2 {
			return model.SecurityAlarmAway
		}
		if value == 1 {
			return model.SecurityAlarmHome
		}
		return model.SecurityAlarmOff
	case model.SafetyAlarm:
		if value == 0 {
			return model.SafetyAlarmEnable
		}
		return model.SafetyAlarmDisable
	case model.LowBattery:
		if value == 1 {
			return model.LowBatteryMessage
		}
		return model.FullBatteryMessage
	case model.Vibration:
		return model.VibrationMessage
	case model.MotionSensor:
		return model.MotionDetect
	default:
		return model.MotionDetect
	}
}

// ConvertUTCToLocalTime convert utc to local time
func ConvertUTCToLocalTime(dateTime time.Time, timezone string) string {
	//init the loc
	location, err := time.LoadLocation(timezone)
	//set timezone,
	if err != nil || location == nil {
		return dateTime.String()
	}
	return dateTime.In(location).String()
}

// MakeDataResponse make data response for simple notification
func MakeDataResponse(notifications []model.Notification) string {
	var data []*notifyProto.NotificationResult

	for _, result := range notifications {
		messageEN := fmt.Sprintf("%s motion detected", result.ThingDisplayName)
		messageVN := fmt.Sprintf("%s phát hiện chuyển động", result.ThingDisplayName)
		if result.Localizes != nil {
			if len(result.Localizes) > 0 {
				messageEN = result.Localizes[0].Message
				if len(result.Localizes) > 1 {
					messageVN = result.Localizes[1].Message
				}
			}
		}
		status := int32(result.Status)
		if result.Status == model.Initial {
			status = 0
		}
		notification := &notifyProto.NotificationResult{
			Id:                 int32(result.ID),
			ThingName:          result.ThingName,
			ThingDisplayName:   result.ThingDisplayName,
			ThingSerial:        result.ThingSerial,
			GatewayName:        result.GatewayName,
			GatewayDisplayName: result.GatewayDisplayName,
			GatewayMacAddress:  result.GatewayMacAddress,
			ZoneName:           result.ZoneName,
			ZoneDisplayName:    result.ZoneDisplayName,
			Status:             status,
			Type:               int32(result.Type),
			Template:           result.Template,
			State:              int32(result.State),
			Description:        messageEN,
			DescriptionVN:      messageVN,
			DateTime:           ConvertUTCToLocalTime(result.CreatedAt, result.Timezone),
			AlertType:          1,
			Acknowledged:       0,
			AlertStatus:        0,
			DeviceId:           int32(result.DeviceID),
		}
		data = append(data, notification)
	}
	b, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(b)
}

// ParseTopic parse mqtt topic
func ParseTopic(prefix string, topic string) (string, error) {
	s1 := strings.Split(topic, "/")
	if s1 == nil || len(s1) < 2 {
		return "", errors.New("invalid:topic")
	}
	if s1[0] != prefix {
		return "", errors.New("invalid:prefix")
	}
	return s1[1], nil
}

// FindValue find a value from list property
func FindValue(properties []*proto.Property, name string, defaultValue string) string {
	if properties == nil || len(properties) == 0 {
		return defaultValue
	}
	for _, n := range properties {
		if n.Name == name && n.Value != "" {
			return n.Value
		}
	}

	return defaultValue
}

// FindModeValue detect mode value in security and safety
func FindModeValue(array []*proto.Property, name string) int {
	for _, n := range array {
		if name == n.Name {
			return ConvertStringToInt(n.Value)
		}
	}
	return 0
}

// GetTimeZone get time zone from locale string
func GetTimeZone(locale string) string {
	bundle := locales.GetBundle(locale)
	if bundle != nil && bundle.DefaultTimezone != "" {
		return bundle.DefaultTimezone
	}
	return TimeZone
}

// IsMn is mn rune
func IsMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// CheckTemplateType detect resource type depend on template
func CheckTemplateType(name string) model.SecurityType {
	switch name {
	case "CO Detector":
		return model.Co
	case "Smoke Detector":
		return model.Smoke
	case "SOS Button":
		return model.SOS
	case "Temperature-Humidity Sensor":
		return model.TempHumd
	case "Zigbee Door Lock":
		return model.DoorLock
	case "OS Locus":
		return model.OSLocus
	case "Bed Sensor":
		return model.BedSensor
	case "Sky Band":
		return model.SkyBand
	case "SkySOS":
		return model.SkySOS
	default:
		return model.Motion
	}
}

// ConvertStringToInt convert string to int
func ConvertStringToInt(text string) int {
	value, err := strconv.Atoi(text)
	if err != nil {
		return 0
	}
	return value
}

// GetNotificationType get notification type
func GetNotificationType(securityType model.SecurityType) model.NotificationType {
	switch securityType {
	case model.Co:
		return model.SafetyBreachCO
	case model.Smoke:
		return model.SafetyBreachSmoke
	case model.TempHumd:
		return model.SafetyBreachTempHumd
	case model.SOS:
		return model.SafetyBreachSOS
	case model.DoorLock:
		return model.SecurityBreach
	case model.OSLocus:
		return model.SafetyBreachSOS
	case model.Motion:
		return model.SecurityBreach
	default:
		return model.SecurityBreach
	}
}

// CheckSecurityState detect mode is security or safety
func CheckSecurityState(templateType model.SecurityType) model.Mode {
	switch templateType {
	case model.Co:
		return model.Safety
	case model.Smoke:
		return model.Safety
	case model.SOS:
		return model.Safety
	case model.TempHumd:
		return model.Safety
	case model.OSLocus:
		return model.Basic
	case model.DoorLock:
		return model.Security
	case model.BedSensor:
		return model.Safety
	case model.SkyBand:
		return model.Safety
	case model.SkySOS:
		return model.Safety
	default:
		return model.Security
	}
}

// ConvertMode convert string to mode
func ConvertMode(mode string) model.Mode {
	switch mode {
	case "safe":
		return model.Safety
	case "safety":
		return model.Safety
	case "security":
		return model.Security
	default:
		return model.Security
	}
}

//-------------PREPARE BODY HELPER--------------------

// PrepareResourceLocale prepare text message for any locale
func PrepareResourceLocale(templateType model.NotificationType, key string, locale string, gatewayName string, deviceName string, zoneName string, date string, timestamp string, fullName string) string {
	bundle := locales.GetBundle(locale)
	if bundle == nil {
		return ""
	}

	switch key {
	case "onsky_security":
		if h, ok := bundle.ServiceHeaders[templateType]; ok {
			return h
		}
		return bundle.DefaultHeader
	case "security_alert":
		if a, ok := bundle.Alerts[templateType]; ok {
			return a
		}
		return bundle.DefaultAlert
	case "please_check":
		if act, ok := bundle.Actions[templateType]; ok {
			return act
		}
		return bundle.DefaultAction
	case "zone":
		return bundle.Labels.ZoneLabel
	case "phone":
		return bundle.Labels.PhoneLabel
	case "gateway_name":
		return gatewayName
	case "zone_name":
		return zoneName
	case "device":
		return bundle.Labels.DeviceLabel
	case "device_name":
		return deviceName
	case "on_date":
		return bundle.Labels.OnDateLabel
	case "date":
		return date
	case "of":
		return bundle.Labels.OfLabel
	case "at":
		return bundle.Labels.AtLabel
	case "full_name":
		return fullName + " "
	case "at_time":
		return bundle.Labels.AtTimeLabel
	case "time":
		return timestamp
	default:
		return ""
	}
}

// PrepareBody prepare body of a message
func PrepareBody(templateType model.NotificationType, locale string, gatewayName string, deviceName string, zoneName string, timezone string, address string, phoneNumber string, fullName string) string {
	now := time.Now()
	if timezone == "" {
		timezone = GetTimeZone(locale)
	}
	//init the loc
	loc, err := time.LoadLocation(timezone)
	//set timezone,
	if err == nil {
		now = now.In(loc)
	} else {
		log.Println("error loc", err)
	}
	date := strconv.Itoa(now.Day()) + "/" + strconv.Itoa(int(now.Month()))
	timestamp := strconv.Itoa(now.Hour()) + ":" + strconv.Itoa(now.Minute())
	//  OnSky Security & Safety service. Emergency SOS sent from location: <Name>, <address>, <Phone>, device name and zone name. Date, Time. Check Now!
	// OnSky Medical Alert Service. Possible Apnea Alert from SkyPad device of [user name] at [address] on [date] at [time]. Phone: [0123456789]. Please check.
	template := "{{onsky_security}}. {{security_alert}}{{of}}{{full_name}}{{at}} {{address}}. {{phone}}:{{phone_number}}, {{device}} {{device_name}}. {{on_date}} {{date}}, {{at_time}} {{time}}. {{please_check}}."
	str := strings.Replace(template, "{{onsky_security}}", PrepareResourceLocale(templateType, "onsky_security", locale, gatewayName, deviceName, zoneName, "", "", fullName), -1)
	// str = strings.Replace(str, "{{zone}}", PrepareResourceLocale(templateType, "zone", locale, gatewayName, deviceName, zoneName, "", ""), -1)
	// str = strings.Replace(str, "{{gateway_name}}", PrepareResourceLocale(templateType, "gateway_name", locale, gatewayName, deviceName, zoneName, "", ""), -1)
	// str = strings.Replace(str, "{{zone_name}}", PrepareResourceLocale(templateType, "zone_name", locale, gatewayName, deviceName, zoneName, "", ""), -1)
	str = strings.Replace(str, "{{address}}", address, -1)
	str = strings.Replace(str, "{{phone}}", PrepareResourceLocale(templateType, "phone", locale, gatewayName, deviceName, zoneName, "", "", fullName), -1)
	str = strings.Replace(str, "{{phone_number}}", phoneNumber, -1)
	str = strings.Replace(str, "{{device}}", PrepareResourceLocale(templateType, "device", locale, gatewayName, deviceName, zoneName, "", "", fullName), -1)
	str = strings.Replace(str, "{{device_name}}", PrepareResourceLocale(templateType, "device_name", locale, gatewayName, deviceName, zoneName, "", "", fullName), -1)
	str = strings.Replace(str, "{{on_date}}", PrepareResourceLocale(templateType, "on_date", locale, gatewayName, deviceName, zoneName, "", "", fullName), -1)
	str = strings.Replace(str, "{{date}}", PrepareResourceLocale(templateType, "date", locale, gatewayName, deviceName, zoneName, date, "", fullName), -1)
	str = strings.Replace(str, "{{at_time}}", PrepareResourceLocale(templateType, "at_time", locale, gatewayName, deviceName, zoneName, "", "", fullName), -1)
	str = strings.Replace(str, "{{time}}", PrepareResourceLocale(templateType, "time", locale, gatewayName, deviceName, zoneName, "", timestamp, fullName), -1)
	str = strings.Replace(str, "{{please_check}}", PrepareResourceLocale(templateType, "please_check", locale, gatewayName, deviceName, zoneName, "", "", fullName), -1)
	str = strings.Replace(str, "{{security_alert}}", PrepareResourceLocale(templateType, "security_alert", locale, gatewayName, deviceName, zoneName, "", "", fullName), -1)
	str = strings.Replace(str, "{{at}}", PrepareResourceLocale(templateType, "at", locale, gatewayName, deviceName, zoneName, "", "", fullName), -1)

	if fullName != "" {
		str = strings.Replace(str, "{{of}}", PrepareResourceLocale(templateType, "of", locale, gatewayName, deviceName, zoneName, "", "", fullName), -1)
		str = strings.Replace(str, "{{full_name}}", PrepareResourceLocale(templateType, "full_name", locale, gatewayName, deviceName, zoneName, "", "", fullName), -1)
	} else {
		str = strings.Replace(str, "{{of}}", " ", -1)
		str = strings.Replace(str, "{{full_name}}", "", -1)
	}
	return str
}

// PrepareMedia prepare a media url
func PrepareMedia(templateType model.NotificationType, locale string) string {
	callType := strings.ToLower(strings.Replace(templateType.String(), " ", "-", -1))
	return fmt.Sprintf("?safety=%s&locale=%s", callType, locale)
}

// Connect to the broker
func Connect(url, clientID, username, password string) (MQTT.Client, error) {
	options := MQTT.NewClientOptions().SetClientID(clientID).AddBroker(url) //! Use for demo
	options.SetUsername(username).SetPassword(password)                     //! Use for demo

	c := MQTT.NewClient(options)
	token := c.Connect()
	if token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}
	return c, nil
}

// Publish to the broker
func Publish(client MQTT.Client, topic, value string) error {
	if token := client.Publish(topic, 1, true, value); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Disconnect to the broker
func Disconnect(client MQTT.Client) {
	client.Disconnect(100)
}
