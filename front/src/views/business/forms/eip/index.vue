<script setup lang="ts">
import { BusinessFormFilter } from '@/typings';
import {
  ref,
} from 'vue';
import FormSelect from '@/views/business/components/form-select.vue';
import aws from './aws.vue';
import azure from './azure.vue';
import gcp from './gcp.vue';
import huawei from './huawei.vue';
import tcloud from './tcloud.vue';
import {
  useBusinessStore,
} from '@/store/business';
import {
  useAccountStore,
} from '@/store/account';

// define emits
const emits = defineEmits(['cancel', 'success']);

// use hooks
const businessStore = useBusinessStore();
const accountStore = useAccountStore();

// define var
const isSubmiting = ref(false);
const formRef = ref(null);
const type = ref('tcloud');
const formData = ref<{
  account_id: string | number,
  region: string | number
}>({
  account_id: '',
  region: ''
});
const componentMap = {
  aws,
  azure,
  gcp,
  huawei,
  tcloud
}

// define method
const handleFormFilter = (value: BusinessFormFilter) => {
  formData.value.account_id = value.account_id;
  formData.value.region = value.region;
  type.value = value.vendor;
}

const handleFormChange = (val: any) => {
  formData.value = val;
}

const handleSubmit = () => {
  isSubmiting.value = true
  formRef
    .value[0]()
    .then(() => {
      return businessStore.addEip(
        accountStore.bizs,
        formData.value
      )
      .then(() => {
        emits('success')
        handleCancel();
      });
    })
    .finally(() => {
      isSubmiting.value = false;
    })
}

const handleCancel = () => {
  emits('cancel')
}
</script>

<template>
  <form-select @change="handleFormFilter"></form-select>
  <component
    class="mb20 pdr20"
    ref="formRef"
    :is="componentMap[type]"
    :region="formData.region"
    :vendor="type"
    @change="handleFormChange"
  />
  <section>
    <bk-button
      theme="primary"
      class="mr10 ml20"
      :loading="isSubmiting"
      @click="handleSubmit"
    >提交创建</bk-button>
    <bk-button
      :disabled="isSubmiting"
      @click="handleCancel"
    >取消</bk-button>
  </section>
</template>

<style lang="scss" scoped>
  .pdr20 {
    padding-right: 20px;
  }
</style>
