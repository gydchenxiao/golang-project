// 动态加载外部js
// @param RESOURCE_LIST 外部地址集合["address"]
export function loadVoLteResourceList (RESOURCE_LIST, success) {
  return new Promise(r => {
    RESOURCE_LIST.reduce((res, el) => res.then(() => loadScript(el)), Promise.resolve()).then(() => {
      r()
    }).catch((error) => {
      console.error('外呼VoLTE sdk 前置 js 资源加载失败:', error.name, error.message)
      return Promise.reject(error)
    })
  })
}

export function loadScript (url) {
  return new Promise((resolve, reject) => {
    const script = document.createElement('script')
    script.onload = () => resolve()
    script.onerror = () => reject(new Error(`Load script from ${url} failed`))
    script.src = url
    const head =
    document.head || document.getElementsByTagName('head')[0]
    ;(document.body || head).appendChild(script)
  })
}

// 专门加载百度地图
export function loadBMap (ak) {
  return new Promise(function (resolve, reject) {
    if (typeof BMap !== 'undefined') {
      resolve(BMap)
      return true
    }
    window.onBMapCallback = function () {
      resolve(BMap)
    }
    let script = document.createElement('script')
    script.type = 'text/javascript'
    script.src = 'http://api.map.baidu.com/api?v=2.0&ak=' + ak + '&__ec_v__=20190126&callback=onBMapCallback'
    script.onerror = reject
    document.head.appendChild(script)
  })
}
