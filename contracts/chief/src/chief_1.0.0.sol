pragma solidity ^0.5.3;

//import "github.com/SmartMeshFoundation/Spectrum/contracts/chief/src/chief_abs_0.5.sol"; // for remix
import "./chief_abs_0.5.sol"; // local

contract ChiefBase {

    string public vsn = "1.0.0";
    address public _owner;
    address[] leaderList;

    modifier owner() {
        require(msg.sender == _owner);
        _;
    }

    //config >>>>
    uint public leaderLimit = 5;
    uint public signerLimit = 17;
    uint public epoch = 6171; // block time 14s, 6171 block = 24H
    uint public volunteerLimit = signerLimit - 1;
    //config <<<<
    //function setEpoch(uint _epoch) public payable owner() {epoch = _epoch;}
    //function setSignerLimit(uint _signerLimit) public payable owner() {signerLimit = _signerLimit;}
    //function setLeaderLimit(uint _leaderLimit) public payable owner() {leaderLimit = _leaderLimit;}
    function setOwner(address newOwner) public payable owner() {_owner = newOwner;}

    function getEpoch() public view returns (uint) {return epoch;}
    function getSignerLimit() public view returns (uint) {return signerLimit;}
    function getLeaderLimit() public view returns (uint) {return leaderLimit;}
    function getVolunteerLimit() public view returns (uint) {return volunteerLimit;}
    function getLeaderList() public view returns (address[] memory) {return leaderList;}
    function getOwner() public view returns (address) {_owner;}

    constructor(uint _epoch,uint _signerLimit) public {
        _owner = msg.sender;
        epoch = _epoch;
        signerLimit = _signerLimit;
    }

    function appendLeader(address addr) public payable owner() {
        if (leaderList.length < leaderLimit) {
            for (uint i=0;i<leaderList.length;i++){
                require(leaderList[0]!=addr);
            }
            leaderList.push(addr);
        }
    }

    function removeLeader(address addr) public payable owner() {
        require(leaderList.length!=1);
        for (uint i=0;i<leaderList.length;i++){
            if(leaderList[0]==addr) {
                for (uint j=i;j<leaderList.length-1;j++) {
                    leaderList[j] = leaderList[j+1];
                }
                i--;
                leaderList.length -= 1;
            }
        }
    }


    // generate a random index for remove signers every epoch
    function getRandomIdx(address addr, uint max) private view returns (uint256) {
        if (max <= 0) {
            return 0;
        }
        bytes32 z = keccak256(abi.encodePacked(int160(addr), block.difficulty, now));
        uint256 random = uint256(z);
        return (random % max);
    }

}

