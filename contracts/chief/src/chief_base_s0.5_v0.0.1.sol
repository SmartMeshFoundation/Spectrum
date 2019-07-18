pragma solidity >=0.5.0 <=0.5.3;

/* for remix
import "github.com/SmartMeshFoundation/Spectrum/contracts/chief/src/poc_s0.5.sol"; // for remix
*/



/* local */
import "./poc_s0.5.sol";

contract ChiefBase_1_0_0 {

    string public vsn = "1.0.0";

    POC_1_0_0 private poc;

    address public _owner;
    address public _tribe;

    address[] public leaderList;

    modifier owner() {
        require(msg.sender == _owner);
        _;
    }

    //config >>>>
    uint public leaderLimit = 5;//5
    uint public signerLimit = 3;
    uint public epoch = 6171; // block time 14s, 6171 block = 24H
    uint public volunteerLimit = signerLimit - 1; //留0号位置给Leader
    //config <<<<
    //function setEpoch(uint _epoch) public payable {epoch = _epoch;}
    //function setSignerLimit(uint _signerLimit) public payable {signerLimit = _signerLimit;}

//    function setLeaderLimit(uint _leaderLimit) public  owner() {leaderLimit = _leaderLimit;}

    function setOwner(address newOwner) public  owner() {_owner = newOwner;}

    function takeEpoch() public view returns (uint) {return epoch;}

    function takeSignerLimit() public view returns (uint) {return signerLimit;}

    function takeLeaderLimit() public view returns (uint) {return leaderLimit;}

    function takeVolunteerLimit() public view returns (uint) {return volunteerLimit;}

    function takeLeaderList() public view returns (address[] memory) {return leaderList;}

    function takeOwner() public view returns (address) {return _owner;}

    constructor(uint _epoch, uint _signerLimit) public {
        _owner = msg.sender;
        epoch = _epoch;
        signerLimit = _signerLimit;
        volunteerLimit = signerLimit - 1;
    }

    function appendLeader(address addr) public owner() {
        if (leaderList.length < leaderLimit) {
            for (uint i = 0; i < leaderList.length; i++) {
                require(leaderList[i] != addr);
            }
            leaderList.push(addr);
        }
    }

    function removeLeader(address addr) public owner() {
        require(leaderList.length != 1);
        for (uint i = 0; i < leaderList.length; i++) {
            if (leaderList[i] == addr) {
                for (uint j = i; j < leaderList.length - 1; j++) {
                    leaderList[j] = leaderList[j + 1];
                }
                i--;
                leaderList.length -= 1;
            }
        }
    }

    function isLeader(address addr) public view returns (bool) {
        for (uint i = 0; i < leaderList.length; i++) {
            if (uint160(addr) != uint160(0) && addr == leaderList[i]) {
                return true;
            }
        }
        return false;
    }

    function init(address pocAddr, address tribeAddr) public {
        require(msg.sender == tribeAddr);
        if (address(poc) == address(0) && pocAddr != address(0)) {
            poc = POC_1_0_0(pocAddr);
        }
        if (tribeAddr != address(0) && _tribe==address(0)) {
            _tribe = tribeAddr;
        }
    }

    function pocChangeOwner(address newOwner, uint256 num) owner() public {
        poc.changeOwner(newOwner, num);
    }

    function pocAddStop(address minerAddress) public returns (int8) {
        require(msg.sender == _tribe);
        return poc.ownerStop(minerAddress);
    }

    // chief -> base -> poc : blacklist options >>>>
    function pocAddBlackList(address minerAddress) public {
        require(msg.sender == _tribe);
        poc.ownerPushBlackList(minerAddress);
    }

    function pocCleanBlackList()  public{
        require(msg.sender == _tribe);
        poc.ownerEmptyBlackList();
    }

    function pocGetBlackList() public view returns(address[] memory) {
        return poc.getBlackList();
    }

    function pocAddLockList(address minerAddress) public {
        require(msg.sender == _tribe);
        poc.ownerPushLockList(minerAddress);
    }

    function pocDelLockList(address minerAddress) public returns (int8) {
        require(msg.sender == _tribe);
        return poc.ownerPopLockList(minerAddress);
    }
    // chief -> base -> poc : locklist options <<<<

}
