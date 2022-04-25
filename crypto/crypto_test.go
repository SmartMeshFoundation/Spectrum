// Copyright 2014 The mesh-chain Authors
// This file is part of the mesh-chain library.
//
// The mesh-chain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The mesh-chain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the mesh-chain library. If not, see <http://www.gnu.org/licenses/>.

package crypto

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"testing"

	"github.com/MeshBoxTech/mesh-chain/crypto/secp256k1"
	"github.com/MeshBoxTech/mesh-chain/crypto/vrf"

	"crypto/elliptic"

	"github.com/MeshBoxTech/mesh-chain/common"
)

var testAddrHex = "970e8128ab834e8eac17ab8e3812f010678cf791"
var testPrivHex = "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"

// These tests are sanity checks.
// They should ensure that we don't e.g. use Sha3-224 instead of Sha3-256
// and that the sha3 library uses keccak-f permutation.
func TestKeccak256Hash(t *testing.T) {
	msg := []byte("abc")
	exp, _ := hex.DecodeString("4e03657aea45a94fc7d47ba826c8d667c0d1e6e33a64a036ec44f58fa12d6c45")
	checkhash(t, "Sha3-256-array", func(in []byte) []byte { h := Keccak256Hash(in); return h[:] }, msg, exp)
}

func TestToECDSAErrors(t *testing.T) {
	if _, err := HexToECDSA("0000000000000000000000000000000000000000000000000000000000000000"); err == nil {
		t.Fatal("HexToECDSA should've returned error")
	}
	if _, err := HexToECDSA("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"); err == nil {
		t.Fatal("HexToECDSA should've returned error")
	}
}

func BenchmarkSha3(b *testing.B) {
	a := []byte("hello world")
	for i := 0; i < b.N; i++ {
		Keccak256(a)
	}
}

func TestVRFByExtra(t *testing.T) {
	/*
		23->24 : 78750960575072602031716003452811896949899840465614369727943970206889973630898
		24->25 : 18412240434434283493026822705484548127888845974056244411367015130186730154773
	*/
	dataList := map[string]string{
		"25": `ae1b76d16042cc7215382e989c4f38ceab16c6c83ef321a71bf84c4f3e3fc7b24c683d54da82f628897d659512255276609b70d5fd859fba21e08cea8a0b0ce3b3cdc29d27aa32145a6dd638638f850a0fe37c6fe74f9eccd0b36d00b6063449045736447641f7598906e77178e92324aae9d0fdae197eae1b110b04422cab74cc373eb6b3a1ca4e5345582c96e7d2a8fac2a79e1dcb3aa0fa644ee13cb88b8651351130fac7a558fb6587b97df89dda04b88d992870523f5c98225e865211fe04017355a9e3a41a7d6394d76cb95217e745ab162070da228315843af8aea27e0d01`,
		"26": `28b4f572ac63791e74e31371e07c4f6fb4c5af1701fe734300e0f31e34abab1580e89c3ac75f49fc03916ea656795583e38694a6aad6adf0b5ad241addea34701c698876087186861c146619369f1080a9d870876c89f7607317107d72a04e9304bb0fafcda8866ef2d60c68e9153e1b56b7a99c2cd72167a92ffbd126db1271529c83f2eacd80fcc5b678330566662f4ec7e5d08e68bc23eb83da0690ab7490aef0ad465fc69153b71575aed20ceb576e457bc54dee7216ab8eb5e429d6e44a182c2c5f9cafcb1705631f2087ccfc029361a4ac1c82f7c2a852d9e29d7014e00b00`,
	}
	for k, data := range dataList {
		buf, _ := hex.DecodeString(data)
		i := buf[:32]
		n := new(big.Int).SetBytes(i)
		t.Log(k, n)
	}

}

func TestVRFF(t *testing.T) {
	key, _ := HexToECDSA("0bcd616498bf7aa08be3aacf5a8e9396dce2977c7e475269c47aa869c1743009")
	vrfPrv := vrf.PrivateKey{key}
	vrfPub := vrf.PublicKey{&key.PublicKey}
	msg, _ := hex.DecodeString("6dde11794aa31fb385b9d8d7ea1cfa03e5a0f389efb9cd90be388a342bdf8d7e")
	i1, p1 := vrfPrv.Evaluate(msg)
	i2, p2 := vrfPrv.Evaluate(msg)
	t.Log(len(i1), len(p1), p1)
	t.Log(len(i2), len(p2), p2)
	v1, _ := vrfPub.ProofToHash(msg, p1)
	v2, _ := vrfPub.ProofToHash(msg, p2)
	t.Log(i1 == v1, v1)
	t.Log(i2 == v2, v2)
	for i := 0; i < 30000; i++ {
		i3, p3 := vrfPrv.Evaluate(msg)
		if len(i3) != 32 {
			t.Log("XXX", len(i3), i)
		}
		if len(p3) != 129 {
			t.Log("YYY", len(p3), i)
		}
	}
}

