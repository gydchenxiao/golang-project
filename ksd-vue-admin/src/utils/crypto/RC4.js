// 引用 crypto-js 加密模块
import CryptoJS from 'crypto-js';

var key = "12345678ASDFG"
export function RC4Encrypt(text) {
    return CryptoJS.RC4.encrypt(text, key).toString();
}

function RC4Decrypt(encryptedData){
    return CryptoJS.RC4.decrypt(encryptedData, key).toString(CryptoJS.enc.Utf8);
}