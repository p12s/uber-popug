import React, { useState } from 'react';
import { Form, Button } from '@douyinfe/semi-ui';
import { signIn } from '../../actions/auth';

const SignIn = () => {
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')

  function handleSignIn() {
    signIn(username, password)
  }

  return (
    <>
      <Form layout='vertical'>
        <Form.Input value={username} onChange={setUsername} label='Login' />
        <Form.Input value={password} onChange={setPassword} label='Password' mode='password' />
        <Button onClick={handleSignIn} theme="solid" type="secondary" size="large">Send</Button>
      </Form>
    </>
  );
}

export default SignIn;
