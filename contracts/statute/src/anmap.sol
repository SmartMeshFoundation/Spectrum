pragma solidity ^0.5.3;


// Address and nodeid mapping
// used for reward txfee
contract Anmap {

    mapping(address => address) anmap;
    mapping(address => address) namap;

    function bindNode(bytes32 hash, uint8 v, bytes32 r, bytes32 s) public pure returns(address) {
        return ecrecover(hash,v,r,s);
    }

    function unbindNode() public {

    }

    function fetch(address node) view public returns (address){
        return namap[node];
    }

}