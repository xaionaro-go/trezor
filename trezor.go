// Package trezor implements master key encryption mechanism
// using open(hardware) device "Trezor One"
package trezor

import (
	"github.com/conejoninja/tesoro/pb/messages"

	"github.com/xaionaro-go/cryptoWallet"
	"github.com/xaionaro-go/cryptoWallet/vendors"
)

type CipherKeyValuer interface {
	CipherKeyValue(path string, isToEncrypt bool, keyName string, data, iv []byte, askOnEncode, askOnDecode bool) ([]byte, messages.MessageType)
}

type Trezor interface {
	cryptoWallet.Wallet
	CipherKeyValuer
}

type trezor struct {
	cryptoWallet.Wallet
}

func New() Trezor {
	result := cryptoWallet.Find(cryptoWallet.Filter{
		VendorID:   &[]uint16{vendors.GetVendorID("satoshilabs")}[0],
		ProductIDs: []uint16{1 /* Trezor One */},
	})
	if len(result) == 0 {
		return nil
	}
	return &trezor{Wallet: result[0]}
}

func (trezor *trezor) CipherKeyValue(path string, isToEncrypt bool, keyName string, data, iv []byte, askOnEncode, askOnDecode bool) ([]byte, messages.MessageType) {
	return trezor.Wallet.(CipherKeyValuer).CipherKeyValue(path, isToEncrypt, keyName, data, iv, askOnEncode, askOnDecode)
}
