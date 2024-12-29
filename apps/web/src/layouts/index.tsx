import { Outlet } from 'react-router-dom';
import { Button, Flex, Layout, Menu, theme } from 'antd';
import { useState } from 'react';
import { AiOutlineMenuFold, AiOutlineMenuUnfold } from 'react-icons/ai';
import { AiOutlineApi } from 'react-icons/ai';
import { TbApiApp } from 'react-icons/tb';

const { Header, Sider, Content } = Layout;

const Index = () => {
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer, borderRadiusLG }
  } = theme.useToken();
  return (
    <Layout style={{ minHeight: '100vh' }}>
      <Sider
        trigger={null}
        collapsible
        collapsed={collapsed}
        theme="light"
        className={
          collapsed ? '-translate-x-[70px]' : 'translate-x-[0px] !border-none'
        }
      >
        <div className="flex items-center px-[22px] h-16">
          <p className="font-semibold text-xl text-gray-600">Controllers</p>
        </div>
        <Menu
          mode="vertical"
          defaultSelectedKeys={['1']}
          className="!border-none"
          items={[
            {
              key: '1',
              label: <span className="text-gray-700 capitalize">nav 1</span>,
              title: 'nav 1',
              icon: <TbApiApp size={15} color="gray" />,
              children: [
                {
                  key: '1-1',
                  icon: <AiOutlineApi />,
                  label: 'nav 1-1'
                },
                {
                  key: '1-2',
                  icon: <AiOutlineApi />,
                  label: 'nav 1-2'
                }
              ]
            }
          ]}
        />
      </Sider>
      <Layout
        className={collapsed ? 'absolute !w-full top-0 left-0 bottom-0' : ''}
      >
        <Header
          style={{ padding: 0, background: colorBgContainer }}
          className="!w-full"
        >
          <Flex justify="space-between" align="center">
            <Flex justify="between" align="center" gap={5} className="h-full">
              <Button
                type="text"
                icon={
                  collapsed ? <AiOutlineMenuUnfold /> : <AiOutlineMenuFold />
                }
                onClick={() => setCollapsed(!collapsed)}
                style={{
                  fontSize: '20px',
                  width: 64,
                  height: 64
                }}
              />
              <h1 className="font-semibold text-gray-700 text-xl">openAxis</h1>
            </Flex>
            <Flex gap={10} className="mr-5">
              <Button icon={<AiOutlineApi size={15} />}>New API</Button>
              <Button icon={<TbApiApp size={15} />}>New Controller</Button>
            </Flex>
          </Flex>
        </Header>
        <Content
          style={{
            margin: '10px 10px',
            padding: 10,
            minHeight: 280,
            background: colorBgContainer,
            borderRadius: borderRadiusLG
          }}
        >
          <Outlet />
        </Content>
      </Layout>
    </Layout>
  );
};

export default Index;
