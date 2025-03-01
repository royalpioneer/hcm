import { defineComponent, reactive, watch, ref, nextTick } from 'vue';
import { RouterLink, RouterView, useRoute, useRouter } from 'vue-router';
import type { RouteRecordRaw } from 'vue-router';
import { Menu, Navigation, Dropdown, Select } from 'bkui-vue';
import { headRouteConfig } from '@/router/header-config';
import Breadcrumb from './breadcrumb';
import workbench from '@/router/module/workbench';
import resource from '@/router/module/resource';
import service from '@/router/module/service';
import business from '@/router/module/business';
import { classes, deleteCookie } from '@/common/util';
import logo from '@/assets/image/logo.png';
import './index.scss';
import { useUserStore, useAccountStore, useCommonStore } from '@/store';
import { useVerify } from '@/hooks';
import { useI18n } from 'vue-i18n';

// import { CogShape } from 'bkui-vue/lib/icon';
// import { useProjectList } from '@/hooks';
// import AddProjectDialog from '@/components/AddProjectDialog';

const { DropdownMenu, DropdownItem } = Dropdown;
const { VERSION } = window.PROJECT_CONFIG;

export default defineComponent({
  setup() {
    const NAV_WIDTH = 240;
    const NAV_TYPE = 'top-bottom';

    const { t } = useI18n();
    const route = useRoute();
    const router = useRouter();
    const userStore = useUserStore();
    const accountStore = useAccountStore();
    const { Option } = Select;

    let topMenuActiveItem = '';
    let menus: RouteRecordRaw[] = [];
    const openedKeys: string[] = [];
    let path = '';
    const curPath = ref('');
    const businessId = ref<number | string>('');
    const businessList = ref<any[]>([]);
    const loading = ref<Boolean>(false);
    const isRouterAlive = ref<Boolean>(true);
    const curYear = ref((new Date()).getFullYear());
    const isMenuOpen = ref<boolean>(true);


    // 获取业务列表
    const getBusinessList = async () => {
      try {
        loading.value = true;
        const res = await accountStore.getBizListWithAuth();
        loading.value = false;
        businessList.value = res?.data;
        if (!businessList.value.length) {      // 没有权限
          router.push({
            name: '403',
            params: {
              id: 'biz_access',
            },
          });
          return;
        }
        businessId.value = accountStore.bizs || res?.data[0].id;   // 默认取第一个业务
        accountStore.updateBizsId(businessId.value); // 设置全局业务id
      } catch (error) {
        console.log(error);
      }
    };

    const changeMenus = (id: string, ...subPath: string[]) => {
      console.log('subPath', subPath, id);
      openedKeys.push(`/${id}`);
      switch (id) {
        case 'business':
          topMenuActiveItem = 'business';
          menus = reactive(business);
          path = '/business/host';
          getBusinessList();    // 业务下需要获取业务列表
          break;
        case 'resource':
          topMenuActiveItem = 'resource';
          menus = reactive(resource);
          path = '/resource/account';
          accountStore.updateBizsId(0); // 初始化业务ID
          break;
        case 'service':
          topMenuActiveItem = 'service';
          menus = reactive(service);
          path = '/service/service-apply';
          break;
        case 'workbench':
          topMenuActiveItem = 'workbench';
          menus = reactive(workbench);
          // path = '/workbench/auto';
          path = '/workbench/audit';
          accountStore.updateBizsId(0); // 初始化业务ID
          break;
        default:
          if (subPath[0] === 'biz_access') {
            topMenuActiveItem = 'business';
            menus = reactive(business);
            path = '/business/host';
          } else {
            topMenuActiveItem = 'resource';
            menus = reactive(resource);
            path = '/resource/account';
          }
          accountStore.updateBizsId(''); // 初始化业务ID
          break;
      }
    };

    watch(
      () => route,
      (val) => {
        const { bizs } = val.query;
        if (bizs) {
          businessId.value = Number(bizs);   // 取地址栏的业务id
          accountStore.updateBizsId(businessId.value); // 设置全局业务id
        }
        curPath.value = route.path;
        const pathArr = val.path.slice(1, val.path.length).split('/');
        changeMenus(pathArr.shift(), ...pathArr);
      },
      {
        immediate: true,
        deep: true,
      },
    );

    watch(() => accountStore.bizs, async (bizs) => {
      if (!bizs) return;
      const commonStore = useCommonStore();
      const { pageAuthData } = commonStore;      // 所有需要检验的查看权限数据
      const bizsPageAuthData = pageAuthData.map((e: any) => {
        // eslint-disable-next-line no-prototype-builtins
        if (e.hasOwnProperty('bk_biz_id')) {
          e.bk_biz_id = bizs;
        }
        return e;
      });
      commonStore.updatePageAuthData(bizsPageAuthData);
      const { getAuthVerifyData } = useVerify();    // 权限中心权限
      await getAuthVerifyData(bizsPageAuthData);
    }, { immediate: true });

    const handleHeaderMenuClick = (id: string, routeName: string): void => {
      if (route.name !== routeName) {
        changeMenus(id);
        router.push({
          path,
        });
      }
    };

    const logout = () => {
      deleteCookie('bk_token');
      deleteCookie('bk_ticket');
      window.location.href = `${window.PROJECT_CONFIG.BK_LOGIN_URL}/?is_from_logout=1&c_url=${window.location.href}`;
    };

    // 选择业务
    const handleChange = async () => {
      accountStore.updateBizsId(businessId.value);    // 设置全局业务id
      // @ts-ignore
      const isbusinessDetail = route.name?.includes('BusinessDetail');
      if (isbusinessDetail) {
        const businessListPath = route.path.split('/detail')[0];
        router.push({
          path: businessListPath,
        });
      } else {
        reload();
      }
    };

    // 切换路由
    const reload = () => {
      isRouterAlive.value = false;
      nextTick(() => {
        isRouterAlive.value = true;
      });
    };

    // 点击
    const handleToggle = (val: any) => {
      isMenuOpen.value = val;
    };

    return () => (
      <main class="flex-column full-page">
        {/* <Header></Header> */}
          <div class="flex-1" >
            {
                <Navigation
                  navigationType={NAV_TYPE}
                  hoverWidth={NAV_WIDTH}
                  defaultOpen={isMenuOpen.value}
                  onToggle={handleToggle}
                >
                  {{
                    'side-header': () => (
                      <div class="left-header flex-row justify-content-between align-items-center">
                        <div class="logo">
                          <img class="logo-icon" src={logo} />
                        </div>
                        <div class="title-text">{t('海垒2.0')}</div>
                      </div>
                    ),
                    header: () => (
                      <header class="bk-hcm-header">
                        <section class="flex-row justify-content-between header-width">
                          {headRouteConfig.map(({ id, route, name, href }) => (
                            <a
                              class={classes({
                                active: topMenuActiveItem === id,
                              }, 'header-title')}
                              key={id}
                              aria-current="page"
                              href={href}
                              onClick={() => handleHeaderMenuClick(id, route)}
                            >
                              {t(name)}
                            </a>
                          ))}
                        </section>
                        <aside class="header-user">
                          <Dropdown
                            trigger='click'
                          >
                            {{
                              default: () => (
                                <span class="cursor-pointer flex-row align-items-center ">
                                  {userStore.username}
                                  <i class={'icon hcm-icon bkhcm-icon-down-shape pl5'}/>
                                </span>
                              ),
                              content: () => (
                                <DropdownMenu>
                                  <DropdownItem onClick={logout}>
                                  {t('退出')}
                                  </DropdownItem>
                                </DropdownMenu>
                              ),
                            }}
                          </Dropdown>
                        </aside>
                      </header>
                    ),
                    menu: () => (
                      <>
                      {topMenuActiveItem === 'business' && isMenuOpen.value
                        ? <Select class="biz-select-warp"
                      v-model={businessId.value}
                      filterable
                      placeholder="请选择业务"
                      onChange={handleChange}>
                        {businessList.value.map(item => (
                            <Option
                                key={item.id}
                                value={item.id}
                                label={item.name}
                            >
                                {item.name}
                            </Option>
                        ))
                        }
                        </Select> : ''}


                      <Menu class="menu-warp" style={`width: ${NAV_WIDTH}px`} uniqueOpen={false} openedKeys={openedKeys} activeKey={route.meta.activeKey as string}>
                        {
                          menus.map(menuItem => (Array.isArray(menuItem.children) ? (
                            <Menu.Submenu
                              key={menuItem.path as string}
                              title={menuItem.name as string}>
                            {{
                              icon: () => <i class={'icon hcm-icon bkhcm-icon-automatic-typesetting menu-icon'}/>,
                              default: () => menuItem.children.map(child => (
                                <RouterLink to={`${child.path}`}>
                                    <Menu.Item key={child.meta.activeKey as string}>
                                      <p class="flex-row flex-1 justify-content-between align-items-center pr16">
                                        <span class="flex-1 text-ov">{child.name as string}</span>
                                      </p>
                                      {/* {route.meta.activeKey} */}
                                    </Menu.Item>
                                  </RouterLink>
                              )),
                            }}
                            </Menu.Submenu>
                          ) : (
                            <RouterLink to={`${menuItem.path}`}>
                              <Menu.Item
                              key={menuItem.meta.activeKey as string}>
                                {/* {menuItem.meta.activeKey} */}
                                {{
                                  // icon: () => <menuItem.icon/>,
                                  default: () => menuItem.name as string,
                                }}
                              </Menu.Item>
                            </RouterLink>
                          )))
                        }
                      </Menu>
                      </>
                    ),
                    default: () => (
                      <>
                        <div class="navigation-breadcrumb">
                            <Breadcrumb></Breadcrumb>
                        </div>
                        <div class={ ['/service/my-apply'].includes(curPath.value) ? 'view-warp no-padding' : 'view-warp'}>
                          {isRouterAlive.value ? <RouterView></RouterView> : ''}
                        </div>
                      </>
                    ),

                    footer: () => (
                      // eslint-disable-next-line max-len
                      <div class="mt20">Copyright © 2012-{ curYear.value } BlueKing - Hybrid Cloud Management System. All Rights Reserved.{ VERSION }</div>
                    ),
                  }}
                </Navigation>
            }
          </div>
        {/* <AddProjectDialog isShow={showAddProjectDialog.value} onClose={toggleAddProjectDialog} /> */}
      </main>
    );
  },
});
