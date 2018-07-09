pragma solidity ^0.4.19;

contract TribeChief_0_0_6 {

    string debugFunc = "";

    string vsn = "0.0.6";

    //config >>>>


    // block time 15s, 5760 block = 24H
    // block time 14s, 6171 block = 24H
    uint epoch = 6171;
    //uint epoch = 16;
    // epoch / signerLimit == 363
    //uint signerLimit = 17;
    uint signerLimit = 3;
    uint volunteerLimit = 300;

    mapping(address => bool) genesisSigner; // genesis signer address
    //config <<<<

    uint blockNumber;

    //signer info
    struct SignerInfo {
        uint score;
        uint number;
    }

    //volunteer object
    struct VolunteerInfo {
        uint weight; // new volunteer weight = 5
        uint number;
    }

    address _owner;

    address[] _genesisSignerList;
    address[] _signerList;
    address[] _volunteerList;
    address[] _blackList;

    // the mapping of the signer score and the block number
    mapping(address => SignerInfo) signersMap;
    // the mapping of the volunteer and block number
    //mapping(address => uint) volunteersMap;
    mapping(address => VolunteerInfo) volunteersMap;
    // the mapping of the blacklist and block number
    mapping(address => uint) blMap;

    function TribeChief_0_0_6(address[] genesisSigners) public {
        _owner = msg.sender;
        uint len = genesisSigners.length;
        if (len > 0) {
            for (uint i = 0; i < len; i++) {
                address g = genesisSigners[i];
                genesisSigner[g] = true;
                _genesisSignerList.push(g);
                pushSigner(g, 3);
            }
        } else {
            // normal default for testing
            // 0x4110bd1ff0b73fa12c259acf39c950277f266787;
            address g1 = uint160(371457687117486736155821182390123011782146942855);
            genesisSigner[g1] = true;
            _genesisSignerList.push(g1);
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
        if (sinner != uint160(0) && blMap[sinner] == 0) {
            _blackList.push(sinner);
            blMap[sinner] = block.number;
        }
    }


    // append a volunteer
    function pushVolunteer(address volunteer, uint weight) private {
        if (weight == 0) {
            if (_volunteerList.length > 0) {
                for (uint i = 0; i < _volunteerList.length; i++) {
                    if (volunteer == _volunteerList[i]) {
                        deleteVolunteer(i);
                        break;
                    }
                }
            }
            pushBlackList(volunteer);
        } else if (weight == 5 && _volunteerList.length < volunteerLimit && volunteersMap[volunteer].number == 0 && blMap[volunteer] == 0 && signersMap[volunteer].number == 0) {
            _volunteerList.push(volunteer);
            volunteersMap[volunteer].weight = weight;
            volunteersMap[volunteer].number = block.number;
        } else if (weight < 5 && volunteersMap[volunteer].number > 0) {
            volunteersMap[volunteer].weight = weight;
            volunteersMap[volunteer].number = block.number;
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

        if (_signerList.length < signerLimit) {
            _signerList.push(signer);
            signersMap[signer].score = score;
            signersMap[signer].number = block.number;
        }
    }

    modifier apply(address _addr) {
        require(_addr != uint160(0));
        require(signersMap[_addr].score > 0);
        _;
    }
    modifier owner(address _addr) {
        require(_addr == _owner);
        _;
    }

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


    uint[] _cleanIdx; //辅助清理
    /*
        这个方法在每个 epoch 时负责清理志愿者列表
    */
    function _cleanVolunteerList() private {
        _cleanIdx.length = 0;
        uint vlen = _volunteerList.length;
        if (vlen > volunteerLimit / 2) {
            for (uint i1 = 0; i1 < vlen; i1++) {
                if (volunteersMap[_volunteerList[i1]].weight < 5) {
                    _cleanIdx.push(i1);
                }
            }
            if (_cleanIdx.length > volunteerLimit / 2) {
                uint total = _cleanIdx.length - (volunteerLimit / 2);
                uint[] memory tiList = new uint[](total);
                for (uint i2 = 0; i2 < _cleanIdx.length; i2++) {
                    uint ti = getRandomIdx(_volunteerList[i2], (_cleanIdx.length - uint(1)));
                    //skip out of range
                    if (ti >= _cleanIdx.length) continue;
                    //skip repeat
                    if (repeatTi(tiList, ti)) continue;
                    tiList[total] = ti;
                    if (total == 0) break;
                    total -= 1;
                }
                for (uint i3 = 0; i3 < tiList.length; i3++) {
                    uint idx = tiList[i3];
                    deleteVolunteer(idx);
                    address volunteer = _volunteerList[idx];
                    pushBlackList(volunteer);
                }
            }
        }
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

    /*
        在志愿者列表中随机选出17个节点替换当前列表,
        在进入这个方法之前，已经判断过志愿者列表尺寸了，所以这里只管随机拼装即可
    */

    //TODO : function generateSignersRule3() private {
    function generateSignersRule3() public {
        address g = _signerList[0];
        for (uint i0 = _signerList.length; i0 > 0; i0--) {
            deleteSigner(i0 - 1);
        }
        if (genesisSigner[g] && _genesisSignerList.length > 1) {
            for (uint i1 = 0; i1 < _genesisSignerList.length; i1++) {
                if (_genesisSignerList[i1] == g) {
                    if (i1 == (_genesisSignerList.length - 1)) {
                        pushSigner(_genesisSignerList[0], 3);
                    } else {
                        pushSigner(_genesisSignerList[i1 + 1], 3);
                    }
                    break;
                }
            }
        } else {
            pushSigner(_genesisSignerList[0], 3);
        }
        uint[] memory tiList = new uint[](signerLimit);
        uint ii = 0;
        for (uint i2 = 0; i2 < _volunteerList.length; i2++) {
            if (ii >= signerLimit) break;
            uint ti = getRandomIdx(_volunteerList[i2], _volunteerList.length - uint(1));
            if (repeatTi(tiList, ti)) continue;
            pushSigner(_volunteerList[ti], 3);
            tiList[ii] = ti;
            ii = ii + 1;
        }
        if (ii < signerLimit) {
            for (uint i3 = 0; i3 < _volunteerList.length; i3++) {
                if (signersMap[_volunteerList[i3]].number == 0) pushSigner(_volunteerList[i3], 3);
                if (_signerList.length >= signerLimit) break;
            }
        }
    }


    /* rule 3 : 出块节点列表已满，候选节点列表大于出块节点列表

        在这个规则生效时，签名节点的分数已经没有意义了，
        此时的规则是每出一轮块就要替换掉全部的出块节点，
        从候选节点列表中按 weight 随机提拔一批新的出块节点到出块节点列表，
        将原出块节点列表移入候选节点列表，并将 weight - 1，
        当 weight == 0 时则移入黑名单，等待下一个 epoch
        假设出块节点列表最大长度 17 ，候选节点列表最大长度与 epoch 相等。每一轮出块，指的就是每 17 个块，每笔交易的确认时间也是 17 块，但是对于交易所来说应该至少经过 34 个块才能确认一笔交易。
    */
    function updateRule3() public {
        blockNumber = block.number;
        uint l = _signerList.length;
        uint signerIdx = uint(blockNumber % l);
        address si = _signerList[signerIdx];
        if (signerIdx > uint(0)) {
            if (msg.sender == si) {
                pushVolunteer(msg.sender, volunteersMap[msg.sender].weight - 1);
            } else {
                pushVolunteer(si, 0);
                delete signersMap[si];
                _signerList[signerIdx] = uint160(0);
            }
        }

        if (signerIdx == uint(l - 1)) {
            generateSignersRule3();
        }
    }

    /*
    rule 1 : 出块节点列表未满
        每个节点3分，每错出或漏出一个块扣1分，0分时被放入黑名单
        在当前 epoch 不再被选拔

    rule 2 : 出块节点列表已满，候选节点列表小于出块节点列表
        此时主要工作是选拔候选节点，为每个被选拔的节点设置 weight = 5，
        出块规则与 “出块节点列表未满” 时的规则相同
    */
    function updateRule1() private {
        debugFunc = "updateRule1";
        blockNumber = block.number;
        uint signerIdx = blockNumber % _signerList.length;
        if (!genesisSigner[_signerList[signerIdx]]) {

            SignerInfo storage signer = signersMap[_signerList[signerIdx]];

            if (msg.sender != _signerList[signerIdx]) {
                if (signer.score > 1) {
                    signer.score -= 1;
                    signer.number = blockNumber;
                } else {
                    // move to blacklist and cannot be selected in this epoch
                    pushVolunteer(_signerList[signerIdx], 0);
                    // vsn-0.0.3
                    // score == 0 , remove on signerList
                    deleteSigner(signerIdx);
                }
            } else {
                signer.score = 3;
            }
        }

        if (_signerList.length < signerLimit && _volunteerList.length > 0) {
            pushSigner(_volunteerList[_volunteerList.length - 1], 3);
            deleteVolunteer(_volunteerList.length - 1);
        }
    }

    //function update(address volunteer) public apply(msg.sender) {
    function update(address volunteer) public {
        debugFunc = "update";
        if (block.number > epoch && block.number % epoch == 0) {
            _cleanBlacklist();
            _cleanVolunteerList();
        }

        if (volunteer != uint160(0)) {
            pushVolunteer(volunteer, 5);
        }

        if (_signerList.length < signerLimit || _volunteerList.length < _signerList.length) {
            updateRule1();
        } else {
            updateRule3();
        }
    }

    function version() public constant returns (string) {
        return vsn;
    }

    function getSignerLimit() public constant returns (uint) {
        return signerLimit;
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

    // ================
    // TEST AND DEBUG
    // ================

    function getDebugData() public constant returns (address, address, bool) {
        uint signerIdx = blockNumber % _signerList.length;
        if (uint160(msg.sender) != uint160(_signerList[signerIdx])) {
        }
        return (msg.sender, _signerList[signerIdx], msg.sender == _signerList[signerIdx]);
    }

    function getVolunteers() public constant returns (
        address[] volunteerList,
        uint[] weightList
    ){
        weightList = new uint[](_volunteerList.length);
        volunteerList = _volunteerList;
        for (uint i = 0; i < _volunteerList.length; i++) {
            weightList[i] = volunteersMap[_volunteerList[i]].weight;
        }
        return;
    }

    function getBlockNumber() public constant returns (uint){
        return block.number;
    }

    function fillSignerForTest() public {
        //TODO : for test >>>>
        address g2 = uint160(371457687117486736155821182390123011782146942856);
        genesisSigner[g2] = true;
        _genesisSignerList.push(g2);
        address g3 = uint160(371457687117486736155821182390123011782146942857);
        genesisSigner[g3] = true;
        _genesisSignerList.push(g3);
        address g4 = uint160(371457687117486736155821182390123011782146942858);
        genesisSigner[g4] = true;
        _genesisSignerList.push(g4);
        //TODO : for test <<<<

        //0xca35b7d915458ef540ade6068dfe2f44e8fa733c
        pushSigner(uint160(1154414090619811796818182302139415280051214250812), 3);
        //0xca35b7d915458ef540ade6068dfe2f44e8fa733d
        pushSigner(uint160(1154414090619811796818182302139415280051214250813), 3);

        blockNumber = block.number;
        fillVolunteerForTest();
    }

    function fillVolunteerForTest() public {
        //0xca35b7d915458ef540ade6068dfe2f44e8fa7330
        uint160 b = uint160(1154414090619811796818182302139415280051214250800);
        uint n = now;
        for (uint i = n; i < n + 10; i++) {
            pushVolunteer(uint160(b + i), 5);
        }
    }

    function pushVolunteerTest5(address volunteer) public {
        pushVolunteer(volunteer, 5);
    }

    function pushVolunteerTest1(address volunteer) public {
        pushVolunteer(volunteer, 1);
    }

    function pushVolunteerTest0(address volunteer) public {
        pushVolunteer(volunteer, 0);
    }
}