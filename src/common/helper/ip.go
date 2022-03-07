package helper

import (
	"net"
	"net/http"

	"github.com/thinkeridea/go-extend/exnet"
	"go.uber.org/zap"
)


func Ip2Long(ip string) uint{
	ip1 := net.ParseIP(ip)
	if ip2,err :=exnet.IP2Long(ip1);err!=nil{
		return 0
	}else{
		return ip2
	}
}

func Ip2Str(ip uint) string{

	ips,_ :=exnet.Long2IPString(ip)

	return ips

}


func ClientIP(req *http.Request)string{

	ip := exnet.ClientPublicIP(req)

	if ip=="" {
		ip = exnet.ClientIP(req)
		
	}

	if ip=="" {
		ip = req.Host
	}

	zap.L().Info("ip",zap.Any("",ip))
	
	return ip
}