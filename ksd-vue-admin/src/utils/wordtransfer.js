
// 这个是把is_deleted====isDeleted
export const B2C = function (sName) {
    let s = sName
    // 处理特殊情况 eg -www
    if (s && s[0] === '-') s = s.substring(1);
    return s.replace(/-(\w)/g, function (_, c, d, e) {
        return c ? c.toUpperCase() : ''
    })
}

// 这个是把 isDeleted ==== is_deleted----error
export const C2B = function (name) {
    let copyName = name.toLowerCase();
    let index = 0;
    let str = '';
    while (index != name.length) {
        if (copyName[index] !== name[index]) {
            str = str + '_' + copyName[index];
        } else {
            str += copyName[index];
        }
        index++;
    }
    return str;
}
