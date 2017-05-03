package rpc

import (
	"context"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGet(t *testing.T) {

	resp, err := DefaultClient.DoRequest(context.Background(), "GET", "https://razeencheng.com")
	if err != nil {
		t.Fatal(err)
	}
	read(resp, t)
}

func TestGetWith(t *testing.T) {
	url := "http://192.168.252.135:8015/api/v1/tools/http2_status"
	data := map[string][]string{
		"domain": []string{"razeencheng.com"},
	}
	resp, err := DefaultClient.DoRequestWithForm(context.Background(), "GET", url, data)
	if err != nil {
		t.Fatal(err)
	}
	read(resp, t)
}

var (
	cert string = `
-----BEGIN CERTIFICATE-----
MIIF/TCCBOWgAwIBAgIQNzDv8u6B/jB8nfD8JgEM5TANBgkqhkiG9w0BAQsFADBE
MQswCQYDVQQGEwJDTjEaMBgGA1UECgwRV29TaWduIENBIExpbWl0ZWQxGTAXBgNV
BAMMEFdvU2lnbiBPViBTU0wgQ0EwHhcNMTcwMTE2MDkxODU4WhcNMjAwMTE2MDkx
ODU4WjCBrzELMAkGA1UEBhMCQ04xLTArBgNVBAoMJOays+WMl+ecgeWFrOWuieWO
heS6pOmAmuitpuWvn+aAu+mYnzEtMCsGA1UECwwk5rKz5YyX55yB5YWs5a6J5Y6F
5Lqk6YCa6K2m5a+f5oC76ZifMRUwEwYDVQQHDAznn7PlrrbluoTluIIxEjAQBgNV
BAgMCeays+WMl+ecgTEXMBUGA1UEAwwOMTEwLjI0OS4yMTguODYwggEiMA0GCSqG
SIb3DQEBAQUAA4IBDwAwggEKAoIBAQCWuy4Wq4lFaET/EFHdPcInuA5JlIAC4Cbn
4r9cF8lIk1q678sKuDGiuhv5hv4Y9a/Opx9nWaaNoaDljoDl52PcPAHMFxa/Lc8c
ihrOEU/Wa+LlX1uRaZGPVD9uqmGY7SzOl3lT4g71K1mkVlLIz8oUnFzcYvaIiv65
upk3/UE4axC72Do2K+w07i5pgNUOW2pf0BswqZJcKl/YYkku20TOLw6Cic+cfr3h
+DSNebXb0VHWevB72SeSxfhILYACMGNU/+Ooru2TduZoMYi86e/lw+hWJapxHtMG
hglbqL0no+0ypECq+wEXLsSy7Opc/Nd2vLIMybyFesmEYZ4ibRv/AgMBAAGjggJ9
MIICeTAMBgNVHRMBAf8EAjAAMDwGA1UdHwQ1MDMwMaAvoC2GK2h0dHA6Ly93b3Np
Z24uY3JsLmNlcnR1bS5wbC93b3NpZ24tb3ZjYS5jcmwwdwYIKwYBBQUHAQEEazBp
MC4GCCsGAQUFBzABhiJodHRwOi8vd29zaWduLW92Y2Eub2NzcC1jZXJ0dW0uY29t
MDcGCCsGAQUFBzAChitodHRwOi8vcmVwb3NpdG9yeS5jZXJ0dW0ucGwvd29zaWdu
LW92Y2EuY2VyMB8GA1UdIwQYMBaAFKETVNxWcywngsrIhO/uvwD9X6tWMB0GA1Ud
DgQWBBTI80yM37wDIz2HP4TQ79GR53RFlTAOBgNVHQ8BAf8EBAMCBaAwggEgBgNV
HSAEggEXMIIBEzAIBgZngQwBAgIwggEFBgwqhGgBhvZ3AgUBDAIwgfQwgfEGCCsG
AQUFBwICMIHkMB8WGEFzc2VjbyBEYXRhIFN5c3RlbXMgUy5BLjADAgEBGoHAVXNh
Z2Ugb2YgdGhpcyBjZXJ0aWZpY2F0ZSBpcyBzdHJpY3RseSBzdWJqZWN0ZWQgdG8g
dGhlIENFUlRVTSBDZXJ0aWZpY2F0aW9uIFByYWN0aWNlIFN0YXRlbWVudCAoQ1BT
KSBpbmNvcnBvcmF0ZWQgYnkgcmVmZXJlbmNlIGhlcmVpbiBhbmQgaW4gdGhlIHJl
cG9zaXRvcnkgYXQgaHR0cHM6Ly93d3cuY2VydHVtLnBsL3JlcG9zaXRvcnkuMB0G
A1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHREEGDAWhwRu+dpWgg4x
MTAuMjQ5LjIxOC44NjANBgkqhkiG9w0BAQsFAAOCAQEAkHpfSOKZ1m0sWzLKFuS6
4DpRCVKKpHqWREOi0vdSRtzkDnZ2E+D+ykYeoZHebg37AR1FFZ5ynk5J7YX875/2
Non1FEOTFbCRTezm9gc47fbVZsjyTN90wG68+HuRwoi7Yp5BZWNQbv6LCrvnb8yY
XNOoinGj8fPHQ4muf48hdw27Kc2oLOF+H+eO2sAZ3ppDe6zT4vwf4a425x+XU15j
Kizr7Y3x8gMdF9kIf70YhrQt9yH1y27h2NgsleUXQ/TvWOaoWNiPUeTSCIor1kpb
veYpzBtnJb3OHiVrUtDbOIi3ICqQPeFDcfMbtd/+52ZfT1Rb30LN9UK2DE1OXQPW
tA==
-----END CERTIFICATE-----`
)

func TestPost(t *testing.T) {
	url := "http://192.168.252.135:8015/api/v1/tools/cert_decode"
	data := map[string][]string{
		"type": []string{"paste"},
		"cert": []string{cert},
	}
	resp, err := DefaultClient.DoRequestWithForm(context.Background(), "POST", url, data)
	if err != nil {
		t.Fatal(err)
	}
	read(resp, t)

}

func read(resp *http.Response, t *testing.T) {
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(body))

}
