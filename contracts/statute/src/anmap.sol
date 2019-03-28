pragma solidity ^0.5.3;

import "../../owned.sol";

// Address and nodeid mapping
// used for reward txfee
contract Anmap is Owned {

    mapping(address => address) anmap;
    mapping(address => address) namap;

    //TODO
    function manager() public onlyOwner() {

    }

    function bind(uint256 nodeSignature) public onlyOwner() {

    }

    function unbind() public {

    }

    function fetch(address node) view public returns (address){
        return namap[node];
    }

}