func TestVRF(t *testing.T) {
	key, err := GenerateKey()
	if err != nil {
		t.Error(err)
		return
	}
	k, err := vrf.NewVRFSigner(key)
	if err != nil {
		t.Error(err)
		return
	}
	pk, err := vrf.NewVRFVerifier(&key.PublicKey)
	if err != nil {
		t.Error(err)
		return
	}
	msg := []byte("data1")
	index, proof := k.Evaluate(msg)
	_index, _proof := k.Evaluate(msg)
	index2, err := pk.ProofToHash(msg, proof)

	if err != nil {
		t.Error(err)
	}
	if index2 != index {
		t.Error("index not equal")
	}

	t.Log(index)
	t.Log(_index)
	t.Log("=========")
	t.Log(proof)
	t.Log(_proof)
}

func TestVRF1(t *testing.T) {
	//data := "89702163480656fb718d9a06d56f719fda0ab2361d42de308ef3589a0a71f8e11a11c9d38a6a0c8e58c6d3a64a5dbd9a914216dc997f48c186daa4d772e613136005baac9928765fc92db4b541b05e5ef1638c56594ef2af2d80c735feb169437e0d56aafd225566059da85922d1a2bea0ebd7da5b63d1e2d846109378a075d67d00"
	msg := []byte("helloworld")
	prv, _ := HexToECDSA("507fd083b5c5af7e645e77a3a3a82708f3af304164e02612ab4b1d5b36c627c6")
	pub := &prv.PublicKey
	np, _ := SimpleVRF2Bytes(prv, msg)
	t.Log("vrfnp", len(np), np)
	t.Log(hex.EncodeToString(np))
	err := SimpleVRFVerify(pub, msg, np)
	t.Log("verify", err)

	buf := elliptic.Marshal(pub.Curve, pub.X, pub.Y)

	var signer common.Address
	copy(signer[:], Keccak256(buf[1:])[12:])

	aa := PubkeyToAddress(*pub)

	t.Log(signer.Hex())
	t.Log(aa.Hex())

	x, y := elliptic.Unmarshal(S256(), buf)
	t.Log(x, pub.X)
	t.Log(y, pub.Y)

	pub2 := ecdsa.PublicKey{S256(), x, y}
	t.Log(PubkeyToAddress(pub2).Hex())

}

/*
私钥 : 0bcd616498bf7aa08be3aacf5a8e9396dce2977c7e475269c47aa869c1743009
原文 : 0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c

签名结果
R : 0xa006360ed68f21443221354903f49dce733afaaeac3b35d420bb2154746c9592
S : 0x6f31ccfa10b449531fd0fff78db52cc02edaabd2c5e9a6abb8fc1cd6df26f442
V : 28

-------------------------
r = sigHex[0:64]
s = sigHex[64:128]
v = sigHex[128:130] + 27
-------------------------
*/
func TestSign2(t *testing.T) {
	key, _ := HexToECDSA("0bcd616498bf7aa08be3aacf5a8e9396dce2977c7e475269c47aa869c1743009")
	msg := Keccak256(common.HexToAddress("0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c").Bytes())

	sig, _ := Sign(msg, key)
	sigHex := hex.EncodeToString(sig)
	pubAddr := PubkeyToAddress(key.PublicKey)
	t.Log("addr =", pubAddr.Hex())
	t.Log("msg =", hex.EncodeToString(msg))
	t.Log("Sign =", sigHex)
	r, _ := hex.DecodeString(sigHex[:64])
	buff := []byte(r)
	fmt.Println("llllll=", len(buff), len(sig))
	fmt.Println(r)

	var r32 [32]byte
	copy(r32[:], r)
	fmt.Println(len(r32), "r32=", r32)

	fmt.Println(sig[:32])
	R := "0x" + sigHex[:64]
	S := "0x" + sigHex[64:128]
	V := 27
	switch sigHex[128:] {
	case "01":
		V = 28
	}
	t.Log("-----------------------------------------")
	t.Log("R:", R)
	t.Log("S:", S)
	t.Log("V:", V)

}

