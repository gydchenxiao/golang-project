// 引用 crypto-js 加密模块
import CryptoJS from 'crypto-js';

export function MD5(text) {
    return CryptoJS.MD5(text).toString()
}


