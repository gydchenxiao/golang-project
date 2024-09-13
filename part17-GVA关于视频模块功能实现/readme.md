## 三种校验的写法

为了测试下面三种方法，就直接把路由写在全局router路由的地方了：src\router\index.js

```js
{
  // form 表单校验
  path: '/form/login',
  name: 'FormLogin',
  component: () => import('@/view/formvalidator/login.vue')
},
{
  // el-form 组件表单校验
  path: '/form/reg',
  name: 'FormReg',
  component: () => import('@/view/formvalidator/reg.vue')
},
{
  // 自定义组件 form 表单校验
  path: '/form/test',
  name: 'FormTest',
  component: () => import('@/view/formvalidator/test.vue')
},
```

测试链接分别是：http://localhost:8080/#/form/login，http://localhost:8080/#/form/reg，http://localhost:8080/#/form/test



### 1. form 表单校验

注意下面的代码的几个问题：

1. @blur="handleValidatorUsername" 所展现的功能？
2. id="password"，id 选择器所展现的功能？
3. v-model="formData.username"，v-model 所展现的功能？
4. @click.prevent="handleSubmit"，后面的 prevent 的功能，还有其他的什么方式吗？
5. 函数用 async 修饰与否有什么不一样的地方吗？
6. **承诺的意义和价值：是得代码优雅，异步请求变成串行执行，报错了就不会向下执行了(异步阻塞)。**



src\view\formvalidator\login.vue

```vue
<template>
    <div class="main">
        <h3>vue3 登陆表单验证</h3>
        <form>
            <div>
                <label class="label">账号</label>
                <input type="text" id="username" @blur="handleValidatorUsername" v-model="formData.username"
                    placeholder="请输入账号" class="input" />
            </div>
            <div>
                <label class="label">密码</label>
                <input tyep="password" id="password" v-model="formData.password" type="text" class="input"
                    placeholder="请输入密码" />
            </div>
            <div>
                <label class="label">确认密码</label>
                <input tyep="password" id="cpassword" v-model="formData.cpassword" type="text" class="input"
                    placeholder="请输入密码" />
            </div>
            <div>
                <!-- 下面的 prevent 的作用就是去掉 form 表单的默认提交能力的。见过三种写法，这是其中的一种写法。 -->
                <button @click.prevent="handleSubmit">保存</button>
            </div>
        </form>
    </div>
</template>

<script setup>
import { reactive } from "vue"
import axios from 'axios'
import Ksd from '@/utils/index.js'

// 响应式数据
const formData = reactive({
    'username': '',
    'password': '',
    'cpassword': '',
})

// 如果没有写 username 的话是会报错的哦
const handleValidatorUsername = () => {
    if (!formData.username) {
        Ksd.error("请输入用户名!!!")
        return;
    }
}

// 验证
const validate = {
    // 下面的 async 使得函数的返回值是一个承诺类型的变量哦。
    // 有问题的用 Promise.reject()，没有问题的用 Promise.resolve()
    async validator() {
        // 获取表单formData的数据（v-model）
        if (!formData.username) {
            Ksd.error("请输入用户名!!!")
            return Promise.reject("username");
        }
        if (!formData.password) {
            Ksd.error("请输入密码!!!")
            return Promise.reject("password");
        }
        if (!formData.cpassword) {
            Ksd.error("请输入确认密码!!!")
            return Promise.reject("cpassword");
        }
        // 密码是否相同
        if (formData.password != formData.cpassword) {
            Ksd.error("输入确认密码和密码不一致!!!")
            return Promise.reject("cpassword");
        }
        return Promise.resolve("success");
    }
}

// 提交 formData 数据的时候得验证一下的哦
const handleSubmit = async () => {
    try {
        const resp = await validate.validator()
        console.log(resp)
        alert("success")

        axios.post("toLogin"); // 异步请求：http:/127.0.0.1:8888/toLogin

    } catch (e) {
        // 使用的是 js 的代码，使用 id 的方式找到对象并聚焦
        document.querySelector("#" + e).focus();
    }
}
</script>

<style lang="css">
.main {
    text-align: center;
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
    margin-top: 10px;
}
</style>
```



### 2. el-form 表单校验

看完下面的代码回答几个问题：

1. :model="loginForm" 的可有可无和下面的 v-model 有什么关系吗？
2. :model="loginForm" 的可有可无和 :rules="rules" 有什么关系吗？
3. 注意自定义验证和官方提供的验证的区别和写法。
4. reactive 和 ref 在响应式数据上的区别？，在使用上的区别？



src\view\formvalidator\reg.vue

