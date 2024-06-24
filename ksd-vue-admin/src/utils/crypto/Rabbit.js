// 引用 crypto-js 加密模块
import CryptoJS from 'crypto-js';
export function rabbitEncrypt(text) {
    return CryptoJS.Rabbit.encrypt(text, "1234567ASDFG").toString();
}

export function rabbitDecrypt(encryptedData) {
    return CryptoJS.Rabbit.decrypt(encryptedData, "1234567ASDFG").toString(CryptoJS.enc.Utf8);
}
