# Định Dạng Thông Báo Đẩy / Ứng Dụng (Notification Message Format)
### Thiết Bị BedSensor, SkySOS và SkyBand

Tài liệu này định nghĩa định dạng thông báo rút gọn chuyên biệt cho hệ thống **Notification (Push Notification / In-App Notification / SMS ngắn)**.

### Đặc điểm cấu trúc:
- **Ngắn gọn, súc tích:** Lược bỏ các trường rườm rà như địa chỉ, tên thiết bị, số điện thoại, ngày giờ để tối ưu độ dài hiển thị trên màn hình khóa điện thoại hoặc banner app.
- **Trực quan:** Nêu rõ sự kiện khẩn cấp được phát hiện và lời kêu gọi hành động kiểm tra tức thì.
- **Hỗ trợ đầy đủ 12 ngôn ngữ** chuẩn quốc tế theo BCP 47:

| Mã BCP 47 | Ngôn ngữ | Quốc gia / Vùng lãnh thổ |
| :--- | :--- | :--- |
| `en-US` | Tiếng Anh | Hoa Kỳ (Mỹ) |
| `vi-VN` | Tiếng Việt | Việt Nam |
| `zh-Hans` | Tiếng Trung (Giản thể) | Trung Quốc |
| `zh-Hant` | Tiếng Trung (Phồn thể) | Đài Loan / Hồng Kông |
| `fr-FR` | Tiếng Pháp | Pháp |
| `de-DE` | Tiếng Đức | Đức |
| `id-ID` | Tiếng Indonesia | Indonesia |
| `ja-JP` | Tiếng Nhật | Nhật Bản |
| `ko-KR` | Tiếng Hàn | Hàn Quốc |
| `es-ES` | Tiếng Tây Ban Nha | Tây Ban Nha |
| `th-TH` | Tiếng Thái | Thái Lan |
| `fil-PH` | Tiếng Filipino | Philippines |

---

## 1. Thiết Bị BedSensor (Cảm Biến Giường / Nệm)

### 1.1. BedSensorSOS (Khẩn cấp SOS)

- **Tiếng Anh (en-US):**
  > `Detect SOS alert from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo SOS từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的SOS紧急警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的SOS緊急警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une alerte SOS depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `SOS-Alarm vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi peringatan SOS dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスからSOSアラートを検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 SOS 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de alerta SOS desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือน SOS จากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng alertong SOS mula sa aparatong OnSky. Pakisuri.`

---

### 1.2. BedSensorHeartStop (Cơn Đau Tim / Ngưng Tim (Heart Attack))

- **Tiếng Anh (en-US):**
  > `Detect heart attack alert from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo cơn đau tim từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的心脏骤停/心脏病发作警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的心臟驟停/心臟病發作警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une alerte de crise cardiaque depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Herzinfarkt-Alarm vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi peringatan serangan jantung dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスから心臓発作アラートを検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 심장마비 의심 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de alerta de ataque cardíaco desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือนภาวะหัวใจวายจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng alerto ng atake sa puso mula sa aparatong OnSky. Pakisuri.`

---

### 1.3. BedSensorBreathStop (Ngưng Thở / Apnea)

- **Tiếng Anh (en-US):**
  > `Detect apnea alert from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo dấu hiệu ngưng thở từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的呼吸暂停警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的呼吸暫停警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une alerte d'apnée depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Schlafapnoe-Alarm vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi peringatan henti napas (apnea) dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスから無呼吸アラートを検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 무호흡 의심 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de alerta de apnea desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือนภาวะหยุดหายใจขณะหลับจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng alerto ng apnea (paghinto ng paghinga) mula sa aparatong OnSky. Pakisuri.`

---

### 1.4. BedSensorTachycardia (Nhịp Tim Nhanh (Tachycardia))

