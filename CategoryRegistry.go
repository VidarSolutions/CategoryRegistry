// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CategoryRegistry

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"parent\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"subCategory\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"SubCatContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parent\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"categoryName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"categoryAddress\",\"type\":\"address\"}],\"name\":\"addCategory\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parent\",\"type\":\"string\"}],\"name\":\"getAllCategories\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"parent\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"categoryName\",\"type\":\"string\"}],\"name\":\"getCategoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRegistryFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_registryFee\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"registryFeeTitle\",\"type\":\"string\"}],\"name\":\"setNonStandardRegistryFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_registryFee\",\"type\":\"string\"}],\"name\":\"setRegistryFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6101406040526000608090815250670de0b6b3a764000060a090815250674563918244f4000060c090815250678ac7230489e8000060e0908152506802b5e3af16b18800006101009081525068056bc75e2d63100000610120908152503480156200006957600080fd5b50604051620021d3380380620021d383398181016040528101906200008f91906200050b565b83600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060805160028190555030600084604051620000ec919062000608565b908152602001604051809103902060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000846040516200014e919062000608565b9081526020016040518091039020600101836040516200016f919062000608565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600083604051620001cd919062000608565b90815260200160405180910390206002018290806001815401808255809150506001900390600052602060002001600090919091909150908162000212919062000861565b5060008360405162000225919062000608565b90815260200160405180910390206003016000815480929190620002499062000977565b919050555060a0516001604051620002619062000a14565b90815260200160405180910390208190555060c0516001604051620002869062000a7b565b90815260200160405180910390208190555060e0516001604051620002ab9062000ae2565b908152602001604051809103902081905550610100516001604051620002d19062000b49565b908152602001604051809103902081905550610120516001604051620002f79062000bb0565b9081526020016040518091039020819055505050505062000bc7565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003548262000327565b9050919050565b620003668162000347565b81146200037257600080fd5b50565b60008151905062000386816200035b565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620003e18262000396565b810181811067ffffffffffffffff82111715620004035762000402620003a7565b5b80604052505050565b60006200041862000313565b9050620004268282620003d6565b919050565b600067ffffffffffffffff821115620004495762000448620003a7565b5b620004548262000396565b9050602081019050919050565b60005b838110156200048157808201518184015260208101905062000464565b60008484015250505050565b6000620004a46200049e846200042b565b6200040c565b905082815260208101848484011115620004c357620004c262000391565b5b620004d084828562000461565b509392505050565b600082601f830112620004f057620004ef6200038c565b5b8151620005028482602086016200048d565b91505092915050565b600080600080608085870312156200052857620005276200031d565b5b6000620005388782880162000375565b945050602085015167ffffffffffffffff8111156200055c576200055b62000322565b5b6200056a87828801620004d8565b935050604085015167ffffffffffffffff8111156200058e576200058d62000322565b5b6200059c87828801620004d8565b9250506060620005af8782880162000375565b91505092959194509250565b600081519050919050565b600081905092915050565b6000620005de82620005bb565b620005ea8185620005c6565b9350620005fc81856020860162000461565b80840191505092915050565b6000620006168284620005d1565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200066957607f821691505b6020821081036200067f576200067e62000621565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620006e97fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620006aa565b620006f58683620006aa565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620007426200073c62000736846200070d565b62000717565b6200070d565b9050919050565b6000819050919050565b6200075e8362000721565b620007766200076d8262000749565b848454620006b7565b825550505050565b600090565b6200078d6200077e565b6200079a81848462000753565b505050565b5b81811015620007c257620007b660008262000783565b600181019050620007a0565b5050565b601f8211156200081157620007db8162000685565b620007e6846200069a565b81016020851015620007f6578190505b6200080e62000805856200069a565b8301826200079f565b50505b505050565b600082821c905092915050565b6000620008366000198460080262000816565b1980831691505092915050565b600062000851838362000823565b9150826002028217905092915050565b6200086c82620005bb565b67ffffffffffffffff811115620008885762000887620003a7565b5b62000894825462000650565b620008a1828285620007c6565b600060209050601f831160018114620008d95760008415620008c4578287015190505b620008d0858262000843565b86555062000940565b601f198416620008e98662000685565b60005b828110156200091357848901518255600182019150602085019450602081019050620008ec565b868310156200093357848901516200092f601f89168262000823565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600062000984826200070d565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620009b957620009b862000948565b5b600182019050919050565b7f6f6e65446f6c6c61720000000000000000000000000000000000000000000000600082015250565b6000620009fc600983620005c6565b915062000a0982620009c4565b600982019050919050565b600062000a2182620009ed565b9150819050919050565b7f66697665446f6c6c617273000000000000000000000000000000000000000000600082015250565b600062000a63600b83620005c6565b915062000a708262000a2b565b600b82019050919050565b600062000a888262000a54565b9150819050919050565b7f74656e446f6c6c61727300000000000000000000000000000000000000000000600082015250565b600062000aca600a83620005c6565b915062000ad78262000a92565b600a82019050919050565b600062000aef8262000abb565b9150819050919050565b7f6669667479446f6c6c6172730000000000000000000000000000000000000000600082015250565b600062000b31600c83620005c6565b915062000b3e8262000af9565b600c82019050919050565b600062000b568262000b22565b9150819050919050565b7f6f6e6548756e64726564446f6c6c617273000000000000000000000000000000600082015250565b600062000b98601183620005c6565b915062000ba58262000b60565b601182019050919050565b600062000bbd8262000b89565b9150819050919050565b60805160a05160c05160e05161010051610120516115ce62000c056000396000505060005050600050506000505060005050600050506115ce6000f3fe6080604052600436106100555760003560e01c806315ca98d51461005a5780636887f514146100855780638dba0f17146100ae578063929c5e6c146100ca578063aa629c8314610107578063c335676514610130575b600080fd5b34801561006657600080fd5b5061006f61016d565b60405161007c9190610891565b60405180910390f35b34801561009157600080fd5b506100ac60048036038101906100a79190610a32565b610177565b005b6100c860048036038101906100c39190610aec565b610296565b005b3480156100d657600080fd5b506100f160048036038101906100ec9190610b77565b6105ce565b6040516100fe9190610bfe565b60405180910390f35b34801561011357600080fd5b5061012e60048036038101906101299190610c19565b610635565b005b34801561013c57600080fd5b5061015760048036038101906101529190610c19565b61074f565b6040516101649190610da3565b60405180910390f35b6000600254905090565b3373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610207576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fe90610e48565b60405180910390fd5b6000821161024a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161024190610eb4565b60405180910390fd5b8160018260405161025b9190610f10565b90815260200160405180910390208190555060018160405161027d9190610f10565b9081526020016040518091039020546002819055505050565b600a6000846040516102a89190610f10565b908152602001604051809103902060030154106102fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f190610fbf565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166000846040516103229190610f10565b9081526020016040518091039020600101836040516103419190610f10565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146103c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103bd9061102b565b60405180910390fd5b6103d1600254610849565b341015610413576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161040a906110e3565b60405180910390fd5b60003073ffffffffffffffffffffffffffffffffffffffff163460405161043990611134565b60006040518083038185875af1925050503d8060008114610476576040519150601f19603f3d011682016040523d82523d6000602084013e61047b565b606091505b50509050806104bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b690611195565b60405180910390fd5b816000856040516104d09190610f10565b9081526020016040518091039020600101846040516104ef9190610f10565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008460405161054b9190610f10565b90815260200160405180910390206002018390806001815401808255809150506001900390600052602060002001600090919091909150908161058e91906113c1565b5060016000856040516105a19190610f10565b908152602001604051809103902060030160008282546105c191906114c2565b9250508190555050505050565b600080836040516105df9190610f10565b9081526020016040518091039020600101826040516105fe9190610f10565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905092915050565b3373ffffffffffffffffffffffffffffffffffffffff16600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146106c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bc90610e48565b60405180910390fd5b60006001826040516106d79190610f10565b9081526020016040518091039020541015610727576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161071e90610eb4565b60405180910390fd5b6001816040516107379190610f10565b90815260200160405180910390205460028190555050565b60606000826040516107619190610f10565b9081526020016040518091039020600201805480602002602001604051908101604052809291908181526020016000905b8282101561083e5783829060005260206000200180546107b1906111e4565b80601f01602080910402602001604051908101604052809291908181526020018280546107dd906111e4565b801561082a5780601f106107ff5761010080835404028352916020019161082a565b820191906000526020600020905b81548152906001019060200180831161080d57829003601f168201915b505050505081526020019060010190610792565b505050509050919050565b600080600454670de0b6b3a76400008461086391906114f6565b61086d9190611567565b905080915050919050565b6000819050919050565b61088b81610878565b82525050565b60006020820190506108a66000830184610882565b92915050565b6000604051905090565b600080fd5b600080fd5b6108c981610878565b81146108d457600080fd5b50565b6000813590506108e6816108c0565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61093f826108f6565b810181811067ffffffffffffffff8211171561095e5761095d610907565b5b80604052505050565b60006109716108ac565b905061097d8282610936565b919050565b600067ffffffffffffffff82111561099d5761099c610907565b5b6109a6826108f6565b9050602081019050919050565b82818337600083830152505050565b60006109d56109d084610982565b610967565b9050828152602081018484840111156109f1576109f06108f1565b5b6109fc8482856109b3565b509392505050565b600082601f830112610a1957610a186108ec565b5b8135610a298482602086016109c2565b91505092915050565b60008060408385031215610a4957610a486108b6565b5b6000610a57858286016108d7565b925050602083013567ffffffffffffffff811115610a7857610a776108bb565b5b610a8485828601610a04565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610ab982610a8e565b9050919050565b610ac981610aae565b8114610ad457600080fd5b50565b600081359050610ae681610ac0565b92915050565b600080600060608486031215610b0557610b046108b6565b5b600084013567ffffffffffffffff811115610b2357610b226108bb565b5b610b2f86828701610a04565b935050602084013567ffffffffffffffff811115610b5057610b4f6108bb565b5b610b5c86828701610a04565b9250506040610b6d86828701610ad7565b9150509250925092565b60008060408385031215610b8e57610b8d6108b6565b5b600083013567ffffffffffffffff811115610bac57610bab6108bb565b5b610bb885828601610a04565b925050602083013567ffffffffffffffff811115610bd957610bd86108bb565b5b610be585828601610a04565b9150509250929050565b610bf881610aae565b82525050565b6000602082019050610c136000830184610bef565b92915050565b600060208284031215610c2f57610c2e6108b6565b5b600082013567ffffffffffffffff811115610c4d57610c4c6108bb565b5b610c5984828501610a04565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610cc8578082015181840152602081019050610cad565b60008484015250505050565b6000610cdf82610c8e565b610ce98185610c99565b9350610cf9818560208601610caa565b610d02816108f6565b840191505092915050565b6000610d198383610cd4565b905092915050565b6000602082019050919050565b6000610d3982610c62565b610d438185610c6d565b935083602082028501610d5585610c7e565b8060005b85811015610d915784840389528151610d728582610d0d565b9450610d7d83610d21565b925060208a01995050600181019050610d59565b50829750879550505050505092915050565b60006020820190508181036000830152610dbd8184610d2e565b905092915050565b600082825260208201905092915050565b7f4f6e6c7920746865206f776e6572206f66207468697320636f6e74726163742060008201527f6d61792063616c6c20746869732066756e6369746f6e00000000000000000000602082015250565b6000610e32603683610dc5565b9150610e3d82610dd6565b604082019050919050565b60006020820190508181036000830152610e6181610e25565b9050919050565b7f496e76616c696420726567697374727920666565000000000000000000000000600082015250565b6000610e9e601483610dc5565b9150610ea982610e68565b602082019050919050565b60006020820190508181036000830152610ecd81610e91565b9050919050565b600081905092915050565b6000610eea82610c8e565b610ef48185610ed4565b9350610f04818560208601610caa565b80840191505092915050565b6000610f1c8284610edf565b915081905092915050565b7f506172656e742063617465676f72792069732066756c6c2c2063686f6f73652060008201527f616e6f746865722063617465676f727920746f2063726561746520796f75722060208201527f73756263617465676f727920756e646572000000000000000000000000000000604082015250565b6000610fa9605183610dc5565b9150610fb482610f27565b606082019050919050565b60006020820190508181036000830152610fd881610f9c565b9050919050565b7f43617465676f727920616c726561647920657869737473000000000000000000600082015250565b6000611015601783610dc5565b915061102082610fdf565b602082019050919050565b6000602082019050818103600083015261104481611008565b9050919050565b7f4e6f7420656e6f75676820415641582073656e7420746f20637265617465207360008201527f686f70206765742063757272656e74206665652062792063616c6c696e67206760208201527f657453686f70466565206f6e207468697320636f6e74726163742e0000000000604082015250565b60006110cd605b83610dc5565b91506110d88261104b565b606082019050919050565b600060208201905081810360008301526110fc816110c0565b9050919050565b600081905092915050565b50565b600061111e600083611103565b91506111298261110e565b600082019050919050565b600061113f82611111565b9150819050919050565b7f4661696c656420746f20636f6c6c65637420415641582e000000000000000000600082015250565b600061117f601783610dc5565b915061118a82611149565b602082019050919050565b600060208201905081810360008301526111ae81611172565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806111fc57607f821691505b60208210810361120f5761120e6111b5565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026112777fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261123a565b611281868361123a565b95508019841693508086168417925050509392505050565b6000819050919050565b60006112be6112b96112b484610878565b611299565b610878565b9050919050565b6000819050919050565b6112d8836112a3565b6112ec6112e4826112c5565b848454611247565b825550505050565b600090565b6113016112f4565b61130c8184846112cf565b505050565b5b81811015611330576113256000826112f9565b600181019050611312565b5050565b601f8211156113755761134681611215565b61134f8461122a565b8101602085101561135e578190505b61137261136a8561122a565b830182611311565b50505b505050565b600082821c905092915050565b60006113986000198460080261137a565b1980831691505092915050565b60006113b18383611387565b9150826002028217905092915050565b6113ca82610c8e565b67ffffffffffffffff8111156113e3576113e2610907565b5b6113ed82546111e4565b6113f8828285611334565b600060209050601f83116001811461142b5760008415611419578287015190505b61142385826113a5565b86555061148b565b601f19841661143986611215565b60005b828110156114615784890151825560018201915060208501945060208101905061143c565b8683101561147e578489015161147a601f891682611387565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006114cd82610878565b91506114d883610878565b92508282019050808211156114f0576114ef611493565b5b92915050565b600061150182610878565b915061150c83610878565b925082820261151a81610878565b9150828204841483151761153157611530611493565b5b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061157282610878565b915061157d83610878565b92508261158d5761158c611538565b5b82820490509291505056fea2646970667358221220d955f13fb6b5944448709bb5fe989124c556f62288f44d27ac64048f32bd9db464736f6c63430008120033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, parent string, subCategory string, SubCatContract common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _owner, parent, subCategory, SubCatContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// GetAllCategories is a free data retrieval call binding the contract method 0xc3356765.
//
// Solidity: function getAllCategories(string parent) view returns(string[])
func (_Api *ApiCaller) GetAllCategories(opts *bind.CallOpts, parent string) ([]string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAllCategories", parent)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetAllCategories is a free data retrieval call binding the contract method 0xc3356765.
//
// Solidity: function getAllCategories(string parent) view returns(string[])
func (_Api *ApiSession) GetAllCategories(parent string) ([]string, error) {
	return _Api.Contract.GetAllCategories(&_Api.CallOpts, parent)
}

// GetAllCategories is a free data retrieval call binding the contract method 0xc3356765.
//
// Solidity: function getAllCategories(string parent) view returns(string[])
func (_Api *ApiCallerSession) GetAllCategories(parent string) ([]string, error) {
	return _Api.Contract.GetAllCategories(&_Api.CallOpts, parent)
}

// GetCategoryAddress is a free data retrieval call binding the contract method 0x929c5e6c.
//
// Solidity: function getCategoryAddress(string parent, string categoryName) view returns(address)
func (_Api *ApiCaller) GetCategoryAddress(opts *bind.CallOpts, parent string, categoryName string) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getCategoryAddress", parent, categoryName)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCategoryAddress is a free data retrieval call binding the contract method 0x929c5e6c.
//
// Solidity: function getCategoryAddress(string parent, string categoryName) view returns(address)
func (_Api *ApiSession) GetCategoryAddress(parent string, categoryName string) (common.Address, error) {
	return _Api.Contract.GetCategoryAddress(&_Api.CallOpts, parent, categoryName)
}

// GetCategoryAddress is a free data retrieval call binding the contract method 0x929c5e6c.
//
// Solidity: function getCategoryAddress(string parent, string categoryName) view returns(address)
func (_Api *ApiCallerSession) GetCategoryAddress(parent string, categoryName string) (common.Address, error) {
	return _Api.Contract.GetCategoryAddress(&_Api.CallOpts, parent, categoryName)
}

// GetRegistryFee is a free data retrieval call binding the contract method 0x15ca98d5.
//
// Solidity: function getRegistryFee() view returns(uint256)
func (_Api *ApiCaller) GetRegistryFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getRegistryFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRegistryFee is a free data retrieval call binding the contract method 0x15ca98d5.
//
// Solidity: function getRegistryFee() view returns(uint256)
func (_Api *ApiSession) GetRegistryFee() (*big.Int, error) {
	return _Api.Contract.GetRegistryFee(&_Api.CallOpts)
}

// GetRegistryFee is a free data retrieval call binding the contract method 0x15ca98d5.
//
// Solidity: function getRegistryFee() view returns(uint256)
func (_Api *ApiCallerSession) GetRegistryFee() (*big.Int, error) {
	return _Api.Contract.GetRegistryFee(&_Api.CallOpts)
}

// AddCategory is a paid mutator transaction binding the contract method 0x8dba0f17.
//
// Solidity: function addCategory(string parent, string categoryName, address categoryAddress) payable returns()
func (_Api *ApiTransactor) AddCategory(opts *bind.TransactOpts, parent string, categoryName string, categoryAddress common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addCategory", parent, categoryName, categoryAddress)
}

// AddCategory is a paid mutator transaction binding the contract method 0x8dba0f17.
//
// Solidity: function addCategory(string parent, string categoryName, address categoryAddress) payable returns()
func (_Api *ApiSession) AddCategory(parent string, categoryName string, categoryAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.AddCategory(&_Api.TransactOpts, parent, categoryName, categoryAddress)
}

// AddCategory is a paid mutator transaction binding the contract method 0x8dba0f17.
//
// Solidity: function addCategory(string parent, string categoryName, address categoryAddress) payable returns()
func (_Api *ApiTransactorSession) AddCategory(parent string, categoryName string, categoryAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.AddCategory(&_Api.TransactOpts, parent, categoryName, categoryAddress)
}

// SetNonStandardRegistryFee is a paid mutator transaction binding the contract method 0x6887f514.
//
// Solidity: function setNonStandardRegistryFee(uint256 _registryFee, string registryFeeTitle) returns()
func (_Api *ApiTransactor) SetNonStandardRegistryFee(opts *bind.TransactOpts, _registryFee *big.Int, registryFeeTitle string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setNonStandardRegistryFee", _registryFee, registryFeeTitle)
}

// SetNonStandardRegistryFee is a paid mutator transaction binding the contract method 0x6887f514.
//
// Solidity: function setNonStandardRegistryFee(uint256 _registryFee, string registryFeeTitle) returns()
func (_Api *ApiSession) SetNonStandardRegistryFee(_registryFee *big.Int, registryFeeTitle string) (*types.Transaction, error) {
	return _Api.Contract.SetNonStandardRegistryFee(&_Api.TransactOpts, _registryFee, registryFeeTitle)
}

// SetNonStandardRegistryFee is a paid mutator transaction binding the contract method 0x6887f514.
//
// Solidity: function setNonStandardRegistryFee(uint256 _registryFee, string registryFeeTitle) returns()
func (_Api *ApiTransactorSession) SetNonStandardRegistryFee(_registryFee *big.Int, registryFeeTitle string) (*types.Transaction, error) {
	return _Api.Contract.SetNonStandardRegistryFee(&_Api.TransactOpts, _registryFee, registryFeeTitle)
}

// SetRegistryFee is a paid mutator transaction binding the contract method 0xaa629c83.
//
// Solidity: function setRegistryFee(string _registryFee) returns()
func (_Api *ApiTransactor) SetRegistryFee(opts *bind.TransactOpts, _registryFee string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setRegistryFee", _registryFee)
}

// SetRegistryFee is a paid mutator transaction binding the contract method 0xaa629c83.
//
// Solidity: function setRegistryFee(string _registryFee) returns()
func (_Api *ApiSession) SetRegistryFee(_registryFee string) (*types.Transaction, error) {
	return _Api.Contract.SetRegistryFee(&_Api.TransactOpts, _registryFee)
}

// SetRegistryFee is a paid mutator transaction binding the contract method 0xaa629c83.
//
// Solidity: function setRegistryFee(string _registryFee) returns()
func (_Api *ApiTransactorSession) SetRegistryFee(_registryFee string) (*types.Transaction, error) {
	return _Api.Contract.SetRegistryFee(&_Api.TransactOpts, _registryFee)
}
