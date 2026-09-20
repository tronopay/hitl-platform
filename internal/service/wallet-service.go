package service

import (
	"errors"

	"github.com/fbsobreira/gotron-sdk/pkg/address"
	"github.com/fbsobreira/gotron-sdk/pkg/keys"
)

type Wallet struct {
	Address    string
	PrivateKey string
}

type WalletService struct {
}

func NewWalletService() *WalletService {
	return &WalletService{}
}

func (s *WalletService) GenerateWallets(count int) ([]Wallet, error) {
	var wallets []Wallet

	for range count {
		if wallet, err := s.generate(); err == nil {
			if wallet != nil {
				wallets = append(wallets, *wallet)
			} else {
				return nil, errors.New("wallet is empty")
			}
		} else {
			return nil, err
		}
	}

	return wallets, nil
}

func (s *WalletService) generate() (*Wallet, error) {
	privateKey, err := keys.GenerateKey()
	if err != nil {
		return nil, err
	}

	addr := address.BTCECPrivkeyToAddress(privateKey)
	walletAddress := address.Address(addr).String()
	walletKey := privateKey.Key.String()

	return &Wallet{
		Address:    walletAddress,
		PrivateKey: walletKey,
	}, nil
}
