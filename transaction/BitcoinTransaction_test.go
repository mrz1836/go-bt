package transaction

import (
	"encoding/hex"
	"fmt"
	"testing"

	"bitbucket.org/simon_ordish/cryptolib"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
)

func TestGetVersion(t *testing.T) {
	const tx = "01000000014c6ec863cf3e0284b407a1a1b8138c76f98280812cb9653231f385a0305fc76f010000006b483045022100f01c1a1679c9437398d691c8497f278fa2d615efc05115688bf2c3335b45c88602201b54437e54fb53bc50545de44ea8c64e9e583952771fcc663c8687dc2638f7854121037e87bbd3b680748a74372640628a8f32d3a841ceeef6f75626ab030c1a04824fffffffff021d784500000000001976a914e9b62e25d4c6f97287dfe62f8063b79a9638c84688ac60d64f00000000001976a914bb4bca2306df66d72c6e44a470873484d8808b8888ac00000000"
	bt, err := NewFromString(tx, false)
	if err != nil {
		t.Error(err)
		return
	}

	res := bt.Version()
	if res != 1 {
		t.Errorf("Expecting 1, got %d", res)
	}
}

func TestElectumTX(t *testing.T) {
	const h = "01000000012284573666f6a0ac538c0559db6a2f378e05808d8e07530ad98123304fd4d8a200000000fb00483045022100e42e7eea232fbe0517e954430115ca7363e47ff00d64657b90d0501f6c696d2b02204f5ba943499d639dccadb70e5573459adca2a74613fcfa860cab742bc43e65914101ff4cad524c53ff0488b21e000000000000000000362f7a9030543db8751401c387d6a71e870f1895b3a62569d455e8ee5f5f5e5f03036624c6df96984db6b4e625b6707c017eb0e0d137cd13a0c989bfa77a4473fd000000004c53ff0488b21e0000000000000000008b20425398995f3c866ea6ce5c1828a516b007379cf97b136bffbdc86f75df14036454bad23b019eae34f10aff8b8d6d8deb18cb31354e5a169ee09d8a4560e8250000000052aefefffffffe709700000000000410270000000000001976a914a84a278b361de016ca93cfa91624db1f926e1b9d88ac10270000000000001976a914a84a278b361de016ca93cfa91624db1f926e1b9d88ac10270000000000001976a914a84a278b361de016ca93cfa91624db1f926e1b9d88ac00fa96000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b18722c40800"

	tx, err := NewFromString(h, true)

	t.Logf("%s\n%s\n%v\n", h, hex.EncodeToString(tx.Hex()), err)

}
func TestElectumTXFullySigned(t *testing.T) {
	const h = "01000000012284573666f6a0ac538c0559db6a2f378e05808d8e07530ad98123304fd4d8a200000000da00483045022100a799bf0a21985a84bedf502c42dd47623c8b7e6f42ee3cf8149370fb0c39dd05022038e79e52cff77769aee76633d80e70eeca1d3186e05f4d6dad9840e196173b62414730440220122472092175cc9a0af67c48e002317c4056b55fbe66d85db616db1bcd12c8f202206b4487a8ac96ac17db116eb41de19d717527e2e3e1f1f908f5be937dbf0e47f24147522103d10369cb9603521e3b2b2f13b71d356e9465867c7c79233e58d85f82dec241942103f2538c34a0991dcbcae32c56c1158822c88a4149e0549363dd9c541a6455114552aefeffffff03a0860100000000001976a914befa24db1fd3b793c37569e68a053b34e40c52c688aca0860100000000001976a914befa24db1fd3b793c37569e68a053b34e40c52c688acf06194000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b187a6c80800"

	tx, err := NewFromString(h, true)

	t.Logf("%s\n%s\n%v\n", h, hex.EncodeToString(tx.Hex()), err)

}
func TestElectumUnsignedTX(t *testing.T) {
	const h = "01000000012284573666f6a0ac538c0559db6a2f378e05808d8e07530ad98123304fd4d8a200000000b40001ff01ff4cad524c53ff0488b21e000000000000000000362f7a9030543db8751401c387d6a71e870f1895b3a62569d455e8ee5f5f5e5f03036624c6df96984db6b4e625b6707c017eb0e0d137cd13a0c989bfa77a4473fd000000004c53ff0488b21e0000000000000000008b20425398995f3c866ea6ce5c1828a516b007379cf97b136bffbdc86f75df14036454bad23b019eae34f10aff8b8d6d8deb18cb31354e5a169ee09d8a4560e8250000000052aefefffffffe7097000000000003a0860100000000001976a914a84a278b361de016ca93cfa91624db1f926e1b9d88aca0860100000000001976a914a84a278b361de016ca93cfa91624db1f926e1b9d88acf06194000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b187f1c60800"

	tx, err := NewFromString(h, true)

	t.Logf("%s\n%s\n%v\n", h, hex.EncodeToString(tx.Hex()), err)

}

