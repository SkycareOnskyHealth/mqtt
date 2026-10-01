# Định Dạng Thông Báo (Message Format) của Thiết Bị BedSensor, SkySOS và SkyBand

Tài liệu này tổng hợp format tin nhắn hoàn chỉnh được sinh ra bởi hàm `PrepareBody` trong package `mqtt` (`mqtt/main.go`) cho 3 dòng thiết bị: **BedSensor**, **SkySOS** và **SkyBand** ở đầy đủ **12 ngôn ngữ** theo chuẩn BCP 47:

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

## 1. Cấu Trúc Message Template Chung

Theo hàm `PrepareBody`:
```text
{{onsky_security}}. {{security_alert}}{{of}}{{full_name}}{{at}} {{address}}. {{phone}}:{{phone_number}}, {{device}} {{device_name}}. {{on_date}} {{date}}, {{at_time}} {{time}}. {{please_check}}.
```

### Các placeholder được quy ước:
- `[ fullName ]`: Họ và tên người dùng.
- `[ address ]`: Địa chỉ vị trí thiết bị.
- `[ phoneNumber ]`: Số điện thoại liên hệ.
- `[ deviceName ]`: Tên thiết bị (ví dụ: `SkyPad`, `SkySOS-01`, `SkyBand-01`...).
- `[ date ]`: Ngày/Tháng (định dạng `D/M`, ví dụ: `18/9`).
- `[ time ]`: Giờ:Phút (định dạng `H:M`, ví dụ: `10:15`).

> **Quy tắc hiển thị tên người dùng:**
> - **Có tên (`fullName != ""`):** Có kèm phần liên từ sở hữu (`of` / `cua` / `de` / `的`...) và tên người dùng.
> - **Không có tên (`fullName == ""`):** Tự động bỏ qua phần sở hữu và tên, trực tiếp đến vị trí / địa chỉ.

---

## 2. Thiết Bị BedSensor (Cảm Biến Giường / Nệm)

