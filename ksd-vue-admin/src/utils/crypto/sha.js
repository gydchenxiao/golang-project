// 引用 crypto-js 加密模块
import CryptoJS from 'crypto-js';

export function SHA1(text) {
    return CryptoJS.SHA1(text).toString();
}


