import React, { useState } from 'react';
import { Form, Button } from '@douyinfe/semi-ui';
import { signUp } from '../../actions/auth';

const SignUp = () => {
  const [name, setName] = useState('')
  const [username, setUsername] = useState('')
  const [password, setPassword] = useState('')

  function handleSignUp() {
    signUp(name, username, password)
  }

  return (
    <>
      <Form layout='vertical'>
        <Form.Input value={name} onChange={setName} label='Name' />
        <Form.Input value={username} onChange={setUsername} label='Login' />
        <Form.Input value={password} onChange={setPassword} label='Password' mode='password' />
        <Button onClick={handleSignUp} theme="solid" type="secondary" size="large">Send</Button>
      </Form>
    </>
  );
}

export default SignUp;
