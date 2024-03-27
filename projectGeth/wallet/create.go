package wallet

import (
	// "fmt"
	"crypto/ecdsa"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"

	// "github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type Allkeys struct {
	Privatekey []byte
	Publickey  []byte
	Address    common.Address
	Decfile    *keystore.Key
	Pvk        *ecdsa.PrivateKey
}

var Mykeys Allkeys

func Create() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "sourav"
	// a, _ := key.NewAccount(password)
	// fmt.Println(a)
	key, _ := ioutil.ReadFile("./wallet/UTC--2024-03-18T19-49-07.729995661Z--66d79f9ea64c1704aad49e120bc3ba1b03a2adb0")
	dec, _ := keystore.DecryptKey(key, password)

	Mykeys.Privatekey = crypto.FromECDSA(dec.PrivateKey)
	Mykeys.Publickey = crypto.FromECDSAPub(&dec.PrivateKey.PublicKey)
	Mykeys.Address = crypto.PubkeyToAddress(dec.PrivateKey.PublicKey)
	Mykeys.Decfile = dec
	Mykeys.Pvk = dec.PrivateKey
}
