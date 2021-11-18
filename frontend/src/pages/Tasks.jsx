import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Typography, Card, CardGroup, Empty, Layout, Nav, Button, Popover } from '@douyinfe/semi-ui';
import { IconHome, IconMoon, IconExit, IconSetting } from '@douyinfe/semi-icons';
import { IllustrationNoContent, IllustrationNoContentDark } from '@douyinfe/semi-illustrations';
import { getAllTask } from '../actions/task';
import { getAccount } from '../actions/auth';

const TASK_STATUS_ASSIGNED = 1;
const TASK_STATUS_COMPLETED = 2;

const Login = () => {
  const dispatch = useDispatch()
  
  useEffect(() => {
    dispatch(getAllTask())
    dispatch(getAccount())

  }, [dispatch])

  const { Header, Footer, Sider, Content } = Layout;
  let currentYear = new Date().getFullYear()

  function switchMode() {
    const body = document.body;
    if (body.hasAttribute('theme-mode')) {
      body.removeAttribute('theme-mode');
    } else {
      body.setAttribute('theme-mode', 'dark');
    }
  }

  function logout() {
    localStorage.removeItem('token')
    window.location.href = '/login';
  }

  const { Paragraph } = Typography;
  const [ spacing ] = useState(12); // setSpacing

  var tasks = useSelector(state => state.task.tasks);

  return (
    <>
    <Layout className='layout'>
        <Sider className='sider'>
          <Nav className='nav-sidebar'
            defaultSelectedKeys={['Home']}
              items={[
                  { itemKey: 'Home', text: 'Tasks', icon: <IconHome size="large" /> },
                  // { itemKey: 'IconInviteStroked', text: 'Billing', icon: <IconInviteStroked size="large" /> },
                  // { itemKey: 'IconLineChartStroked', text: 'Analytics', icon: <IconLineChartStroked size="large" /> },
                  { itemKey: 'Setting', text: 'Setting', icon: <IconSetting size="large" /> },
              ]}
              header={{
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
                <Popover content={ <article style={{ padding: 12 }}>Dark mode</article> }>
                  <Button
                    theme="borderless"
                    onClick={switchMode}
                    icon = {<IconMoon size="large" />}
                    className='icon'
                    color='gray'
                  />
                </Popover>
                {/* <Avatar color='blue' size='small'>YJ</Avatar> */}
                
                <Popover content={ <article style={{ padding: 12 }}>Logout</article> }>
                  <Button
                    onClick={logout}
                    icon = {<IconExit size="large" />}
                    className='icon'
                    color='gray'
                  />
                </Popover>
              </Nav.Footer>
            </Nav>
          </Header>
          <Content className='content'>
            <div className='content-body'>
              { !tasks ?
                <Empty
                  image={<IllustrationNoContent className='empty-image-size' />}
                  darkModeImage={<IllustrationNoContentDark className='empty-image-size' />}
                  description={'No tasks yet'}
                  className='empty-image'
                />
              :
                <CardGroup spacing={spacing}>
                  { tasks.map((task, index) => 
                    <Card 
                      key={index+1}
                      shadows='hover'
                      title={task.id + '. ' + task.description}
                      headerLine={false}
                      style={{ width:260, minHeight: 280 }}
                    >
                    <Paragraph>
                      {task.status === TASK_STATUS_ASSIGNED && <p>Assigned 🔥</p>}
                      {task.status === TASK_STATUS_COMPLETED && <p>Completed 👍</p>}
                    </Paragraph>
                    <Paragraph>{task.public_id}</Paragraph>
                    <Paragraph>{task.created_at}</Paragraph>
                    <Paragraph>{task.description}</Paragraph>
                    
                    <Paragraph>Assigned account id: {task.assigned_account_id}</Paragraph>
                    </Card>
                  )}
                </CardGroup>
              }
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
