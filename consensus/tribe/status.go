package tribe

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/SmartMeshFoundation/Spectrum/core/state"

	"github.com/SmartMeshFoundation/Spectrum/common"
	"github.com/SmartMeshFoundation/Spectrum/core/types"
	"github.com/SmartMeshFoundation/Spectrum/crypto"
	"github.com/SmartMeshFoundation/Spectrum/log"
	"github.com/SmartMeshFoundation/Spectrum/params"
)

func NewTribeStatus() *TribeStatus {
	ts := &TribeStatus{
		Signers:     make([]*Signer, 0),
		SignerLevel: LevelNone,
	}
	return ts
}

func (self *TribeStatus) setTribe(tribe *Tribe) {
	self.tribe = tribe
}

func (self *TribeStatus) getNodekey() *ecdsa.PrivateKey {
	if self.nodeKey == nil {
		panic(errors.New("GetNodekey but nodekey not ready"))
	}
	return self.nodeKey
}

func (self *TribeStatus) GetMinerAddress() common.Address {
	if self.nodeKey == nil {
		panic(errors.New("GetMinerAddress but nodekey not ready"))
	}
	pub := self.nodeKey.PublicKey
	add := crypto.PubkeyToAddress(pub)
	return add
}
func (self *TribeStatus) IsLeader(addr common.Address) bool {
	for _, a := range self.Leaders {
		if a == addr {
			return true
		}
	}
	return false
}
func (self *TribeStatus) GetMinerAddressByChan(rtn chan common.Address) {
	go func() {
		for {
			if self.nodeKey != nil && self.tribe.isInit {
				break
			}
			<-time.After(time.Second)
		}
		pub := self.nodeKey.PublicKey
		rtn <- crypto.PubkeyToAddress(pub)
	}()
}

func (self *TribeStatus) GetSignersFromChiefByHash(hash common.Hash, number *big.Int) ([]*Signer, error) {
	rtn := params.SendToMsgBoxWithHash("GetStatus", hash, number)
	r := <-rtn
	if !r.Success {
		return nil, r.Entity.(error)
	}
	cs := r.Entity.(params.ChiefStatus)
	signers := cs.SignerList
	scores := cs.ScoreList
	sl := make([]*Signer, 0, len(signers))
	for i, signer := range signers {
		score := scores[i]
		sl = append(sl, &Signer{signer, score.Int64()})
	}
	return sl, nil
}

// 在 加载完所有 node.service 后，需要主动调用一次
func (self *TribeStatus) LoadSignersFromChief(hash common.Hash, number *big.Int) error {
	//log.Info(fmt.Sprintf("LoadSignersFromChief hash=%s,number=%s", hash.String(), number))
	cs, err := params.TribeGetStatus(number, hash)
	if err != nil {
		log.Warn("TribeGetStatusError", "err", err, "num", number, "hash", hash.Hex())
		return err
	}
	signers := cs.SignerList
	scores := cs.ScoreList
	sl := make([]*Signer, 0, len(signers))
	for i, signer := range signers {
		score := scores[i]
		sl = append(sl, &Signer{signer, score.Int64()})
	}
	self.LeaderLimit = cs.LeaderLimit
	self.Leaders = cs.LeaderList
	if len(self.Leaders) == 0 && params.IsSIP100Block(number) {
		panic(fmt.Sprintf("LoadSignersFromChief err ,hash=%s,number=%s,cs=%#v", hash.String(), number, cs))
	}
	self.Number = cs.Number.Int64()
	self.blackList = cs.BlackList
	err = self.loadSigners(sl)
	if err != nil {
		return err
	}
	self.Epoch, self.SignerLimit = cs.Epoch, cs.SignerLimit
	go self.resetSignersLevel(hash, number)
	return nil
}

func (self *TribeStatus) resetSignersLevel(hash common.Hash, number *big.Int) {
	m := self.GetMinerAddress()
	for _, s := range self.Signers {
		if s.Address == m {
			self.SignerLevel = LevelSigner
			return
		}
	}
	for _, s := range self.blackList {
		if s == m {
			self.SignerLevel = LevelSinner
			return
		}
	}

	for _, s := range self.Leaders {
		if s == m {
			self.SignerLevel = LevelSigner
			return
		}
	}

	ci := params.GetChiefInfo(number)
	switch ci.Version {
	case "0.0.6":
		// if filterVolunteer return 1 then is volunteer
		rtn := params.SendToMsgBoxForFilterVolunteer(hash, number, m)
		r := <-rtn
		if r.Success {
			if fr := r.Entity.(*big.Int); fr != nil && fr.Int64() == 0 {
				self.SignerLevel = LevelVolunteer
				return
			}
		}
	}
	// default none
	self.SignerLevel = LevelNone
}