//TODO
contract TribeChief_1_0_0 is Chief {

    string vsn = "1.0.0";

    ChiefBase private base;

    constructor(address baseAddress) public {
        base = ChiefBase(baseAddress);
        address[] memory leaderList = base.getLeaderList();
        if (leaderList.length > 0) {
            _signerList[0] = leaderList[0];
        }
    }

    address[] _signerList;
    address[] _blackList;
    address[] _volunteerList;
    address[] history;

    mapping(address => bool) genesisSigner; // genesis signer address
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

    // the mapping of the signer score and the block number
    mapping(address => SignerInfo) signersMap;
    // the mapping of the blacklist and block number
    mapping(address => uint) blMap;

    mapping(address => VolunteerInfo) volunteersMap;


    modifier allow() {
        address _addr = msg.sender;
        require(uint160(_addr) != uint160(0));
        require(signersMap[_addr].score > 0);
        _;
    }

    // append a signer
    function pushSigner(address signer, uint score) private {
        if (_signerList.length < base.getSignerLimit()) {
            _signerList.push(signer);
            signersMap[signer].score = score;
            signersMap[signer].number = block.number;
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

    // append a volunteer
    function pushVolunteer(address volunteer, uint weight) private {
        if (weight == 0) {
            // weight == 0 表示要删除这个候选人，并放到黑名单里冷静一个 epoch
            if (_volunteerList.length > 0) {
                for (uint i = 0; i < _volunteerList.length; i++) {
                    if (volunteer == _volunteerList[i]) {
                        deleteVolunteer(i);
                        break;
                    }
                }
            }
            pushBlackList(volunteer);
        } else if (!genesisSigner[volunteer] && weight == 5 && _volunteerList.length < volunteerLimit && volunteersMap[volunteer].number == 0 && blMap[volunteer] == 0 && signersMap[volunteer].number == 0) {
            //满员或者已经存在于签名人or候选人则不添加
            _volunteerList.push(volunteer);
            volunteersMap[volunteer].weight = weight;
            volunteersMap[volunteer].number = block.number;
        } else if (weight < 5 && volunteersMap[volunteer].number > 0) {
            //这种情况只是为了给 weight - 1 ，所以无条件修改
            volunteersMap[volunteer].weight = weight;
            volunteersMap[volunteer].number = block.number;
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

    // append a sinner to blacklist
    function pushBlackList(address sinner) private {
        if (uint160(sinner) != uint160(0) && blMap[sinner] == 0) {
            _blackList.push(sinner);
            blMap[sinner] = block.number;
        }
    }

    function _cleanBlacklist() private {
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
    function generateSignersRule3() private {
        address g = _signerList[0];
        // 清理旧的列表
        address[] memory sl = new address[](_signerList.length);
        for (uint j = 0; j < sl.length; j++) {
            sl[j] = _signerList[j];
        }
        for (uint i0 = sl.length; i0 > 0; i0--) {
            uint sIndex = i0 - 1;
            deleteSigner(sIndex);

            address signerI = sl[sIndex];
            if (sIndex > 0 && uint160(signerI) != uint160(0)) {
                if (volunteersMap[signerI].weight == 0) {
                    pushVolunteer(signerI, 5);
                }
                pushVolunteer(signerI, volunteersMap[signerI].weight - 1);
            }
        }
        // 顺序选出一个创世签名人放到首位
        if (genesisSigner[g] && _genesisSignerList.length > 1) {
            // 这个循环一定会找到一个 genesisSigner 放到 signers 中
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
        // 随机填满整个 signerList , 走到这个逻辑时 volunteer 一定比 signers 多，所以一定能填满
        // 这个地方循环下来很可能造成 signerList.length < signerLimit 的情况, 后续需要补充
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
        // 如果不满需要补满
        if (ii < signerLimit) {
            for (uint i3 = 0; i3 < _volunteerList.length; i3++) {
                //不存在就放进去
                if (signersMap[_volunteerList[i3]].number == 0) pushSigner(_volunteerList[i3], 3);
                //放满了就退出循环
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
    function updateRule3() private {
        uint l = _signerList.length;
        uint signerIdx = uint(blockNumber % l);
        address si = _signerList[signerIdx];
        //1 : 初始签名人不做处理,不是正常签名人 0 分放回志愿者列表,否则 weight - 1
        if (signerIdx > uint(0)) {
            // 序号对应的不是我，把序号对应的 signer 以 weight=0 扔回志愿者列表 (其实就是删除)
            if (msg.sender != si) {
                pushVolunteer(si, 0);
                //此处还不能直接删除，因为不能破坏列表长度，否则对后续取模逻辑有影响，用 0 占位吧
                delete signersMap[si];
                _signerList[signerIdx] = address(uint160(0));
            }
        }

        //2 : 如果当前块是签名人列表的最后一块，则生成下一轮的列表
        if (signerIdx == uint(l - 1)) {
            //if (volunteersMap[msg.sender].weight == 0) {pushVolunteer(msg.sender, 5);}
            //pushVolunteer(msg.sender, volunteersMap[msg.sender].weight - 1);
            generateSignersRule3();
        }
    }


    function update(address volunteer) public allow() {
        blockNumber = block.number;
        if (blockNumber > base.getEpoch() && blockNumber % base.getEpoch() == 0) {
            _cleanBlacklist();
        }
        if (uint160(volunteer) != uint160(0)) {
            pushVolunteer(volunteer, 5);
        }
        updateRule3();
    }

    function getStatus() public view returns (
    //address[] volunteerList,
        address[] memory signerList,
        address[] memory blackList, // vsn 0.0.3
        uint[] memory scoreList,
        uint[] memory numberList,
        uint totalVolunteer,
        uint number
    ) {
        scoreList = new uint[](_signerList.length);
        numberList = new uint[](_signerList.length);
        for (uint i = 0; i < _signerList.length; i ++) {
            scoreList[i] = signersMap[_signerList[i]].score;
            numberList[i] = signersMap[_signerList[i]].number;
        }
        blackList = _blackList;
        signerList = _signerList;
        // vsn 0.0.3
        number = blockNumber;
        totalVolunteer = _volunteerList.length;
    }

    function version() public view returns (string memory) {vsn;}
    function getSignerLimit() public view returns (uint) {base.getSignerLimit();}
    function getEpoch() public view returns (uint) {base.getEpoch();}
    function getVolunteerLimit() public view returns (uint) {base.getVolunteerLimit();}

    function getVolunteers() public view returns (
        address[] memory volunteerList,
        uint[] memory weightList,
        uint length
    ){
        weightList = new uint[](_volunteerList.length);
        volunteerList = _volunteerList;
        length = volunteerList.length;
    }



    // TODO
    function filterVolunteer(address[] memory volunteers) public view returns (uint[] memory result) {}

}