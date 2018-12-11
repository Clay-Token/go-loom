// +build evm

package cli

import (
	"errors"

	"github.com/loomnetwork/go-loom/crypto"

	"github.com/gogo/protobuf/proto"
	"github.com/loomnetwork/go-loom/auth"
)

func CallContract(defaultAddr string, method string, params proto.Message, result interface{}) error {
	var signerType string

	if TxFlags.PrivFile == "" && TxFlags.HsmConfigFile != "" {
		return errors.New("private key required to call contract")
	}

	var privKey crypto.PrivateKey
	var err error
	if TxFlags.HsmConfigFile != "" {
		privKey, err = crypto.LoadYubiHsmPrivKey(TxFlags.HsmConfigFile)
		if err != nil {
			return err
		}
		signerType = auth.SignerTypeYubiHsm
	} else {
		privKey, err = crypto.LoadECDSA(TxFlags.PrivFile)
		if err != nil {
			return err
		}
	}

	signer := auth.NewSigner(signerType, privKey)
	contract, err := contract(defaultAddr)
	if err != nil {
		return err
	}
	_, err = contract.Call(method, params, signer, result)
	return err
}
