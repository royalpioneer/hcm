<script setup lang="ts">
import type {
  FilterType,
} from '@/typings/resource';
import {
  PropType,
} from 'vue';
// import {
//   useI18n,
// } from 'vue-i18n';

import useQueryList from '../../hooks/use-query-list';
import useColumns from '../../hooks/use-columns';
import useFilter from '@/views/resource/resource-manage/hooks/use-filter';

const props = defineProps({
  filter: {
    type: Object as PropType<FilterType>,
  },
});

// use hooks
// const {
//   t,
// } = useI18n();
const {
  datas,
  pagination,
  isLoading,
  handlePageChange,
  handlePageSizeChange,
  handleSort,
} = useQueryList(props, 'route_tables');

const columns = useColumns('route');

const {
  searchData,
  searchValue,
} = useFilter(props);
</script>

<template>
  <bk-loading
    :loading="isLoading"
  >
    <section
      class="flex-row align-items-center mb20 justify-content-end">
      <bk-search-select
        class="w500 ml10"
        clearable
        :conditions="[]"
        :data="searchData"
        v-model="searchValue"
      />
    </section>
    <bk-table
      class="mt20"
      row-hover="auto"
      remote-pagination
      :pagination="pagination"
      :columns="columns"
      :data="datas"
      show-overflow-tooltip
      @page-limit-change="handlePageSizeChange"
      @page-value-change="handlePageChange"
      @column-sort="handleSort"
    />
  </bk-loading>
</template>

<style lang="scss" scoped>
.w100 {
  width: 100px;
}
.w60 {
  width: 60px;
}
.mt20 {
  margin-top: 20px;
}
</style>
