package tools

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"net/http"
	"net/url"
	"os"
)

// NewEthClientWithProxy generate an eth client with proxy if set.The program will panic because of the invalid proxy.
func NewEthClientWithProxy(endpoint string, proxyUrls ...string) (*ethclient.Client, error) {
	proxyUrl := ""
	switch len(proxyUrls) {
	case 0: // if not set just use http_proxy env params
		httpProxy := os.Getenv("http_proxy")
		if len(httpProxy) == 0 {
			//no http_proxy, just warnning
			log.Warn("new client with proxy", "err", "http_proxy not set")
			break
		}
		log.Info("set proxy to environment", "http_proxy", httpProxy)
		proxyUrl = httpProxy

	case 1: // normal
		proxyUrl = proxyUrls[0]
	default: // should never happen
		panic(1)
	}

	// create http client
	httpClient := &http.Client{
		Transport: func() *http.Transport {
			if len(proxyUrl) == 0 { // no proxy, just return nil
				return nil
			}
			log.Info("use proxy", "proxy", proxyUrl)
			proxyURL, err := url.Parse(proxyUrl)
			if err != nil {
				panic(err)
			}
			return &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		}(),
	}

	c, err := rpc.DialOptions(context.Background(), endpoint, rpc.WithHTTPClient(httpClient))
	if err != nil {
		return nil, err
	}
	client := ethclient.NewClient(c)
	return client, nil
}
