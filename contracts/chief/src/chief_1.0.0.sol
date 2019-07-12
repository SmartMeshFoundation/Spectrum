pragma solidity >=0.5.0 <0.6.0;

/* for remix
import "github.com/SmartMeshFoundation/Spectrum/contracts/chief/src/chief_abs_s0.5.sol"; // for remix
import "github.com/SmartMeshFoundation/Spectrum/contracts/chief/src/chief_base_s0.5_v0.0.1.sol"; // for remix
*/


/* local */
import "./chief_abs_s0.5.sol";
import "./chief_base_s0.5_v0.0.1.sol";

/*
----------
| devnet |
----------   

test nodes :

    leader :

        privatekey: 507fd083b5c5af7e645e77a3a3a82708f3af304164e02612ab4b1d5b36c627c6
        address: 0x9944d0cc757cd9391ee320fc17d5b6f61693f2c5

    deposit :

        privatekey: 543e9d0ddcd02b4bbdb2cecd402da99e9532fface57d8ea74e833c5d413f2daa
        address: 0xd3F50Dfd2887B3818aB9E6c6f846ED1d3bD90B29

        privatekey: 9d7b3b8155ea429cb49cdc556aa7b3367feb91ccf51eb1e9c7e2bac67d939f03
        address: 0x50bf23F1d844465AB69357aa08Af0fEA4B1E62D7



f = "0xbf1736a65f8beadd71b321be585fd883503fdeaa"; personal.unlockAccount(f,"123456")

    sign(f):

        addr: 0xd3F50Dfd2887B3818aB9E6c6f846ED1d3bD90B29
        R: 0x99b087c99fceee44b15373847e38f485b31b5b779ce7b01081f2d61c00c77d19
        S: 0x0f29d84ff759415e969e12a49cf2fc453c1b2aa386d59e0488b11627dcdbd261
        V: 27 

        addr: 0x50bf23F1d844465AB69357aa08Af0fEA4B1E62D7
        R: 0x2edbf406a13757485096ae35d85f1f881f17073e230454ad83a65e003ff2deb8
        S: 0x4629840c12909009ad4092f8da67d53b57a7a2c8fa6abe024166982d65821fcb
        V: 27


ChiefBase : 0xad61f1201f592fbf13d2645f9c59d8d5f82a1837 //
    deploy args : 25,5
    appendLeader args : "0x9944d0cc757cd9391ee320fc17d5b6f61693f2c5"

POC : 0x901c0636c4fc83f353bca2db85e2ace886a9416d
    deploy args : "0xad61f1201f592fbf13d2645f9c59d8d5f82a1837",10,10,50
    deposit args :
        "0x99b087c99fceee44b15373847e38f485b31b5b779ce7b01081f2d61c00c77d19","0x0f29d84ff759415e969e12a49cf2fc453c1b2aa386d59e0488b11627dcdbd261",27
        "0x2edbf406a13757485096ae35d85f1f881f17073e230454ad83a65e003ff2deb8","0x4629840c12909009ad4092f8da67d53b57a7a2c8fa6abe024166982d65821fcb",27


Tribe : 0x6d05f6aa4e19e20cd781fa3def97bbfd0b980534
    deploy args : "0xad61f1201f592fbf13d2645f9c59d8d5f82a1837","0x901c0636c4fc83f353bca2db85e2ace886a9416d"

*/

