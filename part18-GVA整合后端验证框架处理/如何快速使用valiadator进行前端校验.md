# 如何快速使用valiadator进行前端校验



## 01、常用的验证

- element-plus + form + rules (async-validator)
- 自己写form +  rules (async-validator)
- 自己写form +  (validator) (采用这种)



## 02、使用定义form表单+validator.js

1: 定义一个路由addcourse.vue

```vue
<template>

    <form action="" method="post">
        {{ formData }}
        <p>标题：<input type="text" v-model="formData.username"></p>
        <p>描述：<input type="text" v-model="formData.password"></p>
        <p><button @click.prevent="handleAddCourse">添加课程</button></p>
    </form>

</template>

<script setup>
import {ref} from 'vue'
import Ksd from '@/utils/index.js'
// 定义一个数据模式
const formData = ref({
    username:"",
    password:""
})

// 添加课程的方法
const handleAddCourse = ()=>{
    
    // 这里写方法
    if(!formData.value.username){
        Ksd.message("请输入用户名!!!");
        return;
    }

    if(!formData.value.password){
        Ksd.message("请输入密码!!!");
        return;
    }

    alert("success")


}

</script>

<style scoped lang="scss">

</style>
```

上面的代码，其实就可以完成一个最基本的空校验。但是在开发中往往校验会非常的长，所以我们很多会喜欢吧校验和业务分开。就如下的代码：



### 方法的封装

```vue
<template>

    <form action="" method="post">
        {{ formData }}
        <p>标题：<input type="text" v-model="formData.username"></p>
        <p>描述：<input type="text" v-model="formData.password"></p>
        <p><button @click.prevent="handleAddCourse">添加课程</button></p>
    </form>

</template>

<script setup>
import {ref} from 'vue'
import Ksd from '@/utils/index.js'
// 定义一个数据模式
const formData = ref({
    username:"",
    password:""
})


const validator = {
    validate(){
        // 这里写方法
        if(!formData.value.username){
            Ksd.message("请输入用户名!!!");
            return false;
        }

        if(!formData.value.password){
            Ksd.message("请输入密码!!!");
            return false;
        }
        
        return true;
    }
}

// 添加课程的方法
const handleAddCourse = ()=>{
    var flag = validator.validate();
    if(flag){
        alert("success")
    }
}

</script>

<style scoped lang="scss">

</style>
```



### 回调封装

```vue
<template>

    <form action="" method="post">
        {{ formData }}
        <p>标题：<input type="text" v-model="formData.username"></p>
        <p>描述：<input type="text" v-model="formData.password"></p>
        <p><button @click.prevent="handleAddCourse">添加课程</button></p>
    </form>

</template>

<script setup>
import {ref} from 'vue'
import Ksd from '@/utils/index.js'
// 定义一个数据模式
const formData = ref({
    username:"",
    password:""
})


const validator = {
     validate(fn){
        // 这里写方法
        if(!formData.value.username){
            Ksd.message("请输入用户名!!!");
            fn && fn(false); 
            return;
        }

        if(!formData.value.password){
            Ksd.message("请输入密码!!!");
            fn && fn(false);
            return;
        }

        fn && fn(true);
    }
}

// 添加课程的方法
const handleAddCourse = async ()=>{
    validator.validate(function fn(flag){
        if(flag){
            alert("success")
        }else{
            
        }
    })
}




// const validator = {
//     async validate(){
//         // 这里写方法
//         if(!formData.value.username){
//             Ksd.message("请输入用户名!!!");
//             return Promise.reject(false);
//         }

//         if(!formData.value.password){
//             Ksd.message("请输入密码!!!");
//             return Promise.reject(false);
//         }

//         return Promise.resolve(true);
//     }
// }

// // 添加课程的方法
// const handleAddCourse = async ()=>{
//     var valid = await validator.validate();
//     if(valid){
//         alert("success")
//     }
// }

</script>

<style scoped lang="scss">

</style>
```



### 承诺封装