- **Tiếng Anh (en-US):**
  > `Detect irregular heart rhythms - tachycardia. Your heart rate is very fast. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo nhịp tim nhanh từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到心律不齐 - 心动过速。您的心率非常快。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到心律不齊 - 心動過速。您的心率非常快。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'un rythme cardiaque irrégulier - tachycardie. Votre fréquence cardiaque est très rapide. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Herzrhythmusstörungen - Tachykardie erkannt. Ihre Herzfrequenz ist sehr schnell. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi ritme jantung tidak teratur - takikardia. Detak jantung Anda sangat cepat. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `不整脈（頻脈）を検知しました。心拍数が非常に高くなっています。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `부정맥(빈맥)이 감지되었습니다. 심박수가 매우 빠릅니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de ritmos cardíacos irregulares - taquicardia. Su frecuencia cardíaca es muy rápida. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบจังหวะการเต้นของหัวใจผิดปกติ - หัวใจเต้นเร็วมาก กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng irregular na tibok ng puso - tachycardia. Mabilis ang tibok ng iyong puso. Pakisuri.`

---

### 1.5. BedSensorBradycardia (Nhịp Tim Chậm (Bradycardia))

- **Tiếng Anh (en-US):**
  > `Detect irregular heart rhythms - bradycardia. Your heart rate is very slow. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo nhịp tim chậm từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到心律不齐 - 心动过缓。您的心率非常慢。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到心律不齊 - 心動過緩。您的心率非常慢。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'un rythme cardiaque irrégulier - bradycardie. Votre fréquence cardiaque est très lente. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Herzrhythmusstörungen - Bradykardie erkannt. Ihre Herzfrequenz ist sehr langsam. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi ritme jantung tidak teratur - bradikardia. Detak jantung Anda sangat lambat. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `不整脈（徐脈）を検知しました。心拍数が非常に低くなっています。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `부정맥(서맥)이 감지되었습니다. 심박수가 매우 느립니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de ritmos cardíacos irregulares - bradicardia. Su frecuencia cardíaca es muy lenta. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบจังหวะการเต้นของหัวใจผิดปกติ - หัวใจเต้นช้ามาก กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng irregular na tibok ng puso - bradycardia. Mabagal ang tibok ng iyong puso. Pakisuri.`

---

### 1.6. BedSensorSeizure (Co Giật (Seizure))

- **Tiếng Anh (en-US):**
  > `Detect seizures alert from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo cơn co giật từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的癫痫/抽搐发作警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的癲癇/抽搐發作警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une alerte de convulsions depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Krampfanfall-Alarm vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi peringatan kejang dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスから痙攣発作アラートを検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 발작 의심 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de alerta de convulsiones desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือนอาการชักจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng alerto ng pangingisay mula sa aparatong OnSky. Pakisuri.`

---

### 1.7. BedSensorBodyTempHeight (Thân Nhiệt Quá Cao / Sốt Cao)

- **Tiếng Anh (en-US):**
  > `Detect body temperature is too high from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo sốt cao từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的体温过高/发烧警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的體溫過高/發燒警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une température corporelle trop élevée depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Zu hohe Körpertemperatur vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi suhu tubuh terlalu tinggi dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスから体温異常（高熱）を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 고열(체온 과다) 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de temperatura corporal demasiado alta desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือนไข้สูงจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng mataas na lagnat mula sa aparatong OnSky. Pakisuri.`

---

### 1.8. BedSensorRoomTempHeight (Nhiệt Độ Phòng Quá Cao)

- **Tiếng Anh (en-US):**
  > `Detect room temperature is too high from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo nhiệt độ phòng qúa cao từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的室内温度过高警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的室內溫度過高警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une température ambiante trop élevée depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Zu hohe Raumtemperatur vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi suhu ruangan terlalu tinggi dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスから室温異常（高温）を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 실내 온도 과다 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de temperatura ambiente demasiado alta desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือนอุณหภูมิห้องสูงเกินไปจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng sobrang taas na temperatura ng silid mula sa aparatong OnSky. Pakisuri.`

---

### 1.9. BedSensorHumidityHeight (Độ Ẩm Phòng Quá Cao)

- **Tiếng Anh (en-US):**
  > `Detect room humidity is too high from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo độ ẩm phòng quá cao từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的室内湿度过高警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的室內濕度過高警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une humidité ambiante trop élevée depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Zu hohe Luftfeuchtigkeit vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi kelembapan ruangan terlalu tinggi dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスから湿度異常（高湿度）を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 실내 습도 과다 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de humedad ambiente demasiado alta desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือนความชื้นในห้องสูงเกินไปจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng sobrang taas na halumigmig sa silid mula sa aparatong OnSky. Pakisuri.`

