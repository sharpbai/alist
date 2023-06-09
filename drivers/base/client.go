package base

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/alist-org/alist/v3/internal/conf"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/http2"
)

var (
	NoRedirectClient *resty.Client
	RestyClient      *resty.Client
	HttpClient       *http.Client
	Http2Client      *http.Client
)

// var UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
var UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"
var DefaultTimeout = time.Second * 30

func InitClient() {
	NoRedirectClient = resty.New().SetRedirectPolicy(
		resty.RedirectPolicyFunc(func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}),
	).SetTLSClientConfig(&tls.Config{InsecureSkipVerify: conf.Conf.TlsInsecureSkipVerify})
	NoRedirectClient.SetHeader("user-agent", UserAgent)

	RestyClient = NewRestyClient()
	HttpClient = NewHttpClient()
	Http2Client = NewHttp2Client()
}

func NewRestyClient() *resty.Client {
	client := resty.New().
		SetHeader("user-agent", UserAgent).
		SetRetryCount(3).
		SetTimeout(DefaultTimeout).
		SetTLSClientConfig(&tls.Config{InsecureSkipVerify: conf.Conf.TlsInsecureSkipVerify})
	return client
}

func NewHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Hour * 48,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.Conf.TlsInsecureSkipVerify},
		},
	}
}

func NewHttp2Client() *http.Client {
	return &http.Client{
		Timeout: time.Hour * 48,
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: conf.Conf.TlsInsecureSkipVerify},
		},
	}
}
