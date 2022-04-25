package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/MeshBoxTech/mesh-chain/accounts/keystore"
	"github.com/MeshBoxTech/mesh-chain/cmd/utils"
	"github.com/MeshBoxTech/mesh-chain/crypto"
	"github.com/MeshBoxTech/mesh-chain/node"
	"github.com/pborman/uuid"
	"gopkg.in/urfave/cli.v1"
	"io/ioutil"
	"os"
	"path/filepath"
)

const PWD_MIN_LEN = 6

var (
	security        = new(Security)
	securityCommand = cli.Command{
		Before: func(ctx *cli.Context) error {
			fmt.Println("-----------------------")
			if ctx.Bool(utils.TestnetFlag.Name) {
				fmt.Println("-   TESTNET ")
				os.Setenv("TESTNET", "1")
			} else if ctx.Bool(utils.DevnetFlag.Name) {
				fmt.Println("-   DEVNET ")
				os.Setenv("DEVNET", "1")
			} else {
				fmt.Println("-   MAINNET ")
			}
			security.homeDir = node.DefaultNodekeyDir()
			fmt.Println("home_dir :", security.homeDir)
			fmt.Println("-----------------------")
			security.nodekey = filepath.Join(security.homeDir, "nodekey")
			security.nodekeyPrv = filepath.Join(security.homeDir, "nodekey.prv")
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
	homeDir, nodekey, nodekeyPrv string
	l, p                         bool
	prv                          *ecdsa.PrivateKey
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
	var (
		pwd     string
		counter int
		pwdfile = filepath.Join(security.homeDir, "nodekey.pwd")
	)
	data, err := ioutil.ReadFile(self.nodekeyPrv)
	if err != nil {
		fmt.Println("Please set password first.")
		fmt.Println("ERROR :", err)
	}
	kjson, err := hex.DecodeString(string(data))
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
INPUT:
	fmt.Println("Please input password : ")
	fmt.Scanln(&pwd)
	if _, err := keystore.DecryptKey(kjson, pwd); err != nil {
		counter++
		fmt.Println(counter, "❌ Wrong password .")
		if counter < 3 {
			goto INPUT
		}
		return
	}
	os.Remove(pwdfile)
	pwdhex := hex.EncodeToString([]byte(pwd))
	if err := ioutil.WriteFile(pwdfile, []byte(pwdhex), 0666); err == nil {
		fmt.Println("😊 Success.")
	} else {
		fmt.Println("😢 Error :", err)
	}
}

func (self *Security) passwdCmd() {
	var (
		oldpwd, pwd string
	)
	data, err := ioutil.ReadFile(self.nodekeyPrv)
	if err != nil {
		counter := 0
		// set password and move nodekey to nodekey.prv
		if err := self.loadNodekey(self.nodekey); err != nil {
			fmt.Println("ERROR ::", err)
			return
		}
		defer os.Remove(self.nodekey)
	INPUT:
		fmt.Println("Please input password : ")
		fmt.Scanln(&pwd)
		if len(pwd) < PWD_MIN_LEN {
			counter++
			fmt.Println(counter, "❌ The password length needs to be greater than", PWD_MIN_LEN)
			if counter < 3 {
				goto INPUT
			}
			return
		}
		self.storeKey(self.nodekeyPrv, pwd)
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
			counter++
			fmt.Println(counter, "❌ The new password length needs to be greater than", PWD_MIN_LEN)
			if counter < 3 {
				goto NEWPWD
			}
			return
		}
		self.prv = key.PrivateKey
		os.Remove(self.nodekeyPrv)
		self.storeKey(self.nodekeyPrv, pwd)
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
