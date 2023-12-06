'use client';

import {
  Button,
  Column,
  Grid,
  InlineNotification,
  TextInput,
} from '@carbon/react';
import axios from 'axios';
import { useState } from 'react';
import { Address } from 'src/settings';

export default function LoginPage() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const [errorMessage, setErrorMessage] = useState('');

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

        {errorMessage !== '' && (
          <InlineNotification
            aria-label="closes notification"
            kind="error"
            statusIconDescription="notification"
            subtitle={errorMessage}
            onClose={() => {
              setErrorMessage('');
            }}
            style={{ marginBottom: '16px' }}
          />
        )}

        <Button
          style={{ marginBottom: '24px' }}
          onClick={async (event) => {
            axios
              .post(
                Address + '/api/v1/auth',
                {
                  email: email,
                  password: password,
                },
                {
                  // TODO: CORS error happens
                  // withCredentials: true,
                  rejectUnauthorized: false,
                }
              )
              .then((response) => {
                if (response.status === 200) {
                  // TODO: Fix security issues
                  localStorage.setItem('email', response.data.email);
                  localStorage.setItem('token', response.data.token);
                } else {
                  setErrorMessage('Unexpected response from server');
                }
              })
              .catch((err) => {
                if (err.response) {
                  setErrorMessage(err.response.data.message);
                } else {
                  setErrorMessage('Unknown Error');
                  console.error(err);
                }
              });
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
