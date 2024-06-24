// pnpm install node-rsa
import NodeRSA from 'node-rsa'
import crypto from 'crypto'

// 生成一个1024长度的密钥对
const key = new nodeRSA({b: 1024});
// 导出公钥
const publicKey = key.exportKey('public');
// 导出私钥
const privateKey = key.exportKey('private');

// console.log("公钥:\n", publicKey)
// console.log("私钥:\n", privatekey)
// console.log("加密字符串: ", encryptedData)
// console.log("解密字符串: ", decryptedData)

export function rsaEncrypt(secret) {
   return crypto.privateEncrypt(privateKey, Buffer.from(secret))
}

export function rsaDecrypt(encrypt) {
    return crypto.publicDecrypt(publicKey, encrypt);
}
