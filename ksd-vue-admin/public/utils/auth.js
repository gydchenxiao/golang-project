// js-cookie
import Cookies from 'js-cookie'
// crypto-js
import CryptoJS from 'crypto-js'

const TokenKey = 'unilive-token'
export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

export function setCookie(username, password, exdays) {
  return Cookies.set('unilivemanage', CryptoJS.AES.encrypt(username + '_' + password, 'unilive').toString(), {
    expires: new Date().getTime() + 24 * 60 * 60 * 1000 * exdays
  })
}

export function getCookie() {
  var tmp = document.cookie
  if (tmp != undefined && tmp.length > 0 && tmp.indexOf('unilivemanage=') != -1) {
    const arr = document.cookie.split('unilivemanage=')
    return CryptoJS.AES.decrypt(arr[1].split(';')[0], 'unilive').toString(CryptoJS.enc.Utf8)
  }
  return ''
}

export function clearCookie() {
  return Cookies.remove('unilivemanage')
}