func TestElectumAll(t *testing.T) {
	// const h0 = "01000000012284573666f6a0ac538c0559db6a2f378e05808d8e07530ad98123304fd4d8a200000000b40001ff01ff4cad524c53ff0488b21e000000000000000000362f7a9030543db8751401c387d6a71e870f1895b3a62569d455e8ee5f5f5e5f03036624c6df96984db6b4e625b6707c017eb0e0d137cd13a0c989bfa77a4473fd000000004c53ff0488b21e0000000000000000008b20425398995f3c866ea6ce5c1828a516b007379cf97b136bffbdc86f75df14036454bad23b019eae34f10aff8b8d6d8deb18cb31354e5a169ee09d8a4560e8250000000052aefefffffffe7097000000000003a0860100000000001976a914a84a278b361de016ca93cfa91624db1f926e1b9d88aca0860100000000001976a914a84a278b361de016ca93cfa91624db1f926e1b9d88acf06194000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b187f1c60800"
	const h1 = "01000000012284573666f6a0ac538c0559db6a2f378e05808d8e07530ad98123304fd4d8a200000000fb00483045022100ebe01fb1f5e8b188bbe6d408657e6d6eca2330e7ed3aff4d4b3e21669aac96a90220604f1727cfe54c1ae81d960f8c5a4976716e73753c3905dc380274ccb9155b1f4101ff4cad524c53ff0488b21e000000000000000000362f7a9030543db8751401c387d6a71e870f1895b3a62569d455e8ee5f5f5e5f03036624c6df96984db6b4e625b6707c017eb0e0d137cd13a0c989bfa77a4473fd000000004c53ff0488b21e0000000000000000008b20425398995f3c866ea6ce5c1828a516b007379cf97b136bffbdc86f75df14036454bad23b019eae34f10aff8b8d6d8deb18cb31354e5a169ee09d8a4560e8250000000052aefefffffffe7097000000000003a0860100000000001976a914befa24db1fd3b793c37569e68a053b34e40c52c688aca0860100000000001976a914befa24db1fd3b793c37569e68a053b34e40c52c688acf06194000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b187a7c80800"
	const h2 = "01000000012284573666f6a0ac538c0559db6a2f378e05808d8e07530ad98123304fd4d8a200000000da00483045022100ebe01fb1f5e8b188bbe6d408657e6d6eca2330e7ed3aff4d4b3e21669aac96a90220604f1727cfe54c1ae81d960f8c5a4976716e73753c3905dc380274ccb9155b1f414730440220340908f1e828085a64f31aacd6b883f522ce381c5431bd8908eb4c4a67bd802802206aa84488fcb4fcd73e4b007c4ac30dfab2cc23358915b16396b386a615d97ac04147522103d10369cb9603521e3b2b2f13b71d356e9465867c7c79233e58d85f82dec241942103f2538c34a0991dcbcae32c56c1158822c88a4149e0549363dd9c541a6455114552aefeffffff03a0860100000000001976a914befa24db1fd3b793c37569e68a053b34e40c52c688aca0860100000000001976a914befa24db1fd3b793c37569e68a053b34e40c52c688acf06194000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b187a7c80800"

	// h0 is not signed and we can ignore.
	// h1 is signed by simon's key but need liam's signature to be complete
	// h2 is a value bitcoin transaction which is signed.

	const expected = h2

	tx, err := NewFromString(h1, true)
	if err != nil {
		t.Error(err)
		return
	}

	t.Log(tx)
}

