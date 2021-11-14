import React, { useState } from 'react';
import { BrowserRouter } from 'react-router-dom';
import AppRouter from './components/AppRouter';
import './common.css';

import { Card, Typography, CardGroup, Empty, Layout, Nav, Button, Avatar } from '@douyinfe/semi-ui';
import { IconHome, IconMoon, IconLineChartStroked, IconInviteStroked, IconSetting } from '@douyinfe/semi-icons';
import { IllustrationNoContent, IllustrationNoContentDark } from '@douyinfe/semi-illustrations';



function App() {
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

  const { Text } = Typography;
  const [ spacing ] = useState(12); // setSpacing

  return (
    <>
      <BrowserRouter>
        <AppRouter/>
      </BrowserRouter>
    </>
  );
}

export default App;
