// 登录des对称加密
import CryptoJs from 'crypto-js'
var desKey = "kuangxx."    // 密钥
var desIv = "kuangxx."    // 初始向量
function desEncrypt(text) {
    var key = CryptoJs.enc.Utf8.parse(desKey),
        iv = CryptoJs.enc.Utf8.parse(desIv),
        srcs = CryptoJs.enc.Utf8.parse(text),
        // CBC 加密模式，Pkcs7 填充方式
        encrypted = CryptoJs.DES.encrypt(srcs, key, {
            iv: iv,
            mode: CryptoJs.mode.CBC,
            padding: CryptoJs.pad.Pkcs7
        });
    return encrypted.toString();
}

function desDecrypt(srcs) {
    var key = CryptoJs.enc.Utf8.parse(desKey),
        iv = CryptoJs.enc.Utf8.parse(desIv),
        // CBC 加密模式，Pkcs7 填充方式
        decrypted = CryptoJs.DES.decrypt(srcs, key, {
            iv: iv,
            mode: CryptoJs.mode.CBC,
            padding: CryptoJs.pad.Pkcs7
        });
    return decrypted.toString(CryptoJs.enc.Utf8);
}

export default  {
    desEncrypt,
    desDecrypt
}
