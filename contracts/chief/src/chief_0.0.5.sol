// Copyright 2018 The Smartmesh Chain Authors
// This file is part of the Spectrum library.
pragma solidity ^0.4.19;

/*
vsn 0.0.3 改进如下 :
    每个 epoch 增加了一个 黑名单，用来保存被淘汰的 signers ，以保证其不会重复被选拔;
    在每个 epoch 中按打包交易数量来清除 1/3 非 genesisSigners 的节点 ；

vsn 0.0.4 修改 0.0.3 中的 bug
vsn 0.0.5 修正bug , 调整修改 singerLimit 和 volunteerLimit 的逻辑
*/
contract TribeChief_0_0_5 {

    string vsn = "0.0.5";

    //config >>>>
    uint init_epoch = 5760;
    uint init_singerLimit = 17;
    uint init_volunteerLimit = 70;

    uint epoch = init_epoch; // 5760 = 24H ; 11520 = 48H
    mapping(address => bool) genesisSigner; // genesis signer address
    uint singerLimit = init_singerLimit;
    uint volunteerLimit = init_volunteerLimit;
    //config <<<<

    uint blockNumber;

    //signer info
    struct SignerInfo {
        uint score;
        uint number;
    }

    address _owner;

    address[] _signerList;
    address[] _volunteerList;
    address[] _blackList;

    // the mapping of the signer score and the block number
    mapping(address => SignerInfo) signersMap;
    // the mapping of the volunteer and block number
    mapping(address => uint) volunteersMap;
    // the mapping of the blacklist and block number
    mapping(address => uint) blMap;

    function TribeChief_0_0_5(address[] genesisSigners) public {
        _owner = msg.sender;
        uint len = genesisSigners.length;
        if (len > 0) {
            for (uint i = 0; i < len; i++) {
                address g = genesisSigners[i];
                genesisSigner[g] = true;
                pushSigner(g, 3);
            }
        } else {
            // normal default for testing
            address g1 = 0x4110bd1ff0b73fa12c259acf39c950277f266787;
            // nerver delete genesis signer
            genesisSigner[g1] = true;
            pushSigner(g1, 3);
        }
    }

    // delete a blacklist by index
    function deleteBlackList(uint index) private {

        uint blen = _blackList.length;
        if (index < blen) {
            delete blMap[_blackList[index]];
            for (uint i = index; i < blen - 1; i++) {
                _blackList[i] = _blackList[i + 1];
            }
            _blackList.length -= 1;
        }
    }

    // delete a volunteer by index
    function deleteVolunteer(uint index) private {

        uint vlen = _volunteerList.length;
        // _signerList >>>>
        if (index < vlen) {
            delete volunteersMap[_volunteerList[index]];
            for (uint i = index; i < vlen - 1; i++) {
                _volunteerList[i] = _volunteerList[i + 1];
            }
            _volunteerList.length -= 1;
        }
    }

    // delete a signer by index
    function deleteSigner(uint index) private {
        uint slen = _signerList.length;
        if (index < slen) {
            delete signersMap[_signerList[index]];
            for (uint i = index; i < slen - 1; i++) {
                _signerList[i] = _signerList[i + 1];
            }
            _signerList.length -= 1;
        }
    }

    // append a sinner to blacklist
    function pushBlackList(address sinner) private {

        if (blMap[sinner] == 0) {
            _blackList.push(sinner);
            blMap[sinner] = block.number;
        }
    }

    // append a volunteer
    function pushVolunteer(address volunteer) private {
        //满员或者已经存在于签名人or候选人则不添加
        //vsn-0.0.3 : blMap is blacklist map , the new volunteer can not in blacklist
        if (_volunteerList.length < volunteerLimit && volunteersMap[volunteer] == 0 && blMap[volunteer] == 0 && signersMap[volunteer].number == 0) {
            _volunteerList.push(volunteer);
            volunteersMap[volunteer] = block.number;
        }
    }

    // generate a random index for remove signers every epoch
    function getRandomIdx(address addr, uint max) private view returns (uint256) {
        if (max <= 0) {
            return 0;
        }
        uint256 random = uint256(keccak256(addr, block.difficulty, now));
        return (random % max);
    }

    // append a signer
    function pushSigner(address signer, uint score) private {

        if (_signerList.length < singerLimit) {
            _signerList.push(signer);
            signersMap[signer].score = score;
            signersMap[signer].number = block.number;
        }
    }

    modifier apply(address _addr) {
        require(signersMap[_addr].number > 0);
        _;
    }
    modifier owner(address _addr) {
        require(_addr == _owner);
        _;
    }

    // v0.0.5 >>>>
    // before "apply" ,after "owner"
    function setVolunteerLimit(uint n) public owner(msg.sender) {

        require (n >= init_volunteerLimit);

        uint vlen = _volunteerList.length;

        if (n < vlen){
            for (uint i = n; i < vlen; i ++){
                delete volunteersMap[_volunteerList[i]];
            }
            //删除尾部多余
            _volunteerList.length -= (vlen - n);
        }

        volunteerLimit = n;
    }

    function setSingerLimit(uint n) public owner(msg.sender) {

        require (n >= init_singerLimit);

        uint slen = _signerList.length;

        if (n < slen){
            for(uint i = n; i < slen; i ++){
                delete signersMap[_signerList[i]];
            }
            _signerList.length -= (slen - n);
        }

        singerLimit = n;
    }

    function setEpoch(uint n) public owner(msg.sender) {

        require (n >= init_epoch);

        epoch = n;
    }
    // v0.0.5 <<<<

    function repeatTi(uint[] tiList, uint ti) private pure returns (bool) {
        if (tiList.length > 0) {
            for (uint i = 0; i < tiList.length; i++) {
                if (tiList[i] == ti) {
                    return true;
                }
            }
        }
        return false;
    }

    // v0.0.2
    function _cleanVolunteerList() private {
        uint vlen = _volunteerList.length;
        for (uint i1 = 0; i1 < vlen; i1++) {
            delete volunteersMap[_volunteerList[i1]];
        }
        delete _volunteerList;
    }
    // v0.0.4
    function _cleanBlacklist() private {
        // 1 : clean blacklist
        uint blen = _blackList.length;
        for (uint i2 = 0; i2 < blen; i2++) {
            delete blMap[_blackList[i2]];
        }
        delete _blackList;
    }
    // v0.0.4
    function _moveSignersToBlacklist() private {
        uint counter = 0;
        uint[] memory tiList = new uint[](slen);
        uint slen = _signerList.length;
        // target signer idx
        for (uint i3 = 0 ; i3 < slen ; i3++) {
            address _addr = _signerList[i3];
            uint ti = getRandomIdx(_addr, (slen - uint(1)));
            //skip out of range
            if (ti >= slen) {
                continue;
            }
            address _signer = _signerList[ti];
            //skip genesis signer
            if (genesisSigner[_signer])
                continue;
            //skip repeat
            if (repeatTi(tiList, ti))
                continue;
            tiList[counter] = ti;
            if (counter >= (slen / uint(3)))
                break;
            counter += uint(1);
        }
        if (counter > 0) {
            for (uint i4 = 0; i4 < slen; i4++) {
                uint idx = tiList[i4];
                // skip nil , 0 == nil
                if (idx != 0) {
                    pushBlackList(_signerList[tiList[i4]]);
                    deleteSigner(tiList[i4]);
                }
            }
        }
    }

    function update(address volunteer) public apply(msg.sender) {

        blockNumber = block.number;

        // vsn-0.0.2 : every epoch be clean volunteers
        // vsn-0.0.3 : clean blacklist and move 1/3 signers to blacklist excelude genesisSigners
        if (block.number > epoch && block.number % epoch == 0) {
            // ==== vsn-0.0.2 ====
            _cleanVolunteerList();
            // ==== vsn-0.0.3 ====
            // ==== vsn-0.0.4 ====
            // 1 : clean blacklist
            _cleanBlacklist();
            // 2 : move 1/3 signers to blacklist
            _moveSignersToBlacklist();
        }

        // mine
        // 如果当前块 不是 signers[ blockNumber % signers.length ] 出的，就给这个 signer 减分
        // 否则恢复成 3 分
        uint signerIdx = blockNumber % _signerList.length;
        //初始签名人不做处理
        if (!genesisSigner[_signerList[signerIdx]]) {

            SignerInfo storage signer = signersMap[_signerList[signerIdx]];

            // 序号对应的不是我，则扣它一分
            if (msg.sender != _signerList[signerIdx]) {
                if (signer.score > 1) {
                    signer.score -= 1;
                    signer.number = blockNumber;
                } else {
                    // move to blacklist and cannot be selected in this epoch
                    pushBlackList(_signerList[signerIdx]);
                    // vsn-0.0.3
                    // score == 0 , remove on signerList
                    deleteSigner(signerIdx);
                }
            } else {
                // 恢复分数
                signer.score = 3;
            }
        }

        // tag
        if (volunteer != uint160(0)) {
            pushVolunteer(volunteer);
        }

        //是否提拔一个人到签名人
        if (_signerList.length < singerLimit && _volunteerList.length > 0) {

            //将候选人列表首个添加到签名人列表
            pushSigner(_volunteerList[0], 3);

            //删除该候补志愿者
            deleteVolunteer(0);
        }
    }

    function version() public constant returns (string) {
        return vsn;
    }

    function getSignerLimit() public constant returns (uint) {
        return singerLimit;
    }

    function getEpoch() public constant returns (uint) {
        return epoch;
    }

    function getVolunteerLimit() public constant returns (uint) {
        return volunteerLimit;
    }

    function getStatus() public constant returns (
        address[] volunteerList,
        address[] signerList,
        address[] blackList, // vsn 0.0.3
        uint[] memory scoreList,
        uint[] memory numberList,
        uint number
    ) {
        scoreList = new uint[](_signerList.length);
        numberList = new uint[](_signerList.length);
        for (uint i = 0; i < _signerList.length; i ++) {
            scoreList[i] = signersMap[_signerList[i]].score;
            numberList[i] = signersMap[_signerList[i]].number;
        }
        volunteerList = _volunteerList;
        signerList = _signerList;
        blackList = _blackList;
        // vsn 0.0.3
        number = blockNumber;
        return;
    }

}