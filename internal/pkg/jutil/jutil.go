// Package jutil contains func related to JSON.
package jutil

import (
	"encoding/json"

	"github.com/Hidayathamir/go-product/internal/pkg/trace"
	"github.com/sirupsen/logrus"
)

// ToJSONString return JSON string of v, if err return "" and do logging.
func ToJSONString(v any) string {
	jsonByte, err := json.Marshal(v)
	if err != nil {
		logrus.Warn(trace.Wrap(err))
		return ""
	}
	return string(jsonByte)
}
