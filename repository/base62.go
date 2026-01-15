package repository

const base62Chars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func encodeBase62(num uint64) string {
	if num == 0 {
		return "0"
	}

	result := ""
	base := uint64(len(base62Chars))

	for num > 0 {
		result = string(base62Chars[num%base]) + result
		num /= base
	}

	return result
}