---

### 1.10. BedSensorHeartRateHeight (Nhịp Tim Quá Cao)

- **Tiếng Anh (en-US):**
  > `Detect heart rate is too high from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo nhịp tim quá cao từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的心率过高警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的心率過高警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une fréquence cardiaque trop élevée depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Zu hohe Herzfrequenz vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi detak jantung terlalu tinggi dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスから心拍数過多を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 심박수 과다 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de frecuencia cardíaca demasiado alta desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือนอัตราการเต้นของหัวใจสูงเกินไปจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng sobrang bilis na tibok ng puso mula sa aparatong OnSky. Pakisuri.`

---

### 1.11. BedSensorHeartRateLow (Nhịp Tim Quá Thấp)

- **Tiếng Anh (en-US):**
  > `Detect heart rate is too low from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo nhịp tim quá thấp từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的心率过低警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的心率過低警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une fréquence cardiaque trop basse depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Zu niedrige Herzfrequenz vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi detak jantung terlalu rendah dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスから心拍数過少を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 심박수 저하 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de frecuencia cardíaca demasiado baja desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือนอัตราการเต้นของหัวใจต่ำเกินไปจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng sobrang bagal na tibok ng puso mula sa aparatong OnSky. Pakisuri.`

---

### 1.12. BedSensorBedLeaving (Rời Khỏi Giường)

- **Tiếng Anh (en-US):**
  > `Detect user have left the bed from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện người dùng đã rời khỏi giường từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的用户离床警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的用戶離床警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection que l'utilisateur a quitté son lit depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Benutzer hat das Bett verlassen laut OnSky-Gerät. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi pengguna telah meninggalkan tempat tidur dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスが利用者の離床を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 사용자가 침대를 이탈했음을 감지했습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de usuario fuera de la cama desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบผู้ใช้ออกจากเตียงจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita na ang gumagamit ay umalis sa kama mula sa aparatong OnSky. Pakisuri.`

---

### 1.13. BedSensorCrying (Em Bé Khóc)

- **Tiếng Anh (en-US):**
  > `Detect baby crying from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện em bé đang khóc từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的婴儿啼哭警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的嬰兒啼哭警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection de pleurs de bébé depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Babyweinen vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi tangisan bayi dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスが赤ちゃんの泣き声を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 아기 울음소리가 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de llanto de bebé desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบเสียงเด็กร้องไห้จากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng pag-iyak ng sanggol mula sa aparatong OnSky. Pakisuri.`

---

## 2. Thiết Bị SkySOS (Dây Chuyền Khẩn Cấp / Alert Pendant)

### 2.1. SkySOSButtonTriggered (Bấm Nút SOS Khẩn Cấp)

- **Tiếng Anh (en-US):**
  > `Detect SOS Alert from OnSky Alert Pendant. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo SOS từ OnSky Alert Pendant. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky Alert Pendant的SOS警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky Alert Pendant的SOS警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une alerte SOS depuis l'OnSky Alert Pendant. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `SOS-Alarm von OnSky Alert Pendant erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi Peringatan SOS dari OnSky Alert Pendant. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSky Alert PendantからSOSアラートを検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky Alert Pendant에서 SOS 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de alerta SOS desde OnSky Alert Pendant. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือน SOS จาก OnSky Alert Pendant กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng Alertong SOS mula sa OnSky Alert Pendant. Pakisuri.`

---

### 2.2. SkySOSFallDetection (Phát Hiện Té Ngã)

- **Tiếng Anh (en-US):**
  > `Detect Fall Alert from OnSky Alert Pendant. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo té ngã từ OnSky Alert Pendant. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky Alert Pendant的跌倒警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky Alert Pendant的跌倒警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une alerte de chute depuis l'OnSky Alert Pendant. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Sturz-Alarm von OnSky Alert Pendant erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi Peringatan Jatuh dari OnSky Alert Pendant. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSky Alert Pendantが転倒を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky Alert Pendant에서 낙상 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de alerta de caída desde OnSky Alert Pendant. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือนการหกล้มจาก OnSky Alert Pendant กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng Alertong Pagkahulog mula sa OnSky Alert Pendant. Pakisuri.`

