package shelf

import (
	"regexp"
	"sync"
)

var (
	regCache sync.Map
)

func getOrCreateReg(pattern string) (reg *regexp.Regexp, ok bool) {
	if IsBlank(pattern) {
		return nil, false
	}

	obj, ok := regCache.Load(pattern)
	if !ok {
		reg, err := regexp.Compile(pattern)
		if err != nil {
			// TODO LOG
			return nil, false
		}
		regCache.Store(pattern, reg)
		return reg, true
	} else {
		reg := obj.(*regexp.Regexp)
		return reg, true
	}
}

func FindRegMatched(pattern, value string) string {
	reg, ok := getOrCreateReg(pattern)
	if !ok {
		return value
	}
	matched, ok := FindMatchedString(reg, value)
	if ok {
		value = matched
	}
	return value
}

func RemoveRegsMatched(patterns, value string) string {
	if IsNotBlank(patterns) {
		patterns := Split(patterns, ';')
		for _, pattern := range patterns {
			reg, ok := getOrCreateReg(pattern)
			if !ok {
				continue
			}
			value = reg.ReplaceAllString(value, "")
		}
	}
	return value
}

func FindMatchedString(reg *regexp.Regexp, str string) (string, bool) {
	match := reg.FindSubmatch([]byte(str))
	if len(match) > 1 {
		return string(match[1]), true
	}
	return "", false
}
