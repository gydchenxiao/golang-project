# Vue3+异步验证器实现的表单验证示例代码

官网：https://github.com/yiminghe/async-validator

## 验证

- 前端js验证
- 后端验证
- 国际化（动态路由）cn.js / en.js

## 你在做表单处理的时候，你必须要知道的事情

- button/input type=submit/button都具有提交表单的能力。因为在前后端的开发中，提交form都是使用ajax处理，也就是通过js进行处理，那么button/submit提交表单表能力就必要去掉，否则就造成同步提交，而这种提交毫无意义。而且会造成奇怪刷新或者跳转。

  ```js
   <button @click.prevent="handleSubmit">保存</button>
  ```

- 为什么用js验证还要使用服务端验证

  ```js
  <template>
      <div class="main">
          <h3>vue3 表单验证</h3>
          <form>
              <div>
                  <label class="label">账号</label>
                  <input  type="text" id="username"  v-model="formData.username" placeholder="请输入账号" class="input" />
              </div>
              <div>
                  <label class="label">密码</label>
                  <input  tyep="password" id="password" v-model="formData.password" type="text" class="input"  placeholder="请输入密码"  />
              </div>
              <div>
                  <button @click.prevent="handleSubmit">保存</button>
              </div>
          </form>
      </div>
  </template>
  <script setup>
  import { reactive } from "vue"
  import Ksd from '@/utils/index.js'
  
  // 响应式数据
  const formData = reactive({
      'username': '',
      'password': ''
  })
  
  // 验证
  const validate = {
      async validator(){
          // 获取表单formData的数据（v-model）
          if(!formData.username){
              Ksd.error("请输入用户名!!!")
              return Promise.reject("username");
          }
          if(!formData.password){
              Ksd.error("请输入密码!!!")
              return Promise.reject("password");
          }
          return Promise.resolve("success");
      }
  }
  
  const handleSubmit = async ()=>{
     try{
        const resp = await validate.validator()
        console.log(resp)
        alert("success")
     }catch(e){
        document.querySelector("#"+e).focus();
     }
  }
  </script>
  <style lang="css">
  .main{
      text-align:center;
  }
  .label {
      padding-right: 10px;
      padding-left: 10px;
      display: inline-block;
      box-sizing: border-box;
      width: 100px;
      text-align: right;
  }
  .input {
      width: 200px;
      height: 30px;
      text-indent: 1em;
      margin-top:10px;
  }
  </style>
  ```

  为什么要用服务端校验

  1： js校验只是一种简单的防护手段，很多时候，我们可以通过流程的工具debug可以看到你提供的服务请求地址，我可以拿到你请求，通过专业攻防软件或者服务语言（go/java）等通过这些软件和语言本身就同请求和参数来模拟表单请求提交。请问你js校验还用码？

  2：==所以真正的校验其实还是服务端才是最安全的。js校验只是一层防护和配合。也是一种减少服务一种机制，而把所有校验都交给服务端，造成服务端的请求压力。==

  3：在这种添加，修改的时候，数据的合法性是非常必要，但是在做业务的时候，不建议验证和业务一起开发，除非非常简单。而是把业务功能和验证分阶段来完成。



## 关于开发中什么时候去做验证

- 验证是为了保证数据的合法性的一种校验机制。真正的验证：前端（js）和接口端(go/java)都要进行校验。

**常用的校验分为：**

参考这个文件：https://github.com/validatorjs/validator.js/blob/master/src/index.js