---

### 2.3. SkySOSGeofenceEnter (Đi Vào Vùng An Toàn (Geofence Enter))

- **Tiếng Anh (en-US):**
  > `Detect user entering the Safety Zone by OnSky Alert Pendant. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện người dùng vào vùng an toàn từ OnSky Alert Pendant. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到用户进入OnSky Alert Pendant设定的安全区域。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到用戶進入OnSky Alert Pendant設定的安全區域。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection de l'entrée de l'utilisateur dans la zone de sécurité par l'OnSky Alert Pendant. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Benutzer betritt die Sicherheitszone laut OnSky Alert Pendant. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi pengguna memasuki Zona Aman oleh OnSky Alert Pendant. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSky Alert Pendantがセーフティゾーンへの進入を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky Alert Pendant에서 사용자가 안전 구역에 진입했음을 감지했습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de usuario entrando en la Zona Segura por OnSky Alert Pendant. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบผู้ใช้เข้าสู่พื้นที่ปลอดภัยโดย OnSky Alert Pendant กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita na ang gumagamit ay pumasok sa Safety Zone mula sa OnSky Alert Pendant. Pakisuri.`

---

### 2.4. SkySOSGeofenceExit (Đi Ra Khỏi Vùng An Toàn (Geofence Exit))

- **Tiếng Anh (en-US):**
  > `Detect user exiting the Safety Zone by OnSky Alert Pendant. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện người dùng rời khỏi vùng an toàn từ OnSky Alert Pendant. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到用户离开OnSky Alert Pendant设定的安全区域。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到用戶離開OnSky Alert Pendant設定的安全區域。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection de la sortie de l'utilisateur de la zone de sécurité par l'OnSky Alert Pendant. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Benutzer verlässt die Sicherheitszone laut OnSky Alert Pendant. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi pengguna keluar dari Zona Aman oleh OnSky Alert Pendant. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSky Alert Pendantがセーフティゾーンからの退出を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky Alert Pendant에서 사용자가 안전 구역을 벗어났음을 감지했습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de usuario saliendo de la Zona Segura por OnSky Alert Pendant. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบผู้ใช้ออกจากพื้นที่ปลอดภัยโดย OnSky Alert Pendant กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita na ang gumagamit ay lumabas sa Safety Zone mula sa OnSky Alert Pendant. Pakisuri.`

---

### 2.5. SkySOSDeviceMoving (Người Dùng Đang Di Chuyển (Device Moving))

- **Tiếng Anh (en-US):**
  > `Detect user moving by OnSky Alert Pendant.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện người dùng đang di chuyển từ OnSky Alert Pendant.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky Alert Pendant的用户移动状态。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky Alert Pendant的用戶移動狀態。`

- **Tiếng Pháp (fr-FR):**
  > `Détection de mouvements de l'utilisateur par l'OnSky Alert Pendant.`

- **Tiếng Đức (de-DE):**
  > `Benutzerbewegung von OnSky Alert Pendant erkannt.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi pengguna sedang bergerak oleh OnSky Alert Pendant.`

- **Tiếng Nhật (ja-JP):**
  > `OnSky Alert Pendantが利用者の移動を検知しました。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky Alert Pendant에서 사용자의 이동이 감지되었습니다.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de movimiento del usuario por OnSky Alert Pendant.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบผู้ใช้กำลังเคลื่อนไหวโดย OnSky Alert Pendant`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita na ang gumagamit ay kumikilos mula sa OnSky Alert Pendant.`

---

### 2.6. SkySOSDeviceStopped (Người Dùng Đã Dừng Lại (Device Stopped))

- **Tiếng Anh (en-US):**
  > `Detect user stopped by OnSky Alert Pendant.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện người dùng đã dừng lại từ OnSky Alert Pendant.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky Alert Pendant的用户静止/停止移动状态。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky Alert Pendant的用戶靜止/停止移動狀態。`

- **Tiếng Pháp (fr-FR):**
  > `Détection de l'arrêt de l'utilisateur par l'OnSky Alert Pendant.`

