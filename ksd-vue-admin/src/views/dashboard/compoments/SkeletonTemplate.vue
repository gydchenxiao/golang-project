<template>
  <el-skeleton animated :loading="skLoading">
    <template #template>
      <div class="gva-card-box">
        <el-row :gutter="12">
          <el-col :span="6" :xs="12" v-for="i in 4" :key="i">
            <el-card shadow="hover">
              <el-skeleton-item variant="div" style="height: 20px" />
              <div style="margin-top: 15px;"><el-skeleton-item variant="h3" style="width: 60%;height: 30px" />
              </div>
              <div style="margin-top: 10px;"><el-skeleton-item variant="div" style="width: 40%; height: 20px" />
              </div>
            </el-card>
          </el-col>
        </el-row>
      </div>
      <div class="gva-card-box">
        <div class="gva-card gva-top-card" style="padding:2px;">
          <div style="display: flex;flex-direction: column;padding:20px;width: 100%;">
            <el-skeleton-item variant="div" style="width: 40%; height: 20px;" />
            <el-skeleton-item variant="div" style="width: 50%; height: 20px;margin-top: 10px;" />
            <el-skeleton-item variant="div" style="width: 60%; height: 20px;margin-top: 10px;" />
            <el-skeleton-item variant="div" style="width: 70%; height: 20px;margin-top: 10px;" />
          </div>
          <el-skeleton-item variant="image" style="width: 640px; height: 240px" />
        </div>
      </div>
      <div class="gva-card-box">
        <div class="gva-card gva-top-card">
          <el-skeleton-item variant="div" v-for="i in 6" :key="i" style="width: 240px; height: 70px" />
        </div>
      </div>
      <div class="gva-card-box">
        <div class="gva-card">
          <div class="card-header"><span>
              <el-skeleton-item variant="div" style="width: 140px; height: 50px" /></span>
          </div>
          <div class="echart-box">
            <div class="el-row" style="margin-left: -10px; margin-right: -10px;">
              <div class="el-col el-col-24 el-col-xs-24 el-col-sm-18 is-guttered">
                <div class="dashboard-line-box">
                  <el-skeleton-item variant="div" style="width: 100%; height: 450px" />
                </div>
              </div>
              <div class="el-col el-col-24 el-col-xs-24 el-col-sm-6 is-guttered" style="padding:20px;">
                <div class="commit-table">
                  <div class="commit-table-title"><el-skeleton-item variant="div" style="width: 140px; height: 50px" />
                  </div>
                  <div class="log">
                    <div class="log-item" v-for="i in 11" :key="i">
                      <el-skeleton-item variant="div" style="width: 100%; height: 20px" />
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
      <!--  -->
      <slot></slot>
    </template>
  </el-skeleton>
</template>

<script setup>
import { useSkeletonStore } from '@/stores/skeleton.js'
const skeletonStore = useSkeletonStore()
const skLoading = computed(() => skeletonStore.skLoading)


</script>