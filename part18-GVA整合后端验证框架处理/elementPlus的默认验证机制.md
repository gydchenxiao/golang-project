# elementPlus的默认验证机制



1 : 准备空模板

```vue
<template>
    <el-form ref="ruleFormRef" :model="ruleForm" status-icon :rules="rules" label-width="120px" class="demo-ruleForm">
        <el-form-item label="Password" prop="pass">
            <el-input v-model="ruleForm.pass" type="password" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Confirm" prop="checkPass">
            <el-input v-model="ruleForm.checkPass" type="password" autocomplete="off" />
        </el-form-item>
        <el-form-item label="Age" prop="age">
            <el-input v-model.number="ruleForm.age" />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="submitForm(ruleFormRef)">Submit</el-button>
            <el-button @click="resetForm(ruleFormRef)">Reset</el-button>
        </el-form-item>
    </el-form>
</template>
  
<script  setup>
import { reactive, ref } from 'vue'

const ruleFormRef = ref()
const checkAge = (rule, value, callback) => {
    if (!value) {
        return callback(new Error('Please input the age'))
    }
    setTimeout(() => {
        if (!Number.isInteger(value)) {
            callback(new Error('Please input digits'))
        } else {
            if (value < 18) {
                callback(new Error('Age must be greater than 18'))
            } else {
                callback()
            }
        }
    }, 1000)
}

const validatePass = (rule, value, callback) => {
    if (value === '') {
        callback(new Error('Please input the password'))
    } else {
        if (ruleForm.checkPass !== '') {
            if (!ruleFormRef.value) return
            ruleFormRef.value.validateField('checkPass', () => null)
        }
        callback()
    }
}
const validatePass2 = (rule, value, callback) => {
    if (value === '') {
        callback(new Error('Please input the password again'))
    } else if (value !== ruleForm.pass) {
        callback(new Error("Two inputs don't match!"))
    } else {
        callback()
    }
}

const ruleForm = reactive({
    pass: '',
    checkPass: '',
    age: '',
})

const rules = reactive({
    pass: [{ validator: validatePass, trigger: 'blur' }],
    checkPass: [{ validator: validatePass2, trigger: 'blur' }],
    age: [{ validator: checkAge, trigger: 'blur' }],
})

const submitForm = (formEl) => {
    if (!formEl) return
    formEl.validate((valid) => {
        if (valid) {
            console.log('submit!')
        } else {
            console.log('error submit!')
            return false
        }
    })
}

const resetForm = (formEl) => {
    if (!formEl) return
    formEl.resetFields()
}
</script>
  
```

```html
<el-form
      ref="ruleFormRef"
      :model="ruleForm"
      status-icon
      :rules="rules"
      label-width="120px"
      class="demo-ruleForm"
    >
```

 **ref="ruleFormRef"**    

这个其实就获取当前表单的dom对象，就类似于 document.getElementById(“formid”) ，但是在vue使用

```js
const ruleFormRef = ref(null) // document.getElementById(“formid”)
```

那么ruleFormRef对象有什么用处，- 表单提交，校验，重置

```js
var formDom = document.getElementById(“formid”) 
formDom.submit()
formDom.reset()
```

那么在VUE3的验证主要是用来ref=“ruleFormRef”是用来做如下事情的：

```js
# html
<el-form ref="loginFormRef" >
    
# js部分    
const loginFormRef = ref()

// 提交表单
const submitForm = () => {
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
```

  **:model="loginForm"**

注意目的 / 作用：给后面的 rules 验证使用的，而不是给下面的 v-model 使用的。

数据模型，其实就是表单定义的响应式数据，也是发送给服务端的数据。一般来说，提交都具有结构性和对象性，最好定义的时候都使用reactive定义

```js
// 定义表单的响应式数据，建议都使用reactive
const loginForm = reactive({
    pass: '',
    checkPass: '',
    age: '',
})

```

**rules**

具体的校验规则

```js
<el-form-item label="密码" prop="pass">
```

这里的prop=“pass”一定和rules和loginform的响应式数据某个属性一定要一致，而上面的prop=“pass”还有一层含义，就是如果出错了的信息，会显示在输入框下方



完整代码如下：

```js
<template>
    {{ loginForm }}
    <el-form ref="loginFormRef" :model="loginForm" status-icon :rules="rules" label-width="120px" class="demo-ruleForm">
        <el-form-item label="密码" prop="pass">
            <el-input v-model="loginForm.pass" type="password" autocomplete="off" />
        </el-form-item>
        <el-form-item label="确认密码" prop="checkPass">
            <el-input v-model="loginForm.checkPass" type="password" autocomplete="off" />
        </el-form-item>
        <el-form-item label="年龄" prop="age">
            <el-input v-model.number="loginForm.age" />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="submitForm">注册</el-button>
            <el-button @click="resetForm">Reset</el-button>
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

// 验证规则
const rules = reactive({
    pass:{ required: true, message: '请输入密码', trigger: 'blur' },
    checkPass:{ required: true, message: '请输入确认密码', trigger: 'blur' },
    age:{ required: true, message: '请输入年龄', trigger: 'blur' },
})


// 提交表单
const submitForm = () => {
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



## elementform验证真正验证规则—async-validator

https://github.com/yiminghe/async-validator