func TestSign(t *testing.T) {
	key, _ := HexToECDSA(testPrivHex)
	addr := common.HexToAddress(testAddrHex)

	msg := Keccak256([]byte("foobar"))
	sig, err := Sign(msg, key)
	t.Log("Sign =", sig)
	t.Log("Sign.len =", len(sig))
	t.Log("Sign.hex =", hex.EncodeToString(sig))
	if err != nil {
		t.Errorf("Sign error: %s", err)
	}
	recoveredPub, err := Ecrecover(msg, sig)
	if err != nil {
		t.Errorf("ECRecover error: %s", err)
	}
	pubKey := ToECDSAPub(recoveredPub)
	recoveredAddr := PubkeyToAddress(*pubKey)
	if addr != recoveredAddr {
		t.Errorf("Address mismatch: want: %x have: %x", addr, recoveredAddr)
	}

	// should be equal to SigToPub
	recoveredPub2, err := SigToPub(msg, sig)
	if err != nil {
		t.Errorf("ECRecover error: %s", err)
	}
	recoveredAddr2 := PubkeyToAddress(*recoveredPub2)
	if addr != recoveredAddr2 {
		t.Errorf("Address mismatch: want: %x have: %x", addr, recoveredAddr2)
	}
}

func TestInvalidSign(t *testing.T) {
	if _, err := Sign(make([]byte, 1), nil); err == nil {
		t.Errorf("expected sign with hash 1 byte to error")
	}
	if _, err := Sign(make([]byte, 33), nil); err == nil {
		t.Errorf("expected sign with hash 33 byte to error")
	}
}

func TestNewContractAddress(t *testing.T) {
	key, _ := HexToECDSA(testPrivHex)
	addr := common.HexToAddress(testAddrHex)
	genAddr := PubkeyToAddress(key.PublicKey)
	// sanity check before using addr to create contract address
	checkAddr(t, genAddr, addr)

	caddr0 := CreateAddress(addr, 0)
	caddr1 := CreateAddress(addr, 1)
	caddr2 := CreateAddress(addr, 2)
	checkAddr(t, common.HexToAddress("333c3310824b7c685133f2bedb2ca4b8b4df633d"), caddr0)
	checkAddr(t, common.HexToAddress("8bda78331c916a08481428e4b07c96d3e916d165"), caddr1)
	checkAddr(t, common.HexToAddress("c9ddedf451bc62ce88bf9292afb13df35b670699"), caddr2)
}

func TestLoadECDSAFile(t *testing.T) {
	keyBytes := common.FromHex(testPrivHex)
	fileName0 := "test_key0"
	fileName1 := "test_key1"
	checkKey := func(k *ecdsa.PrivateKey) {
		checkAddr(t, PubkeyToAddress(k.PublicKey), common.HexToAddress(testAddrHex))
		loadedKeyBytes := FromECDSA(k)
		if !bytes.Equal(loadedKeyBytes, keyBytes) {
			t.Fatalf("private key mismatch: want: %x have: %x", keyBytes, loadedKeyBytes)
		}
	}

	ioutil.WriteFile(fileName0, []byte(testPrivHex), 0600)
	defer os.Remove(fileName0)

	key0, err := LoadECDSA(fileName0)
	if err != nil {
		t.Fatal(err)
	}
	checkKey(key0)

	// again, this time with SaveECDSA instead of manual save:
	err = SaveECDSA(fileName1, key0)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(fileName1)

	key1, err := LoadECDSA(fileName1)
	if err != nil {
		t.Fatal(err)
	}
	checkKey(key1)
}

