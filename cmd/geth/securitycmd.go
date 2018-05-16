package main

import (
	"gopkg.in/urfave/cli.v1"
	"fmt"
	"github.com/SmartMeshFoundation/Spectrum/node"
	"os"
	"io/ioutil"
	"path/filepath"
	"encoding/hex"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/SmartMeshFoundation/Spectrum/accounts/keystore"
	"crypto/ecdsa"
	"github.com/pborman/uuid"
)

const PWD_MIN_LEN = 6

var (
	security        = new(Security)
	securityCommand = cli.Command{
		Before: func(ctx *cli.Context) error {
			fmt.Println("-----------------------")
			if ctx.Bool("testnet") {
				fmt.Println("-   TESTNET ")
				os.Setenv("TESTNET", "1")
			} else {
				fmt.Println("-   MAINNET ")
			}
			security.homeDir = node.DefaultNodekeyDir()
			fmt.Println("home_dir :", security.homeDir)
			fmt.Println("-----------------------")
			return nil
		},
		Action:    security.run,
		Name:      "security",
		Usage:     "encrypt / decrypt nodekey, params mutual exclusion, only one choice",
		ArgsUsage: "<TODO>",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "testnet,t",
				Usage: "",
			},
			cli.BoolFlag{
				Name:        "unlock,l",
				Usage:       "unlock nodekey with password",
				Destination: &security.l,
			},
			cli.BoolFlag{
				Name:        "passwd,p",
				Usage:       "set or reset password",
				Destination: &security.p,
			},
		},
		Category:    "BLOCKCHAIN COMMANDS",
		Description: `TODO`,
	}
)

type Security struct {
	homeDir string
	l, p    bool
	prv     *ecdsa.PrivateKey
}

func (self *Security) run(ctx *cli.Context) error {
	if security.l && security.p {
		fmt.Println("================================================================================")
		fmt.Println(" INPUT : unlock=", security.l, " ; passwd=", security.p)
		fmt.Println(" ERROR : params mutual exclusion, only one choice")
		fmt.Println("================================================================================")
	}
	switch {
	case security.l:
		self.unlockCmd()
	case security.p:
		self.passwdCmd()
	}
	return nil
}

func (self *Security) unlockCmd() {
	fmt.Println("unlock -->", security.l)
}

func (self *Security) passwdCmd() {
	var (
		oldpwd, pwd string
		nodekey     = filepath.Join(security.homeDir, "nodekey")
		nodekeyPrv  = filepath.Join(security.homeDir, "nodekey.prv")
	)
	data, err := ioutil.ReadFile(nodekeyPrv)
	if err != nil {
		counter := 0
		// set password and move nodekey to nodekey.prv
		if err := self.loadNodekey(nodekey); err != nil {
			fmt.Println("ERROR ::", err)
			return
		}
		defer os.Remove(nodekey)
	INPUT:
		fmt.Println("Please input password : ")
		fmt.Scanln(&pwd)
		if len(pwd) < PWD_MIN_LEN {
			counter ++
			fmt.Println(counter, "❌ The password length needs to be greater than", PWD_MIN_LEN)
			if counter < 3 {
				goto INPUT
			}
			return
		}
		self.storeKey(nodekeyPrv, pwd)
	} else {
		// reset password
		counter := 0
		fmt.Println("Please input old password : ")
		fmt.Scanln(&oldpwd)
		kjson, err := hex.DecodeString(string(data))
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}

		key, err := keystore.DecryptKey(kjson, oldpwd)
		if err != nil {
			fmt.Println("ERROR:", err)
			return
		}
	NEWPWD:
		fmt.Println("Please input new password : ")
		fmt.Scanln(&pwd)
		if len(pwd) < PWD_MIN_LEN {
			counter ++
			fmt.Println(counter, "❌ The new password length needs to be greater than", PWD_MIN_LEN)
			if counter < 3 {
				goto NEWPWD
			}
			return
		}
		self.prv = key.PrivateKey
		os.Remove(nodekeyPrv)
		self.storeKey(nodekeyPrv, pwd)
	}

}

func (self *Security) loadNodekey(path string) error {
	f, err := os.Open(path)
	if err != nil {
		prv, err := crypto.GenerateKey()
		if err != nil {
			return err
		}
		self.prv = prv
	}
	buf := make([]byte, 512)
	if n, err := f.Read(buf); err == nil {
		k := string(buf[0:n])
		data, _ := hex.DecodeString(k)
		prv, err := crypto.ToECDSA(data)
		if err != nil {
			return err
		}
		self.prv = prv
	}
	return nil
}

func (self *Security) storeKey(path, passphrase string) {
	id := uuid.NewRandom()
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(self.prv.PublicKey),
		PrivateKey: self.prv,
	}
	keyjson, _ := keystore.EncryptKey(key, passphrase, keystore.StandardScryptN, keystore.StandardScryptP)
	hexkey := hex.EncodeToString(keyjson)
	if err := ioutil.WriteFile(path, []byte(hexkey), 0666); err == nil {
		fmt.Println("😊 Success.")
	} else {
		fmt.Println("😢 Error :", err)
	}
}