func TestElectumSignedTX(t *testing.T) {
	const h = "01000000012284573666f6a0ac538c0559db6a2f378e05808d8e07530ad98123304fd4d8a200000000fa00473044022041682b268531cf6209577deae34b92fdc83d9ef6e3abc190d4952e927761efd502201696256fba4dd6b05e44ed3871abbd1bc11356aea5ddc36816ca779f68cca6fa4101ff4cad524c53ff0488b21e000000000000000000362f7a9030543db8751401c387d6a71e870f1895b3a62569d455e8ee5f5f5e5f03036624c6df96984db6b4e625b6707c017eb0e0d137cd13a0c989bfa77a4473fd000000004c53ff0488b21e0000000000000000008b20425398995f3c866ea6ce5c1828a516b007379cf97b136bffbdc86f75df14036454bad23b019eae34f10aff8b8d6d8deb18cb31354e5a169ee09d8a4560e8250000000052aefefffffffe7097000000000003a0860100000000001976a914a84a278b361de016ca93cfa91624db1f926e1b9d88aca0860100000000001976a914a84a278b361de016ca93cfa91624db1f926e1b9d88acf06194000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b187f1c60800"

	tx, err := NewFromString(h, true)

	t.Logf("%s\n%s\n%v\n", h, hex.EncodeToString(tx.Hex()), err)

}

func TestSign(t *testing.T) {
	// const h0 = "01000000012284573666f6a0ac538c0559db6a2f378e05808d8e07530ad98123304fd4d8a200000000b40001ff01ff4cad524c53ff0488b21e000000000000000000362f7a9030543db8751401c387d6a71e870f1895b3a62569d455e8ee5f5f5e5f03036624c6df96984db6b4e625b6707c017eb0e0d137cd13a0c989bfa77a4473fd000000004c53ff0488b21e0000000000000000008b20425398995f3c866ea6ce5c1828a516b007379cf97b136bffbdc86f75df14036454bad23b019eae34f10aff8b8d6d8deb18cb31354e5a169ee09d8a4560e8250000000052aefefffffffe7097000000000003a0860100000000001976a914a00de33a188cfc48ccc2cb3f79401cc03a5f05dd88aca0860100000000001976a914a00de33a188cfc48ccc2cb3f79401cc03a5f05dd88acf06194000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b18737c90800"
	const h1 = "01000000012284573666f6a0ac538c0559db6a2f378e05808d8e07530ad98123304fd4d8a200000000fa004730440220682543040b1ccb9fa0e64d8f17836012012a7f4e69846bfd5907ded0481b629102206dd0776b3d8dbf1c245dccb0924ec730f506efb5aa6a2108e90df114d2652a4a4101ff4cad524c53ff0488b21e000000000000000000362f7a9030543db8751401c387d6a71e870f1895b3a62569d455e8ee5f5f5e5f03036624c6df96984db6b4e625b6707c017eb0e0d137cd13a0c989bfa77a4473fd000000004c53ff0488b21e0000000000000000008b20425398995f3c866ea6ce5c1828a516b007379cf97b136bffbdc86f75df14036454bad23b019eae34f10aff8b8d6d8deb18cb31354e5a169ee09d8a4560e8250000000052aefefffffffe7097000000000003a0860100000000001976a914a00de33a188cfc48ccc2cb3f79401cc03a5f05dd88aca0860100000000001976a914a00de33a188cfc48ccc2cb3f79401cc03a5f05dd88acf06194000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b18737c90800"
	const h2 = "01000000012284573666f6a0ac538c0559db6a2f378e05808d8e07530ad98123304fd4d8a200000000da004730440220682543040b1ccb9fa0e64d8f17836012012a7f4e69846bfd5907ded0481b629102206dd0776b3d8dbf1c245dccb0924ec730f506efb5aa6a2108e90df114d2652a4a41483045022100da2fd4319259ef9a8a5afa05fd2c401da039573d86e533343dc44e861318495f022051e381907be2856654d348b7e073d6024312da40e48fcb30c3b90628ccd699774147522103d10369cb9603521e3b2b2f13b71d356e9465867c7c79233e58d85f82dec241942103f2538c34a0991dcbcae32c56c1158822c88a4149e0549363dd9c541a6455114552aefeffffff03a0860100000000001976a914a00de33a188cfc48ccc2cb3f79401cc03a5f05dd88aca0860100000000001976a914a00de33a188cfc48ccc2cb3f79401cc03a5f05dd88acf06194000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b18737c90800"
	/*
				Serialized Pre-image:
				010000002f7eb8cbf174a619fd4c7671ad332330d3453287e5374c0301b9828a21974cb418606b350cd8bf565266bc352f0caddcf01e8fa789dd8a15386327cf8cabe1982284573666f6a0ac538c0559db6a2f378e05808d8e07530ad98123304fd4d8a20000000047522103d10369cb9603521e3b2b2f13b71d356e9465867c7c79233e58d85f82dec241942103f2538c34a0991dcbcae32c56c1158822c88a4149e0549363dd9c541a6455114552aefe70970000000000feffffffea9a6871bbcb55ed435ec86ee22166b893b2947f4089b3c35e343c23cc85593637c9080041000000
		  	010000000117a9140e95261082d65c384a6106f114474bc0784ba67e8703a0860100000000001976a914a00de33a188cfc48ccc2cb3f79401cc03a5f05dd88aca0860100000000001976a914a00de33a188cfc48ccc2cb3f79401cc03a5f05dd88acf06194000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b18737c9080001000000
	*/
	/*
		   da = 436 chars
		   00
			 47
			 30440220682543040b1ccb9fa0e64d8f17836012012a7f4e69846bfd5907ded0481b629102206dd0776b3d8dbf1c245dccb0924ec730f506efb5aa6a2108e90df114d2652a4a41
			 48
			 3045022100da2fd4319259ef9a8a5afa05fd2c401da039573d86e533343dc44e861318495f022051e381907be2856654d348b7e073d6024312da40e48fcb30c3b90628ccd6997741
			 47
			 52
			 21
			 03d10369cb9603521e3b2b2f13b71d356e9465867c7c79233e58d85f82dec24194
			 21
			 03f2538c34a0991dcbcae32c56c1158822c88a4149e0549363dd9c541a64551145
			 52
			 ae

		   feffffff
		   03
		   a0860100000000001976a914a00de33a188cfc48ccc2cb3f79401cc03a5f05dd88aca0860100000000001976a914a00de33a188cfc48ccc2cb3f79401cc03a5f05dd88acf06194000000000017a914289225b74059ef9b1024b9f04a294d32e5cf97b18737c90800"

	*/

	const liamxprv = "xprv9s21ZrQH143K3ShHqGb2ago1pjts78QvhAtYUbe1kPraUtjkxaftf28Pc6LdHKBAzi2jAH3EhQWgibbJxMFDW1yS8ZrPy172LEvwddxV55D"

	privKey, _ := cryptolib.NewPrivateKey(liamxprv)

	tx, err := NewFromString(h1, true)
	if err != nil {
		t.Error(err)
	}
	tx.Sign(privKey)

	// t.Logf("%s\n%s\n%v\n", h, hex.EncodeToString(tx.Hex()), err)

}