func TestValidateSignatureValues(t *testing.T) {
	check := func(expected bool, v byte, r, s *big.Int) {
		if ValidateSignatureValues(v, r, s, false) != expected {
			t.Errorf("mismatch for v: %d r: %d s: %d want: %v", v, r, s, expected)
		}
	}
	minusOne := big.NewInt(-1)
	one := common.Big1
	zero := common.Big0
	secp256k1nMinus1 := new(big.Int).Sub(secp256k1_N, common.Big1)

	// correct v,r,s
	check(true, 0, one, one)
	check(true, 1, one, one)
	// incorrect v, correct r,s,
	check(false, 2, one, one)
	check(false, 3, one, one)

	// incorrect v, combinations of incorrect/correct r,s at lower limit
	check(false, 2, zero, zero)
	check(false, 2, zero, one)
	check(false, 2, one, zero)
	check(false, 2, one, one)

	// correct v for any combination of incorrect r,s
	check(false, 0, zero, zero)
	check(false, 0, zero, one)
	check(false, 0, one, zero)

	check(false, 1, zero, zero)
	check(false, 1, zero, one)
	check(false, 1, one, zero)

	// correct sig with max r,s
	check(true, 0, secp256k1nMinus1, secp256k1nMinus1)
	// correct v, combinations of incorrect r,s at upper limit
	check(false, 0, secp256k1_N, secp256k1nMinus1)
	check(false, 0, secp256k1nMinus1, secp256k1_N)
	check(false, 0, secp256k1_N, secp256k1_N)

	// current callers ensures r,s cannot be negative, but let's test for that too
	// as crypto package could be used stand-alone
	check(false, 0, minusOne, one)
	check(false, 0, one, minusOne)
}

func checkhash(t *testing.T, name string, f func([]byte) []byte, msg, exp []byte) {
	sum := f(msg)
	if !bytes.Equal(exp, sum) {
		t.Fatalf("hash %s mismatch: want: %x have: %x", name, exp, sum)
	}
}

func checkAddr(t *testing.T, addr0, addr1 common.Address) {
	if addr0 != addr1 {
		t.Fatalf("address mismatch: want: %x have: %x", addr0, addr1)
	}
}

// test to help Python team with integration of libsecp256k1
// skip but keep it after they are done
func TestPythonIntegration(t *testing.T) {
	kh := "289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032"
	k0, _ := HexToECDSA(kh)

	msg0 := Keccak256([]byte("foo"))
	sig0, _ := Sign(msg0, k0)

	msg1 := common.FromHex("00000000000000000000000000000000")
	sig1, _ := Sign(msg0, k0)

	t.Logf("msg: %x, privkey: %s sig: %x\n", msg0, kh, sig0)
	t.Logf("msg: %x, privkey: %s sig: %x\n", msg1, kh, sig1)
}

func TestToECDSA(t *testing.T) {
	target := "0adec74681c34173416ae564849e53e52034156e238084948b53cd1f7a753d6cc90763b5024f26998ee192d9662232d4c68abd588347626c2e8cd95e5c1ca38b"
	nodekey := "b07f62f843adf1622811a79945944748fe5e59174d1ceffd14ca90a35c5813a7"
	k, _ := hex.DecodeString(nodekey)
	prv, _ := ToECDSA(k)
	// TODO 将一个公钥变成一个 nodeid
	pub := prv.PublicKey
	// 变成一个 Address
	address := PubkeyToAddress(pub)
	t.Log("address =", address.Hex())

	pbytes := elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	nodeid := pbytes[1:]
	t.Log("---------------")
	t.Log(target)
	t.Log(hex.EncodeToString(nodeid))
	t.Logf("%x\n", nodeid)
	t.Log(hex.EncodeToString(nodeid) == target)
}

func TestFromECDSA(t *testing.T) {
	prv, _ := GenerateKey()
	buf := FromECDSA(prv)
	s := hex.EncodeToString(buf)
	t.Log("key", s)
	t.Log(PubkeyToAddress(prv.PublicKey).Hex())
	//0xd91c07c6be04111aecc721e57e11a8a0b1c73fc1
}

func TestForPOCDepositArgs(t *testing.T) {
	prvkey1, _ := HexToECDSA("543e9d0ddcd02b4bbdb2cecd402da99e9532fface57d8ea74e833c5d413f2daa")
	prvkey2, _ := HexToECDSA("9d7b3b8155ea429cb49cdc556aa7b3367feb91ccf51eb1e9c7e2bac67d939f03")

	keys := []*ecdsa.PrivateKey{prvkey1, prvkey2}

	msg := Keccak256(common.HexToAddress("0xbf1736a65f8beadd71b321be585fd883503fdeaa").Bytes())

	for _, key := range keys {
		t.Log("========================================")
		sig, _ := Sign(msg, key)
		sigHex := hex.EncodeToString(sig)
		pubAddr := PubkeyToAddress(key.PublicKey)
		t.Log("msg =", hex.EncodeToString(msg))
		t.Log("Sign =", sigHex)
		r, _ := hex.DecodeString(sigHex[:64])

		var r32 [32]byte
		copy(r32[:], r)
		fmt.Println(len(r32), "r32=", r32)

		fmt.Println(sig[:32])
		R := "0x" + sigHex[:64]
		S := "0x" + sigHex[64:128]
		V := 27
		switch sigHex[128:] {
		case "01":
			V = 28
		}
		t.Log("-----------------------------------------")
		t.Log("addr:", pubAddr.Hex())
		t.Log("R:", R)
		t.Log("S:", S)
		t.Log("V:", V)
	}

}

