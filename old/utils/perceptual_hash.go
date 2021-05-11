package utils

func IsImage(mimeType string) bool {
	switch mimeType {
	case "image/jpeg":
		return true
	case "image/png":
		return true
	case "image/webp":
		return true
	default:
		return false
	}
}
