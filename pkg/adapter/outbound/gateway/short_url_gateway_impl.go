package gateway

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/canmor/go_ms_clean_arch/pkg/app/usecase"
	"io"
	"net/http"
)

type shortURLGatewayImpl struct {
}

func NewShortURLGatewayImpl() usecase.ShortURL {
	return shortURLGatewayImpl{}
}

func (s shortURLGatewayImpl) Create(long string) (string, error) {
	client := http.Client{}
	out, _ := json.Marshal(map[string]interface{}{"url": long})
	resp, err := client.Post("https://api.s.url/short", "application/json", bytes.NewReader(out))
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("failed to create short URL, status: %s", resp.Status)
	}
	in, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var respBody map[string]string
	err = json.Unmarshal(in, &respBody)
	if err != nil {
		return "", err
	}
	return respBody["shortcut"], nil
}
