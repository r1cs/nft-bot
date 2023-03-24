pragma solidity ^0.8.0;
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ISeaDrop} from "./seadrop/src/interfaces/ISeaDrop.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/MultiCall.sol";

interface Reader {
    function totalSupply() external view returns (uint64);
}

contract BotGuard is Multicall, OwnableUpgradeable {
    function initialize(
        string memory _domain,
        string memory _targetDomain,
        address _faucetToken
    ) public initializer {
        __Ownable_init();
    }

    function _requireTotalSupply(
        uint64 _gotTotalSupply,
        uint64 _expectedTotalSupply,
        uint64 _maxNum
    ) internal pure returns (uint64) {
        require(
            _gotTotalSupply >= _expectedTotalSupply &&
                _gotTotalSupply < _expectedTotalSupply + _maxNum,
            "ERR_SUPPLY"
        );
        return _maxNum - (_gotTotalSupply - _expectedTotalSupply);
    }

    function expectMintTotalSupply(
        address _mintContract,
        bytes memory _params,
        address _nftContract,
        uint64 _expectedTotalSupply,
        uint256 _mintPrice,
        uint64 _maxNum
    ) public payable onlyOwner {
        require(msg.value == _mintPrice * _maxNum);
        uint64 _totalSupply = Reader(_nftContract).totalSupply();
        /// @dev require total supply is right
        uint64 _nftNum = _requireTotalSupply(
            _totalSupply,
            _expectedTotalSupply,
            _maxNum
        );
        /// @dev now calc the actual value need to used and refund
        uint256 _value = _mintPrice * _nftNum;
        if (_nftNum < _maxNum) {
            uint256 _refound = msg.value - _value;
            payable(msg.sender).transfer(_refound);
        }
        /// @dev now try to mint public
        (bool _success, bytes memory _ret) = _mintContract.call{value: _value}(
            _params
        );
        require(_success, string(_ret));
    }

    function call894739472894(
        address _target,
        bytes memory params
    ) public payable onlyOwner {
        Address.functionCallWithValue(_target, params, msg.value);
    }

    function staticCall897984374(
        address _target,
        bytes memory _params
    ) public view returns (bytes memory) {
        return Address.functionStaticCall(_target, _params);
    }
}