func TestConvertXPriv(t *testing.T) {
	const xprv = "xprv9s21ZrQH143K2beTKhLXFRWWFwH8jkwUssjk3SVTiApgmge7kNC3jhVc4NgHW8PhW2y7BCDErqnKpKuyQMjqSePPJooPJowAz5BVLThsv6c"
	const expected = "5f86e4023a4e94f00463f81b70ff951f83f896a0a3e6ed89cf163c152f954f8b"

	r, _ := cryptolib.NewPrivateKey(xprv)

	t.Logf("%x", r.PrivateKey)
}

func TestSignRedeemScript(t *testing.T) {
	var redeemScript, _ = hex.DecodeString("524c53ff0488b21e000000000000000000362f7a9030543db8751401c387d6a71e870f1895b3a62569d455e8ee5f5f5e5f03036624c6df96984db6b4e625b6707c017eb0e0d137cd13a0c989bfa77a4473fd000000004c53ff0488b21e0000000000000000008b20425398995f3c866ea6ce5c1828a516b007379cf97b136bffbdc86f75df14036454bad23b019eae34f10aff8b8d6d8deb18cb31354e5a169ee09d8a4560e8250000000052ae")
	const expected = "3044022041682b268531cf6209577deae34b92fdc83d9ef6e3abc190d4952e927761efd502201696256fba4dd6b05e44ed3871abbd1bc11356aea5ddc36816ca779f68cca6fa"

	const xprv = "xprv9s21ZrQH143K2beTKhLXFRWWFwH8jkwUssjk3SVTiApgmge7kNC3jhVc4NgHW8PhW2y7BCDErqnKpKuyQMjqSePPJooPJowAz5BVLThsv6c"
	const privHex = "5f86e4023a4e94f00463f81b70ff951f83f896a0a3e6ed89cf163c152f954f8b"

	pkBytes, err := hex.DecodeString(privHex)
	if err != nil {
		fmt.Println(err)
		return
	}
	privKey, pubKey := btcec.PrivKeyFromBytes(btcec.S256(), pkBytes)

	// Sign a message using the private key.
	messageHash := chainhash.DoubleHashB(redeemScript)
	signature, err := privKey.Sign(messageHash)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Serialize and display the signature.
	fmt.Printf("Serialized Signature: %x\n", signature.Serialize())

	// Verify the signature for the message using the public key.
	verified := signature.Verify(messageHash, pubKey)
	fmt.Printf("Signature Verified? %v\n", verified)
}