### 2.1. BedSensorSOS (Khẩn cấp SOS)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Possible SOS Urgency Alert from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Canh bao co tin hieu cap cuu cua nguoi dung duoc gui tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备发出疑似SOS紧急警报 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備發出疑似SOS緊急警報 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Alerte d'urgence SOS possible de l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Möglicher SOS-Notfallalarm vom OnSky-Gerät von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Kemungkinan Peringatan Darurat SOS dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスからSOS緊急通報の可能性がありますの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 SOS 긴급 알림 가능성 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Posible alerta de urgencia SOS del dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. อาจมีการแจ้งเตือนฉุกเฉิน SOS จากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Posibleng SOS Emergency Alert mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.2. BedSensorHeartStop (Ngưng Tim / Heart Attack)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Possible Heart Attack Alert from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Canh bao tim nguoi dung co dau hieu ngung dap duoc gui tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备发出疑似心脏骤停警报 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備發出疑似心臟驟停警報 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Alerte d'arrêt cardiaque possible de l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Möglicher Herzstillstand-Alarm vom OnSky-Gerät von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Kemungkinan Peringatan Serangan Jantung dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスから心停止の兆候を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 심정지 의심 알림 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Posible alerta de paro cardíaco del dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. อาจมีการแจ้งเตือนภาวะหัวใจหยุดเต้นจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Posibleng Heart Attack Alert mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.3. BedSensorBreathStop (Ngưng Thở / Apnea)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Possible Apnea Alert from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Canh bao phoi nguoi dung co dau hieu ngung tho duoc gui tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备发出疑似呼吸暂停警报 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備發出疑似呼吸暫停警報 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Alerte d'apnée possible de l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Möglicher Atemstillstand-Alarm (Apnoe) vom OnSky-Gerät von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Kemungkinan Peringatan Apnea (Henti Napas) dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスから無呼吸の兆候を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 무호흡 의심 알림 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Posible alerta de apnea del dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. อาจมีการแจ้งเตือนภาวะหยุดหายใจขณะหลับจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Posibleng Apnea Alert mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.4. BedSensorTachycardia (Nhịp Tim Nhanh / Tachycardia)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Detect irregular heart rhythms - tachycardia from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Phat hien nhip tim bat thuong - nhip tim nhanh tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. 从OnSky设备检测到心律不齐 - 心动过速 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. 從OnSky設備檢測到心律不齊 - 心動過速 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Détection de rythme cardiaque irrégulier - tachycardie par l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Herzrhythmusstörungen - Tachykardie vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Mendeteksi ritme jantung tidak teratur - takikardia dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスから不整脈（頻脈）を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 부정맥(빈맥) 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Detección de ritmos cardíacos irregulares - taquicardia desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. ตรวจพบจังหวะการเต้นของหัวใจผิดปกติ - หัวใจเต้นเร็วผิดปกติ (tachycardia) จากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Nakakita ng irregular na tibok ng puso - tachycardia mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.5. BedSensorBradycardia (Nhịp Tim Chậm / Bradycardia)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Detect irregular heart rhythms - bradycardia from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Phat hien nhip tim bat thuong - nhip tim cham tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. 从OnSky设备检测到心律不齐 - 心动过缓 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. 從OnSky設備檢測到心律不齊 - 心動過緩 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Détection de rythme cardiaque irrégulier - bradycardie par l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Herzrhythmusstörungen - Bradykardie vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Mendeteksi ritme jantung tidak teratur - bradikardia dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスから不整脈（徐脈）を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 부정맥(서맥) 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Detección de ritmos cardíacos irregulares - bradicardia desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. ตรวจพบจังหวะการเต้นของหัวใจผิดปกติ - หัวใจเต้นช้าผิดปกติ (bradycardia) จากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Nakakita ng irregular na tibok ng puso - bradycardia mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.6. BedSensorSeizure (Co Giật / Seizure)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Possible Seizures Alert from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Canh bao nguoi dung co dau hieu co giat duoc gui tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备发出疑似癫痫发作警报 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備發出疑似癲癇發作警報 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Alerte de convulsions possible de l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Möglicher Krampfanfall-Alarm vom OnSky-Gerät von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Kemungkinan Peringatan Kejang dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスから痙攣発作の兆候を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 발작 의심 알림 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Posible alerta de convulsiones del dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. อาจมีการแจ้งเตือนอาการชักจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Posibleng Seizures Alert mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.7. BedSensorAbnormalVitalSigns (Dấu Hiệu Sinh Tồn Bất Thường)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Alert detect abnormal Vital Signs from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Canh bao phat hien Dau Hieu Sinh Ton bat thuong duoc gui tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备检测到生命体征异常警报 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備檢測到生命體徵異常警報 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Alerte : Détection de signes vitaux anormaux de l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Alarm: Abnormale Vitalwerte vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Peringatan: Mendeteksi Tanda Vital abnormal dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. 警告: OnSkyデバイスからバイタルサインの異常を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. 경고: OnSky 기기에서 비정상적인 생체 징후 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Alerta: Detección de signos vitales anormales del dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. คำเตือน: ตรวจพบสัญญาณชีพผิดปกติจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Alerto: Nakakita ng hindi normal na Vital Signs mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.8. BedSensorBodyTempHeight (Thân Nhiệt Quá Cao)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Detecting body temperature is too high from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Phat hien nhiet do co the qua cao tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备检测到体温过高 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備檢測到體溫過高 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Température corporelle trop élevée détectée par l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Zu hohe Körpertemperatur vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Mendeteksi suhu tubuh terlalu tinggi dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスから体温異常（高温）を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 체온 과다 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Detección de temperatura corporal demasiado alta desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. ตรวจพบอุณหภูมิร่างกายสูงเกินไปจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Nakakita ng mataas na temperatura ng katawan mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.9. BedSensorRoomTempHeight (Nhiệt Độ Phòng Quá Cao)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Detecting room temperature is too high from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Phat hien nhiet do phong qua cao tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备检测到室内温度过高 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備檢測到室內溫度過高 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Température ambiante trop élevée détectée par l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Zu hohe Raumtemperatur vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Mendeteksi suhu ruangan terlalu tinggi dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスから室温異常（高温）を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 실내 온도 과다 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Detección de temperatura ambiente demasiado alta desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. ตรวจพบอุณหภูมิห้องสูงเกินไปจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Nakakita ng mataas na temperatura ng silid mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.10. BedSensorHumidityHeight (Độ Ẩm Phòng Quá Cao)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Detecting room humidity is too high from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Phat hien do am qua cao tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备检测到室内湿度过高 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備檢測到室內濕度過高 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Humidité ambiante trop élevée détectée par l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Zu hohe Luftfeuchtigkeit vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Mendeteksi kelembapan ruangan terlalu tinggi dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスから湿度異常（高湿度）を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 실내 습도 과다 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Detección de humedad ambiente demasiado alta desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. ตรวจพบความชื้นในห้องสูงเกินไปจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Nakakita ng mataas na halumigmig sa silid mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.11. BedSensorHeartRateHeight (Nhịp Tim Quá Cao)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Detecting heart rate is too high from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Phat hien nhip tim qua cao tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备检测到心率过高 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備檢測到心率過高 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Fréquence cardiaque trop élevée détectée par l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Zu hohe Herzfrequenz vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Mendeteksi detak jantung terlalu tinggi dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスから心拍数過多を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 심박수 과다 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Detección de frecuencia cardíaca demasiado alta desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. ตรวจพบอัตราการเต้นของหัวใจสูงเกินไปจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Nakakita ng napakabilis na tibok ng puso mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.12. BedSensorHeartRateLow (Nhịp Tim Quá Thấp)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Detecting heart rate is too low from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Phat hien nhip tim qua thap tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备检测到心率过低 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備檢測到心率過低 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Fréquence cardiaque trop basse détectée par l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Zu niedrige Herzfrequenz vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Mendeteksi detak jantung terlalu rendah dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスから心拍数過少を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 심박수 저하 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Detección de frecuencia cardíaca demasiado baja desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. ตรวจพบอัตราการเต้นของหัวใจต่ำเกินไปจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Nakakita ng napakababang tibok ng puso mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.13. BedSensorBedLeaving (Rời Khỏi Giường)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Detecting user have left the bed from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Phat hien nguoi dung da roi khoi giuong tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备检测到用户已离开床位 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備檢測到用戶已離開床位 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. L'utilisateur a quitté son lit selon l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Benutzer hat das Bett verlassen laut OnSky-Gerät von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Mendeteksi pengguna telah meninggalkan tempat tidur dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスが利用者の離床を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 사용자가 침대를 이탈했음을 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Detección de usuario fuera de la cama desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. ตรวจพบผู้ใช้ออกจากเตียงจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Nakakita na ang gumagamit ay umalis sa kama mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 2.14. BedSensorCrying (Em Bé Khóc)

