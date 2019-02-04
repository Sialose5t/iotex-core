// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package address

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/iotexproject/iotex-core/address/bech32"
	"github.com/iotexproject/iotex-core/pkg/hash"
	"github.com/iotexproject/iotex-core/pkg/log"
)

// V1 is a singleton and defines V1 address metadata
var V1 = v1{
	AddressLength: 20,
}

type v1 struct {
	// AddressLength indicates the byte length of an address
	AddressLength int
}

// New constructs an address struct
func (v *v1) New(pkHash hash.PKHash) *AddrV1 {
	return &AddrV1{
		pkHash: pkHash,
	}
}

// Bech32ToAddress decodes an encoded address string into an address struct
func (v *v1) Bech32ToAddress(encodedAddr string) (*AddrV1, error) {
	payload, err := v.decodeBech32(encodedAddr)
	if err != nil {
		return nil, err
	}
	return v.BytesToAddress(payload)
}

// BytesToAddress converts a byte array into an address struct
func (v *v1) BytesToAddress(bytes []byte) (*AddrV1, error) {
	if len(bytes) != v.AddressLength {
		return nil, errors.Wrapf(ErrInvalidAddr, "invalid address length in bytes: %d", len(bytes))
	}
	var pkHash hash.PKHash
	copy(pkHash[:], bytes)
	return &AddrV1{
		pkHash: pkHash,
	}, nil
}

// Bech32ToPKHash returns the public key hash from an encoded address
func (v *v1) Bech32ToPKHash(encodedAddr string) (hash.PKHash, error) {
	addr, err := v.Bech32ToAddress(encodedAddr)
	if err != nil {
		return hash.ZeroPKHash, errors.Wrap(err, "failed to decode encoded address to address")
	}
	return addr.PublicKeyHash(), nil
}

func (v *v1) decodeBech32(encodedAddr string) ([]byte, error) {
	hrp, grouped, err := bech32.Decode(encodedAddr)
	if hrp != prefix() {
		return nil, errors.Wrapf(err, "hrp %s and address prefix %s don't match", hrp, prefix())
	}
	// Group the payload into 8 bit groups.
	payload, err := bech32.ConvertBits(grouped[:], 5, 8, false)
	if err != nil {
		return nil, errors.Wrapf(err, "error when converting 5 bit groups into the payload")
	}
	return payload, nil
}

// AddrV1 is V1 address format to be used on IoTeX blockchain and subchains. It is composed of
// 20 bytes: hash derived from the the public key:
type AddrV1 struct {
	pkHash hash.PKHash
}

// Bech32 encodes an address struct into a a Bech32 encoded address string
// The encoded address string will start with "io" for mainnet, and with "it" for testnet
func (addr *AddrV1) Bech32() string {
	payload := addr.pkHash[:]
	// Group the payload into 5 bit groups.
	grouped, err := bech32.ConvertBits(payload, 8, 5, true)
	if err != nil {
		log.L().Panic("Error when grouping the payload into 5 bit groups.", zap.Error(err))
		return ""
	}
	encodedAddr, err := bech32.Encode(prefix(), grouped)
	if err != nil {
		log.L().Panic("Error when encoding bytes into a base32 string.", zap.Error(err))
		return ""
	}
	return encodedAddr
}

// Bytes converts an address struct into a byte array
func (addr *AddrV1) Bytes() []byte {
	return addr.pkHash[:]
}

// Payload returns the payload, which is the public key hash
func (addr *AddrV1) Payload() []byte {
	return addr.pkHash[:]
}

// PublicKeyHash returns the public key hash
func (addr *AddrV1) PublicKeyHash() hash.PKHash {
	return addr.pkHash
}
