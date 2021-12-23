// ***********************************************************************************************
// ***                               G O L A N D   S T U D I O S                               ***
// ***********************************************************************************************
// * Auth: ColeCai
// * Date: 2021/12/23 10:06:49
// * File: des.go
// * Proj: go-tools
// * Pack: crypto
// * Ides: GoLand
// *----------------------------------------------------------------------------------------------
// * Functions:
// * 	-DesCBCEncrypt(decrypted, desKey, iv []byte, padding PaddingT) ([]byte, error)
// * 	-DesCBCDecrypt(encrypted, desKey, iv []byte, padding PaddingT) ([]byte, error)
// * - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

package crypto

import "crypto/des"

// ***********************************************************************************************
// * SUMMARY:
// * WARNING:
// * 	-deskey and iv length must equal 8. if not DesCBCEncrypt will panic.
// * HISTORY:
// *    -create: 2021/12/23 10:06:54 ColeCai.
// ***********************************************************************************************
func DesCBCEncrypt(decrypted []byte, desKey, iv []byte, padding PaddingT) ([]byte, error) {
	ciphers, err := des.NewCipher(desKey)
	if err != nil {
		return nil, err
	}
	return CBCEncrypt(ciphers, decrypted, iv, padding)
}

// ***********************************************************************************************
// * SUMMARY:
// * WARNING:
// * 	-deskey and iv length must equal 8. if not DesCBCDecrypt will panic.
// * HISTORY:
// *    -create: 2021/12/23 10:09:12 ColeCai.
// ***********************************************************************************************
func DesCBCDecrypt(encrypted []byte, desKey, iv []byte, padding PaddingT) ([]byte, error) {
	ciphers, err := des.NewCipher(desKey)
	if err != nil {
		return nil, err
	}
	return CBCDecrypt(ciphers, encrypted, iv, padding)
}