- **Tiếng Anh (en-US):**
  > `OnSky Medical Alert service. Detecting baby crying from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dich vu y te OnSky. Phat hien em be dang khoc tu thiet bi OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky医疗警报服务. OnSky设备检测到婴儿哭闹 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky醫療警報服務. OnSky設備檢測到嬰兒哭鬧 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte médicale OnSky. Pleurs de bébé détectés par l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Medizinischer Notrufdienst. Babyweinen vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan Medis OnSky. Mendeteksi tangisan bayi dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyメディカルアラートサービス. OnSkyデバイスが赤ちゃんの泣き声を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 응급 의료 알림 서비스. OnSky 기기에서 아기 울음소리 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta médica OnSky. Detección de llanto de bebé desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือนทางการแพทย์ OnSky. ตรวจพบเสียงเด็กร้องไห้จากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Medikal na Alerto ng OnSky. Nakakita ng pag-iyak ng sanggol mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

---

## 3. Thiết Bị SkySOS (Dây Chuyền Khẩn Cấp / Alert Pendant)

### 3.1. SkySOSButtonTriggered (Bấm Nút SOS Khẩn Cấp)

- **Tiếng Anh (en-US):**
  > `OnSky Alert service. Possible SOS Emergency Alert from OnSky Alert Pendant of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dịch vụ cảnh báo OnSky. Canh bao khan cap SOS tu OnSky Alert Pendant cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky警报服务. OnSky Alert Pendant发出疑似SOS紧急警报 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky警報服務. OnSky Alert Pendant發出疑似SOS緊急警報 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte OnSky. Alerte d'urgence SOS possible de l'OnSky Alert Pendant de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Warndienst. Möglicher SOS-Notfallalarm von OnSky Alert Pendant von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan OnSky. Kemungkinan Peringatan Darurat SOS dari OnSky Alert Pendant dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyアラートサービス. OnSky Alert PendantからSOS緊急アラートを検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 알림 서비스. OnSky Alert Pendant에서 SOS 긴급 알림 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta OnSky. Posible alerta de emergencia SOS de OnSky Alert Pendant de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือน OnSky. อาจมีการแจ้งเตือนฉุกเฉิน SOS จาก OnSky Alert Pendant ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Alerto ng OnSky. Posibleng SOS Emergency Alert mula sa OnSky Alert Pendant ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 3.2. SkySOSFallDetection (Phát Hiện Té Ngã)

