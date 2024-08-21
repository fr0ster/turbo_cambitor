package futures_rest_api_test

import (
	"testing"

	futures "github.com/fr0ster/turbo-cambitor/rest_api/binance/futures"
	"github.com/stretchr/testify/assert"
)

func TestListenKey(t *testing.T) {
	ra := futures.New(sign, true)
	listenKey, err := ra.ListenKey()
	assert.Nil(t, err)
	assert.NotEmpty(t, listenKey)
}

func TestKeepAliveListenKey(t *testing.T) {
	ra := futures.New(sign, true)
	listenKey, err := ra.KeepAliveListenKey()
	assert.Nil(t, err)
	assert.NotEmpty(t, listenKey)
}
