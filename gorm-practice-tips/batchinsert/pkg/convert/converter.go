package convert

import (
	"bytes"
	"strings"
)

var (
	commonInit           []string
	uncommonInitReplacer *strings.Replacer
)

func init() {
	var commonInitForReplacer []string
	var uncommonInitForReplacer []string

	for _, initialism := range commonInit {
		commonInitForReplacer = append(commonInitForReplacer, initialism, strings.Title(strings.ToLower(initialism)))
		uncommonInitForReplacer = append(uncommonInitForReplacer, strings.Title(strings.ToLower(initialism)), initialism)
	}

	uncommonInitReplacer = strings.NewReplacer(uncommonInitForReplacer...)
}

// StringToBool ...
func StringToBool(s string) bool {
	if s == "1" || s == "true" || s == "True" {
		return true
	}

	return false
}

// UnderlineToBig ...
func UnderlineToBig(name string) string {
	if name == "" {
		return ""
	}

	temp := strings.Split(name, "_")
	var s string

	for _, v := range temp {
		vv := []rune(v)
		if len(vv) > 0 {
			if bool(vv[0] >= 'a' && vv[0] <= 'z') {
				vv[0] -= 32
			}

			s += string(vv)
		}
	}

	s = uncommonInitReplacer.Replace(s)

	return s
}

// RemoveRepInSliceOfString ...
func RemoveRepInSliceOfString(slc []string) []string {
	var result []string
	tempMap := map[string]byte{}

	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0

		if len(tempMap) != l {
			result = append(result, e)
		}
	}

	return result
}

func isASCIIUpper(c byte) bool {
	return 'A' <= c && c <= 'Z'
}

func isASCIIDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// BigToUnderline ...
func BigToUnderline(s string) string {
	if s == "" {
		return ""
	}

	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' {
		t = append(t, 'X')
		i++
	}

	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && isASCIIUpper(s[i+1]) {
			continue
		}

		if isASCIIDigit(c) {
			t = append(t, c)
			continue
		}

		if isASCIIUpper(c) {
			c ^= ' '
		}

		t = append(t, c)

		for i+1 < len(s) && isASCIIUpper(s[i+1]) {
			i++
			t = append(t, '_')
			t = append(t, bytes.ToLower([]byte{s[i]})[0])
		}
	}

	return string(t)
}
