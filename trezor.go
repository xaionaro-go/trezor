// Package trezor implements master key encryption mechanism
// using open(hardware) device "Trezor One"
package trezor

import (
	"github.com/xaionaro-go/cryptoWallet"
	"github.com/xaionaro-go/cryptoWallet/interfaces"
)

func New() cryptoWalletInterfaces.Trezor {
	result := cryptoWallet.Find(cryptoWallet.Filter{})
	if len(result) == 0 {
		return nil
	}
	trezorInstance, ok := result[0].(cryptoWalletInterfaces.Trezor)
	if !ok {
		return nil
	}
	return trezorInstance
}
