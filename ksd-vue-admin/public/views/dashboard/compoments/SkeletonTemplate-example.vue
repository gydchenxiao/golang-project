<template>
    <el-skeleton animated :loading="true">
        <template #template>
            <!--card的骨架屏-->
            <div class="gva-card-box statebox">
                <el-row :gutter="12">
                    <el-col :span="6" :xs="12" v-for="i in 4" :key="i">
                        <el-card shadow="hover">
                            <div class="tit"><el-skeleton-item variant="div" style="width: 40%;height: 20px" /></div>
                            <div class="num">
                                <el-skeleton-item variant="div" style="width: 60%;height:30px" />
                            </div>
                            <div class="info">
                                <el-skeleton-item variant="div" style="width: 80%;height:20px" />
                            </div>
                        </el-card>
                    </el-col>
                </el-row>
            </div>

            <!--统计报表的骨架-->
            <div class="gva-card-box">
                <div class="gva-card">
                    <div class="card-header"><span><el-skeleton-item variant="h3" style="width:6%;height:30px;" /></span>
                    </div>
                    <div class="echart-box">
                        <div class="el-row" style="margin-left: -10px; margin-right: -10px;">
                            <div class="el-col el-col-24 el-col-xs-24 el-col-sm-18 is-guttered">
                                <div class="dashboard-line-box">
                                    <el-skeleton-item variant="h3" style="height:400px;" />
                                </div>
                            </div>
                            <div class="el-col el-col-24 el-col-xs-24 el-col-sm-6 is-guttered"
                                style="padding-right: 10px; padding-left: 10px;">
                                <div class="commit-table">
                                    <div class="commit-table-title"><el-skeleton-item variant="h3"
                                            style="width:20%;height:30px;" /> </div>
                                    <div class="log">
                                        <div class="log-item" v-for="i in 10" :key="i">
                                            <el-skeleton-item variant="h3" style="height:22px;" />
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </template>
        <template #default>
            <slot></slot>
        </template>
    </el-skeleton>
</template>
<script setup>
import { useMenuTabStore } from '@/stores/menuTab.js'
const menuTabStore = useMenuTabStore()
const skLoading = computed(() => menuTabStore.skLoading)

onMounted(() => {
    setTimeout(() => {
        menuTabStore.skLoading = false;
    }, 1000)
})
</script>