- **Tiếng Anh (en-US):**
  > `OnSky Alert service. Detect fall from OnSky Alert Pendant of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dịch vụ cảnh báo OnSky. Phat hien te nga tu OnSky Alert Pendant cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky警报服务. OnSky Alert Pendant检测到跌倒 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky警報服務. OnSky Alert Pendant檢測到跌倒 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte OnSky. Chute détectée par l'OnSky Alert Pendant de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Warndienst. Sturz erkannt von OnSky Alert Pendant von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan OnSky. Mendeteksi jatuh dari OnSky Alert Pendant dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyアラートサービス. OnSky Alert Pendantが転倒を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 알림 서비스. OnSky Alert Pendant에서 낙상 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta OnSky. Detección de caída de OnSky Alert Pendant de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือน OnSky. ตรวจพบการหกล้มจาก OnSky Alert Pendant ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Alerto ng OnSky. Nakakita ng pagkahulog mula sa OnSky Alert Pendant ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 3.3. SkySOSGeofenceEnter (Đi Vào Vùng An Toàn)

- **Tiếng Anh (en-US):**
  > `OnSky Alert service. Detect user entering Safety Zone by OnSky device Alert Pendant of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dịch vụ cảnh báo OnSky. Phat hien nguoi dung di vao vung an toan tu OnSky Alert Pendant cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky警报服务. OnSky Alert Pendant检测到用户进入安全区域 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky警報服務. OnSky Alert Pendant檢測到用戶進入安全區域 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte OnSky. L'utilisateur entre dans la zone de sécurité selon l'OnSky Alert Pendant de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Warndienst. Benutzer betritt Sicherheitszone laut OnSky Alert Pendant von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan OnSky. Mendeteksi pengguna memasuki Zona Aman oleh OnSky Alert Pendant dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyアラートサービス. OnSky Alert Pendantがセーフティゾーンへの進入を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 알림 서비스. OnSky Alert Pendant에서 사용자가 안전 구역에 진입했음을 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta OnSky. Detección de usuario entrando en la Zona Segura por OnSky Alert Pendant de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือน OnSky. ตรวจพบผู้ใช้เข้าสู่พื้นที่ปลอดภัยโดย OnSky Alert Pendant ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Alerto ng OnSky. Nakakita na ang gumagamit ay pumasok sa Safety Zone mula sa OnSky Alert Pendant ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 3.4. SkySOSGeofenceExit (Đi Ra Khỏi Vùng An Toàn)