func TestVRFForS256Fenbu(t *testing.T) {
	if testing.Short() {
		return
	}
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("key=%s\n", hex.EncodeToString(FromECDSA(key)))
	var n = 380
	var batch = 10000000
	var k100 []*ecdsa.PrivateKey
	for i := 0; i < n; i++ {
		k, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Printf("k%d=%s\n", i, hex.EncodeToString(FromECDSA(k)))
		k100 = append(k100, k)
	}
	m := make(map[int]int)
	var lastVrf []byte
	for i := 0; i < batch; i++ {
		var tk *ecdsa.PrivateKey
		if i%17 == 0 || len(lastVrf) == 0 {
			tk = key
		} else {
			num := new(big.Int).SetBytes(lastVrf)
			num = num.Mod(num, big.NewInt(int64(n)))
			tk = k100[num.Int64()]
			m[int(num.Int64())] = m[int(num.Int64())] + 1
		}
		msg := append(big.NewInt(int64(i-1)).Bytes(), lastVrf...)
		vrfnp, err := SimpleVRF2Bytes(tk, msg)
		if err != nil {
			t.Error(err)
			return
		}
		lastVrf = vrfnp
	}
	var avg int
	for k, v := range m {
		avg += v
		fmt.Printf("%d,%d\n", k, v)
	}
	fmt.Printf("total=%d", avg)
	fmt.Printf("avg=%d", avg/n)
}

func TestVRFForS256Fenbu2(t *testing.T) {
	if testing.Short() {
		return
	}
	key, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("key=%s\n", hex.EncodeToString(FromECDSA(key)))
	var n = 380
	var batch = 1000000
	var k100 []*ecdsa.PrivateKey
	for i := 0; i < n; i++ {
		k, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)
		if err != nil {
			t.Error(err)
			return
		}
		fmt.Printf("k%d=%s\n", i, hex.EncodeToString(FromECDSA(k)))
		k100 = append(k100, k)
	}
	var signers [17]*ecdsa.PrivateKey
	var nextRoundSigners [17]*ecdsa.PrivateKey
	var emptySigners [17]*ecdsa.PrivateKey
	signers[0] = key
	copy(signers[1:17], k100[0:16])
	copy(nextRoundSigners[:], signers[:])
	find := func(num *big.Int) *ecdsa.PrivateKey {
		num = num.Mod(num, big.NewInt(int64(n)))
		k := num.Int64()
		var ck *ecdsa.PrivateKey
		for {
		nextRound:
			ck = k100[k]
			for i := 0; i < 17; i++ {
				if signers[i] == ck {
					k++
					k = k % int64(n)
					goto nextRound
				}
			}
			for i := 0; i < 17; i++ {
				if nextRoundSigners[i] == ck {
					k++
					k = k % int64(n)
					goto nextRound
				}
			}
			break //没有和已选择的冲突
		}
		return ck
	}
	m := make(map[int]int)
	var lastVrf []byte
	for i := 0; i < batch; i++ {
		var tk *ecdsa.PrivateKey
		if i%17 == 0 || len(lastVrf) == 0 {
			copy(signers[:], nextRoundSigners[:])
			copy(nextRoundSigners[:], emptySigners[:])
			tk = signers[0]
			nextRoundSigners[0] = tk //第0个不变
		} else {
			num := new(big.Int).SetBytes(lastVrf)
			tk = find(num)
			m[int(num.Int64())] = m[int(num.Int64())] + 1
		}
		msg := append(big.NewInt(int64(i-1)).Bytes(), lastVrf...)
		vrfnp, err := SimpleVRF2Bytes(tk, msg)
		if err != nil {
			t.Error(err)
			return
		}
		lastVrf = vrfnp
	}
	var avg int
	for k, v := range m {
		avg += v
		fmt.Printf("%d,%d\n", k, v)
	}
	fmt.Printf("total=%d", avg)
	fmt.Printf("avg=%d", avg/n)
}