```vue
<template>
    <el-form ref="loginFormRef" :model="loginForm" status-icon :rules="rules" label-width="120px" class="demo-ruleForm">
        <el-form-item label="密码" prop="pass">
            <el-input v-model="loginForm.pass" placeholder="请输入密码" type="password" autocomplete="off" />
        </el-form-item>
        <el-form-item label="确认密码" prop="checkPass">
            <el-input v-model="loginForm.checkPass" placeholder="请输入确认密码" type="password" autocomplete="off" />
        </el-form-item>
        <el-form-item label="年龄" prop="age">
            <el-input v-model.number="loginForm.age" placeholder="请输入年龄" />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="submitForm">注册</el-button>
            <el-button @click="resetForm">重置 </el-button>
            <el-button @click="clearValidForm">清除验证</el-button>
        </el-form-item>
    </el-form>
</template>
  
<script  setup>
import { reactive, ref } from 'vue'

// 定义一个表单的ref
const loginFormRef = ref()

// 定义表单的响应式数据，建议都使用reactive
const loginForm = reactive({
    pass: '',
    checkPass: '',
    age: '',
})


// 2. 自定义验证（使用自定义的函数）
// rule 验证规则
// value 你填写表的值 
// callback 就是验证的错误和正确的提示
const validateCheckPass = (rule, value, callback) => {
    if (value !== loginForm.pass) {
        callback(new Error("确认密码和密码不一致!"))
    } else {
        callback()
    }
}
const validateCheckPassLen = (rule, value, callback) => {
    if (value.length < 3 || value.length > 20) {
        callback(new Error("密码长度是3到20位!"))
    } else {
        callback()
    }
}
// 1. 验证规则（使用官方的 required, message, trigger 等属性）
// 注意： 上面的 :model="loginForm" 就是为了后面的 :rules="rules" 使用的哦，而不是为了下面的 v-model，可以通过去掉 :model="loginForm" 来试试看。
//        下面 rules 中的 pass，checkPass, age 就回去 loginForm 对象中看有没有
const rules = reactive({
    pass: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { validator: validateCheckPassLen, trigger: 'blur' }
    ],
    checkPass: [
        { required: true, message: '请输入确认密码', trigger: 'blur' },
        { validator: validateCheckPassLen, trigger: 'blur' },
        { validator: validateCheckPass, trigger: 'blur' }
    ],
    age: { required: true, message: '请输入年龄', trigger: 'blur' },
})


// 提交表单
const submitForm = () => {
    // 拿到 el-form 组件对象，通过组件对象调用组件暴露出来的校验函数
    loginFormRef.value.validate((valid) => {
        // 如果为valid = true 校验全部都是正确的
        if (valid) {
            console.log('submit!')
        } else {
            console.log('error submit!')
            return false
        }
    })
}

// 重置表单
const resetForm = () => {
    loginFormRef.value.resetFields()
}

// 清除验证
const clearValidForm = () => {
    loginFormRef.value.clearValidate()
}
</script>
```



### 3. 自定义组件表单校验

自定义表单验证组件：src\view\formvalidator\KsdForm.vue

```vue
<template>
    <form ref="formRef">
        <slot></slot>
    </form>
</template>

<script setup>
import { ref } from "vue"

// 自定义属性
const props = defineProps({
    model: {
        type: Object,
        default: {}
    },
    rules: {
        type: Object,
        default: {}
    }
})

const formRef = ref();
// 重置表单
const resetFields = () => {
    formRef.value.reset();
}

// 验证方法
const validate = async (callback) => {
    var modelData = props.model; //{username:"xxx","password":"xxx"}
    var rulesData = props.rules; //{"username":{required:true}}
    // 遍历规则-开始和数据模型中的数据开始进行校验
    for (const key in rulesData) {
        // 获取到验证的属性
        var field = key;
        // 获取验证属性的规则
        var rule = rulesData[key]
        // 获取需要验证的数据模型
        for (var ckey in rule) {
            var message = rule["message"]
            // 获取
            if ((ckey == "required" && rule[ckey]) && !modelData[field]) {
                callback && callback(false, message)
                return Promise.resolve(message)
            }
        }
    }
    callback && callback(true, "success")
    return Promise.resolve(true);
}

// 清除验证

// 单个属性验证

// 暴露方法
defineExpose({
    resetFields,
    validate
})
</script>
```

测试上面的自定义表单验证组件：src\view\formvalidator\test.vue

```vue
<template>
    <div class="main">
        <h3>vue3 表单验证</h3>
        <ksd-form ref="videoForm" :model="formData" :rules="rules">
            <div>
                <label class="label">账号</label>
                <input type="text" id="username" @blur="handleValidatorUsername" v-model="formData.username"
                    placeholder="请输入账号" class="input" />
            </div>
            <div>
                <label class="label">密码</label>
                <input tyep="password" id="password" v-model="formData.password" type="text" class="input"
                    placeholder="请输入密码" />
            </div>
            <div>
                <label class="label">确认密码</label>
                <input tyep="password" id="cpassword" v-model="formData.cpassword" type="text" class="input"
                    placeholder="请输入密码" />
            </div>
            <div>
                <button @click.prevent="handleSubmit">保存</button>
                <button @click.prevent="handleReset">重置</button>
            </div>
        </ksd-form>
    </div>
</template>
<script setup>
import { reactive, ref } from "vue"
import KsdForm from './KsdForm.vue'
import Ksd from '@/utils/index.js'
const videoForm = ref({});
const formData = reactive({ username: "", password: "" })
const rules = reactive({
    'username': { required: true, message: "请输入用户名" },
    'password': { required: true, message: "请输入密码" },
})

// 重置
const handleReset = () => {
    videoForm.value.resetFields();
}

// 提交
const handleSubmit = () => {
    videoForm.value.validate((valid, msg) => {
        if (valid) {
            alert("success")
        } else {
            alert(msg)
        }
    })
}

</script>
<style lang="css">
.main {
    text-align: center;
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
    margin-top: 10px;
}
</style>
```

看完上面的 test.vue 测试自定义表单验证组件的代码，回答下面的几个问题：

1. 自定义的和element-puls官方的比较注意其原理。

2. 自定义的暴露了对外提供的接口 ----> 官方的也是哦。

   