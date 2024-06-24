/**
 * @param {string} path
 * @returns {Boolean}
 */
export function isExternal(path) {
  return /^(https?:|mailto:|tel:)/.test(path)
}

export function rulesObj(info) {
  return {
    name: [{ required: true, message: `请输入${info}名称`, trigger: 'blur' }],
    sortValue: [{ required: true, message: '请输入排序号', trigger: 'blur' }],
    level: [{ required: true, message: `请输入${info}等级`, trigger: 'blur' }],
    isWork: [{ required: true, message: '请选择是否开售', trigger: 'change' }],
    alias: [{ required: true, message: `请输入${info}别名`, trigger: 'blur' }],
    categoryCode: [
      { required: true, message: `请输入${info}编号`, trigger: 'blur' },
      { pattern: /^[a-zA-Z0-9]{1,10}$/, message: `${info}编号必须为数字和字母` }
    ],
    parentCode: [{ required: true, message: `请输入${info}类型`, trigger: 'blur' }],
    isShow: [{ required: true, message: '请选择是否删除', trigger: 'change' }],
    gameCode: [
      { required: true, message: `请输入${info}编号`, trigger: 'blur' },
      { pattern: /^[a-zA-Z0-9]{1,10}$/, message: `${info}编号必须为数字和字母` }
    ],
    startlottoTimes: [{ required: true, message: `请输入${info}别名`, trigger: 'blur' }]
  }
}

export function MenmerAddMonery() {
  return {
    handleNum: [
      { required: true, message: '请输入金额!', trigger: 'blur' },
      {
        pattern: /^(([1-9][0-9]*)|(([0]\.\d{1,3}|[1-9][0-9]*\.\d{1,3})))$/,
        message: '金额为数字值或最多保留三位小数',
        trigger: 'blur'
      }
    ],
    addAmount: [
      { required: true, message: '请输入增加金额!', trigger: 'blur' },
      {
        pattern: /^(([1-9][0-9]*)|(([0]\.\d{1,3}|[1-9][0-9]*\.\d{1,3})))$/,
        message: '增加金额为数字值或最多保留三位小数',
        trigger: 'blur'
      }
    ],
    isIdMoney: [{ required: true, message: '会员ID与金额不能为空!', trigger: 'blur' }],
    damaRatio: [
      { required: true, message: '请输入打码倍数!', trigger: 'blur' },
      { pattern: /^[0-9]*$/, message: '打码倍数为数字值', trigger: 'blur' }
    ]
  }
}

export function MenmerDamaoRules() {
  return {
    reducedamaliang: [
      { required: true, message: '请输入减少打码量!', trigger: 'blur' },
      { pattern: /^[1-9]\d*$/, message: '减少打码量必须是大于0的整数!', trigger: 'blur' }
    ],
    adddamaliang: [
      { required: true, message: '请输入增加打码量!', trigger: 'blur' },
      { pattern: /^[1-9]\d*$/, message: '增加打码量必须是大于0的整数!', trigger: 'blur' }
    ],
    isIdMoney: [{ required: true, message: '会员ID与打码量不能为空!', trigger: 'blur' }],
    remarkreson: [{ required: true, message: '请填写操作原因!', trigger: 'blur' }]
  }
}

export const ThirdCategory = () => {
  return {
    name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }],
    categoryId: [{ required: true, message: '请输入分类编号', trigger: 'blur' }],
    infoHost: [{ required: true, message: '请输入基础接口', trigger: 'blur' }],
    dataHost: [{ required: true, message: '请输入数据接口', trigger: 'blur' }],
    md5Key: [{ required: true, message: '请输入md5密钥', trigger: 'blur' }],
    aesKey: [{ required: true, message: '请输入AesKey加密参数', trigger: 'blur' }],
    iv: [{ required: true, message: '请输入IV向量', trigger: 'blur' }],
    backUrl: [
      { required: true, message: '请输入商户地址', trigger: 'blur' },
      { pattern: /(http|https):\/\/([\w.]+\/?)\s*/, message: '地址格式不正确', trigger: 'blur' }
    ],
    domain: [
      { required: true, message: '请输入动态游戏名称', trigger: 'blur' },
      { pattern: /(http|https):\/\/([\w.]+\/?)\s*/, message: '地址格式不正确', trigger: 'blur' }
    ],
    sort: [
      { required: true, message: '请输入排序', trigger: 'blur' },
      { pattern: /^\d+$/, message: '请输入整数', trigger: ['blur'] }
    ],
    isWork: [{ required: true, message: '请选择是否开售', trigger: 'blur' }],
    // type: [{ required: true, message: '请选择游戏分类', trigger: 'blur' }],
    gameType: [{ required: true, message: '请选择分类', trigger: 'blur' }],
    platformCode: [{ required: true, message: '请选择第三方平台', trigger: 'blur' }],
    gameWallet: [{ required: true, message: '请选择分类钱包', trigger: 'blur' }]
  }
}
export const ThirdManage = () => {
  return {
    platformCode: [{ required: true, message: '请选择平台', trigger: 'blur' }],
    name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }],
    categoryId: [{ required: true, message: '请选择游戏类型', trigger: 'blur' }],
    gameId: [{ required: true, message: '请输入游戏编号', trigger: 'blur' }],
    icon: [{ required: true, message: '请上传logo', trigger: 'blur' }],
    sort: [
      { required: true, message: '请输入排序', trigger: 'blur' },
      { pattern: /^\d+$/, message: '请输入整数', trigger: ['blur'] }
    ],
    oddType: [{ required: true, message: '请输入限红Id', trigger: 'blur' }],
    isWork: [{ required: true, message: '请选择是否开售', trigger: 'blur' }],
    isShow: [{ required: true, message: '请选择是否显示', trigger: 'blur' }],
    countryCodeList: [{ required: true, message: '请选择至少一个地区', trigger: 'blur' }]
  }
}
