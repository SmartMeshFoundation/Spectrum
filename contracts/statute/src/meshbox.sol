pragma solidity >=0.5.0 <=0.5.3;

import "./owned.sol";

contract MeshBox is Owned {

    mapping(address => uint256) meshboxAddress;

    function addAddress(address[] memory _owners, uint256 version) public onlyOwner() {
        uint len = _owners.length;
        for (uint i = 0; i < len; i ++) {
            meshboxAddress[_owners[i]] = version;
        }
    }

    function delAddress(address[] memory _owners) public onlyOwner() {
        uint len = _owners.length;
        for (uint i = 0; i < len; i ++) {
            meshboxAddress[_owners[i]] = 0;
        }
    }

    function existAddress(address _owner) view public returns (uint256){
        return meshboxAddress[_owner];
    }
}