```vue
<template>

    <form action="" method="post">
        {{ formData }}
        <p>标题：<input type="text" v-model="formData.username"></p>
        <p>描述：<input type="text" v-model="formData.password"></p>
        <p><button @click.prevent="handleAddCourse">添加课程</button></p>
    </form>

</template>

<script setup>
import {ref} from 'vue'
import Ksd from '@/utils/index.js'
// 定义一个数据模式
const formData = ref({
    username:"",
    password:""
})


const validator = {
     async validate(){
        // 这里写方法
        if(!formData.value.username){
            Ksd.message("请输入用户名!!!");
            return Promise.reject({field:"username",msg:"请输入用户名"});
        }

        if(!formData.value.password){
            Ksd.message("请输入密码!!!");
            return Promise.reject({field:"password",msg:"请输入密码"});
        }

        return Promise.resolve(true);
    }
}

// 添加课程的方法 es6 
const handleAddCourse = ()=>{
    validator.validate().then(flag=>{
        alert("success")
    }).catch(e=>{
        alert(e)
    })

    validator.validate().then(flag=>{
        alert("success")
    },e => {
        alert(e)
    })
}

// es7 async +await
const handleAddCourse2 = async ()=>{
    try{
        const flag = await validator.validate();
        if(flag){
            alert("success")
        }
    }catch(e){
        alert(e)
    }
}




// const validator = {
//     async validate(){
//         // 这里写方法
//         if(!formData.value.username){
//             Ksd.message("请输入用户名!!!");
//             return Promise.reject(false);
//         }

//         if(!formData.value.password){
//             Ksd.message("请输入密码!!!");
//             return Promise.reject(false);
//         }

//         return Promise.resolve(true);
//     }
// }

// // 添加课程的方法
// const handleAddCourse = async ()=>{
//     var valid = await validator.validate();
//     if(valid){
//         alert("success")
//     }
// }

</script>

<style scoped lang="scss">

</style>
```

## 验证器封装和收集

```vue
<template>

    <form action="" method="post">
        {{ formData }}
        <p>标题：<input type="text" v-model.trim="formData.title"></p>
        <p>描述：<input type="text" v-model.trim="formData.desc"></p>
        <p>金额：<input type="text" v-model.trim="formData.money"></p>
        <p>邮箱：<input type="text" v-model.trim="formData.email"></p>
        <p><button @click.prevent="handleAddCourse">添加课程</button></p>
    </form>

</template>

<script setup>
import {ref} from 'vue'
import Ksd from '@/utils/index.js'
// 定义一个数据模式
const formData = ref({
    title:"",
    desc:"",
    money:0,
    email:""
})

// const handleParseFloat = ()=>{
//     formData.value.money = parseFloat(formData.value.money) || 0;
// }

const validator = {
     async validate(){
        // 这里写方法
        if(!formData.value.title){
            Ksd.message("请输入标题!!!");
            return Promise.reject({field:"title",msg:"请输入标题"});
        }

        if(!formData.value.desc){
            Ksd.message("请输入描述!!!");
            return Promise.reject({field:"desc",msg:"请输入描述"});
        }

        if(!formData.value.money){
            Ksd.message("请输入金额!!!");
            return Promise.reject({field:"money",msg:"请输入金额"});
        }

        if(isNaN(formData.value.money)){
            Ksd.message("金额必须是数字!!!");
            return Promise.reject({field:"money",msg:"金额必须是数字"});
        }

        if(!formData.value.email){
            Ksd.message("请输入邮箱!!!");
            return Promise.reject({field:"email",msg:"请输入邮箱"});
        }

        // 邮箱的合法性----正则
        // 邮箱的合法性----正则 test 返回只是 true/false .match返回是数组，用于快速把使用正则表达式去把某个字符串中想要的部分快速抓取出来。
        // aaabbbccc ---/xxxx/.match("aaabbbccc")===["aaa","bbb","cccc"]
        if(!/^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/.test(formData.value.email)){
            Ksd.message("邮箱不合法!!!");
            return Promise.reject({field:"email",msg:"邮箱不合法"});
        }

        return Promise.resolve(true);
    }
}

// es7 async +await
const handleAddCourse = async ()=>{
    try{
        const flag = await validator.validate();
        if(flag){
            alert("success")
        }
    }catch(e){
        alert(JSON.stringify(e))
    }
}




// const validator = {
//     async validate(){
//         // 这里写方法
//         if(!formData.value.username){
//             Ksd.message("请输入用户名!!!");
//             return Promise.reject(false);
//         }

//         if(!formData.value.password){
//             Ksd.message("请输入密码!!!");
//             return Promise.reject(false);
//         }

//         return Promise.resolve(true);
//     }
// }

// // 添加课程的方法
// const handleAddCourse = async ()=>{
//     var valid = await validator.validate();
//     if(valid){
//         alert("success")
//     }
// }

</script>

<style scoped lang="scss">

</style>
```

