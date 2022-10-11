package response

import (
	"encoding/json"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func Response(status int, data interface{}, rw http.ResponseWriter) {
	respBytes, err := json.Marshal(data)
	if err != nil {
		logger.Error(err)
		status = http.StatusInternalServerError
	}
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(status)
	rw.Write(respBytes)
}
