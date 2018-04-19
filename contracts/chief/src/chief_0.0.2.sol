// Copyright 2018 The Smartmesh Chain Authors
// This file is part of the Spectrum library.
pragma solidity ^0.4.19;


contract TribeChief {

    string vsn = "0.0.2";

    //配置 >>>>
    uint epoch = 11520; // 48H
    mapping(address => bool) genesisSigner; //创世签名节点
    uint singerLimit = 3;
    uint volunteerLimit = 4;
    //配置 <<<<

    uint blockNumber;

    //签名人信息
    struct SignerInfo {
        uint score;
        uint number;
    }

    address _owner;

    //签名人列表
    address[] _signerList;

    //候补人列表
    address[] _volunteerList;

    //签名人分数和变更区块号的映射
    mapping(address => SignerInfo) signersMap;

    //候选人变更区块号的映射
    mapping(address => uint) volunteersMap;

    // 删除一个候补志愿者
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

    // 删除一个签名人
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

    // 增加一个志愿者
    function pushVolunteer(address volunteer) private {

        //满员或者已经存在于签名人or候选人则不添加
        if (_volunteerList.length < volunteerLimit && volunteersMap[volunteer] == 0 && signersMap[volunteer].number == 0) {
            _volunteerList.push(volunteer);
            volunteersMap[volunteer] = block.number;
        }
    }

    function pushSigner(address signer, uint score) private {

        if (_signerList.length < singerLimit) {
            _signerList.push(signer);
            signersMap[signer].score = score;
            signersMap[signer].number = block.number;
        }
    }

    function TribeChief(address[] genesisSigners) public {
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
            address g2 = 0xca012e2facf405885293d8d3ba3f14fae1e58b71;
            address g3 = 0xadb3ea3ad356199206ca817b04fd668cc5196df2;

            // nerver delete genesis signer
            genesisSigner[g1] = true;

            pushSigner(g1, 3);
            pushSigner(g2, 3);
            pushSigner(g3, 3);
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

    function setVolunteerLimit(uint n) public apply(msg.sender) {
        volunteerLimit = n;
    }

    function setSingerLimit(uint n) public apply(msg.sender) {
        singerLimit = n;
    }

    function update(address volunteer) public apply(msg.sender) {

        blockNumber = block.number;

        // every epoch be clean volunteers
        if (block.number > epoch && block.number % epoch == 0) {

            uint vlen = _volunteerList.length;
            for (uint i = 0; i < vlen; i++) {
                delete volunteersMap[_volunteerList[i]];
            }

            delete _volunteerList;
        }

        // mine
        // 如果当前块 不是 signers[ blockNumber % signers.length ] 出的，就给这个 signer 减分
        // 否则恢复成 3 分

        // 序号
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
                    // 0 分时就删除了
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

    function version() constant returns (string) {
        return vsn;
    }
    function getSignerLimit() constant returns (uint) {
        return singerLimit;
    }
    function getVolunteerLimit() constant returns (uint) {
        return volunteerLimit;
    }

    function getStatus() constant returns (
        address[] volunteerList,
        address[] signerList,
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
        number = blockNumber;
        return;
    }

}
