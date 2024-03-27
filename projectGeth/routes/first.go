package routes

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	contract "nft/mycontract"
	"nft/wallet"

	"github.com/ethereum/go-ethereum/common"
)

func CallContractTokenName(w http.ResponseWriter, r *http.Request) {
	var addr common.Address
	json.NewDecoder(r.Body).Decode(&addr)
	value, err := contract.GetTokenName(addr)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(value)

}
func CallContractTokenSymbol(w http.ResponseWriter, r *http.Request) {
	var addr common.Address
	json.NewDecoder(r.Body).Decode(&addr)
	name, err := contract.GetTokenSymbol(addr)
	if err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(name)
}

func CallContractTotalSupply(w http.ResponseWriter, r *http.Request) {
	var addr common.Address
	json.NewDecoder(r.Body).Decode(&addr)
	value, err := contract.GetTotalTokenSupply(addr)
	if err != nil {
		log.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(value)
}
func CallContractTokenValueById(w http.ResponseWriter, r *http.Request) {
	type token struct {
		Addr common.Address
		Id   *big.Int
	}
	var myval token
	json.NewDecoder(r.Body).Decode(&myval)
	value, err := contract.GetTokenIdValue(myval.Id, myval.Addr)
	if err != nil {
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(value)
}
func CallContractTokenAddressById(w http.ResponseWriter, r *http.Request) {
	type token struct {
		Addr common.Address
		Id   *big.Int
	}
	var mytoken token
	json.NewDecoder(r.Body).Decode(&mytoken)
	value, err := contract.GetAddressOfaTokenId(mytoken.Id, mytoken.Addr)
	if err != nil {
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(value)
}
func CallContarctTokenMint(w http.ResponseWriter, r *http.Request) {
	type token struct {
		TokenName string
	}
	pvk := wallet.Mykeys.Decfile.PrivateKey
	var mytoken token
	json.NewDecoder(r.Body).Decode(&mytoken)
	value, err := contract.CallTokenMint(pvk, mytoken.TokenName)
	if err != nil {
		log.Println("in token mint :", err)
		return
	}
	json.NewEncoder(w).Encode(value)
}
func CallContractChangeTokenOwner(w http.ResponseWriter, r *http.Request) {
	type token struct {
		Id *big.Int
	}
	var mytoken token
	json.NewDecoder(r.Body).Decode(&mytoken)
	pvk := wallet.Mykeys.Decfile.PrivateKey
	value, err := contract.ChangeTokenOwner(pvk, mytoken.Id)
	if err != nil {
		log.Println(err)
		return
	}

	json.NewEncoder(w).Encode(value)
}
func CallContractWalletTotalToken(w http.ResponseWriter, r *http.Request) {
	var addr common.Address
	json.NewDecoder(r.Body).Decode(&addr)
	value, err := contract.GetTotalTokenOfaWallet(addr)
	if err != nil {
		log.Println(err)

	}
	json.NewEncoder(w).Encode(value)
}
func CallContractOwner(w http.ResponseWriter, r *http.Request) {
	var addr common.Address
	json.NewDecoder(r.Body).Decode(&addr)
	value, err := contract.GetContractOwner(addr)
	if err != nil {
		log.Println(err)
		return
	}
	json.NewEncoder(w).Encode(value)
}
