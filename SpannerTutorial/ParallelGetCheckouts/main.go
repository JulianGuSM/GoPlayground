package main

import (
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger := initLogger()

	req, err := http.NewRequest("GET", "http://localhost:8080/connectors/v2/checkouts?app_platform=shopify&app_key=guo-api-test&organization_id=0360ab9dc10d40298c86bb8e8aa48bce", nil)
	if err != nil {
		logger.Error("error: ", zap.Error(err))
	}
	req.Header.Set("am-api-key", "ojli2BhpVANaT1j1Qd5oOGHn2s5OkU3h1")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("error: ", zap.Error(err))
	}
	if resp.StatusCode != 200 {
		logger.Error("error", zap.String("status code", resp.Status))
	}
}

func initLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	return logger
}
