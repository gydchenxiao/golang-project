<template>
    <div class="page admin-box" element-loading-text="正在加载中">
        <SkeletonTemplate>
            <relation-state></relation-state>
            <div class="gva-card-box">
                <div class="gva-card gva-top-card">
                    <div class="gva-top-card-left">
                        <div class="gva-top-card-left-title">早安，管理员，请开始一天的工作吧</div>
                        <div class="gva-top-card-left-dot">今日晴，0℃ - 10℃，天气寒冷，注意添加衣物。</div>
                        <div class="gva-top-card-left-rows">
                            <div class="el-row">
                                <div class="el-col el-col-8 el-col-xs-24 el-col-sm-8">
                                    <div class="flex-center"><i class="el-icon dashboard-icon"><svg
                                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024">
                                                <path fill="currentColor"
                                                    d="M384 96a32 32 0 0 1 64 0v786.752a32 32 0 0 1-54.592 22.656L95.936 608a32 32 0 0 1 0-45.312h.128a32 32 0 0 1 45.184 0L384 805.632V96zm192 45.248a32 32 0 0 1 54.592-22.592L928.064 416a32 32 0 0 1 0 45.312h-.128a32 32 0 0 1-45.184 0L640 218.496V928a32 32 0 1 1-64 0V141.248z">
                                                </path>
                                            </svg></i> 今日流量 (1231231) </div>
                                </div>
                                <div class="el-col el-col-8 el-col-xs-24 el-col-sm-8">
                                    <div class="flex-center"><i class="el-icon dashboard-icon"><svg
                                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024">
                                                <path fill="currentColor"
                                                    d="M628.736 528.896A416 416 0 0 1 928 928H96a415.872 415.872 0 0 1 299.264-399.104L512 704l116.736-175.104zM720 304a208 208 0 1 1-416 0 208 208 0 0 1 416 0z">
                                                </path>
                                            </svg></i> 总用户数 (24001) </div>
                                </div>
                                <div class="el-col el-col-8 el-col-xs-24 el-col-sm-8">
                                    <div class="flex-center"><i class="el-icon dashboard-icon"><svg
                                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024">
                                                <path fill="currentColor"
                                                    d="M736 504a56 56 0 1 1 0-112 56 56 0 0 1 0 112zm-224 0a56 56 0 1 1 0-112 56 56 0 0 1 0 112zm-224 0a56 56 0 1 1 0-112 56 56 0 0 1 0 112zM128 128v640h192v160l224-160h352V128H128z">
                                                </path>
                                            </svg></i> 好评率 (99%) </div>
                                </div>
                            </div>
                        </div>
                        <div>
                            <div class="gva-top-card-left-item"> 使用教学： <a target="view_window"
                                    href="https://www.bilibili.com/video/BV1Rg411u7xH/"
                                    style="color: rgb(64, 158, 255);">https://www.bilibili.com/video/BV1Rg411u7xH</a></div>
                            <div class="gva-top-card-left-item"> 插件仓库： <a target="view_window"
                                    href="https://plugin.gin-vue-admin.com/#/layout/home"
                                    style="color: rgb(64, 158, 255);">https://plugin.gin-vue-admin.com</a></div>
                        </div>
                    </div><img src="https://demo.gin-vue-admin.com/assets/dashboard-70e55b71.png" class="gva-top-card-right"
                        alt="">
                </div>
            </div>
            <div class="gva-card-box">
                <div class="el-card is-always-shadow gva-card quick-entrance">
                    <div class="el-card__header">
                        <div class="card-header"><span>快捷入口</span></div>
                    </div>
                    <div class="el-card__body" style="">
                        <div class="el-row" style="margin-left: -10px; margin-right: -10px;">
                            <div class="el-col el-col-4 el-col-xs-8 is-guttered quick-entrance-items"
                                style="padding-right: 10px; padding-left: 10px;">
                                <div class="quick-entrance-item">
                                    <div class="quick-entrance-item-icon"
                                        style="background-color: rgba(255, 156, 110, 0.3);"><i class="el-icon"><svg
                                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024"
                                                style="color: rgb(255, 156, 110);">
                                                <path fill="currentColor"
                                                    d="M544 768v128h192a32 32 0 1 1 0 64H288a32 32 0 1 1 0-64h192V768H192A128 128 0 0 1 64 640V256a128 128 0 0 1 128-128h640a128 128 0 0 1 128 128v384a128 128 0 0 1-128 128H544zM192 192a64 64 0 0 0-64 64v384a64 64 0 0 0 64 64h640a64 64 0 0 0 64-64V256a64 64 0 0 0-64-64H192z">
                                                </path>
                                            </svg></i></div>
                                    <p>用户管理</p>
                                </div>
                            </div>
                            <div class="el-col el-col-4 el-col-xs-8 is-guttered quick-entrance-items"
                                style="padding-right: 10px; padding-left: 10px;">
                                <div class="quick-entrance-item">
                                    <div class="quick-entrance-item-icon"
                                        style="background-color: rgba(105, 192, 255, 0.3);"><i class="el-icon"><svg
                                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024"
                                                style="color: rgb(105, 192, 255);">
                                                <path fill="currentColor"
                                                    d="M600.704 64a32 32 0 0 1 30.464 22.208l35.2 109.376c14.784 7.232 28.928 15.36 42.432 24.512l112.384-24.192a32 32 0 0 1 34.432 15.36L944.32 364.8a32 32 0 0 1-4.032 37.504l-77.12 85.12a357.12 357.12 0 0 1 0 49.024l77.12 85.248a32 32 0 0 1 4.032 37.504l-88.704 153.6a32 32 0 0 1-34.432 15.296L708.8 803.904c-13.44 9.088-27.648 17.28-42.368 24.512l-35.264 109.376A32 32 0 0 1 600.704 960H423.296a32 32 0 0 1-30.464-22.208L357.696 828.48a351.616 351.616 0 0 1-42.56-24.64l-112.32 24.256a32 32 0 0 1-34.432-15.36L79.68 659.2a32 32 0 0 1 4.032-37.504l77.12-85.248a357.12 357.12 0 0 1 0-48.896l-77.12-85.248A32 32 0 0 1 79.68 364.8l88.704-153.6a32 32 0 0 1 34.432-15.296l112.32 24.256c13.568-9.152 27.776-17.408 42.56-24.64l35.2-109.312A32 32 0 0 1 423.232 64H600.64zm-23.424 64H446.72l-36.352 113.088-24.512 11.968a294.113 294.113 0 0 0-34.816 20.096l-22.656 15.36-116.224-25.088-65.28 113.152 79.68 88.192-1.92 27.136a293.12 293.12 0 0 0 0 40.192l1.92 27.136-79.808 88.192 65.344 113.152 116.224-25.024 22.656 15.296a294.113 294.113 0 0 0 34.816 20.096l24.512 11.968L446.72 896h130.688l36.48-113.152 24.448-11.904a288.282 288.282 0 0 0 34.752-20.096l22.592-15.296 116.288 25.024 65.28-113.152-79.744-88.192 1.92-27.136a293.12 293.12 0 0 0 0-40.256l-1.92-27.136 79.808-88.128-65.344-113.152-116.288 24.96-22.592-15.232a287.616 287.616 0 0 0-34.752-20.096l-24.448-11.904L577.344 128zM512 320a192 192 0 1 1 0 384 192 192 0 0 1 0-384zm0 64a128 128 0 1 0 0 256 128 128 0 0 0 0-256z">
                                                </path>
                                            </svg></i></div>
                                    <p>角色管理</p>
                                </div>
                            </div>
                            <div class="el-col el-col-4 el-col-xs-8 is-guttered quick-entrance-items"
                                style="padding-right: 10px; padding-left: 10px;">
                                <div class="quick-entrance-item">
                                    <div class="quick-entrance-item-icon"
                                        style="background-color: rgba(179, 127, 235, 0.3);"><i class="el-icon"><svg
                                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024"
                                                style="color: rgb(179, 127, 235);">
                                                <path fill="currentColor"
                                                    d="M160 448a32 32 0 0 1-32-32V160.064a32 32 0 0 1 32-32h256a32 32 0 0 1 32 32V416a32 32 0 0 1-32 32H160zm448 0a32 32 0 0 1-32-32V160.064a32 32 0 0 1 32-32h255.936a32 32 0 0 1 32 32V416a32 32 0 0 1-32 32H608zM160 896a32 32 0 0 1-32-32V608a32 32 0 0 1 32-32h256a32 32 0 0 1 32 32v256a32 32 0 0 1-32 32H160zm448 0a32 32 0 0 1-32-32V608a32 32 0 0 1 32-32h255.936a32 32 0 0 1 32 32v256a32 32 0 0 1-32 32H608z">
                                                </path>
                                            </svg></i></div>
                                    <p>菜单管理</p>
                                </div>
                            </div>
                            <div class="el-col el-col-4 el-col-xs-8 is-guttered quick-entrance-items"
                                style="padding-right: 10px; padding-left: 10px;">
                                <div class="quick-entrance-item">
                                    <div class="quick-entrance-item-icon"
                                        style="background-color: rgba(255, 214, 102, 0.3);"><i class="el-icon"><svg
                                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024"
                                                style="color: rgb(255, 214, 102);">
                                                <path fill="currentColor"
                                                    d="M320 256a64 64 0 0 0-64 64v384a64 64 0 0 0 64 64h384a64 64 0 0 0 64-64V320a64 64 0 0 0-64-64H320zm0-64h384a128 128 0 0 1 128 128v384a128 128 0 0 1-128 128H320a128 128 0 0 1-128-128V320a128 128 0 0 1 128-128z">
                                                </path>
                                                <path fill="currentColor"
                                                    d="M512 64a32 32 0 0 1 32 32v128h-64V96a32 32 0 0 1 32-32zm160 0a32 32 0 0 1 32 32v128h-64V96a32 32 0 0 1 32-32zm-320 0a32 32 0 0 1 32 32v128h-64V96a32 32 0 0 1 32-32zm160 896a32 32 0 0 1-32-32V800h64v128a32 32 0 0 1-32 32zm160 0a32 32 0 0 1-32-32V800h64v128a32 32 0 0 1-32 32zm-320 0a32 32 0 0 1-32-32V800h64v128a32 32 0 0 1-32 32zM64 512a32 32 0 0 1 32-32h128v64H96a32 32 0 0 1-32-32zm0-160a32 32 0 0 1 32-32h128v64H96a32 32 0 0 1-32-32zm0 320a32 32 0 0 1 32-32h128v64H96a32 32 0 0 1-32-32zm896-160a32 32 0 0 1-32 32H800v-64h128a32 32 0 0 1 32 32zm0-160a32 32 0 0 1-32 32H800v-64h128a32 32 0 0 1 32 32zm0 320a32 32 0 0 1-32 32H800v-64h128a32 32 0 0 1 32 32z">
                                                </path>
                                            </svg></i></div>
                                    <p>代码生成器</p>
                                </div>
                            </div>
                            <div class="el-col el-col-4 el-col-xs-8 is-guttered quick-entrance-items"
                                style="padding-right: 10px; padding-left: 10px;">
                                <div class="quick-entrance-item">
                                    <div class="quick-entrance-item-icon"
                                        style="background-color: rgba(255, 133, 192, 0.3);"><i class="el-icon"><svg
                                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024"
                                                style="color: rgb(255, 133, 192);">
                                                <path fill="currentColor"
                                                    d="M805.504 320 640 154.496V320h165.504zM832 384H576V128H192v768h640V384zM160 64h480l256 256v608a32 32 0 0 1-32 32H160a32 32 0 0 1-32-32V96a32 32 0 0 1 32-32zm318.4 582.144 180.992-180.992L704.64 510.4 478.4 736.64 320 578.304l45.248-45.312L478.4 646.144z">
                                                </path>
                                            </svg></i></div>
                                    <p>表单生成器</p>
                                </div>
                            </div>
                            <div class="el-col el-col-4 el-col-xs-8 is-guttered quick-entrance-items"
                                style="padding-right: 10px; padding-left: 10px;">
                                <div class="quick-entrance-item">
                                    <div class="quick-entrance-item-icon"
                                        style="background-color: rgba(92, 219, 211, 0.3);"><i class="el-icon"><svg
                                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024"
                                                style="color: rgb(92, 219, 211);">
                                                <path fill="currentColor"
                                                    d="M512 512a192 192 0 1 0 0-384 192 192 0 0 0 0 384zm0 64a256 256 0 1 1 0-512 256 256 0 0 1 0 512zm320 320v-96a96 96 0 0 0-96-96H288a96 96 0 0 0-96 96v96a32 32 0 1 1-64 0v-96a160 160 0 0 1 160-160h448a160 160 0 0 1 160 160v96a32 32 0 1 1-64 0z">
                                                </path>
                                            </svg></i></div>
                                    <p>关于我们</p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <div class="gva-card-box">
                <div class="gva-card">
                    <div class="card-header"><span>数据统计</span></div>
                    <div class="echart-box">
                        <div class="el-row" style="margin-left: -10px; margin-right: -10px;">
                            <div class="el-col el-col-24 el-col-xs-24 el-col-sm-18 is-guttered">
                                <div class="dashboard-line-box">
                                    <order-stat-charts></order-stat-charts>
                                </div>
                            </div>
                            <div class="el-col el-col-24 el-col-xs-24 el-col-sm-6 is-guttered"
                                style="padding-right: 10px; padding-left: 10px;">
                                <div class="commit-table">
                                    <div class="commit-table-title"> 更新日志 </div>
                                    <div class="log">
                                        <div class="log-item">
                                            <div class="flex-1 flex key-box"><span data-v-144ac47f=""
                                                    class="key top">1</span>
                                            </div>
                                            <div class="flex-5 flex message">feat: update ci add
                                                CGO_ENABLED=0 (#1497)</div>
                                            <div class="flex-3 flex form">2023-08-04</div>
                                        </div>
                                        <div class="log-item">
                                            <div class="flex-1 flex key-box"><span data-v-144ac47f=""
                                                    class="key top">2</span>
                                            </div>
                                            <div class="flex-5 flex message">集成tailwindcss 替换新的登陆页面和版权信息组件
                                                (#1499)

                                                * 集成tailwindcss

                                                * fix: 修改登录，init 页面

                                                * Update package.json

                                                * 升级可升级的第三方库到新版本

                                                * 细节调整

                                                ---------

                                                Co-authored-by: bypanghu &lt;bypanghu@163.com&gt;
                                                Co-authored-by: task &lt;121913992@qq.com&gt;</div>
                                            <div class="flex-3 flex form">2023-08-04</div>
                                        </div>
                                        <div class="log-item">
                                            <div class="flex-1 flex key-box"><span data-v-144ac47f=""
                                                    class="key top">3</span>
                                            </div>
                                            <div class="flex-5 flex message">fix: macos和linux下代码定位命令不对 #1495
                                                (#1496)</div>
                                            <div class="flex-3 flex form">2023-08-02</div>
                                        </div>
                                        <div class="log-item">
                                            <div class="flex-1 flex key-box"><span data-v-144ac47f="" class="key">4</span>
                                            </div>
                                            <div class="flex-5 flex message">Rename outer.go to other.go
                                                (#1490)</div>
                                            <div class="flex-3 flex form">2023-08-02</div>
                                        </div>
                                        <div class="log-item">
                                            <div class="flex-1 flex key-box"><span data-v-144ac47f="" class="key">5</span>
                                            </div>
                                            <div class="flex-5 flex message">feat: use description to name
                                                apiGroup (#1494)</div>
                                            <div class="flex-3 flex form">2023-08-01</div>
                                        </div>
                                        <div class="log-item">
                                            <div class="flex-1 flex key-box"><span data-v-144ac47f="" class="key">6</span>
                                            </div>
                                            <div class="flex-5 flex message">Merge pull request #1483 from
                                                ChengDaqi2023/oscs_fix_cis9at8au51vj78hfkt0

                                                fix(sec): upgrade golang.org/x/image to 0.5.0</div>
                                            <div class="flex-3 flex form">2023-07-21</div>
                                        </div>
                                        <div class="log-item">
                                            <div class="flex-1 flex key-box"><span data-v-144ac47f="" class="key">7</span>
                                            </div>
                                            <div class="flex-5 flex message">update golang.org/x/image
                                                v0.0.0-20210220032944-ac19c3e999fb to 0.5.0</div>
                                            <div class="flex-3 flex form">2023-07-20</div>
                                        </div>
                                        <div class="log-item">
                                            <div class="flex-1 flex key-box"><span data-v-144ac47f="" class="key">8</span>
                                            </div>
                                            <div class="flex-5 flex message">修复marked问题</div>
                                            <div class="flex-3 flex form">2023-07-19</div>
                                        </div>
                                        <div class="log-item">
                                            <div class="flex-1 flex key-box"><span data-v-144ac47f="" class="key">9</span>
                                            </div>
                                            <div class="flex-5 flex message">增加多图片自动生成功能</div>
                                            <div class="flex-3 flex form">2023-07-18</div>
                                        </div>
                                        <div class="log-item">
                                            <div class="flex-1 flex key-box"><span data-v-144ac47f="" class="key">10</span>
                                            </div>
                                            <div class="flex-5 flex message">修复element2.3.8以上i18n的问题</div>
                                            <div class="flex-3 flex form">2023-07-17</div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </SkeletonTemplate>
    </div>
</template>

<script setup>
import OrderStatCharts from './compoments/OrderStatCharts.vue'; // 百度图库插件；页面的样式在src\assets\base.css中有设置的哦
import RelationState from './compoments/RelationState.vue'; // 动态数字
import SkeletonTemplate from './compoments/SkeletonTemplate.vue'; // 组件
// import { LoadData } from '@/api/sysmenus.js'

// axios 异步请求后端，获取当前用户的菜单数据
// const handleLoadMenusData = async () => {
//     const menuDatas = await LoadData()
//     console.log('dashboard.vue: menuDatas', menuDatas)
// }

// onMounted(() => {
//     handleLoadMenusData()
// })
</script> 

<style lang="scss" scoped>
.page {
    background: #f0f2f5;
}

.page .gva-card-box {
    padding: 12px 16px
}

.page .gva-card-box+.gva-card-box {
    padding-top: 0
}

.page .gva-card {
    box-sizing: border-box;
    background-color: #fff;
    border-radius: 2px;
    height: auto;
    padding: 26px 30px;
    overflow: hidden;
    box-shadow: 0 0 7px 1px rgba(0, 0, 0, .03)
}

.page .gva-top-card {
    height: 260px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    color: #777
}

.page .gva-top-card-left {
    height: 100%;
    display: flex;
    flex-direction: column
}

.page .gva-top-card-left-title {
    font-size: 22px;
    color: #343844
}

.page .gva-top-card-left-dot {
    font-size: 16px;
    color: #6b7687;
    margin-top: 24px
}

.page .gva-top-card-left-rows {
    margin-top: 18px;
    color: #6b7687;
    width: 600px;
    align-items: center
}

.page .gva-top-card-left-item {
    margin-top: 14px
}

.page .gva-top-card-left-item+.gva-top-card-left-item {
    margin-top: 24px
}

.page .gva-top-card-right {
    height: 600px;
    width: 600px;
    margin-top: 28px
}

.page .el-card__header {
    padding: 0;
    border-bottom: none
}

.page .card-header {
    padding-bottom: 20px;
    border-bottom: 1px solid #e8e8e8
}

.page .quick-entrance-title {
    height: 30px;
    font-size: 22px;
    color: #333;
    width: 100%;
    border-bottom: 1px solid #eee
}

.page .quick-entrance-items {
    display: flex;
    align-items: center;
    justify-content: center;
    text-align: center;
    color: #333
}

.page .quick-entrance-items .quick-entrance-item {
    padding: 16px 28px;
    margin-top: -16px;
    margin-bottom: -16px;
    border-radius: 4px;
    transition: all .2s;
    cursor: pointer;
    height: auto;
    text-align: center
}

.page .quick-entrance-items .quick-entrance-item:hover {
    box-shadow: 0 0 7px rgba(217, 217, 217, .55)
}

.page .quick-entrance-items .quick-entrance-item-icon {
    width: 50px;
    height: 50px !important;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    margin: 0 auto
}

.page .quick-entrance-items .quick-entrance-item-icon i {
    font-size: 24px
}

.page .quick-entrance-items .quick-entrance-item p {
    margin-top: 10px
}

.page .echart-box {
    padding: 14px
}

.dashboard-icon {
    font-size: 20px;
    color: #55a0f8;
    width: 30px;
    height: 30px;
    margin-right: 10px;
    display: flex;
    align-items: center
}

.flex-center {
    display: flex;
    align-items: center
}

@media (max-width: 750px) {
    .gva-card {
        padding: 20px 10px !important
    }

    .gva-card .gva-top-card {
        height: auto
    }

    .gva-card .gva-top-card-left-title {
        font-size: 20px !important
    }

    .gva-card .gva-top-card-left-rows {
        margin-top: 15px;
        align-items: center
    }

    .gva-card .gva-top-card-right {
        display: none
    }

    .gva-card .gva-middle-card-item {
        line-height: 20px
    }

    .gva-card .dashboard-icon {
        font-size: 18px
    }
}

.commit-table {
    background-color: #fff;
    height: 400px
}

.commit-table-title {
    font-weight: 600;
    margin-bottom: 12px
}

.commit-table .log-item {
    display: flex;
    justify-content: space-between;
    margin-top: 14px
}

.commit-table .log-item .key-box {
    justify-content: center
}

.commit-table .log-item .key {
    display: inline-flex;
    justify-content: center;
    align-items: center;
    width: 20px;
    height: 20px;
    border-radius: 50%;
    background: #F0F2F5;
    text-align: center;
    color: rgba(0, 0, 0, .65)
}

.commit-table .log-item .key.top {
    background: #314659;
    color: #fff
}

.commit-table .log-item .message {
    color: rgba(0, 0, 0, .65)
}

.commit-table .log-item .form {
    color: rgba(0, 0, 0, .65);
    margin-left: 12px
}

.commit-table .log-item .flex {
    line-height: 20px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap
}

.commit-table .log-item .flex-1 {
    flex: 1
}

.commit-table .log-item .flex-2 {
    flex: 2
}

.commit-table .log-item .flex-3 {
    flex: 3
}

.commit-table .log-item .flex-4 {
    flex: 4
}

.commit-table .log-item .flex-5 {
    flex: 5
}
</style>