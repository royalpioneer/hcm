<script setup lang="ts">
import type {
  FilterType,
} from '@/typings/resource';
import {
  PropType,
  defineExpose,
  h,
  computed,
} from 'vue';
import {
  Button,
  Message,
  InfoBox,
} from 'bkui-vue';
import {
  useResourceStore,
} from '@/store/resource';

import useColumns from '../../hooks/use-columns';
import useQueryList from '../../hooks/use-query-list';
import useFilter from '@/views/resource/resource-manage/hooks/use-filter';

const props = defineProps({
  filter: {
    type: Object as PropType<FilterType>,
  },
  isResourcePage: {
    type: Boolean,
  },
  authVerifyData: {
    type: Object as PropType<any>,
  },
});

const resourceStore = useResourceStore();
const columns = useColumns('subnet');
const {
  datas,
  pagination,
  isLoading,
  handlePageChange,
  handlePageSizeChange,
  handleSort,
} = useQueryList(props, 'subnets');

const hostSearchData = computed(() => {
  return [
    ...searchData.value,
    ...[{
      name: '所属vpc',
      id: 'vpc_id',
    }, {
      name: '云地域',
      id: 'region',
    }],
  ];
});

const {
  searchData,
  searchValue,
} = useFilter(props);

const emit = defineEmits(['auth']);

// 抛出请求数据的方法，新增成功使用
const fetchComponentsData = () => {
  handlePageChange(1);
};

const handleDeleteSubnet = (data: any) => {
  if (data.vendor === 'gcp') {
    const subnetIds = [data.id];
    const getRelateNum = (type: string, field = 'subnet_id', op = 'in') => {
      return resourceStore
        .list(
          {
            page: {
              count: true,
            },
            filter: {
              op: 'and',
              rules: [{
                field,
                op,
                value: subnetIds,
              }],
            },
          },
          type,
        );
    };
    Promise
      .all([
        getRelateNum('cvms', 'subnet_ids', 'json_overlaps'),
        getRelateNum('network_interfaces'),
      ])
      .then(([cvmsResult, networkResult]: any) => {
        if (cvmsResult?.data?.count || networkResult?.data?.count) {
          const getMessage = (result: any, name: string) => {
            if (result?.data?.count) {
              return `${result?.data?.count}个${name}，`;
            }
            return '';
          };
          Message({
            theme: 'error',
            message: `该子网（name：${data.name}，id：${data.id}）关联${getMessage(cvmsResult, 'CVM')}${getMessage(networkResult, '网络接口')}不能删除`,
          });
        } else {
          handledelete(data);
        }
      });
  } else {
    resourceStore
      .countSubnetIps(data.id as string)
      .then((res: any) => {
        if (res?.data?.used_ip_count) {
          Message({
            theme: 'error',
            message: `该子网（name：${data.name}，id：${data.id} IPv4已经被使用${res?.data?.used_ip_count}不能删除`,
          });
        } else {
          handledelete(data);
        }
      });
  }
};

const handledelete = (data: any) => {
  InfoBox({
    title: '请确认是否删除',
    subTitle: `将删除【${data.name || '--'}】`,
    theme: 'danger',
    headerAlign: 'center',
    footerAlign: 'center',
    contentAlign: 'center',
    onConfirm() {
      resourceStore
        .deleteBatch(
          'subnets',
          {
            ids: [data.id],
          },
        ).then(() => {
          Message({
            theme: 'success',
            message: '删除成功',
          });
        });
    },
  });
};

const renderColumns = [
  ...columns,
  {
    label: '操作',
    render({ data }: any) {
      return h(h(
        'span',
        {
          onClick() {
            emit('auth', props.isResourcePage ? 'iaas_resource_delete' : 'biz_iaas_resource_delete');
          },
        },
        [
          h(
            Button,
            {
              text: true,
              theme: 'primary',
              disabled: !props.authVerifyData?.permissionAction[props.isResourcePage ? 'iaas_resource_delete' : 'biz_iaas_resource_delete'],
              onClick() {
                handleDeleteSubnet(data);
              },
            },
            [
              '删除',
            ],
          ),
        ],
      ));
    },
  },
];

defineExpose({ fetchComponentsData });
</script>

<template>
  <bk-loading
    :loading="isLoading"
  >
    <section
      class="flex-row align-items-center"
      :class="isResourcePage ? 'justify-content-end' : 'justify-content-between'">
      <slot>
      </slot>
      <bk-search-select
        class="w500 ml10"
        clearable
        :conditions="[]"
        :data="hostSearchData"
        v-model="searchValue"
      />
    </section>

    <bk-table
      class="mt20"
      row-hover="auto"
      remote-pagination
      :pagination="pagination"
      :columns="renderColumns"
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