```js
import toDate from './lib/toDate';
import toFloat from './lib/toFloat';
import toInt from './lib/toInt';
import toBoolean from './lib/toBoolean';
import equals from './lib/equals';
import contains from './lib/contains';
import matches from './lib/matches';

import isEmail from './lib/isEmail';
import isURL from './lib/isURL';
import isMACAddress from './lib/isMACAddress';
import isIP from './lib/isIP';
import isIPRange from './lib/isIPRange';
import isFQDN from './lib/isFQDN';
import isDate from './lib/isDate';
import isTime from './lib/isTime';

import isBoolean from './lib/isBoolean';
import isLocale from './lib/isLocale';

import isAlpha, { locales as isAlphaLocales } from './lib/isAlpha';
import isAlphanumeric, { locales as isAlphanumericLocales } from './lib/isAlphanumeric';
import isNumeric from './lib/isNumeric';
import isPassportNumber from './lib/isPassportNumber';
import isPort from './lib/isPort';
import isLowercase from './lib/isLowercase';
import isUppercase from './lib/isUppercase';

import isIMEI from './lib/isIMEI';

import isAscii from './lib/isAscii';
import isFullWidth from './lib/isFullWidth';
import isHalfWidth from './lib/isHalfWidth';
import isVariableWidth from './lib/isVariableWidth';
import isMultibyte from './lib/isMultibyte';
import isSemVer from './lib/isSemVer';
import isSurrogatePair from './lib/isSurrogatePair';

import isInt from './lib/isInt';
import isFloat, { locales as isFloatLocales } from './lib/isFloat';
import isDecimal from './lib/isDecimal';
import isHexadecimal from './lib/isHexadecimal';
import isOctal from './lib/isOctal';
import isDivisibleBy from './lib/isDivisibleBy';

import isHexColor from './lib/isHexColor';
import isRgbColor from './lib/isRgbColor';
import isHSL from './lib/isHSL';

import isISRC from './lib/isISRC';

import isIBAN, { locales as ibanLocales } from './lib/isIBAN';
import isBIC from './lib/isBIC';

import isMD5 from './lib/isMD5';
import isHash from './lib/isHash';
import isJWT from './lib/isJWT';

import isJSON from './lib/isJSON';
import isEmpty from './lib/isEmpty';

import isLength from './lib/isLength';
import isByteLength from './lib/isByteLength';

import isUUID from './lib/isUUID';
import isMongoId from './lib/isMongoId';

import isAfter from './lib/isAfter';
import isBefore from './lib/isBefore';

import isIn from './lib/isIn';

import isLuhnNumber from './lib/isLuhnNumber';
import isCreditCard from './lib/isCreditCard';
import isIdentityCard from './lib/isIdentityCard';

import isEAN from './lib/isEAN';
import isISIN from './lib/isISIN';
import isISBN from './lib/isISBN';
import isISSN from './lib/isISSN';
import isTaxID from './lib/isTaxID';

import isMobilePhone, { locales as isMobilePhoneLocales } from './lib/isMobilePhone';

import isEthereumAddress from './lib/isEthereumAddress';

import isCurrency from './lib/isCurrency';

import isBtcAddress from './lib/isBtcAddress';

import { isISO6346, isFreightContainerID } from './lib/isISO6346';
import isISO6391 from './lib/isISO6391';
import isISO8601 from './lib/isISO8601';
import isRFC3339 from './lib/isRFC3339';
import isISO31661Alpha2 from './lib/isISO31661Alpha2';
import isISO31661Alpha3 from './lib/isISO31661Alpha3';
import isISO4217 from './lib/isISO4217';

import isBase32 from './lib/isBase32';
import isBase58 from './lib/isBase58';
import isBase64 from './lib/isBase64';
import isDataURI from './lib/isDataURI';
import isMagnetURI from './lib/isMagnetURI';
import isMailtoURI from './lib/isMailtoURI';

import isMimeType from './lib/isMimeType';

import isLatLong from './lib/isLatLong';
import isPostalCode, { locales as isPostalCodeLocales } from './lib/isPostalCode';

import ltrim from './lib/ltrim';
import rtrim from './lib/rtrim';
import trim from './lib/trim';
import escape from './lib/escape';
import unescape from './lib/unescape';
import stripLow from './lib/stripLow';
import whitelist from './lib/whitelist';
import blacklist from './lib/blacklist';
import isWhitelisted from './lib/isWhitelisted';

import normalizeEmail from './lib/normalizeEmail';

import isSlug from './lib/isSlug';
import isLicensePlate from './lib/isLicensePlate';
import isStrongPassword from './lib/isStrongPassword';

import isVAT from './lib/isVAT';

const version = '13.9.0';

const validator = {
  version,
  toDate,/*日期*/
  toFloat,/*浮点数*/
  toInt,/*整数*/
  toBoolean,/*布尔类型*/
  equals,/*判断两个是否相同*/
  contains, /*包括*/
  matches,/*匹配*/
  isEmail,/*邮箱*/
  isURL,/*URL*/
  isMACAddress, /*mac*/
  isIP, /*ip*/
  isIPRange,
  isFQDN,
  isBoolean,
  isIBAN,
  isBIC,
  isAlpha,
  isAlphaLocales,
  isAlphanumeric,
  isAlphanumericLocales,
  isNumeric,
  isPassportNumber,
  isPort,
  isLowercase,
  isUppercase,
  isAscii,
  isFullWidth,
  isHalfWidth,
  isVariableWidth,
  isMultibyte,
  isSemVer,
  isSurrogatePair,
  isInt,
  isIMEI,
  isFloat,
  isFloatLocales,
  isDecimal,
  isHexadecimal,
  isOctal,
  isDivisibleBy,
  isHexColor,
  isRgbColor,
  isHSL,
  isISRC,
  isMD5,
  isHash,
  isJWT,
  isJSON,
  isEmpty,
  isLength,
  isLocale,
  isByteLength,
  isUUID,
  isMongoId,
  isAfter,
  isBefore,
  isIn,
  isLuhnNumber,
  isCreditCard,
  isIdentityCard,
  isEAN,
  isISIN,
  isISBN,
  isISSN,
  isMobilePhone,
  isMobilePhoneLocales,
  isPostalCode,
  isPostalCodeLocales,
  isEthereumAddress,
  isCurrency,
  isBtcAddress,
  isISO6346,
  isFreightContainerID,
  isISO6391,
  isISO8601,
  isRFC3339,
  isISO31661Alpha2,
  isISO31661Alpha3,
  isISO4217,
  isBase32,
  isBase58,
  isBase64,
  isDataURI,
  isMagnetURI,
  isMailtoURI,
  isMimeType,
  isLatLong,
  ltrim,
  rtrim,
  trim,
  escape,
  unescape,
  stripLow,
  whitelist,
  blacklist,
  isWhitelisted,
  normalizeEmail,
  toString,
  isSlug,
  isStrongPassword,
  isTaxID,
  isDate,
  isTime,
  isLicensePlate,
  isVAT,
  ibanLocales,
};

export default validator;
```



