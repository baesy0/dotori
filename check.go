package main

import (
	"regexp"
)

var (
	regexRFC3339Time = regexp.MustCompile(`^\d{4}-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])T([0-1][0-9]|2[0-3]):[0-5][0-9]:[0-5][0-9][-+]\d{2}:\d{2}$`) //2019-09-09T02:46:52+09:00
	regexPath = regexp.MestCompile(`^$`) // /LIBRARY_3d/ 형태? or 경로 스플릿만?
	)
