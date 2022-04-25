package tribe

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	mrand "math/rand"
	"testing"
	"time"

	"github.com/MeshBoxTech/mesh-chain/accounts/keystore"
	"github.com/MeshBoxTech/mesh-chain/crypto"
)

func TestNormal(t *testing.T) {
	prv, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	t.Log(prv)
	t.Log(prv.PublicKey)
	pub := prv.PublicKey
	pk := crypto.PubkeyToAddress(prv.PublicKey)
	t.Log(pk.Hex())
	m := elliptic.Marshal(crypto.S256(), pub.X, pub.Y)
	t.Log("m=", m)
	x, y := elliptic.Unmarshal(crypto.S256(), m)
	t.Log("x,y=", x, y)

}

func TestDelay(t *testing.T) {
	for signerLen := 1; signerLen < 11; signerLen += 1 {
		wiggle := wiggleTime
		delay := time.Duration(mrand.Int63n(int64(wiggle)))
		t.Log(signerLen, wiggle, "\t", delay)
	}
	for signerLen := 10; signerLen < 110; signerLen += 10 {
		wiggle := wiggleTime
		delay := time.Duration(mrand.Int63n(int64(wiggle)))
		t.Log(signerLen, wiggle, "\t", delay)
	}
}

func TestNodekey(t *testing.T) {
	prv, _ := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	k := hex.EncodeToString(crypto.FromECDSA(prv))
	t.Log(k)

	json := `{"address":"4110bd1ff0b73fa12c259acf39c950277f266787","crypto":{"cipher":"aes-128-ctr","ciphertext":"3c12a4bb0048503ea6874c9357b046f3585cdb3af3219618e56220e69d01e5f8","cipherparams":{"iv":"e3e516858fa657437ab53aca9bd48b30"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"63795963bf80c8acff4c4c7f2f33f83b2ebf8f82747c057992cab92ba80beca6"},"mac":"12672d255fbbf0ef9374d9f37da7d25b39ea65b9dec3e63ff6055950dbb48103"},"id":"d541af26-402e-4cdb-841d-8bb700c09887","version":3}`
	key, _ := keystore.DecryptKey([]byte(json), "123456")
	prv2 := key.PrivateKey
	t.Log(crypto.SaveECDSA("/tmp/nodekey", prv2))

	prv3, _ := crypto.HexToECDSA("4755839cb237126c4317a858d85d3837745615b1f67eb20de260416658c43b19")
	t.Log(&prv2)
	t.Log(&prv3)

}

func TestHexToPrv(t *testing.T) {
	h := "4755839cb237126c4317a858d85d3837745615b1f67eb20de260416658c43b19"
	prv, e := crypto.HexToECDSA(h)
	t.Log(e)
	pub := prv.PublicKey
	pbytes := elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	nodeid := pbytes[1:]
	t.Log(crypto.PubkeyToAddress(pub).Hex())
	t.Log(h)
	t.Log(hex.EncodeToString(nodeid))
	//f9cf4fdc53ce5ecf84e7c26fbe262bddd7cbf3d17593328f74816a1c646a0ccfac9a85d81f7d51b59bc02dc8f0e8c5dada4db081efd79698a820faf9384773c0
}

func TestKeyAddressExange(t *testing.T) {
	nodekey := "6188edbba7ba30e7c29f314fea016937816598c8d43a903dabea58bef189d09a"
	prv, _ := crypto.HexToECDSA(nodekey)
	target := "bff277e91e8adeff5c39bb5969f7c8d9be27db76cb61080dc1f1b98601029eb69912731717580dfc140fe331a7953fcba211d59b1d2b086fb669dbdedbf1f717"
	pub := prv.PublicKey
	pbytes := elliptic.Marshal(pub.Curve, pub.X, pub.Y)
	t.Log(pbytes)
	nodeid := pbytes[1:]
	t.Log("---------------")
	t.Log(target)
	t.Log(hex.EncodeToString(nodeid))
	t.Logf("%x\n", nodeid)
	t.Log(hex.EncodeToString(nodeid) == target)
	pub2 := new(ecdsa.PublicKey)
	pub2.Curve = pub.Curve
	tb, _ := hex.DecodeString(target)
	tb = append([]byte{4}, tb...)
	t.Log(tb)
	pub2.X, pub2.Y = elliptic.Unmarshal(pub.Curve, tb)
	t.Log(pub)
	t.Log(pub2)
}

func TestExtra(t *testing.T) {
	extra := []byte("0xd68301000083736d6386676f312e31328664617277696e00000000000000000009b5bd71d0cababfc178e89f63210d322675d53d16bdb58ce5741e06d78f60bd4e2e517089c35d9e278bd6bbf52b6f77746ff314c6cf3237a27209a3e922035300")
	t.Log(len(extra))
}


