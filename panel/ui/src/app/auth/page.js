'use client';

import { Button, Column, Grid, TextInput } from '@carbon/react';
import axios from 'axios';
import { useState } from 'react';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  return (
    <Grid className="login-page" fullWidth>
      <Column lg={4} md={4} sm={4} className="login-page_main">
        <h1 style={{ marginBottom: '4px' }}>Log in</h1>
        <p style={{ marginBottom: '48px', color: '#454545' }}>
          Please enter your email and password to log in to the PlutoEngine
          panel.
        </p>

        <TextInput
          id={'email'}
          labelText={'Email'}
          style={{ marginBottom: '16px' }}
          onChange={(event) => {
            setEmail(event.target.value);
          }}
        />

        <TextInput
          id={'password'}
          type={'password'}
          labelText={'Password'}
          style={{ marginBottom: '16px' }}
          onChange={(event) => {
            setPassword(event.target.value);
          }}
        />

        <Button
          style={{ marginBottom: '24px' }}
          onClick={async (event) => {
            // TODO
          }}
        >
          Login
        </Button>

        <div style={{ borderTop: '1px solid #D2D2D2' }}></div>
        <p style={{ marginBottom: '16px', marginTop: '16px' }}>
          Alternative logins
        </p>
        {/* TODO: Implement it */}
        <Button kind="tertiary">Log in with PublicKey</Button>
      </Column>
    </Grid>
  );
}
