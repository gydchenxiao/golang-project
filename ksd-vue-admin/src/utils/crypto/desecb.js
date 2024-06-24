// 登录des对称加密
import CryptoJs from 'crypto-js'

const desEncrypt = function(key="maya0810maya^%.m",text) {
    var l = CryptoJS.enc.Utf8.parse(text);
    var e = CryptoJS.enc.Utf8.parse(key);
    var a = CryptoJS.DES.encrypt(l, e, {
        mode: CryptoJS.mode.ECB,
        padding: CryptoJS.pad.Pkcs7
    })
    return a.toString()  // 此方式返回base64
}

const desDecrypt= function(key="maya0810maya^%.m", text) {
    var e = CryptoJS.enc.Utf8.parse(key);
    var a = CryptoJS.DES.decrypt(text, e, {
        mode: CryptoJS.mode.ECB,
        padding: CryptoJS.pad.Pkcs7
    });
    return CryptoJS.enc.Utf8.stringify(a).toString()
}


export default  {
    desEncrypt,
    desDecrypt
}