- **Tiếng Đức (de-DE):**
  > `Benutzer hat angehalten laut OnSky Alert Pendant.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi pengguna telah berhenti oleh OnSky Alert Pendant.`

- **Tiếng Nhật (ja-JP):**
  > `OnSky Alert Pendantが利用者の停止を検知しました。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky Alert Pendant에서 사용자가 멈추었음을 감지했습니다.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de usuario detenido por OnSky Alert Pendant.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบผู้ใช้หยุดเคลื่อนไหวโดย OnSky Alert Pendant`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita na ang gumagamit ay huminto mula sa OnSky Alert Pendant.`

---

## 3. Thiết Bị SkyBand (Vòng Đeo Tay Thông Minh)

### 3.1. SkyBandSpo2Low (Nồng Độ Oxy Trong Máu Thấp (SpO2))

- **Tiếng Anh (en-US):**
  > `Detect low peripheral oxygen saturation (spo2) from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện độ bão hòa oxy trong máu (spo2) thấp từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的血氧饱和度(SpO2)偏低警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的血氧飽和度(SpO2)偏低警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une faible saturation en oxygène (SpO2) depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Niedrige Sauerstoffsättigung (SpO2) vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi saturasi oksigen darah (SpO2) rendah dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスから低血中酸素飽和度(SpO2)を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 낮은 혈중 산소포화도(SpO2)가 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de baja saturación de oxígeno (SpO2) desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบระดับออกซิเจนในเลือด (SpO2) ต่ำจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng mababang oxygen saturation (SpO2) mula sa aparatong OnSky. Pakisuri.`

---

### 3.2. SkyBandHeartRateLow (Nhịp Tim Thấp)

- **Tiếng Anh (en-US):**
  > `Detect low heart rate from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện nhịp tim thấp từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的心率偏低警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的心率偏低警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une fréquence cardiaque faible depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Niedrige Herzfrequenz vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi detak jantung rendah dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスから低い心拍数を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 낮은 심박수가 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de frecuencia cardíaca baja desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบอัตราการเต้นของหัวใจต่ำจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng mababang tibok ng puso mula sa aparatong OnSky. Pakisuri.`

---

### 3.3. SkyBandHeartRateHeight (Nhịp Tim Cao)

- **Tiếng Anh (en-US):**
  > `Detect high heart rate from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện nhịp tim cao từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的心率偏高警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的心率偏高警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une fréquence cardiaque élevée depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Hohe Herzfrequenz vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi detak jantung tinggi dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスから高い心拍数を検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 높은 심박수가 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de frecuencia cardíaca alta desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบอัตราการเต้นของหัวใจสูงจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng mataas na tibok ng puso mula sa aparatong OnSky. Pakisuri.`

---

## 4. Cảnh Báo Mặc Định (Default Alert)

### 4.1. DefaultAlert (Cảnh Báo Mặc Định / Chung)

- **Tiếng Anh (en-US):**
  > `Detect default alert from OnSky device. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Phát hiện cảnh báo mặc định từ thiết bị OnSky. Vui lòng kiểm tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `检测到来自OnSky设备的默认警报。请检查。`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `檢測到來自OnSky設備的默認警報。請檢查。`

- **Tiếng Pháp (fr-FR):**
  > `Détection d'une alerte par défaut depuis l'appareil OnSky. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `Standard-Alarm vom OnSky-Gerät erkannt. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Mendeteksi peringatan default dari perangkat OnSky. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyデバイスからデフォルトアラートを検知しました。ご確認ください。`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 기기에서 기본 알림이 감지되었습니다. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Detección de alerta predeterminada desde el dispositivo OnSky. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `ตรวจพบการแจ้งเตือนเริ่มต้นจากอุปกรณ์ OnSky กรุณาตรวจสอบ`

- **Tiếng Filipino (fil-PH):**
  > `Nakakita ng default na alerto mula sa aparatong OnSky. Pakisuri.`

---

