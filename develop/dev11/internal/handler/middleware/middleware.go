package middleware

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func NewErrorResponse(w http.ResponseWriter, statusCode int, message error) {
	logrus.Error(message)
	output, _ := json.Marshal(map[string]interface{}{"error": message.Error()})
	http.Error(w, string(output), statusCode)
}

func ValidTime(t string) (string, int, error) {
	layout := "2006-01-02"
	day, err := time.Parse(layout, t)
	if err != nil {
		return t, 400, err
	}
	return day.Format(layout), 200, nil
}
