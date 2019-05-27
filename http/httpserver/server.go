package httpserver

import (
	"crypto/x509"
	"errors"
	"fmt"
	"gitee.com.egcode.plugins/eglog"
	"gitee.com.egcode.plugins/ssl"
	"github.com/facebookgo/grace/gracehttp"
	"io/ioutil"
	"net/http"
	"os"
)

func NewServer(c Config) error {
	server := &http.Server{
		Addr:    c.ListenAddr,
		Handler: c.Handler,
	}

	if len(c.TLS.CertFile) == 0 && len(c.TLS.KeyFile) == 0 {
		pid := os.Getpid()

		eglog.Info().Msgf("Serving %s with pid %d .", server.Addr, pid)
		if err := gracehttp.Serve(server); err != nil {
			eglog.Error().Msg(err.Error())
			return err
		}
		return nil
	}

	ca, err := ioutil.ReadFile(c.TLS.CAFile)
	if nil != err {
		return fmt.Errorf("read server tls file failed. err:%v", err)
	}

	if false == x509.NewCertPool().AppendCertsFromPEM(ca) {
		return errors.New("append cert from pem failed")
	}

	tlsC, err := ssl.ServerTslConfVerityClient(c.TLS.CAFile,
		c.TLS.CertFile,
		c.TLS.KeyFile,
		c.TLS.Password)
	if err != nil {
		return fmt.Errorf("generate tls config failed. err: %v", err)
	}
	tlsC.BuildNameToCertificate()

	server.TLSConfig = tlsC
	pid := os.Getpid()
	eglog.Info().Msgf("start secure server on %s,you can exec 'kill -USR2 %d' to restart server!", c.ListenAddr, pid)
	if err := gracehttp.Serve(server); err != nil {
		eglog.Error().Msgf(err.Error())
		return err
	}
	return nil
}
