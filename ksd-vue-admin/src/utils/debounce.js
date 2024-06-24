let timeout = null

/**
 * 防抖原理：一定时间内，只有最后一次操作，再过wait毫秒后才执行函数
 * 在指定的时间内，无论你执行多少次轮回，只要还没有停止那么就继续拦截。
 * 一句话：在有效（3000）以内，都是无效请求，只有超过3000会执行一次
 * 
 * 防抖：
 * 1：在指定的时间轮回执行，比如：1000，第一秒：会执行，第二秒：会执行
 *
 * @param {Function} func 要执行的回调函数
 * @param {Number} wait 延时的时间
 * @param {Boolean} immediate 是否立即执行
 * @return null
 */
function debounce(func, wait = 500, immediate = false) {
  // 清除定时器
  if (timeout !== null) clearTimeout(timeout)
  // 立即执行，此类情况一般用不到
  if (immediate) {
    const callNow = !timeout
    timeout = setTimeout(() => {
      timeout = null
    }, wait)
    if (callNow) typeof func === 'function' && func()
  } else {
    // 设置定时器，当最后一次操作后，timeout不会再被清除，所以在延时wait毫秒后执行func回调方法
    timeout = setTimeout(() => {
      typeof func === 'function' && func()
    }, wait)
  }
}

export default debounce
