'use client';

import {
  Column,
  Content,
  Grid,
  InlineNotification,
  Theme,
} from '@carbon/react';
import MainHeader from '@/components/MainHeader/MainHeader';
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Address } from '@/settings';

export default function HomePage() {
  const [errorMessage, setErrorMessage] = useState('');
  const [statistics, setStatistics] = useState({});

  function loadData() {
    axios
      .get(Address + '/api/v1/statistics', {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
      })
      .then(function (response) {
        if (response.status !== 200) {
          setErrorMessage('Unexpected response from server');
          return;
        }
        setStatistics(response.data);
      })
      .catch(function (error) {
        if (error.response) {
          if (error.response.status === 401) {
            window.location.href = '/auth';
            return;
          }
          setErrorMessage(error.response.data.message);
        } else {
          setErrorMessage('Unknown Error');
          console.error(error);
        }
      });
  }

  useEffect(() => {
    const interval = setInterval(() => {
      loadData();
    }, 3000);
    return () => clearInterval(interval);
  }, []);

  const secToDuration = (sec) => {
    if (!sec) return 'Unavailable';
    let hours = Math.floor(sec / 3600);
    let minutes = Math.floor((sec - hours * 3600) / 60);
    let seconds = sec - hours * 3600 - minutes * 60;
    if (hours < 10) {
      hours = '0' + hours;
    }
    if (minutes < 10) {
      minutes = '0' + minutes;
    }
    if (seconds < 10) {
      seconds = '0' + seconds;
    }
    return hours + ':' + minutes + ':' + seconds;
  };

  function formatBytes(bytes, decimals = 2) {
    if (!+bytes) return '0 Bytes';

    const k = 1024;
    const dm = decimals < 0 ? 0 : decimals;
    const sizes = [
      'Bytes',
      'KiB',
      'MiB',
      'GiB',
      'TiB',
      'PiB',
      'EiB',
      'ZiB',
      'YiB',
    ];

    const i = Math.floor(Math.log(bytes) / Math.log(k));

    return `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))} ${sizes[i]}`;
  }

  function hasLowMem(m) {
    if (m < 200000000) {
      return 'red';
    } else if (m < 500000000) {
      return '#FFB534';
    } else {
      return 'black';
    }
  }

  function avgToColor(avg) {
    if (avg > 1 && avg < 2) {
      return '#FFB534';
    } else if (avg > 2) {
      return 'red';
    } else {
      return '#557C55';
    }
  }

  if (typeof window !== 'undefined')
    if (!localStorage.getItem('token')) window.location.assign('/auth');

  return (
    <>
      <div>
        <Theme theme="g100">
          <MainHeader />
        </Theme>
        <Content>
          <Grid className="panel-page" fullWidth>
            <Column
              lg={16}
              md={8}
              sm={4}
              className="panel-page_header"
              style={{ marginBottom: '48px' }}
            >
              <h1 className="panel-page__heading">Dashboard</h1>

              {errorMessage !== '' && (
                <InlineNotification
                  aria-label="closes notification"
                  kind="error"
                  statusIconDescription="notification"
                  subtitle={errorMessage}
                  onClose={() => {
                    setErrorMessage('');
                  }}
                  style={{ marginTop: '16px' }}
                />
              )}
            </Column>

            <Column
              lg={4}
              md={8}
              sm={4}
              style={{
                marginBottom: '48px',
                background: '#fbfbfb',
                padding: '20px',
              }}
            >
              <div
                style={{
                  display: 'flex',
                  flexDirection: 'column',
                }}
              >
                <h3>Memory Usage</h3>
                <p style={{ color: '#616161', marginBottom: '15px' }}>
                  4 fields found
                </p>
                <div
                  style={{
                    display: 'flex',
                    flexDirection: 'column',
                  }}
                >
                  <p
                    style={{
                      fontSize: '40px',
                      color: hasLowMem(statistics.free_memory),
                    }}
                  >
                    {formatBytes(statistics.free_memory)}
                  </p>
                  <p style={{ fontSize: '18px' }}>
                    {formatBytes(statistics.total_memory)}
                  </p>
                </div>
              </div>
            </Column>

            <Column lg={4} md={8} sm={4} style={{ marginBottom: '48px' }}>
              <div
                style={{
                  display: 'flex',
                  flexDirection: 'column',
                }}
              >
                <div
                  style={{
                    display: 'flex',
                    flexDirection: 'column',
                    marginBottom: '10px',
                    background: '#fbfbfb',
                    padding: '20px',
                  }}
                >
                  <h3>Load Average</h3>
                  <div
                    style={{
                      display: 'flex',
                      flexDirection: 'row',
                      justifyContent: 'space-evenly',
                      paddingTop: '11px',
                    }}
                  >
                    {statistics.load_average_1 && (
                      <p
                        style={{
                          fontSize: '25px',
                          fontWeight: 'bold',
                          color: avgToColor(statistics.load_average_1),
                        }}
                      >
                        {statistics.load_average_1.toFixed(1)}
                      </p>
                    )}
                    {statistics.load_average_5 && (
                      <p
                        style={{
                          fontSize: '25px',
                          fontWeight: 'bold',
                          color: avgToColor(statistics.load_average_5),
                        }}
                      >
                        {statistics.load_average_5.toFixed(1)}
                      </p>
                    )}
                    {statistics.load_average_15 && (
                      <p
                        style={{
                          fontSize: '25px',
                          fontWeight: 'bold',
                          color: avgToColor(statistics.load_average_15),
                        }}
                      >
                        {statistics.load_average_15.toFixed(1)}
                      </p>
                    )}
                  </div>
                </div>

                <div
                  style={{
                    display: 'flex',
                    flexDirection: 'column',
                    background: '#fbfbfb',
                    padding: '20px',
                  }}
                >
                  <h3>Uptime</h3>
                  <div
                    style={{
                      display: 'flex',
                      flexDirection: 'column',
                    }}
                  >
                    <p style={{ fontSize: '40px' }}>
                      {secToDuration(statistics.uptime)}
                    </p>
                  </div>
                </div>
              </div>
            </Column>

            <Column
              lg={4}
              md={8}
              sm={4}
              style={{
                marginBottom: '48px',
                background: '#fbfbfb',
                padding: '20px',
              }}
            >
              <div
                style={{
                  display: 'flex',
                  flexDirection: 'column',
                }}
              >
                <h3>Concurrent Tasks</h3>
                <p style={{ color: '#616161', marginBottom: '15px' }}>
                  Running Goroutines
                </p>
                <div
                  style={{
                    display: 'flex',
                    flexDirection: 'row',
                    justifyContent: 'start',
                    alignItems: 'baseline',
                  }}
                >
                  <p style={{ fontSize: '80px' }}>
                    {statistics.running_goroutines}
                  </p>
                  <p>G</p>
                </div>
              </div>
            </Column>
          </Grid>

          <Grid>
            <Column
              lg={4}
              md={8}
              sm={4}
              style={{
                marginBottom: '48px',
                background: '#fbfbfb',
                padding: '20px',
              }}
            >
              <div
                style={{
                  display: 'flex',
                  flexDirection: 'column',
                }}
              >
                <h3>Connected Clients</h3>
                <p style={{ color: '#616161', marginBottom: '15px' }}>
                  Accepted and authenticated
                </p>
                <div
                  style={{
                    display: 'flex',
                    flexDirection: 'row',
                    justifyContent: 'start',
                    alignItems: 'baseline',
                  }}
                >
                  <p style={{ fontSize: '80px' }}>
                    {statistics.connected_clients}
                  </p>
                  <p style={{ color: '#616161', fontSize: '13px' }}>
                    Connection
                  </p>
                </div>
              </div>
            </Column>

            <Column
              lg={4}
              md={8}
              sm={4}
              style={{
                marginBottom: '48px',
                background: '#fbfbfb',
                padding: '20px',
              }}
            >
              <div
                style={{
                  display: 'flex',
                  flexDirection: 'column',
                }}
              >
                <h3>Waiting Clients</h3>
                <p style={{ color: '#616161', marginBottom: '15px' }}>
                  Waiting to complete authentication
                </p>
                <div
                  style={{
                    display: 'flex',
                    flexDirection: 'row',
                    justifyContent: 'start',
                    alignItems: 'baseline',
                  }}
                >
                  <p style={{ fontSize: '80px' }}>
                    {statistics.waiting_clients}
                  </p>
                  <p style={{ color: '#616161', fontSize: '13px' }}>Request</p>
                </div>
              </div>
            </Column>

            <Column
              lg={4}
              md={8}
              sm={4}
              style={{
                marginBottom: '48px',
                background: '#fbfbfb',
                padding: '20px',
              }}
            >
              <div
                style={{
                  display: 'flex',
                  flexDirection: 'column',
                }}
              >
                <h3>Active Pipelines</h3>
                <p style={{ color: '#616161', marginBottom: '15px' }}>
                  Ready to handle requests
                </p>
                <div
                  style={{
                    display: 'flex',
                    flexDirection: 'row',
                    justifyContent: 'start',
                    alignItems: 'baseline',
                  }}
                >
                  <p style={{ fontSize: '80px' }}>
                    {statistics.active_pipelines}
                  </p>
                  <p style={{ color: '#616161', fontSize: '13px' }}>
                    Pipelines
                  </p>
                </div>
              </div>
            </Column>
          </Grid>
        </Content>
      </div>
    </>
  );
}
