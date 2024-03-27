package main

import (
	"log"
	"net/http"

	abi "nft/abi"
	user "nft/mycontract"
	routes "nft/routes"
	"nft/wallet"

	"github.com/ethereum/go-ethereum/common"
	// "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	log.Println("servcer start")
	client, err := ethclient.Dial("https://eth-sepolia.g.alchemy.com/v2/lh4oR7wrbDIA2c2qBJMo5g5NDk_YtuKB")
	if err != nil {
		log.Panic("in sep", err)
	}
	defer client.Close()
	addr := common.HexToAddress("0x149406be6ED4B534a26B14Ff0109394A70D9A2D9")
	contract, conerr := abi.NewMycontract(addr, client)
	if conerr != nil {
		log.Println(conerr)
		return
	}
	user.Contract = contract
	user.Myclient = client
	wallet.Create()

	address := wallet.Mykeys.Address

	log.Println(address)

	router.HandleFunc("/token-name", routes.CallContractTokenName).Methods("POST")
	router.HandleFunc("/token-symbol", routes.CallContractTokenSymbol).Methods("POST")
	router.HandleFunc("/total-supply", routes.CallContractTotalSupply).Methods("POST")
	router.HandleFunc("/token-value-by-id", routes.CallContractTokenValueById).Methods("POST")
	router.HandleFunc("/token-address-by-id", routes.CallContractTokenAddressById).Methods("POST")
	router.HandleFunc("/mint-token", routes.CallContarctTokenMint).Methods("POST")
	router.HandleFunc("/change-token-owner", routes.CallContractChangeTokenOwner).Methods("POST")
	router.HandleFunc("/user-total-token", routes.CallContractWalletTotalToken).Methods("POST")
	router.HandleFunc("/owner-of-the-contract", routes.CallContractOwner).Methods("POST")

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(router)))

	// http.ListenAndServe(":8080", router)

}