## 自己封装的校验器

```js

const validator = {
    // 邮箱的合法性----正则
    // 邮箱的合法性----正则 test 返回只是 true/false .match返回是数组，用于快速把使用正则表达式去把某个字符串中想要的部分快速抓取出来。
    // aaabbbccc ---/xxxx/.match("aaabbbccc")===["aaa","bbb","cccc"]
    isEmail(value){
        return /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/.test(value)
    },
    isNumber(val){
        return /^\d*$/.test(val)
    }
}

export default validator
```

## 第三方的js校验器

https://github.com/validatorjs/validator.js

1: 下载和安装

```js
npm install validator
```

2: 使用

```js
import validator from 'validator';
```

或者直接导入你想要验证器

```js
import { isEmail ,isNumber} from 'validator/lib/isEmail';
```

3: 用法

```js
validator.isEmail(val)
validator.isNumber(val)
```

或者

```js
isEmail(val)
isNumber(val)
```

具体使用

```vue
<template>

    <form action="" method="post">
        {{ formData }}
        <p>标题：<input type="text" v-model.trim="formData.title"></p>
        <p>描述：<input type="text" v-model.trim="formData.desc"></p>
        <p>金额：<input type="text" v-model.trim="formData.money"></p>
        <p>邮箱：<input type="text" v-model.trim="formData.email"></p>
        <p><button @click.prevent="handleAddCourse">添加课程</button></p>
    </form>

</template>

<script setup>
import {ref} from 'vue'
import Ksd from '@/utils/index.js'
import Validator from 'validator'
// 定义一个数据模式
const formData = ref({
    title:"",
    desc:"",
    money:0,
    email:""
})

const validator = {
     async validate(){
        // 这里写方法
        if(Validator.isEmpty(formData.value.title)){
            Ksd.message("请输入标题!!!");
            return Promise.reject({field:"title",msg:"请输入标题"});
        }

        if(Validator.isEmpty(formData.value.desc)){
            Ksd.message("请输入描述!!!");
            return Promise.reject({field:"desc",msg:"请输入描述"});
        }

        if(Validator.isEmpty(formData.value.money)){
            Ksd.message("请输入金额!!!");
            return Promise.reject({field:"money",msg:"请输入金额"});
        }

        if(!Validator.isNumber(formData.value.money)){
            Ksd.message("金额必须是数字!!!");
            return Promise.reject({field:"money",msg:"金额必须是数字"});
        }

        if(Validator.isEmpty(formData.value.email)){     
            Ksd.message("请输入邮箱!!!");
            return Promise.reject({field:"email",msg:"请输入邮箱"});
        }

        if(!Validator.isEmail(formData.value.email)){
            Ksd.message("邮箱不合法!!!");
            return Promise.reject({field:"email",msg:"邮箱不合法"});
        }
        
        return Promise.resolve(true);
    }
}

// es7 async +await
const handleAddCourse = async ()=>{
    try{
        const flag = await validator.validate();
        if(flag){
            alert("success")
        }
    }catch(e){
        alert(JSON.stringify(e))
    }
}




// const validator = {
//     async validate(){
//         // 这里写方法
//         if(!formData.value.username){
//             Ksd.message("请输入用户名!!!");
//             return Promise.reject(false);
//         }

//         if(!formData.value.password){
//             Ksd.message("请输入密码!!!");
//             return Promise.reject(false);
//         }

//         return Promise.resolve(true);
//     }
// }

// // 添加课程的方法
// const handleAddCourse = async ()=>{
//     var valid = await validator.validate();
//     if(valid){
//         alert("success")
//     }
// }

</script>

<style scoped lang="scss">

</style>
```









