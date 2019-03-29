pragma solidity ^0.5.3;

import "./owned.sol";


// Address and nodeid mapping
// used for reward txfee
contract Anmap is Owned {

    uint smt = 1000000000000000000;
    uint defBalance = 50; // default is 50 only for test;
    uint minBalance = defBalance * smt;

    mapping(address => address) anmap;
    mapping(address => address) namap;

    constructor(uint _minBalance) public {
        resetMinBalance(_minBalance);
    }

    function resetMinBalance(uint _minBalance) public onlyOwner() {
        if (_minBalance > defBalance) {
            minBalance = _minBalance * smt;
        }
    }

    function getMinBalance() public view returns (uint) {
        return minBalance / smt;
    }

    //"0x70aEfe8d97EF5984B91b5169418f3db283F65a29",28,"0xa006360ed68f21443221354903f49dce733afaaeac3b35d420bb2154746c9592","0x6f31ccfa10b449531fd0fff78db52cc02edaabd2c5e9a6abb8fc1cd6df26f442"
    function bind(address nodePub, uint8 v, bytes32 r, bytes32 s) public {

        bytes32 hash = keccak256(abi.encodePacked(msg.sender));
        address signer = ecrecover(hash, v, r, s);

        require(msg.sender.balance >= minBalance);
        require(nodePub == signer);
        require(anmap[msg.sender] == address(0));
        require(namap[nodePub] == address(0));

        anmap[msg.sender] = nodePub;
        namap[nodePub] = msg.sender;
    }

    function unbind() public {
        address np = anmap[msg.sender];
        require(np != address(0));
        delete anmap[msg.sender];
        delete namap[np];
    }

    function unbindBySig(address nodePub, uint8 v, bytes32 r, bytes32 s) public {
        bytes32 hash = keccak256(abi.encodePacked(msg.sender));
        address signer = ecrecover(hash, v, r, s);
        require(nodePub == signer);

        address addr = namap[nodePub];

        require(addr != address(0));
        delete anmap[addr];
        delete namap[nodePub];
    }

    function fetchInNamap(address node) view public returns (address a, uint b, uint m) {
        a = namap[node];
        m = getMinBalance();
        if (a != address(0)) {
            b = a.balance / smt;
        }
    }

    function fetchInAnmap(address addr) view public returns (address n, uint b, uint m) {
        n = anmap[addr];
        m = getMinBalance();
        if (n != address(0)) {
            b = n.balance / smt;
        }
    }

}
