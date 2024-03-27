package mycontract

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"

	abi "nft/abi"
	"nft/wallet"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Myclient *ethclient.Client
var Contract *abi.Mycontract
var gasLimit = 300000

func CollectEverything(addr common.Address) (uint64, *big.Int, *big.Int, error) {
	nonce, ner := Myclient.PendingNonceAt(context.TODO(), addr)
	if ner != nil {
		log.Println(ner)
		return 0, nil, nil, ner

	}
	gasPrice, ger := Myclient.SuggestGasPrice(context.TODO())
	if ger != nil {
		log.Println(ger)
		return 0, nil, nil, ger

	}
	chainId, cer := Myclient.NetworkID(context.TODO())
	if cer != nil {
		log.Println(cer)
		return 0, nil, nil, cer
	}
	return nonce, gasPrice, chainId, nil

}
func GetTokenIdValue(id *big.Int, addr common.Address) (string, error) {
	// add := common.HexToAddress(addr)
	value, err := Contract.GetIdValue(&bind.CallOpts{From: addr}, id)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return value, nil

}
func GetTotalTokenOfaWallet(user common.Address) (*big.Int, error) {
	value, err := Contract.GetTotalBalanceByAddress(&bind.CallOpts{From: user})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return value, nil
}
func GetAddressOfaTokenId(id *big.Int, user common.Address) (string, error) {
	value, err := Contract.GetAddressById(&bind.CallOpts{From: user}, id)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return value.Hex(), nil

}
func GetTotalTokenSupply(user common.Address) (*big.Int, error) {
	value, err := Contract.TokenCounter(&bind.CallOpts{From: user})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return value, nil
}

func GetTokenName(user common.Address) (string, error) {
	name, err := Contract.Name(&bind.CallOpts{From: user})
	if err != nil {
		log.Println(err)
		return "", err
	}
	return name, nil
}
func GetTokenSymbol(user common.Address) (string, error) {
	symbol, err := Contract.Symbol(&bind.CallOpts{From: user})
	if err != nil {
		log.Println(err)
		return "", nil
	}
	return symbol, nil

}
func GetContractOwner(user common.Address) (string, error) {
	owner, err := Contract.Owner(&bind.CallOpts{From: user})
	if err != nil {
		log.Println(err)
		return "", err
	}
	return owner.Hex(), nil

}

func CallTokenMint(user *ecdsa.PrivateKey, name string) (*types.Transaction, error) {

	// pvk, pverr := crypto.HexToECDSA(user)

	// if pverr != nil {
	// 	log.Println(pverr)
	// 	return nil, pverr
	// }
	addr := crypto.PubkeyToAddress(user.PublicKey)

	nonce, gasPrice, chainId, oterr := CollectEverything(addr)
	if oterr != nil {
		log.Println(oterr)
		return nil, oterr
	}

	auth, err := bind.NewKeyedTransactorWithChainID(user, chainId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	auth.GasLimit = uint64(gasLimit)
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))
	value, terr := Contract.MintLizToken(auth, name)
	if terr != nil {
		log.Println(terr)
		return nil, terr
	}
	return value, nil
}
func ChangeTokenOwner(user *ecdsa.PrivateKey, id *big.Int) (*types.Transaction, error) {
	// pvk, err := crypto.HexToECDSA(user)
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }
	pvk := wallet.Mykeys.Decfile.PrivateKey
	addr := crypto.PubkeyToAddress(pvk.PublicKey)
	nonce, gasPrice, chainId, otherr := CollectEverything(addr)
	if otherr != nil {
		log.Println(otherr)
		return nil, otherr
	}
	auth, auerr := bind.NewKeyedTransactorWithChainID(pvk, chainId)
	if auerr != nil {
		log.Println(auerr)
		return nil, auerr
	}
	auth.GasLimit = uint64(gasLimit)
	auth.GasPrice = gasPrice
	auth.Nonce = big.NewInt(int64(nonce))
	value, verr := Contract.ChangeTokenOwner(auth, id)
	if verr != nil {
		log.Println(verr)
		return nil, verr
	}
	return value, nil

}
