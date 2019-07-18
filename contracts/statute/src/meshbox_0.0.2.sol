pragma solidity >=0.5.0 <=0.5.3;

//import "github.com/SmartMeshFoundation/Spectrum/contracts/statute/src/owned.sol"; // for remix

import "./owned.sol";

contract MeshBox_0_0_2 is Owned {

    string vsn = "0.0.2";

    mapping(address => uint256) meshboxAddress;

    address[] meshboxList; //0.0.2

    // ["0xa9a461abfac6a9a058a6ef2198fd2780e7d4f7e2","0xa9a461abfac6a9a058a6ef2198fd2780e7d4f7e3","0xa9a461abfac6a9a058a6ef2198fd2780e7d4f7e4","0xa9a461abfac6a9a058a6ef2198fd2780e7d4f7e2","0xa9a461abfac6a9a058a6ef2198fd2780e7d4f7e5"],1
    function addAddress(address[] memory _owners, uint weight) public onlyOwner() {
        uint len = _owners.length;
        for (uint i = 0; i < len; i ++) {
            address _o = _owners[i];
            if (meshboxAddress[_o] == 0) {
                meshboxList.push(_o);
            }
            meshboxAddress[_o] = weight;
        }
    }

    //0.0.2
    function cleanMeshboxList(address addr) private {
        meshboxAddress[addr] = 0;
        for (uint i=0;i<meshboxList.length;i++) {
            if (meshboxList[i] == addr) {
                for (uint j=i;j<meshboxList.length-1;j++) {
                    meshboxList[j] = meshboxList[j+1];
                }
                i--;
                meshboxList.length -= 1;
            }
        }
    }

    function delAddress(address[] memory _owners) public onlyOwner() {
        uint len = _owners.length;
        for (uint i = 0; i < len; i ++) {
            cleanMeshboxList(_owners[i]);
        }
    }

    function existAddress(address _owner) view public returns (uint256){
        return meshboxAddress[_owner];
    }

    // vsn-0.0.2 >>>>>
    function version() public view returns (string memory) {return vsn;}
    function getMeshboxList() view public returns (address[] memory) {return meshboxList;}
    // vsn-0.0.2 <<<<<

}

