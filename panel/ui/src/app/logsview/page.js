'use client';

import {
  Column,
  Content,
  Grid,
  ListItem,
  OrderedList,
  Theme,
} from '@carbon/react';
import MainHeader from '@/components/MainHeader/MainHeader';
import { useEffect, useRef, useState } from 'react';
import useWebSocket from 'react-use-websocket';
import 'ws';

export default function LogviewPage() {
  const connection = useRef(null);

  useEffect(() => {
    const socket = new WebSocket('wss://panel.localhost/api/v1/logs/bind');

    // Connection opened
    socket.addEventListener('open', (event) => {
      console.log('connected');
      // socket.send("Connection established")
    });

    // Listen for messages
    socket.addEventListener('message', (event) => {
      console.log('Message from server ', event.data);
    });

    connection.current = socket;

    return () => connection.close();
  }, []);

  const [logs, setLogs] = useState([]);

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
            {logs.map((value, index) => {
              return <p key={index}>{value}</p>;
            })}
          </Column>
        </Grid>
      </Content>
    </div>
  );
}
