package helper

import  "github.com/mileusna/useragent"

// 解析用户代理
func ParseUserAgent(userAgentStr string)ua.UserAgent{

	return ua.Parse(userAgentStr)
}