//每一块都会调用
func (self *TribeStatus) loadSigners(sl []*Signer) error {
	self.Signers = append(self.Signers[:0], sl...)
	return nil
}

//InTurnForCalcDiffcultyChief100 计算规则参考inTurnForCalcChief100
func (self *TribeStatus) InTurnForCalcDiffcultyChief100(signer common.Address, parent *types.Header) *big.Int {
	return self.inTurnForCalcDifficultyChief100(parent.Number.Int64()+1, parent.Hash(), signer)
}

/*
inTurnForCalcDifficultyChief100 计算如果当前出块节点是signer的话,它对应的难度是多少.
signers:[0,...,16] 0号对应的是常委会节点,1-16对应的是普通出块节点
场景1:
1. 当前应该出块节点应该是3,如果signer是3,那么难度就是6.
2. 如果singers[0]对应的是常委2, 这时候常委2出块,难度是5,常委3出块难度是4,...,常委1出块难度则是1
场景2:当前出块节点应该是singers[0],也就是某个常委会节点
1. 如果signers[0] 出块,那么难度就是6
2. 假设signers[0]是常委2,那么常委3替他出块难度是5,常委4出块就是4,...常委1出块难度则是2

这里的number参数主要是选定合约版本,而parentHash则是用来选择读取哪个block时候的合约状态
*/
func (self *TribeStatus) inTurnForCalcDifficultyChief100(number int64, parentHash common.Hash, signer common.Address) *big.Int {
	var (
		signers, _ = self.GetSignersFromChiefByHash(parentHash, big.NewInt(number)) //self.GetSigners()
		sl         = len(signers)
	)
	//	log.Info(fmt.Sprintf("singers=%v,signer=%s,leaders=%v,number=%d,parentHash=%s", signers, signer.String(), self.Leaders, number, parentHash.String()))
	if idx, _, err := self.fetchOnSigners(signer, signers); err == nil {
		// main
		if sl > 0 && number%int64(sl) == idx.Int64() {
			return big.NewInt(diff)
		}
		// second
		if idx.Int64() == 0 {
			return big.NewInt(diff - 1)
		}

	} else if sl > 0 {
		if leaders, err := leaderSort(signers[0].Address, self.Leaders); err == nil {
			for i, leader := range leaders {
				if signer == leader && number%int64(sl) == 0 {
					return big.NewInt(diff - int64(i+1))
				} else if signer == leader {
					return big.NewInt(diff - int64(i+2))
				}
			}
		}
	}
	return diffNoTurn
}

//InTurnForVerifyDifficultyChief100: 计算规则参考inTurnForCalcChief100
func (self *TribeStatus) InTurnForVerifyDifficultyChief100(number int64, parentHash common.Hash, signer common.Address) *big.Int {
	return self.inTurnForCalcDifficultyChief100(number, parentHash, signer)
}

/*
假设list=[1,2,3,4,5]
first=3,那么返回[4,5,1,2]
如果first=2,返回[3,4,5,1]
如果first=5,返回[1,2,3,4]
不允许返回错误,是因为考虑到运行过程中leader可能会被删除,从而导致找不到leader
*/
func leaderSort(first common.Address, list []common.Address) ([]common.Address, error) {
	for i, o := range list {
		if first == o {
			return append(list[i+1:], list[:i]...), nil
		}
	}

	return list, nil
}

//InTurnForCalcDifficulty 在0.6版本yiqian 计算难度
func (self *TribeStatus) InTurnForCalcDifficulty(signer common.Address, parent *types.Header) *big.Int {
	number := parent.Number.Int64() + 1
	signers := self.Signers
	if idx, _, err := self.fetchOnSigners(signer, signers); err == nil {
		sl := len(signers)
		if params.IsSIP002Block(big.NewInt(number)) {
			if sl > 0 && number%int64(sl) == idx.Int64() {
				return diffInTurnMain
			} else if sl > 0 && (number+1)%int64(sl) == idx.Int64() {
				return diffInTurn
			}
		} else {
			if sl > 0 && number%int64(sl) == idx.Int64() {
				return diffInTurn
			}
		}
	}

	return diffNoTurn
}

//0.6版本之前校验难度
func (self *TribeStatus) InTurnForVerifyDiffculty(number int64, parentHash common.Hash, signer common.Address) *big.Int {

	if ci := params.GetChiefInfo(big.NewInt(number)); ci != nil {
		switch ci.Version {
		case "1.0.0":
			//TODO max value is a var ???
			return self.InTurnForVerifyDifficultyChief100(number, parentHash, signer)
		}
	}

	var signers []*Signer
	if number > 3 {
		var err error
		signers, err = self.GetSignersFromChiefByHash(parentHash, big.NewInt(number))
		if err != nil {
			log.Warn("InTurn:GetSignersFromChiefByNumber:", "err", err)
		}
	} else {
		return diffInTurn
	}
	if idx, _, err := self.fetchOnSigners(signer, signers); err == nil {
		sl := len(signers)
		if params.IsSIP002Block(big.NewInt(number)) {
			if sl > 0 && number%int64(sl) == idx.Int64() {
				return diffInTurnMain
			} else if sl > 0 && (number+1)%int64(sl) == idx.Int64() {
				return diffInTurn
			}
		} else {
			if sl > 0 && number%int64(sl) == idx.Int64() {
				return diffInTurn
			}
		}
	}
	return diffNoTurn
}

