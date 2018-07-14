// Package trezor implements master key encryption mechanism
// using open(hardware) device "Trezor One"
package trezor

import (
	"github.com/xaionaro-go/cryptoWallet"
	"github.com/xaionaro-go/cryptoWallet/interfaces"
	"github.com/xaionaro-go/cryptoWallet/vendors"
)

func New() cryptoWalletInterfaces.Trezor {
	result := cryptoWallet.Find(cryptoWallet.Filter{
		VendorID:   &[]uint16{vendors.GetVendorID("satoshilabs")}[0],
		ProductIDs: []uint16{1 /* Trezor One */},
	})
	if len(result) == 0 {
		result := cryptoWallet.Find(cryptoWallet.Filter{
			VendorID:   &[]uint16{vendors.GetVendorID("interbiometrics")}[0],
			ProductIDs: []uint16{0x53C1 /* Trezor T */},
		})
	}
	if len(result) == 0 {
		return nil
	}
	return result[0].(cryptoWalletInterfaces.Trezor)
}