contract TribeChief_1_0_0 is Chief {

    string vsn = "1.0.0";

    struct BlackMember {
        address addr;
        uint score;
        uint number;
    }

    ChiefBase_1_0_0 private base;


    address[] public  _signerList; //当前轮的17块的出块人列表(包含Leader)
    address[] public  _nextRoundSignerList; //下一轮17块出块人列表,不包含Leader
    uint  public  blockNumber; //当前块



    // the mapping of the signer score and the block number
    mapping(address => uint)   signersMap; //标记某个地址在当前轮中是否是signer

    constructor(address baseAddress, address pocAddress,uint startBlockNumber) public {
        base = ChiefBase_1_0_0(baseAddress);
        base.init(pocAddress, address(this));
        blockNumber=startBlockNumber; //从此块开始分叉

        address[] memory leaderList = base.takeLeaderList();
        require(leaderList.length > 0);

        signersMap[leaderList[0]] = 1;
        _signerList.push(leaderList[0]);

        for (uint i = _signerList.length; i < base.takeSignerLimit(); i++) {
            // placeholder
            _signerList.push(address(0));
        }

    }

    modifier allow() {
        address _addr = msg.sender;
        require(uint160(_addr) != uint160(0));
        require(signersMap[_addr] > 0 || base.isLeader(_addr));
        _;
    }

    function pushNextRoundSigner(address addr)  private {
        if (_nextRoundSignerList.length < base.takeVolunteerLimit()) {
            _nextRoundSignerList.push(addr);
        } else{
            revert("next round signer too much");
        }
    }

    // append a signer
    function pushSigner(address signer ) private {
        if (_signerList.length < base.takeSignerLimit()) {
                _signerList.push(signer);
            signersMap[signer] = 1;
        } else{
            revert("too many signer");
        }
    }

    // clean all of list and return _signerList[0]
    function clean_all_signer_and_get0()  private returns (address) {
        address signer0=_signerList[0]; //至少会有一个signer
        // clean all of signerList
        for (uint i = 0;i<_signerList.length;i++) {
            signersMap[_signerList[i]]=0;
            _signerList[i]=address(0);
        }
        _signerList.length=0;
        return signer0;
    }

    function genSigners_set_leader(address signer0 ) private {
        address[] memory leaders = base.takeLeaderList();
        for (uint i = 0; i < leaders.length; i++) {
            address l = leaders[i];
            if (signer0 == l) {
                if (i == leaders.length - 1) {
                    pushSigner(leaders[0]);
                } else {
                    pushSigner(leaders[i + 1]);
                }
               return;
            }
        }
        revert("signer0 must exist in leader list");
    }

    // push volunteerList to signerList
    function genSigners_v2s() private {
        for (uint i = 0; i < _nextRoundSignerList.length; i++) {
            address v = _nextRoundSignerList[i];
            pushSigner(v);
        }
        for (uint i = _signerList.length; i < base.takeSignerLimit(); i++) {
            // placeholder
            _signerList.push(address(0));
        }
    }

    // clean volunteerList
    function genSigners_clean_next_round_signers() private {
        for (uint i =0; i< _nextRoundSignerList.length; i++) {
            _nextRoundSignerList[i]=address(0);
        }
        _nextRoundSignerList.length=0;
    }

    function genSigners_clean_blackList() private {
        base.pocCleanBlackList();
    }

    function genSigners() private  {
        address signer0=clean_all_signer_and_get0();
        require(signer0!=address(0),"signer0 must not be zero");
        // generate
        genSigners_set_leader(signer0);
        // move next round candidate signers to current signer list
        genSigners_v2s();
        // clear next round candidate signers.
        genSigners_clean_next_round_signers();
    }
    //合约外部调用唯一的入口
    function update(address volunteer) public allow() {

        blockNumber = block.number;

        uint l = base.takeSignerLimit();
        uint signerIdx = uint(blockNumber % l); //轮到signerIdx出块
        address si = _signerList[signerIdx];

        //1 : not leader signer, and `sender!=si` move `si` to waitList in POC contract
        if (signerIdx > uint(0)) { //leader不是选出来的

            if (uint160(volunteer) != uint160(0)) {
                pushNextRoundSigner(volunteer);
            }
            //轮到si出块,si没有,那么会惩罚si
            if (si != address(0) && msg.sender != si) {
                // move si to stopList in POC contract
                if (base.pocAddStop(si) > 0) {
                    // move to blackList
                    base.pocAddBlackList(si);
                }
                delete signersMap[si];
                // mark 0  
                _signerList[signerIdx] = address(0);
            }
        }

        //2 : last block, reset signerList
        if (signerIdx == (l - 1)) {
            genSigners(); //一轮结束,选择下一轮出块人
        }
        //拉黑持续时间为一个epoch
        if (block.number%getEpoch()==0){
            // clean blackList
            genSigners_clean_blackList();
        }
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
            scoreList[i] = 0;
            numberList[i] = signersMap[_signerList[i]];
        }
        // TODO
        blackList = base.pocGetBlackList();

        signerList = _signerList;
        // vsn 0.0.3
        number = blockNumber;
        totalVolunteer = _nextRoundSignerList.length;
    }

    function version() public view returns (string memory) {return vsn;}

    function getSignerLimit() public view returns (uint) {return base.takeSignerLimit();}

    function getEpoch() public view returns (uint) {return base.takeEpoch();}

    function getVolunteerLimit() public view returns (uint) {return base.takeVolunteerLimit();}
    //为了保持兼容性,所以名字保持不变
    function getVolunteers() public view returns (
        address[] memory volunteerList,
        uint[] memory weightList,
        uint length
    ){
        (volunteerList,length)=getNextRoundSignerList();
        weightList = new uint[](length);
    }
    function getNextRoundSignerList() public view returns (
    address[] memory nextRoundSignerList,
    uint length
    ) {
        nextRoundSignerList=_nextRoundSignerList;
        length=nextRoundSignerList.length;
    }

    // TODO
    function filterVolunteer(address[] memory volunteers) public view returns (uint[] memory result) {}
//    function test_update(address nextsigner) public{
//        update(nextsigner);
//    }
//    function test_push_signer(uint256 i,address signer) public{
//            _signerList[i]=signer;
//            signersMap[signer] = 1;
//    }
}