func (self *TribeStatus) genesisSigner(header *types.Header) (common.Address, error) {
	extraVanity := extraVanityFn(header.Number)
	signer := common.Address{}
	copy(signer[:], header.Extra[extraVanity:])
	self.loadSigners([]*Signer{{signer, 3}})
	return signer, nil
}

//address对应的signer以及其在signers中的下标
func (self *TribeStatus) fetchOnSigners(address common.Address, signers []*Signer) (*big.Int, *Signer, error) {
	if signers == nil {
		signers = self.Signers
	}
	if l := len(signers); l > 0 {
		for i := 0; i < l; i++ {
			if s := signers[i]; s.Address == address {
				return big.NewInt(int64(i)), s, nil
			}
		}
	}
	return nil, nil, errors.New("not_found")
}

func verifyVrfNum(parent, header *types.Header) (err error) {
	var (
		np  = header.Extra[:extraVanityFn(header.Number)]
		sig = header.Extra[len(header.Extra)-extraSeal:]
		msg = append(parent.Number.Bytes(), parent.Extra[:32]...)
	)
	pubbuf, err := ecrecoverPubkey(header, sig)
	if err != nil {
		//panic(err) //这地方不能panic,否则一个节点出一个恶意的块,所以的节点就全崩了.
		return err
	}
	x, y := elliptic.Unmarshal(crypto.S256(), pubbuf)
	pubkey := ecdsa.PublicKey{crypto.S256(), x, y}
	err = crypto.SimpleVRFVerify(&pubkey, msg, np)
	log.Debug("[verifyVrfNum]", "err", err, "num", header.Number, "vrfn", new(big.Int).SetBytes(np[:32]), "parent", header.ParentHash.Bytes())
	return
}

/*
validateSigner:
1. 验证出块时间符合规则,具体规则见GetPeriodChief100描述
2.
*/
func (self *TribeStatus) validateSigner(parentHeader, header *types.Header, signer common.Address) bool {
	var (
		err        error
		signers    []*Signer
		number     = header.Number.Int64()
		parentHash = header.ParentHash
	)
	//if number > 1 && self.Number != parentNumber {
	if number <= CHIEF_NUMBER {
		return true
	}

	signers, err = self.GetSignersFromChiefByHash(parentHash, big.NewInt(number))
	if err != nil {
		log.Warn("TribeStatus.ValidateSigner : GetSignersFromChiefByNumber :", "err", err)
	}

	if params.IsSIP002Block(header.Number) {
		// second time of verification block time
		period := self.tribe.GetPeriod(header, signers)
		pt := parentHeader.Time.Uint64()
		if pt+period > header.Time.Uint64() {
			log.Error("[ValidateSigner] second time verification block time error", "num", header.Number, "pt", pt, "period", period, ", pt+period=", pt+period, " , ht=", header.Time.Uint64())
			log.Error("[ValidateSigner] second time verification block time error", "err", ErrInvalidTimestampSIP002)
			return false
		}
	}

	if params.IsSIP100Block(header.Number) && header.Coinbase == common.HexToAddress("0x") {
		log.Error("error_signer", "num", header.Number.String(), "miner", header.Coinbase.Hex(), "signer", signer.Hex())
		return false
	}

	idx, _, err := self.fetchOnSigners(signer, signers)
	if params.IsSIP100Block(header.Number) {
		if err == nil {
			// 轮到谁出就谁出的块
			idx_m := number % int64(len(signers))
			if idx_m == idx.Int64() {
				return true
			}
			// 其他只能有常委会节点替代
			if idx.Int64() == 0 {
				return true
			}
		} else {
			// other leader
			for _, leader := range self.Leaders {
				if signer == leader { //有没有测试过多个常委会节点同时出块的情况呢?
					return true
				}
			}
		}
	} else if err == nil {
		return true
	}
	return false
}