func TestIsCoinbase(t *testing.T) {
	const coinbase = "01000000010000000000000000000000000000000000000000000000000000000000000000ffffffff4303bfea07322f53696d6f6e204f726469736820616e642053747561727420467265656d616e206d61646520746869732068617070656e2f9a46434790f7dbdea3430000ffffffff018a08ac4a000000001976a9148bf10d323ac757268eb715e613cb8e8e1d1793aa88ac00000000"
	bt1, err := NewFromString(coinbase, false)
	if err != nil {
		t.Error(err)
		return
	}

	cb1 := bt1.IsCoinbase()
	if cb1 == false {
		t.Errorf("Expecting true, got %t", cb1)
	}

	const tx = "01000000014c6ec863cf3e0284b407a1a1b8138c76f98280812cb9653231f385a0305fc76f010000006b483045022100f01c1a1679c9437398d691c8497f278fa2d615efc05115688bf2c3335b45c88602201b54437e54fb53bc50545de44ea8c64e9e583952771fcc663c8687dc2638f7854121037e87bbd3b680748a74372640628a8f32d3a841ceeef6f75626ab030c1a04824fffffffff021d784500000000001976a914e9b62e25d4c6f97287dfe62f8063b79a9638c84688ac60d64f00000000001976a914bb4bca2306df66d72c6e44a470873484d8808b8888ac00000000"
	bt2, err := NewFromString(tx, false)
	if err != nil {
		t.Error(err)
		return
	}

	cb2 := bt2.IsCoinbase()
	if cb2 == true {
		t.Errorf("Expecting false, got %t", cb2)
	}
}

/*
48
30
45
02
21
00f4de422896e461da647b21d800a4ca9ace98dbd08c2dc9b8e049c93197c314f5
02
20
68836c3dfa6650ebeff73b1e3caa8761cd107ed13d6cc713856ebde3f874dd41
41

21
02aea77c449eeeef2746562e56ad053202755f9844276e3f0c684f9d59cdb9458d
ac OP_CHECKSIG

*/
// func TestMyTransaction(t *testing.T) {
// 	fromTx, err := NewFromString("02000000010000000000000000000000000000000000000000000000000000000000000000ffffffff0d510101092f45423132382e302fffffffff0100f2052a01000000232102aea77c449eeeef2746562e56ad053202755f9844276e3f0c684f9d59cdb9458dac00000000")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	toTx, err := NewFromString("02000000019bb2dea27bcff46bca60e46ba2fdce706a8eb9d22c9b05e54166b8f9ac57d6de0000000049483045022100f4de422896e461da647b21d800a4ca9ace98dbd08c2dc9b8e049c93197c314f5022068836c3dfa6650ebeff73b1e3caa8761cd107ed13d6cc713856ebde3f874dd4141feffffff0200ca9a3b000000001976a9143c134f3ccd097be40242efd6fb370fc62501afe788ac00196bee000000001976a914c3d737cb0d93ded96a35d240aa3f01b34edc4e5d88ac65000000")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	t.Errorf("%x\n%x\n", fromTx.GetOutputs()[0].GetOutputScript(), toTx.GetInputs()[0].GetInputScript())
// }
