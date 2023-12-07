'use client';

import {
  Column,
  Content,
  Grid,
  InlineNotification,
  Theme,
} from '@carbon/react';
import MainHeader from '@/components/MainHeader/MainHeader';
import { useEffect, useRef, useState } from 'react';
import 'ws';

export default function LogviewPage() {
  // TODO: Move it to the settings.
  const MAX_LOGS = 20;

  const [logs, setLogs] = useState([]);
  const [showErrorNotification, setShowErrorNotification] = useState(false);

  const connection = useRef(null);
  useEffect(() => {
    const socket = new WebSocket(
      'wss://panel.localhost/api/v1/logs/bind/' + localStorage.getItem('token')
    );

    socket.addEventListener('open', (event) => {
      setShowErrorNotification(false);
    });

    socket.addEventListener('error', (event) => {
      setShowErrorNotification(true);
      console.log(event);
    });

    socket.addEventListener('message', (event) => {
      setLogs((prevLogs) => {
        if (prevLogs.length > MAX_LOGS) {
          prevLogs.shift();
        }
        return [...prevLogs, JSON.parse(event.data.toString())];
      });
      console.log(event.data);
    });

    connection.current = socket;

    return () => connection.current.close();
  }, []);

  return (
    <div>
      <Theme theme="g100">
        <MainHeader />
      </Theme>
      <Content>
        <Grid className="create-page" fullWidth>
          <Column
            lg={16}
            md={8}
            sm={4}
            className="create-page_header"
            style={{ marginBottom: '48px' }}
          >
            {showErrorNotification ? (
              <InlineNotification
                title="An error occured"
                subtitle="Please see console logs to debug or refresh this page to try again. Make sure your login token is not expired."
              />
            ) : null}

            {logs.length ? (
              logs.map((value, index) => {
                const d = new Date(value.created_at);

                let levelColor;
                switch (value.level) {
                  case 'Info':
                    levelColor = '#0077b6';
                    break;
                  case 'Warning':
                    levelColor = '#ee9b00';
                    break;
                  case 'Error':
                    levelColor = '#ef233c';
                    break;
                  case 'Panic':
                    levelColor = '#d90429';
                    break;
                  default:
                    levelColor = 'black';
                    break;
                }

                return (
                  <div
                    key={index}
                    style={{
                      marginBottom: '10px',
                    }}
                  >
                    <Grid className="logs" fullWidth>
                      <Column lg={1} md={8} sm={1}>
                        <p>
                          {d.getHours()}:{d.getMinutes()}:{d.getSeconds()}
                        </p>
                      </Column>
                      <Column lg={1} md={8} sm={1}>
                        <p style={{ fontWeight: 'bold', color: levelColor }}>
                          {value.level}
                        </p>
                      </Column>
                      <Column lg={8} md={8} sm={4}>
                        <p>{value.message}</p>
                      </Column>
                      <Column lg={6} md={8} sm={4}>
                        <p style={{ marginLeft: '10px' }}>
                          {JSON.stringify(value.extra)}
                        </p>
                      </Column>
                    </Grid>
                  </div>
                );
              })
            ) : !showErrorNotification ? (
              <p>No log published yet.</p>
            ) : null}
          </Column>
        </Grid>
      </Content>
    </div>
  );
}
