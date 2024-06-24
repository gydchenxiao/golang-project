// 登录des对称加密
import CryptoJS from 'crypto-js'

var desKey = "0123456789ABCDEF"    // 密钥
var desIv = "0123456789ABCDEF"    // 偏移量
export function tripleDesEncrypt(text) {
    var key = CryptoJS.enc.Utf8.parse(desKey),
        iv = CryptoJS.enc.Utf8.parse(desIv),
        srcs = CryptoJS.enc.Utf8.parse(text),
        // ECB 加密方式，Iso10126 填充方式
        encrypted = CryptoJS.TripleDES.encrypt(srcs, key, {
            iv: iv,
            mode: CryptoJS.mode.ECB,
            padding: CryptoJS.pad.Iso10126
        });
    return encrypted.toString();
}

export function tripleDesDecrypt(encryptedData) {
    var key = CryptoJS.enc.Utf8.parse(desKey),
        iv = CryptoJS.enc.Utf8.parse(desIv),
        srcs = encryptedData,
        // ECB 加密方式，Iso10126 填充方式
        decrypted = CryptoJS.TripleDES.decrypt(srcs, key, {
            iv: iv,
            mode: CryptoJS.mode.ECB,
            padding: CryptoJS.pad.Iso10126
        });
    return decrypted.toString(CryptoJS.enc.Utf8);
}
