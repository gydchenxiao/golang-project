import CryptoJS from 'crypto-js';

export const base64Encode = (text) => {
    var srcs = CryptoJS.enc.Utf8.parse(text);
    var encodeData = CryptoJS.enc.Base64.stringify(srcs);
    return encodeData
}

export const base64Decode = (encodeData) => {
    var srcs = CryptoJS.enc.Base64.parse(encodeData);
    var decodeData = srcs.toString(CryptoJS.enc.Utf8);
    return decodeData
}