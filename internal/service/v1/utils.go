package service

import (
	"net"
)

type utilsService struct{}

func newUtilsService() utils {
	return &utilsService{}
}

func (u *utilsService) GetLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localIP := conn.LocalAddr().(*net.UDPAddr).IP
	return localIP.String(), nil
}