- **Tiếng Anh (en-US):**
  > `OnSky Alert service. Detect user exiting Safety Zone by OnSky device Alert Pendant of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dịch vụ cảnh báo OnSky. Phat hien nguoi dung di ra khoi vung an toan tu OnSky Alert Pendant cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky警报服务. OnSky Alert Pendant检测到用户离开安全区域 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky警報服務. OnSky Alert Pendant檢測到用戶離開安全區域 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte OnSky. L'utilisateur quitte la zone de sécurité selon l'OnSky Alert Pendant de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Warndienst. Benutzer verlässt Sicherheitszone laut OnSky Alert Pendant von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan OnSky. Mendeteksi pengguna keluar dari Zona Aman oleh OnSky Alert Pendant dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyアラートサービス. OnSky Alert Pendantがセーフティゾーンからの退出を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 알림 서비스. OnSky Alert Pendant에서 사용자가 안전 구역을 벗어났음을 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta OnSky. Detección de usuario saliendo de la Zona Segura por OnSky Alert Pendant de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือน OnSky. ตรวจพบผู้ใช้ออกจากพื้นที่ปลอดภัยโดย OnSky Alert Pendant ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Alerto ng OnSky. Nakakita na ang gumagamit ay lumabas sa Safety Zone mula sa OnSky Alert Pendant ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

---

## 4. Thiết Bị SkyBand (Vòng Đeo Tay Thông Minh)

### 4.1. SkyBandSpo2Low (Nồng Độ Oxy Trong Máu Thấp (SpO2))

- **Tiếng Anh (en-US):**
  > `OnSky Alert service. Detect low peripheral oxygen saturation (spo2) from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dịch vụ cảnh báo OnSky. Phát hiện độ bão hòa oxy trong máu (spo2) thấp từ thiết bị OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky警报服务. OnSky设备检测到血氧饱和度(SpO2)偏低 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky警報服務. OnSky設備檢測到血氧飽和度(SpO2)偏低 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte OnSky. Saturation pulsée en oxygène (SpO2) faible détectée par l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Warndienst. Niedrige Sauerstoffsättigung (SpO2) vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan OnSky. Mendeteksi saturasi oksigen (SpO2) rendah dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyアラートサービス. OnSkyデバイスが低血中酸素飽和度(SpO2)を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 알림 서비스. OnSky 기기에서 낮은 혈중 산소포화도(SpO2) 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta OnSky. Detección de baja saturación periférica de oxígeno (SpO2) desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือน OnSky. ตรวจพบระดับออกซิเจนในเลือด (SpO2) ต่ำจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Alerto ng OnSky. Nakakita ng mababang oxygen saturation (SpO2) mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 4.2. SkyBandHeartRateLow (Nhịp Tim Thấp)

