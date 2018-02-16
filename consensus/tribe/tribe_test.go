package tribe

import (
	"testing"
	"crypto/ecdsa"
	"crypto/rand"
	mrand "math/rand"
	"github.com/SmartMeshFoundation/SMChain/crypto"
	"crypto/elliptic"
	"github.com/hashicorp/golang-lru"
	"time"
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

func TestLRU(t *testing.T) {
	lru, e := lru.NewARC(10)
	t.Log(e)
	for i := 0; i < 12; i++ {
		lru.Add(i, 3)
		lru.Get(0)
		t.Log(lru.Keys())
	}
}

func TestDelay(t *testing.T) {
	for signerLen := 1; signerLen < 11; signerLen+=1 {
		wiggle := time.Duration(signerLen/2+1) * wiggleTime
		delay := time.Duration(mrand.Int63n(int64(wiggle)))
		t.Log(signerLen,wiggle,"\t",delay)
	}
	for signerLen := 10; signerLen < 110; signerLen+=10 {
		wiggle := time.Duration(signerLen/2+1) * wiggleTime
		delay := time.Duration(mrand.Int63n(int64(wiggle)))
		t.Log(signerLen,wiggle,"\t",delay)
	}
}
