'use strict';
/*
* Copyright IBM Corp All Rights Reserved
*
* SPDX-License-Identifier: Apache-2.0
*/
/*
 * Chaincode Invoke
 */

var hfc = require('fabric-client');
var FabricCAServices = require('fabric-ca-client');
var FabricCAClient = FabricCAServices.FabricCAClient;
var X509 = require('x509');
var	tlsOptions = {
	trustedRoots: [],
	verify: false
};

	var options = {
    wallet_path: path.join(__dirname, './credtest'),
    user_id: 'test',
    channel_id: 'mychannel',
    chaincode_id: 'fabcar',
    network_url: 'grpc://localhost:7051',
};


var path = require('path');
var LocalMSP = require('fabric-ca-client/lib/msp/msp.js');
var idModule = require('fabric-ca-client/lib/msp/identity.js');
var SigningIdentity = idModule.SigningIdentity;
var Signer = idModule.Signer;
var User = require('fabric-client/lib/User.js');
//FabricCAServices.getConfigSetting('crypto-keysize', '256');
//FabricCAServices.setConfigSetting('crypto-hash-algo', 'SHA2');
var client;
var eresult;
var pubkey2;
var keyvalpath;
var member;
var cryptoSuite;
var client2;
var channel = {};

Promise.resolve().then(() => {
keyvalpath = path.join(__dirname,'./credtest');
client2 =new hfc();
return hfc.newDefaultKeyValueStore({ path: keyvalpath });
        }).catch((err) => {
console.log(err);
}).then((wallet) =>{
member = new User('test');
cryptoSuite = hfc.newCryptoSuite();
cryptoSuite.setCryptoKeyStore(hfc.newCryptoKeyStore({path: keyvalpath}));
member.setCryptoSuite(cryptoSuite);
client2.setStateStore(wallet);

return new FabricCAServices('http://localhost:7054',tlsOptions,null,cryptoSuite);

}).catch((err)=>{
console.log(err);
}).then((caservice) =>{

return caservice.enroll({
enrollmentID:'admin',
enrollmentSecret:'adminpw'
});

}).catch((err)=>{
console.log(err);
}).then((admin2) => {
return member.setEnrollment(admin2.key, admin2.certificate,'Org1MSP');
}).catch((err)=>{

console.log(err);
}).then(()=>{
if(member.isEnrolled()){
console.log('successfully user enroll');
}else{
console.log('failed user enroll');
}
return client2.setUserContext(member);

}).then((user)=>{

client2.setCryptoSuite(cryptoSuite);
return client2.getUserContext('test',true);
}).then((user)=>{
console.log(user.getName());
    console.log("Check user is enrolled, and set a query URL in the network");
    if (user === undefined || user.isEnrolled() === false) {
        console.error("User not defined, or not enrolled - error");
    }
    channel = client2.newChannel(options.channel_id);
    channel.addPeer(client2.newPeer(options.network_url));
    return;
}).then(() => {
    console.log("Make query");
    var transaction_id = client2.newTransactionID();
    console.log("Assigning transaction_id: ", transaction_id._transaction_id);

    // queryCar - requires 1 argument, ex: args: ['CAR4'],
    // queryAllCars - requires no arguments , ex: args: [''],
    const request = {
        chaincodeId: options.chaincode_id,
        txId: transaction_id,
        fcn: 'queryDoc',
        args: ['5c4ac3271c6dd860ddf7ae310a7990191cefbb33']
    };
    return channel.queryByChaincode(request);
}).then((query_responses) => {
    console.log("returned from query");
    if (!query_responses.length) {
        console.log("No payloads were returned from query");
    } else {
        console.log("Query result count = ", query_responses.length)
    }
    if (query_responses[0] instanceof Error) {
        console.error("error from query = ", query_responses[0]);
    }
    console.log("Response is ", query_responses[0].toString());
}).catch((err) => {
    console.error("Caught Error", err);
});