/*
VerifySignerBalance: 在chief1.0之前直接通过账号余额来判断是否具有出块资格,chief1.0之后只能通过抵押到poc合约中才具有资格.
*/
//func (self *TribeStatus) VerifySignerBalance(state *state.StateDB, addr common.Address, header *types.Header) error {
//	// SIP100 skip this verify
//	if params.IsSIP100Block(header.Number) {
//		return nil
//	}
//	var (
//		pnum, num *big.Int
//	)
//	if addr == common.HexToAddress("0x") {
//		if _addr, err := ecrecover(header, self.tribe); err == nil {
//			addr = _addr
//		} else {
//			return err
//		}
//	}
//	if header != nil {
//		num = header.Number
//		pnum = new(big.Int).Sub(num, big.NewInt(1))
//	} else {
//		return errors.New("params of header can not be null")
//	}
//	// skip when v in meshbox.sol
//	if params.IsReadyMeshbox(pnum) && params.MeshboxExistAddress(addr) {
//		return nil
//	}
//
//	return nil
//
//}

// every block
// sync download or mine
// check chief tx
func (self *TribeStatus) ValidateBlock(state *state.StateDB, parent, block *types.Block, validateSigner bool) error {
	if block.Number().Int64() <= CHIEF_NUMBER {
		return nil
	}
	err := self.LoadSignersFromChief(parent.Hash(), parent.Number())
	if err != nil {
		log.Error(fmt.Sprintf("[ValidateBlock] LoadSignersFromChief ,parent=%s,current=%s,currentNumber=%s", parent.Hash().String(), block.Hash().String(), block.Number()))
		return err
	}
	header := block.Header()
	number := header.Number.Int64()

	//number := block.Number().Int64()
	// add by liangc : seal call this func must skip validate signer 因为这时候签名都还没准备好
	if validateSigner {
		signer, err := ecrecover(header, self.tribe)
		// verify signer
		if err != nil {
			return err
		}
		// verify difficulty 就算是
		if !params.IsBeforeChief100block(header.Number) {
			difficulty := self.InTurnForVerifyDiffculty(number, header.ParentHash, signer)
			if difficulty.Cmp(header.Difficulty) != 0 {
				log.Error("** verifySeal ERROR **", "head.diff", header.Difficulty.String(), "target.diff", difficulty.String(), "err", errInvalidDifficulty)
				return errInvalidDifficulty
			}
		}
		// verify vrf num
		if params.IsSIP100Block(header.Number) {
			err = verifyVrfNum(parent.Header(), header)
			if err != nil {
				log.Error("vrf_num_fail", "num", number, "err", err)
				return err
			}
		}
		if !self.validateSigner(parent.Header(), header, signer) {
			return errUnauthorized
		}
	}
	// check first tx , must be chief.tx , and onely one chief.tx in tx list
	if block != nil && block.Transactions().Len() == 0 {
		return ErrTribeNotAllowEmptyTxList
	}

	var total = 0
	for i, tx := range block.Transactions() {
		from := types.GetFromByTx(tx)
		/*
			must verify tx.from ==signer:
			otherwise:
			if miner A minging the block#16,then A can make chief.update tx fail,
			so signerList will never update, A will make sure he can mine block for every round.
		*/
		if tx.To() != nil && params.IsChiefAddress(*tx.To()) && params.IsChiefUpdate(tx.Data()) {
			if i != 0 {
				return ErrTribeChiefTxMustAtPositionZero
			}
			if validateSigner {
				signer, err := ecrecover(header, self.tribe)
				// verify signer
				if err != nil {
					return err
				}
				if from == nil || *from != signer {
					return ErrTribeChiefTxSignerAndBlockSignerNotMatch
				}
			}
			//verify volunteer
			if state != nil {
				if params.IsSIP100Block(header.Number) {
					// TODO SIP100 check volunteer by vrfnp
					volunteerHex := common.Bytes2Hex(tx.Data()[4:])
					volunteer := common.HexToAddress(volunteerHex)
					vrfn := header.Extra[:32]
					if !params.VerifyMiner(header.ParentHash, volunteer, vrfn) {
						return errors.New("verify_volunteer_fail")
					}
				}
				//else {
				//	volunteerHex := common.Bytes2Hex(tx.Data()[4:])
				//	volunteer := common.HexToAddress(volunteerHex)
				//	if volunteer != common.HexToAddress("0x") {
				//		log.Debug("<<TribeStatus.ValidateBlock>> verify_volunteer =>", "num", number, "v", volunteer.Hex())
				//		if err := self.VerifySignerBalance(state, volunteer, parent.Header()); err != nil {
				//			return err
				//		}
				//	}
				//}
			}
			total++
		}
	}
	if total == 0 {
		return ErrTribeMustContainChiefTx
	}

	log.Debug("ValidateBlockp-->", "num", block.NumberU64(), "check_signer", validateSigner)
	return nil
}

func (self *TribeStatus) String() string {
	if b, e := json.Marshal(self); e != nil {
		return "error: " + e.Error()
	} else {
		return "status: " + string(b)
	}
}
