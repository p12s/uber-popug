import React from 'react';
import { Card, CardGroup, Empty, Layout, Nav, Button, Avatar } from '@douyinfe/semi-ui';
import { IconHome, IconMoon, IconLineChartStroked, IconInviteStroked, IconSetting } from '@douyinfe/semi-icons';
import { IllustrationNoContent, IllustrationNoContentDark } from '@douyinfe/semi-illustrations';

const Login = () => {
  return (
    <>
    <Layout className='layout'>
        <Sider className='sider'>
          <Nav className='nav-sidebar'
            defaultSelectedKeys={['Home']}
              items={[
                  { itemKey: 'Home', text: 'Home', icon: <IconHome size="large" /> },
                  { itemKey: 'IconInviteStroked', text: 'Billing', icon: <IconInviteStroked size="large" /> },
                  { itemKey: 'IconLineChartStroked', text: 'Analytics', icon: <IconLineChartStroked size="large" /> },
                  { itemKey: 'Setting', text: 'Setting', icon: <IconSetting size="large" /> },
              ]}
              header={{
                  logo: <img src="//lf1-cdn-tos.bytescm.com/obj/ttfe/ies/semi/webcast_logo.svg" alt='Uber Popug Jira'/>,
                  text: 'Popug Jira'
              }}
              footer={{
                  collapseButton: true,
              }}
          />
        </Sider>
        <Layout>
          <Header className='header'>
            <Nav mode='horizontal'>
              <Nav.Footer>
                <Button
                  theme="borderless"
                  onClick={switchMode}
                  icon = {<IconMoon size="large" />}
                  className='icon'
                  color='gray'
                />
                <Avatar color='blue' size='small'>YJ</Avatar>
              </Nav.Footer>
            </Nav>
          </Header>
          <Content className='content'>
            <div className='content-body'>
              <Empty
                  image={<IllustrationNoContent className='empty-image-size' />}
                  darkModeImage={<IllustrationNoContentDark className='empty-image-size' />}
                  description={'No content yet'}
                  className='empty-image'
              />
              <CardGroup spacing={spacing}>
                {
                  new Array(8).fill(null).map((v,idx)=>(
                    <Card 
                      key={idx}
                      shadows='hover'
                      title='Card title'
                      headerLine={false}
                      style={{ width:260 }}
                      headerExtraContent={
                        <Text link>
                          More
                        </Text>
                      }
                    >
                      <Text>Card content</Text>
                    </Card>
                  ))
                }
              </CardGroup>
            </div>
          </Content>
          <Footer className="footer">
            <span className="footer-copyright">
              <span>UberPopug Inc. {currentYear}</span>
            </span>
          </Footer>
        </Layout>  
      </Layout>
    </>
  );
}

export default Login;