## vue3的表单验证

这里我们使用 [async-validator](https://link.juejin.cn/?target=https%3A%2F%2Fgithub.com%2Fyiminghe%2Fasync-validator) 这是个异步验证表单的插件，在github上有 5k+ 的star，使用的也很广泛，比如 `Ant.design`，`Element UI`， `Naive UI` 等都在使用这个插件，甚至与有些Nodejs后端项目也在使用这个。

先安装一下这个插件，在命令行输入

```
pnpm install async-validator
```

这里 `async-validator` 版本是 `4.1.1`

### 1.表单代码

打开项目中的 App.vue 文件，删除多余的文件内容，输入标题 **vue3 表单验证**，并添加一些初始代码

```
<template>
    <div class="main">
        <h3>vue3 表单验证-注册</h3>
        <form>
            <div>
                <label class="label">账号</label>
                <input  type="text" id="username"  placeholder="请输入账号" class="input" />
            </div>
            <div>
                <label class="label">密码</label>
                <input  tyep="password" id="password"  type="text" class="input"  placeholder="请输入密码"  />
            </div>
            <div>
                <button @click.prevent="handleSubmitReg">注册</button>
            </div>
        </form>
    </div>
</template>
<script setup>
const handleSubmitReg= async ()=>{
  
}
</script>
<style lang="css">
.main{
    text-align:center;
}
.label {
    padding-right: 10px;
    padding-left: 10px;
    display: inline-block;
    box-sizing: border-box;
    width: 100px;
    text-align: right;
}
.input {
    width: 200px;
    height: 30px;
    text-indent: 1em;
    margin-top:10px;
}
</style>
```

是不是看起来有点丑，别急，我们加点css代码，简单的美化一下

```
<template>
    <div class="main">
        <h3>Vue3表单验证</h3>

        <form class="form-box">
            <div class="form-group ">
                <label class="label">账号</label>
                <input type="text" class="input" placeholder="请输入账号"  />
            </div>
            <div class="form-group">
                <label class="label">密码</label>
                <input  tyep="password" type="text"  placeholder="请输入密码" class="input" />
            </div>

            <div class="form-group">
                <button class="btn ">保存</button>
            </div>
        </form>
    </div>
</template>

<style scoped>
.main {
    text-align: center;
}
.btn{
    margin: 0;
    line-height: 1;
    padding: 15px;
    height: 30px;
    width: 60px;
    font-size: 14px;
    border-radius: 4px;
    color: #fff;
    background-color: #2080f0;
    white-space: nowrap;
    outline: none;
    position: relative;
    border: none;
    display: inline-flex;
    flex-wrap: nowrap;
    flex-shrink: 0;
    align-items: center;
    justify-content: center;
    user-ｓｅｌｅｃｔ: none;
    text-align: center;
    cursor: pointer;
    text-decoration: none;
}
.form-box{
    width: 500px;
    max-width: 100%;
    margin: 0 auto;
    padding: 10px;
}
.form-group{
    margin: 10px;
    padding: 10px 15px 10px 0
}
.label {
    padding-right: 10px;
    padding-left: 10px;
    display: inline-block;
    box-sizing: border-box;
    width: 110px;
    text-align: right;
}

.input {
    width: calc(100% - 120px);
    height: 28px;
}
</style>
```



### 2.添加验证



#### 2-1. 初始化

引入`ref` 属性和 `async-validator`，这里我们给每个 input 框添加 `v-model` 绑定属性,

```
// html
<input type="text" v-model="form.account" class="input" placeholder="请输入账号"  />
<input  tyep="password" v-model="form.password" type="text"  placeholder="请输入密码" class="input" />
// ｓｃｒｉｐｔ
import { ref } from "vue"
import Schema from 'async-validator';

const form = ref({
    account: null,
    password: null,
})
```

根据表单的情况，我们定义一个对象，这个对象里面存储了需要校验的对象和校验不通过时的信息

```
const rules = {
    account: { required: true, message: '请输入账号' },
    password: { required: true, message: '请输入密码' }
}
```

实例化 Schema, 将 rules 传入 Schema，得到一个 validator

```
const validator = new Schema(rules)
```

验证单个表单我们使用 **失去焦点事件**, 定义一个函数，将这个函数添加到 account input上的失焦事件上

```
// html
<input v-model="account" type="text" class="input" @blur="handleBlurAccount" placeholder="请输入账号"  />
// ｓｃｒｉｐｔ
const handleBlurAccount = () => {}
```

接着将实例化后的校验器函数写到 handleBlurAccount 中

```
const handleBlurAccount = () => {
    validator.validate({account: form.value.account}, (errors, fields) => {
        if (errors && fields.account) {
            console.log(fields.account[0].message);
            return errors
        }
    })
}
```

在account 的 input 中测试，我们可以看到在控制台打印出了 **请输入账号** 等字

同样的，我们给密码框也添加如下代码

```
//html
<input v-model="form.password" tyep="password" type="text" @blur="handleBlurPassword"  placeholder="请输入密码" class="input" />
//ｓｃｒｉｐｔ
const handleBlurPassword = () => {
    validator.validate({password: form.value.password}, (errors, fields) => {
        if (errors && fields.password) {
            console.log(errors, fields);
            console.log(fields.password[0].message);
            return errors
        }
    })
}
```



#### 2-2. 多个表单的验证

当然这里校验的只是单个input的，我们接下来说说多个表单的校验，定义一个点击事件为submit，将submit事件添加到button上，当然不要忘记阻止浏览器默认事件

```
const submit = (e) => {
    e.preventDefault();
    validator.validate(form.value, (errors, fields) => {
        if (errors) {
            for(let key of errors) {
                console.log(key.message);
            }
            return errors
        }
    })
}
```



#### 2-3. Promise方式验证

了上面的方式，`async-validator` 还提供 Promise 的方式，我们把 submit 函数中的代码修改为如下

```
const submit = (e) => {
    e.preventDefault();
    validator.validate(form.value, (errors, fields) => {
        if (errors) {
            for(let key of errors) {
                console.log(key.message);
            }
            return errors
        }
    })
}
```

点击保存，同样的，我们可以看到控制台已经打印了错误信息，说明我们写的是合适的





#### 2-4. 正则验证

当然有时候我们会输入邮箱，电话号码等表单，这时候我们就需要添加正则来进行验证了,我们先添加两个表单，并添加失焦事件, 正则验证需要用到 `async-validator` 的属性 pattern，我们将符合要求的正则添加到 rules ，代码如下所示

```
<div class="form-group ">
    <label class="label">电话号码</label>
    <input v-model="form.phone" type="text" class="input" @blur="handleBlurPhone"
                    placeholder="请输入电话号码" />
</div>

<div class="form-group ">
    <label class="label">邮箱</label>
    <input v-model="form.email" type="text" class="input" @blur="handleBlurEmail"
                    placeholder="请输入邮箱" />
</div>


const form = ref({
    account: null,
    email: null,
    password: null,
})

const rules = {
    account: { required: true, message: '请输入账号' },
    phone: {
        required: true,
        pattern: /^1\d{10}$/,
        message: "请输入电话号码"
    },
    email: {
        required: true,
        pattern: /^([a-zA-Z0-9]+[_|_|\-|.]?)*[a-zA-Z0-9]+@([a-zA-Z0-9]+[_|_|.]?)*[a-zA-Z0-9]+\.[a-zA-Z]{2,6}$/,
        message: "请输入邮箱"
    },
    password: { required: true, message: '请输入密码' }
}

const handleBlurPhone = () => {
    validator.validate({ phone: form.value.phone }, (errors, fields) => {
        if (errors && fields.phone) {
            console.log(errors, fields);
            console.log(fields.phone[0].message);
            return errors
        }
    })
}
const handleBlurEmail = () => {
    validator.validate({ email: form.value.email }, (errors, fields) => {
        if (errors && fields.email) {
            console.log(errors, fields);
            console.log(fields.email[0].message);
            return errors
        }
    })
}
```

当然，测试是没有问题的



#### 2-5. 长度控制

假如你要控制表单输入内容的长度，可以使用属性 min 和 max，我们用 account 这个表单作为例子，我们 rules 对象的 account 中添加这两个属性，比如要求账号最少5个字符，最多10个字符，如下

```
account: { required: ``true``, min:5, max:10, message: ``'请输入账号'` `}
```

我们还可以使用 input 的原生属性 maxLength="10" 来控制用户的输入



#### 2-6. 多个验证条件

当我们有多个验证条件的时候，我们可以把 rules 的验证条件写成一个数组，我们还是用 account 这个表单作为例子，比如 账号要求必须用中文，且账号最少5个字符，最多10个字符，代码如下

```
account: [
    { required: true, min:5, max:10, message: '请输入账号' },
    { required: true, pattern: /[\u4e00-\u9fa5]/, message: '请输入中文账号' }
],
```



#### 2-5. 自定义验证

有时候，我们会有使用自定义验证函数的情况，以满足特殊验证情况，这时候，我们可以这样做

```
field:{
    required: true,
    validator(rule, value, callback){
      return value === '';
    },
    message: '值不等于 "".',
}
```

到这里，vue3的表单验证功能雏形已经基本出来了，下面我们对验证功能进行完善



### 3.优化完善

之前的表单验证虽然已经做出了，但是校验的提示信息是在控制台，这个很不友好，用户也看不到提示，所以这里我们完善下这部分功能

首先我们在 label 边加一个 "*" 表示必填，并且添加样式,给一个红色，醒目一些

```
<label class="label">
    <span>账号</span>
    <span class="asterisk"> *</span>
</label>
.asterisk{
    color: #d03050;
}
```

我们考虑到 `rules` 对象中 `required` 属性的作用，这里使用 vue 的条件判断语句 `v-if` 来判断，先定义一个函数，名字就叫 `getRequired`，然后将 `rules.account`,作为参数传进去，这里要重点说明一下，如果考虑封装验证方法，这里可以不用传参，不多说，后面讲到了，我们再说，先看代码

```
 <span class="asterisk" v-if="getRequired(rules.account)"> *</span>
const getRequired = (condition) => {
    if(Object.prototype.toString.call(condition) === "[object Object]") {
        return condition.required
    } else if (Object.prototype.toString.call(condition) === "[object Array]") {
        let result = condition.some(item => item.required)
        return result
    }
    return false
}
```

因为 `rules.account`, 有可能是对象或者数组，这里我们加一个判断区别下，如果传递进来的是对象，我们直接将属性`required`返回回去，至于`required`属性是否存在，这里没有必要多判断。 如果传递进来的是数组，我们使用 some 函数获取下结果，然后再返回.

修改 `rules.account` 的 `required` 值为false，星号消失，这里只要有一个`required` 值为true，那么这个星号就显示

我们接着来添加错误信息的显示与隐藏

我们定义一个对象 `modelControl`，这个对象里面动态存储错误信息，

```
const modelControl = ref({})
```

接着给 `account` 的 `input` 框添加一个自定义属性 `prop`, 属性值是 `account`, 再加一个div显示错误提示信息

```
<div class="form-group">
    <label class="label">
        <span>账号</span>
        <span class="asterisk" v-if="getRequired(rules.account)"> *</span>
    </label>
    <input v-model="form.account" type="text" maxLength="10" class="input" prop="account" @blur="handleBlurAccount" placeholder="请输入账号" />
    <div class="input feedback" v-if="modelControl['account']">{{modelControl['account']}}</div>
</div>

.feedback{
    color: #d03050;
    font-size:14px;
    margin-top: 3px;
    text-align:left;
    margin-left:110px;
}
```

为了动态的显示和隐藏错误信息，我们需要修改失焦事件 和 submit 事件,在事件执行的时候，动态的将值赋予或清除，代码如下

```
const handleBlurAccount = (e) => {
    const prop = e.target.attributes.prop.value
    if (!prop) {
        return false
    }
    validator.validate({ account: form.value.account }, (errors, fields) => {
        if (errors && fields.account) {
            console.log(errors, fields);
            console.log(fields.account[0].message);

            modelControl.value[prop] = fields[prop][0].message
            return errors
        }
        modelControl.value[prop] = null
    })
}
validator.validate(form.value).then((value) => {
        // 校验通过
    console.log(value);
}).catch(({ errors, fields }) => {
    console.log(errors, fields);
    for(let key in fields) {
        modelControl.value[key] = fields[key][0].message
    }
    console.log(modelControl);
    return errors
})
```

到这里 表单的动态验证功能基本算是完成了，但是我们发现，每次错误信息的展示都会使得input框跳动，所以还得调整下样式

```
.form-group {
    margin: 2px;
    padding: 10px 15px 3px 0;
    height:57px;
    transition: color .3s ease;
}
```