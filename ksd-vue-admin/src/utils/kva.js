const KVA = {
  alert(title, content, options) {
    // 默认值
    var defaultOptions = { icon: 'warning', confirmButtonText: '确定', cancelButtonText: '取消' }
    // 用户传递和默认值就行覆盖处理
    var opts = { ...defaultOptions, ...options }
    return ElMessageBox.alert(content, title, {
      //确定按钮文本
      confirmButtonText: opts.confirmButtonText,
      // 内容支持html
      dangerouslyUseHTMLString: true,
      // 是否支持拖拽
      draggable: true,
      // 修改图标
      type: opts.icon
    })
  },
  confirm(title, content, options) {
    // 默认值
    var defaultOptions = { icon: 'warning', confirmButtonText: '确定', cancelButtonText: '取消' }
    // 用户传递和默认值就行覆盖处理
    var opts = { ...defaultOptions, ...options }
    // 然后提示
    return ElMessageBox.confirm(content, title, {
      //确定按钮文本
      confirmButtonText: opts.confirmButtonText,
      //取消按钮文本
      cancelButtonText: opts.cancelButtonText,
      // 内容支持html
      dangerouslyUseHTMLString: true,
      // 是否支持拖拽
      draggable: true,
      // 修改图标
      type: opts.icon
    })
  },
  prompt(title, content, options) {
    // 默认值
    var defaultOptions = { confirmButtonText: '确定', cancelButtonText: '取消' }
    // 用户传递和默认值就行覆盖处理
    var opts = { ...defaultOptions, ...options }
    return ElMessageBox.prompt(content, title, {
      //确定按钮文本
      confirmButtonText: opts.confirmButtonText,
      //取消按钮文本
      cancelButtonText: opts.cancelButtonText,
      // 内容支持html
      dangerouslyUseHTMLString: true,
      // 是否支持拖拽
      draggable: true,
      // 输入框的正则验证
      inputPattern: opts.pattern,
      // 验证的提示内容
      inputErrorMessage: opts.message || '请输入正确的内容'
    })
  },
  message(message, type, duration = 3000) {
    //永远保持只有一个打开状态
    ElMessage.closeAll()
    return ElMessage({
      showClose: true,
      dangerouslyUseHTMLString: true,
      message,
      duration,
      type
    })
  },
  success(message) {
    return this.message(message, 'success')
  },
  warning(message) {
    return this.message(message, 'warning')
  },
  error(message) {
    return this.message(message, 'error')
  },
  info(message) {
    return this.message(message, 'info')
  },
  notifyError(message) {
    return this.notify('提示', message, 3000, { type: 'error', position: 'tr' })
  },
  notifySuccess(message) {
    return this.notify('提示', message, 3000, { type: 'success', position: 'tr' })
  },
  notify(title, message, duration = 3000, options) {
    // 默认值
    var defaultOptions = { type: 'info', position: 'tr' }
    // 用户传递和默认值就行覆盖处理
    var opts = { ...defaultOptions, ...options }
    //永远保持只有一个打开状态
    ElNotification.closeAll()
    var positionMap = {
      tr: 'top-right',
      tl: 'top-left',
      br: 'bottom-right',
      bl: 'bottom-left'
    }
    return ElNotification({
      title,
      message,
      duration: duration,
      type: opts.type,
      position: positionMap[opts.position || 'tr'],
      dangerouslyUseHTMLString: true
    })
  }
}

export default KVA
