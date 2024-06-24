<template>
    <div class="gva-card-box statebox">
        <el-row :gutter="12">
            <el-col :span="6" :xs="12">
                <el-card shadow="hover">
                    <div class="tit">总销售额</div>
                    <div class="num">
                        <AnimatedNumber :from="0" :to="12560" duration="3000">
                            <template #default="{ option, item }">
                                <span>￥{{ item.number }}</span>
                            </template>
                        </AnimatedNumber>
                    </div>
                    <div class="info">
                        <AnimatedNumber :from="0" :to="8869" duration="3000">
                            <template #default="{ option, item }">
                                <span>日销售额：￥{{ item.number }}</span>
                            </template>
                        </AnimatedNumber>
                    </div>
                </el-card>
            </el-col>
            <el-col :span="6" :xs="12">
                <el-card shadow="hover">
                    <div class="tit">用户注册</div>
                    <div class="num">
                        <AnimatedNumber :from="0" :to="8846" duration="3000" />
                    </div>
                    <div class="info">
                        <AnimatedNumber :from="0" :to="1423" duration="3000">
                            <template #default="{ option, item }">
                                <span>日注册量：{{ item.number }}</span>
                            </template>
                        </AnimatedNumber>
                    </div>
                </el-card>
            </el-col>
            <el-col :span="6" :xs="12">
                <el-card shadow="hover">
                    <div class="tit">访问量</div>
                    <div class="num">
                        <AnimatedNumber :from="0" :to="6560" duration="3000" />
                    </div>
                    <div class="info">
                        <AnimatedNumber :from="0" :to="2423" duration="3000">
                            <template #default="{ option, item }">
                                <span>日访问量：{{ item.number }}</span>
                            </template>
                        </AnimatedNumber>
                    </div>
                </el-card>
            </el-col>
            <el-col :span="6" :xs="12">
                <el-card shadow="hover">
                    <div class="tit">支付笔数</div>
                    <div class="num">
                        <AnimatedNumber :from="0" :to="12589" duration="3000" />
                    </div>
                    <div class="info">
                        <AnimatedNumber :from="0" :to="80" duration="3000">
                            <template #default="{ option, item }">
                                <span>转化率：{{ item.number }} %</span>
                            </template>
                        </AnimatedNumber>
                    </div>
                </el-card>
            </el-col>
        </el-row>
    </div>
</template>
<script setup>
</script>
<style lang="scss" scoped>
.statebox {
    .el-card {
        padding: 10px;
        margin-bottom: 5px;
    }

    .tit {
        height: 22px;
        color: rgba(0, 0, 0, .45);
        font-size: 14px;
        line-height: 22px;
    }

    .num {
        height: 38px;
        margin-top: 10px;
        color: rgba(0, 0, 0, .75);
        font-size: 32px;
        line-height: 38px;
        white-space: nowrap;
        font-weight: bold;
        text-overflow: ellipsis;
        word-break: break-all;
    }

    .info {
        position: relative;
        width: 100%;
        margin-top: 20px;
    }
}</style>