- **Tiếng Anh (en-US):**
  > `OnSky Alert service. Detect low heart rate from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dịch vụ cảnh báo OnSky. Phát hiện nhip tim thấp từ thiết bị OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky警报服务. OnSky设备检测到心率偏低 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky警報服務. OnSky設備檢測到心率偏低 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte OnSky. Fréquence cardiaque faible détectée par l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Warndienst. Niedrige Herzfrequenz vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan OnSky. Mendeteksi detak jantung rendah dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyアラートサービス. OnSkyデバイスが低い心拍数を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 알림 서비스. OnSky 기기에서 낮은 심박수 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta OnSky. Detección de frecuencia cardíaca baja desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือน OnSky. ตรวจพบอัตราการเต้นของหัวใจต่ำจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Alerto ng OnSky. Nakakita ng mababang tibok ng puso mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

### 4.3. SkyBandHeartRateHeight (Nhịp Tim Cao)

- **Tiếng Anh (en-US):**
  > `OnSky Alert service. Detect high heart rate from OnSky device of [ fullName ] at [ address ]. phone:[ phoneNumber ], device [ deviceName ]. on [ date ], at [ time ]. Please check.`

- **Tiếng Việt (vi-VN):**
  > `Dịch vụ cảnh báo OnSky. Phát hiện nhịp tim cao từ thiết bị OnSky cua [ fullName ] tai [ address ]. SDT:[ phoneNumber ], thiet bi [ deviceName ]. Vao ngay [ date ], luc [ time ]. Vui long kiem tra.`

- **Tiếng Trung Giản thể (zh-Hans):**
  > `OnSky警报服务. OnSky设备检测到心率偏高 的 [ fullName ] 在 [ address ]. 电话:[ phoneNumber ], 设备 [ deviceName ]. 于日期 [ date ], 时间 [ time ]. 请检查.`

- **Tiếng Trung Phồn thể (zh-Hant):**
  > `OnSky警報服務. OnSky設備檢測到心率偏高 的 [ fullName ] 在 [ address ]. 電話:[ phoneNumber ], 設備 [ deviceName ]. 於日期 [ date ], 時間 [ time ]. 請檢查.`

- **Tiếng Pháp (fr-FR):**
  > `Service d'alerte OnSky. Fréquence cardiaque élevée détectée par l'appareil OnSky de [ fullName ] à [ address ]. téléphone:[ phoneNumber ], appareil [ deviceName ]. le [ date ], à [ time ]. Veuillez vérifier.`

- **Tiếng Đức (de-DE):**
  > `OnSky Warndienst. Hohe Herzfrequenz vom OnSky-Gerät erkannt von [ fullName ] bei [ address ]. Telefon:[ phoneNumber ], Gerät [ deviceName ]. am [ date ], um [ time ]. Bitte prüfen.`

- **Tiếng Indonesia (id-ID):**
  > `Layanan Peringatan OnSky. Mendeteksi detak jantung tinggi dari perangkat OnSky dari [ fullName ] di [ address ]. telepon:[ phoneNumber ], perangkat [ deviceName ]. pada tanggal [ date ], pukul [ time ]. Silakan periksa.`

- **Tiếng Nhật (ja-JP):**
  > `OnSkyアラートサービス. OnSkyデバイスが高い心拍数を検知しましたの[ fullName ] 場所: [ address ]. 電話:[ phoneNumber ], デバイス [ deviceName ]. 日付 [ date ], 時刻 [ time ]. ご確認ください.`

- **Tiếng Hàn (ko-KR):**
  > `OnSky 알림 서비스. OnSky 기기에서 높은 심박수 감지의 [ fullName ] 위치: [ address ]. 전화번호:[ phoneNumber ], 기기 [ deviceName ]. 일자 [ date ], 시간 [ time ]. 확인해 주세요.`

- **Tiếng Tây Ban Nha (es-ES):**
  > `Servicio de alerta OnSky. Detección de frecuencia cardíaca alta desde el dispositivo OnSky de [ fullName ] en [ address ]. teléfono:[ phoneNumber ], dispositivo [ deviceName ]. el [ date ], a las [ time ]. Por favor revise.`

- **Tiếng Thái (th-TH):**
  > `บริการแจ้งเตือน OnSky. ตรวจพบอัตราการเต้นของหัวใจสูงจากอุปกรณ์ OnSky ของ [ fullName ] ที่ [ address ]. เบอร์โทรศัพท์:[ phoneNumber ], อุปกรณ์ [ deviceName ]. ในวันที่ [ date ], เวลา [ time ]. กรุณาตรวจสอบ.`

- **Tiếng Filipino (fil-PH):**
  > `Serbisyong Alerto ng OnSky. Nakakita ng mataas na tibok ng puso mula sa aparatong OnSky ni [ fullName ] sa [ address ]. telepono:[ phoneNumber ], aparato [ deviceName ]. noong [ date ], nang [ time ]. Pakisuri.`

---

