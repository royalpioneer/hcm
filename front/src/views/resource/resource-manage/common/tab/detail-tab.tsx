import {
  defineComponent,
  PropType,
  ref,
} from 'vue';

import './detail-tab.scss';

type Tab = {
  name: string,
  value: string
};

export default defineComponent({
  props: {
    tabs: Array as PropType<Tab[]>,
    active: String as PropType<any>,
  },

  setup(props) {
    const activeTab = ref(props.active || props.tabs[0].value);

    return {
      activeTab,
    };
  },

  render() {
    return <>
      <bk-tab
        v-model:active={this.activeTab}
        type="card"
        class="detail-tab-main"
      >
        {
          this.tabs.map((tab) => {
            return <>
              <bk-tab-panel
                class="g-scroller"
                name={tab.value}
                label={tab.name}
                key={tab.name}
              >
                {
                  this.$slots.default(this.activeTab)
                }
              </bk-tab-panel>
            </>;
          })
        }
      </bk-tab>
    </>;
  },
});
