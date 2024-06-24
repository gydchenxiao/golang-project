export const moment = (date = new Date()) => {
  class Moment {
    constructor(date) {
      this.date = date
    }
    rezo(num) {
      return num < 10 ? `0${num}` : num
    }
    before(day) {
      let datetime = this.date.getTime()
      datetime -= day * 60 * 60 * 24 * 1000
      this.date = new Date(datetime)
      return this
    }
    format(str) {
      const datetime = this.date
      const YYYY = datetime.getFullYear()
      const MM = datetime.getMonth() + 1
      const DD = datetime.getDate()
      const hh = datetime.getHours()
      const mm = datetime.getMinutes()
      const ss = datetime.getSeconds()
      const obj = {
        YYYY,
        MM,
        DD,
        hh,
        mm,
        ss
      }
      for (const key in obj) {
        const reg = new RegExp(key)
        const val = this.rezo(obj[key])
        str = str.replace(reg, () => {
          return val
        })
      }
      return str
    }
  }
  return new Moment(date)
}
