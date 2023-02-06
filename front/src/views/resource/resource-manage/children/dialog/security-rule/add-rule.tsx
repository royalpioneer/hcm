import {
  defineComponent,
  ref,
  watch,
  nextTick,
} from 'vue';
import { Table, Input, Select, Button } from 'bkui-vue';
import { ACTION_STATUS, GCP_PROTOCOL_LIST, IP_TYPE_LIST } from '@/constants';
import Confirm from '@/components/confirm';
import {
  useI18n,
} from 'vue-i18n';
import StepDialog from '@/components/step-dialog/step-dialog';
import {
  useResourceStore,
} from '@/store/resource';
const { Option } = Select;


export default defineComponent({
  components: {
    StepDialog,
  },

  props: {
    title: {
      type: String,
    },
    isShow: {
      type: Boolean,
    },
    vendor: {
      type: String,
    },
    loading: {
      type: Boolean,
    },
  },

  emits: ['update:isShow', 'submit'],

  setup(props, { emit }) {
    const {
      t,
    } = useI18n();

    const resourceStore = useResourceStore();

    const cloudTargetSecurityGroup = [{
      id: 'cloud_target_security_group_id',
      name: '安全组',
    }];

    const securityRuleId = ref('');

    const renderSourceAddressSlot = (data: any) => {
      if (data.ipv4_cidr) {
        return <Input v-model={ data.ipv4_cidr }></Input>;
      } if (data.ipv6_cidr) {
        return <Input v-model={ data.ipv6_cidr }></Input>;
      } if (data.cloud_target_security_group_id) {
        return <Input v-model={ data.cloud_target_security_group_id }></Input>;
      }
      return <Input v-model={ data.ipv4_cidr }></Input>;
    };
    const columnsData = [
      { label: t('优先级'),
        field: 'priority',
        render: ({ data }: any) => <Input class="mt25" v-model={ data.priority }></Input>,
      },
      { label: t('策略'),
        field: 'action',
        render: ({ data }: any) => {
          return (
            <Select class="mt25" v-model={data.action}>
                {ACTION_STATUS.map(ele => (
                <Option value={ele.id} label={ele.name} key={ele.id} />
                ))}
          </Select>
          );
        },
      },
      { label: t('协议端口'),
        field: 'port',
        render: ({ data }: any) => {
          return (
                <>
                <Select v-model={data.protocol}>
                    {GCP_PROTOCOL_LIST.map(ele => (
                    <Option value={ele.id} label={ele.name} key={ele.id} />
                    ))}
                </Select>
                <Input v-model={ data.port }></Input>
                </>
          );
        },
      },
      { label: t('类型'),
        field: 'ethertype',
        render: ({ data }: any) => {
          return (
                <>
                <Select v-model={data.ethertype}>
                    {IP_TYPE_LIST.map(ele => (
                    <Option value={ele.id} label={ele.name} key={ele.id} />
                    ))}
                </Select>
                </>
          );
        },
      },
      { label: t('源地址'),
        field: 'id',
        render: ({ data }: any) => {
          return (
                  <>
                  <Select v-model={data.sourceAddress}>
                      {([...IP_TYPE_LIST, ...cloudTargetSecurityGroup]).map(ele => (
                      <Option value={ele.id} label={ele.name} key={ele.id} />
                      ))}
                  </Select>
                  {
                   renderSourceAddressSlot(data)
                  }
                  </>
          );
        },
      },
      { label: t('描述'),
        field: 'memo',
        render: ({ data }: any) => <Input class="mt25" v-model={ data.memo }></Input>,
      },
      { label: t('操作'),
        field: 'operate',
        render: ({ data, row }: any) => {
          return (
                <div class="mt20">
                <Button text theme="primary" onClick={() => {
                  hanlerCopy(data);
                }}>{t('复制')}</Button>
                <Button text theme="primary" class="ml20" onClick={() => {
                  handlerDelete(data, row);
                }}>{t('删除')}</Button>
                </div>
          );
        },
      },
    ];
    const tableData = ref<any>([{}]);
    const columns = ref<any>(columnsData);
    const steps = [
      {
        component: () => <>
            <Table
              class="mt20"
              row-hover="auto"
              columns={columns.value}
              data={tableData.value}
            />
            {securityRuleId.value ? '' : <Button text theme="primary" class="ml20 mt20" onClick={handlerAdd}>{t('新增一条规则')}</Button>}
          </>,
      },
    ];


    // watch(
    //   () => props.vendor,
    //   (v) => {
    //     if (v === 'tcloud' || v === 'aws') {
    //       nextTick(() => {
    //         columns.value = columns.value.filter((e: any) => {
    //           return e.field !== 'priority' && e.field !== 'type';
    //         });
    //       });
    //     }
    //   },
    //   {
    //     immediate: true,
    //   },
    // );

    watch(
      () => props.isShow,
      (v) => {
        if (!v) {
          tableData.value = [{}];
          return;
        }
        columns.value = columnsData;  // 初始化表表格列
        if (props.vendor === 'tcloud' || props.vendor === 'aws') {    // 腾讯云、aws不需要优先级和类型
          nextTick(() => {
            columns.value = columns.value.filter((e: any) => {
              return e.field !== 'priority' && e.field !== 'type';
            });
          });
        }

        // @ts-ignore
        securityRuleId.value = resourceStore.securityRuleDetail?.id;
        if (securityRuleId.value) { // 如果是编辑 则需要将详细数据展示成列表数据
          const sourceAddressData = [...IP_TYPE_LIST, ...cloudTargetSecurityGroup]
            .filter((e: any) => resourceStore.securityRuleDetail[e.id]);
          tableData.value = [{ ...resourceStore.securityRuleDetail, ...{ sourceAddress: sourceAddressData[0].id } }];
          columns.value = columns.value.filter((e: any) => {    // 编辑不能进行复制和删除操作
            return e.field !== 'operate';
          });
        }
      },
      {
        immediate: true,
      },
    );

    // 方法
    const handleClose = () => {
      emit('update:isShow', false);
    };

    const handleConfirm = () => {
      tableData.value.forEach((e: any) => {
        e[e.sourceAddress] = e.ipv4_cidr || e.ipv6_cidr || e.cloud_target_security_group_id;
        if (e.sourceAddress !== 'ipv4_cidr') {
          delete e.ipv4_cidr;
        }
        delete e.sourceAddress;
      });
      // @ts-ignore
      if (securityRuleId.value) {  // 更新
        const {
          id,
          protocol,
          port,
          ipv4_cidr,
          action,
          memo,
        } = tableData.value[0];
        emit('submit', {
          id,
          protocol,
          port,
          ipv4_cidr,
          action,
          memo,
        });
      } else {
        emit('submit', tableData.value);  // 新增
      }
    };

    // 新增
    const handlerAdd = () => {
      tableData.value.push({});
    };

    // 删除
    const handlerDelete = (data: any, row: any) => {
      const index = row.__$table_row_index;
      Confirm('确定删除', '删除之后不可恢复', () => {
        tableData.value.splice(index, 1);
      });
    };

    // 复制
    const hanlerCopy = (data: any) => {
      tableData.value.push(data);
    };

    return {
      steps,
      handleClose,
      handleConfirm,
    };
  },

  render() {
    return <>
        <step-dialog
          title={this.title}
          loading={this.loading}
          isShow={this.isShow}
          steps={this.steps}
          onConfirm={this.handleConfirm}
          onCancel={this.handleClose}
        >
        </step-dialog>
      </>;
  },
});

