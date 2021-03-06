import React from 'react';
import SignUp from './SignUp';
import SignIn from './SignIn';
import { Tabs, TabPane } from '@douyinfe/semi-ui';
import './SignForms.css';
import parrotImage from '../../assets/parrots.jpg'

const SignForms = () => {
  return (
    <>
      <div className="sign-form">
        <img className="sign-form-img" alt="Popug Jira" src={parrotImage} />
        <Tabs type="line">
          <TabPane tab="Sign In" itemKey="1">
            <SignIn/>
          </TabPane>
          <TabPane tab="Sign Up" itemKey="2">
            <SignUp/>
          </TabPane>
        </Tabs>
      </div>
    </>
  );
}

export default SignForms;
