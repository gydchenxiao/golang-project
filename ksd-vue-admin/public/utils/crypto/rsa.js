// npm install node-rsa
import NodeRSA from 'node-rsa'

var key = new NodeRSA({b: 512});                    //生成512位秘钥
var publicKey = key.exportKey('pkcs8-public');    //导出公钥
var privatekey = key.exportKey('pkcs8-private');  //导出私钥

// console.log("公钥:\n", publicKey)
// console.log("私钥:\n", privatekey)
// console.log("加密字符串: ", encryptedData)
// console.log("解密字符串: ", decryptedData)

export function rsaEncrypt(text) {
    const pubKey = new NodeRSA(publicKey,'pkcs8-public');
    var encryptedData = pubKey.encrypt(text, 'base64');
    return encryptedData
}

export function rsaDecrypt(text) {
    const priKey = new NodeRSA(privatekey,'pkcs8-private');
    var decryptedData = priKey.decrypt(encryptedData, 'utf8');
    return decryptedData
}
