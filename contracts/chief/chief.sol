// Copyright 2018 The Smartmesh Chain Authors
// This file is part of the SMChain library.
pragma solidity ^0.4.19;

contract TribeChief {

    address private owner;
    //配置 >>>>
    uint256 epoch = 5; // TODO 调试环境设置为 50 , 生产环境为 30000
    int blockPeriod = 15; // 出块时间
    mapping(address => bool) genesisSigner; //创世签名节点
    mapping(int => address) recents; //当前 epoch 最新的块的签名者集合
    //配置 <<<<
    uint constant singerLimit = 111;
    uint constant volunteerLimit = 333;
    uint256 blockNumber;

    //用偏移量来找 signer 对应的 score
    struct Signers {
        address[] signerList; //签名人列表
        int[] scoreList; //签名人分数
        //只用于查询，没有其他用途
        mapping(address => int) scoremap;
        uint256[] numberList; //日志，记录本次列表变动在哪个块上
    }

    //候补志愿者集合，每个 epoch 清理一次
    struct Volunteers {
        address[] volunteerList;
        uint256[] numberList; //日志，记录本次列表变动在哪个块上
        mapping(address => uint) vmap; //用来去重
    }

    //签名人集合
    Signers signers;
    //候补志愿者集合
    Volunteers volunteers;

    function getVolunteerIdx(address volunteer) private returns (uint, bool){
        uint l = volunteers.volunteerList.length;
        for (uint i = 0; i < l - 1; i++) {
            if (volunteer == volunteers.volunteerList[i]) {
                return (i, true);
            }
        }
        return (0, false);
    }

    function verifyVolunteer(address volunteer) private returns (bool) {
        //TODO
        return true;
    }

    // 删除一个候补志愿者
    function deleteVolunteer(address volunteer) private {
        uint l = volunteers.volunteerList.length;
        uint idx = volunteers.vmap[volunteer];
        // signerList >>>>
        if (idx > 0) {
            uint _idx = idx - 1;
            for (uint i = _idx; i < l - 1; i++) {
                volunteers.volunteerList[i] = volunteers.volunteerList[i + 1];
                volunteers.numberList[i] = volunteers.numberList[i + 1];
            }
            delete volunteers.volunteerList[l - 1];
            delete volunteers.numberList[l - 1];
            volunteers.volunteerList.length -= 1;
            volunteers.numberList.length -= 1;
            delete volunteers.vmap[volunteer];
        }
        // signerList <<<<
    }

    //删除时，要把 signerlist 、 scorelist 、 scoremap 都删一遍
    function deleteSigner(uint index) private {
        uint len = signers.signerList.length;
        if (index >= len) return;

        // scoremap >>>>
        delete signers.scoremap[signers.signerList[index]];
        // scoremap <<<<
        // >>>>
        for (uint i = index; i < len - 1; i++) {
            signers.signerList[i] = signers.signerList[i + 1];
            signers.scoreList[i] = signers.scoreList[i + 1];
            signers.numberList[i] = signers.numberList[i + 1];
        }
        delete signers.signerList[len - 1];
        delete signers.scoreList[len - 1];
        delete signers.numberList[len - 1];
        signers.signerList.length -= 1;
        signers.scoreList.length -= 1;
        signers.numberList.length -= 1;
        // <<<<
    }


    // 增加一个志愿者
    function pushVolunteer(address volunteer) private returns (uint, bool){
        if (volunteers.volunteerList.length >= volunteerLimit) {
            return (0, false);
        }
        deleteVolunteer(volunteer);
        volunteers.volunteerList.push(volunteer);
        volunteers.numberList.push(block.number);
        volunteers.vmap[volunteer] = volunteers.volunteerList.length;
        return (volunteers.volunteerList.length, true);
    }

    function pushSigner(address signer, int score) private returns (uint, bool){
        Signers self = signers;
        if (self.signerList.length >= singerLimit) {
            return (0, false);
        }
        self.signerList.push(signer);
        self.scoreList.push(score);
        self.numberList.push(block.number);
        self.scoremap[signer] = score;
        return (self.signerList.length, true);
    }

    function TribeChief(address[] genesisSigners) public {
        owner = msg.sender;
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
            address g2 = 0xadb3ea3ad356199206ca817b04fd668cc5196df2;
            address g3 = 0xb94b3aa41609e3f59cbaff3c2c298c6cc4c50b81;
            genesisSigner[g1] = true;
            genesisSigner[g2] = true;
            genesisSigner[g3] = true;
            pushSigner(g1, 3);
            pushSigner(g2, 3);
            pushSigner(g3, 3);
        }
    }

    modifier apply(address _addr) {
        require(signers.scoremap[_addr] > 0);
        _;
    }

    //function update(address volunteer) public apply(msg.sender) {
    function update(address volunteer) public {
        blockNumber = block.number;
        //if (block.number > epoch && block.number % epoch == 0) {
        //    delete volunteers.volunteerList;
        //    delete volunteers.numberList;
        //}
        //// tag
        if (volunteer != uint160(0) && verifyVolunteer(volunteer)) {
            pushVolunteer(volunteer);
        }

        int score = signers.scoreList[0];
        if (score > 0) {
            score = score - 1;
            signers.scoreList[0] = score;
            signers.numberList[0] = block.number;
        } else {
            deleteSigner(0);
        }
    }

    function getStatus() constant returns (
        address[] volunteerList,
        address[] signerList,
        int[] scoreList,
        uint256[] numberList,
        uint256 number
    ) {
        volunteerList = volunteers.volunteerList;
        signerList = signers.signerList;
        scoreList = signers.scoreList;
        numberList = signers.numberList;
        number = blockNumber;
        return;
    }

}
