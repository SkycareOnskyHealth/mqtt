package locales

import "strings"

var registry = make(map[string]*Bundle)

// DefaultLocale is the fallback language when requested locale is not found
const DefaultLocale = "en-US"

// Register registers a language Bundle into the registry
func Register(b *Bundle) {
	registry[strings.ToLower(b.Code)] = b
}

// GetAllBundles returns all registered bundles (used for unit tests)
func GetAllBundles() map[string]*Bundle {
	return registry
}

// GetBundle searches for a matching Bundle by locale string.
// Supports hyphen (-), underscore (_), and case-insensitive formats.
// If exact match is not found, tries matching by language prefix (e.g. "ja" -> "ja-jp").
// If still not found, automatically falls back to DefaultLocale ("en-US").
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

	// Try matching by language prefix (e.g. "ja" matches "ja-jp")
	parts := strings.Split(clean, "-")
	langPrefix := parts[0]
	for k, b := range registry {
		if k == langPrefix || strings.HasPrefix(k, langPrefix+"-") {
			return b
		}
	}

	// Fallback to en-US
	if def, ok := registry[strings.ToLower(DefaultLocale)]; ok {
		return def
	}

	// If en-US is not registered, return any available bundle in registry
	for _, b := range registry {
		return b
	}
	return nil
}
