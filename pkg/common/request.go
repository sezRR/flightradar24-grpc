package common

import "net/http"

func SetHeaders(req *http.Request) {
	headers := map[string]string{
		"accept":                "*/*",
		"accept-language":       "en-US,en;q=0.9,tr;q=0.8,zh-CN;q=0.7,zh;q=0.6",
		"content-type":          "application/grpc-web+proto",
		"fr24-device-id":        "web-1hkr68cai-yPAHWRoTuhSRbai2ijUyF",
		"fr24-platform":         "web-25.010.1124",
		"origin":                "https://www.flightradar24.com",
		"priority":              "u=1, i",
		"referer":               "https://www.flightradar24.com/",
		"sec-ch-ua":             `"Google Chrome";v="131", "Chromium";v="131", "Not_A Brand";v="24"`,
		"sec-ch-ua-mobile":      "?0",
		"sec-ch-ua-platform":    `"Windows"`,
		"sec-fetch-dest":        "empty",
		"sec-fetch-mode":        "cors",
		"sec-fetch-site":        "same-site",
		"user-agent":            "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36",
		"x-envoy-retry-grpc-on": "unavailable",
		"x-grpc-web":            "1",
		"x-user-agent":          "grpc-web-javascript/0.1",
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
}
