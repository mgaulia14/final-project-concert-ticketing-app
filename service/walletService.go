package service

import (
	"errors"
	"final-project-ticketing-api/database"
	"final-project-ticketing-api/repository"
	"final-project-ticketing-api/structs"
)

func GetWalletInfoByCustomerId(customerId int) (structs.Wallet, error) {
	var result structs.Wallet
	err, result := repository.GetWalletInfo(database.DBConnection, customerId)
	if err != nil {
		return result, err
	}
	return result, nil
}

func TopUpWallet(request structs.WalletTopUp) (structs.Wallet, []error) {
	var result structs.Wallet
	var err []error
	err1, wallet := repository.GetWalletByAccountNumber(database.DBConnection, request.AccountNumber)
	if err1 != nil {
		err = append(err, err1)
		return result, err
	}
	wallet, err = prepareRequestWallet(request, wallet)
	if err != nil {
		return wallet, err
	}
	wallet, err = repository.TopUpBalance(database.DBConnection, wallet)
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func prepareRequestWallet(request structs.WalletTopUp, wallet structs.Wallet) (structs.Wallet, []error) {
	request, err := validateRequestWallet(request)
	if err != nil {
		return wallet, err
	}

	// calculate balance
	wallet.Balance = wallet.Balance + request.Balance
	return wallet, nil
}

func validateRequestWallet(request structs.WalletTopUp) (structs.WalletTopUp, []error) {
	var err []error
	if request.Balance > 20000000 {
		err = append(err, errors.New("top up max 20.000.000"))
	}
	if len(err) > 0 {
		return request, err
	}
	return request, nil
}
