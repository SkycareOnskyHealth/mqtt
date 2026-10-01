package locales

import "strings"

var registry = make(map[string]*Bundle)

// DefaultLocale là ngôn ngữ fallback khi không tìm thấy locale yêu cầu
const DefaultLocale = "en-US"

// Register đăng ký một Bundle ngôn ngữ vào registry
func Register(b *Bundle) {
	registry[strings.ToLower(b.Code)] = b
}

// GetAllBundles trả về toàn bộ bundle đã đăng ký (dùng cho unit test)
func GetAllBundles() map[string]*Bundle {
	return registry
}

// GetBundle tìm kiếm Bundle phù hợp theo chuỗi locale đầu vào.
// Hỗ trợ định dạng có gạch nối (-), gạch dưới (_) hoặc chữ hoa chữ thường.
// Nếu không tìm thấy chính xác, thử tìm theo tiền tố ngôn ngữ (ví dụ "ja" -> "ja-jp").
// Nếu vẫn không thấy, tự động fallback về DefaultLocale ("en-US").
func GetBundle(locale string) *Bundle {
	if locale == "" {
		if def, ok := registry[strings.ToLower(DefaultLocale)]; ok {
			return def
		}
		return nil
	}

	clean := strings.ToLower(strings.ReplaceAll(locale, "_", "-"))
	if b, ok := registry[clean]; ok {
		return b
	}

	// Thử tìm theo tiền tố ngôn ngữ (ví dụ "ja" khớp "ja-jp")
	parts := strings.Split(clean, "-")
	langPrefix := parts[0]
	for k, b := range registry {
		if k == langPrefix || strings.HasPrefix(k, langPrefix+"-") {
			return b
		}
	}

	// Fallback mặc định về en-US
	if def, ok := registry[strings.ToLower(DefaultLocale)]; ok {
		return def
	}

	// Nếu en-US cũng chưa kịp đăng ký, lấy bất kỳ bundle nào có trong registry
	for _, b := range registry {
		return b
	}
	return nil
}

