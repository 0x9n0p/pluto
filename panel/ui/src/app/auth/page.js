'use client';

import {
  Button,
  Column,
  Content,
  Grid,
  InlineNotification, Modal,
  TextInput,
  Theme,
} from '@carbon/react';
import axios from 'axios';
import React, { useState } from 'react';
import { Address, getValueFromLocalStorage } from 'src/settings';
import { UnAuthenticatedHeader } from '@/components/MainHeader/MainHeader';

export default function LoginPage() {
  const [openSetupServer, setOpenSetupServer] = useState(getValueFromLocalStorage('base_host') === null && (typeof window !== 'undefined'));
  const [host, setHost] = useState(getValueFromLocalStorage('base_host') ? getValueFromLocalStorage('base_host') : '');
  const [hostAddressError, setHostAddressError] = useState('');

  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const [errorMessage, setErrorMessage] = useState('');

  return (
    <div>
      <Theme theme='g100'>
        <UnAuthenticatedHeader />
      </Theme>
      <Content>
        {openSetupServer ? (
          <Modal
            open
            preventCloseOnClickOutside={true}
            isFullWidth
            modalHeading='Setup Server'
            modalLabel='Server information'
            primaryButtonText='Continue'
            secondaryButtonText='Cancel'
            onRequestSubmit={(event) => {
              if (host === '') {
                setHostAddressError('This field is required');
                return;
              }
              localStorage.setItem('base_host', host);
              setOpenSetupServer(false);
            }}
            onRequestClose={(e) => {
              let h = getValueFromLocalStorage('base_host');
              if (h !== null && h !== '')
                setOpenSetupServer(false);
            }}
          >
            <div
              style={{
                padding: '20px',
              }}
            >
              <TextInput
                invalid={!!hostAddressError}
                invalidText={hostAddressError}
                required={true}
                id='text-input-1'
                labelText='Host Address'
                placeholder='e.g. localhost, 127.0.0.1:443'
                helperText='Domain name or IP address'
                defaultValue={getValueFromLocalStorage('base_host') ? getValueFromLocalStorage('base_host') : ''}
                onChange={(event) => {
                  setHost(event.target.value);
                }}
              />
            </div>
          </Modal>
        ) : null}

        <Grid className='login-page' fullWidth>
          <Column lg={4} md={4} sm={4} className='login-page_main'>
            <h1 style={{ marginBottom: '4px', marginTop: '48px' }}>Log in</h1>
            <p style={{ marginBottom: '48px', color: '#454545' }}>
              Please enter your email and password to log in to the PlutoEngine
              panel.
            </p>

            <TextInput
              id={'email'}
              labelText={'Email'}
              placeholder='Please enter your email address'
              style={{ marginBottom: '16px' }}
              onChange={(event) => {
                setEmail(event.target.value);
              }}
            />

            <TextInput
              id={'password'}
              type={'password'}
              labelText={'Password'}
              placeholder='Please enter your password'
              style={{ marginBottom: '16px' }}
              onChange={(event) => {
                setPassword(event.target.value);
              }}
            />

            {errorMessage !== '' && (
              <InlineNotification
                aria-label='closes notification'
                kind='error'
                statusIconDescription='notification'
                subtitle={errorMessage}
                onClose={() => {
                  setErrorMessage('');
                }}
                style={{ marginBottom: '16px' }}
              />
            )}

            <div
              style={{
                display: 'flex',
              }}
            >
              <Button
                style={{ marginBottom: '24px' }}
                onClick={async (event) => {
                  axios
                    .post(
                      Address() + '/api/v1/auth',
                      {
                        email: email,
                        password: password,
                      },
                      {
                        // TODO: CORS error happens
                        // withCredentials: true,
                        rejectUnauthorized: false,
                      },
                    )
                    .then((response) => {
                      if (response.status === 200) {
                        // TODO: Fix security issues
                        localStorage.setItem('email', response.data.email);
                        localStorage.setItem('token', response.data.token);
                        window.location.href = '/';
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
              <Button
                kind='ghost'
                style={{ marginBottom: '24px', marginLeft: '10px' }}
                onClick={async (event) => {
                  setOpenSetupServer(true);
                }}
              >
                Setup Server
              </Button>
            </div>

            <div style={{ borderTop: '1px solid #D2D2D2' }}></div>
            <p style={{ marginBottom: '16px', marginTop: '16px' }}>
              Alternative logins
            </p>
            {/* TODO: Implement it */}
            <Button kind='tertiary'>Log in with PublicKey</Button>
          </Column>
        </Grid>
      </Content>
    </div>
  );
}
