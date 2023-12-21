'use client';

import {
  Breadcrumb,
  BreadcrumbItem,
  Button,
  Column,
  Content,
  Grid,
  InlineNotification,
  OverflowMenu,
  OverflowMenuItem,
  StructuredListBody,
  StructuredListCell,
  StructuredListHead,
  StructuredListRow,
  StructuredListWrapper,
  Theme,
} from '@carbon/react';
import MainHeader from '@/components/MainHeader/MainHeader';
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Address } from '@/settings';
import useLocalStorage from '../../hooks/useLocalStorage';
import { useRouter } from 'next/navigation';

export default function LandingPage() {
  const [errorMessage, setErrorMessage] = useState('');
  const [pipelines, setPipelines] = useLocalStorage('Pipelines', []);
  const router = useRouter();

  const syncPipelines = () => {
    axios
      .get(Address() + '/api/v1/pipelines', {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
      })
      .then(function (response) {
        if (response.status !== 200) {
          setErrorMessage('Unexpected response from server');
          return;
        }
        setPipelines(response.data);
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
  };

  useEffect(() => {
    syncPipelines();
  }, []);

  if (typeof window !== 'undefined')
    if (!localStorage.getItem('token')) window.location.assign('/auth');

  return (
    <>
      <div>
        <Theme theme="g100">
          <MainHeader />
        </Theme>
        <Content>
          <Grid className="pipelines-page" fullWidth>
            <Column
              lg={16}
              md={8}
              sm={4}
              className="create-page_header"
              style={{ marginBottom: '48px' }}
            >
              <Breadcrumb>
                <BreadcrumbItem>
                  <a href="/">Home</a>
                </BreadcrumbItem>
              </Breadcrumb>
              <Grid fullWidth>
                <Column md={4} lg={{ span: 7, offset: 0 }} sm={4}>
                  <h1 className="pipelines-page__heading">Pipelines</h1>
                </Column>
                <Column md={4} lg={{ span: 1, offset: 13 }} sm={4}>
                  <Button
                    onClick={(e) => {
                      window.location.href = '/pipelines/create';
                    }}
                  >
                    New pipeline
                  </Button>
                </Column>
              </Grid>
            </Column>
            <Column lg={16} md={8} sm={4} className="pipelines-page_header">
              {errorMessage !== '' && (
                <InlineNotification
                  aria-label="closes notification"
                  kind="error"
                  statusIconDescription="notification"
                  subtitle={errorMessage}
                  onClose={() => {
                    setErrorMessage('');
                  }}
                  style={{ marginBottom: '16px', maxWidth: '500px' }}
                />
              )}

              {/*// TODO:*/}
              {/*//  Action: Rename, Edit, Delete*/}
              {pipelines.length ? (
                <StructuredListWrapper selection={true}>
                  <StructuredListHead>
                    <StructuredListRow head>
                      <StructuredListCell head>Name</StructuredListCell>
                      <StructuredListCell head>Processors</StructuredListCell>
                      <StructuredListCell head>Status</StructuredListCell>
                      <StructuredListCell head>Created At</StructuredListCell>
                    </StructuredListRow>
                  </StructuredListHead>
                  <StructuredListBody>
                    {pipelines &&
                      pipelines.map((item, index) => (
                        <StructuredListRow
                          key={item.name}
                          onClick={(e) => {
                            console.log('cliecked');
                          }}
                        >
                          <StructuredListCell noWrap>
                            {item?.name}
                          </StructuredListCell>
                          <StructuredListCell>
                            {item.processors?.length}
                          </StructuredListCell>
                          <StructuredListCell>Active</StructuredListCell>
                          <StructuredListCell>
                            {new Date(item?.saved_at).toLocaleString()}
                          </StructuredListCell>
                          <StructuredListCell>
                            <div
                              style={{
                                display: 'flex',
                                justifyContent: 'end',
                              }}
                            >
                              <OverflowMenu
                                flipped={document?.dir === 'rtl'}
                                menuOffset={{ left: -60 }}
                                aria-label="overflow-menu"
                              >
                                <OverflowMenuItem itemText="Remake" disabled />
                                <OverflowMenuItem itemText="Rename" disabled />
                                <OverflowMenuItem
                                  itemText="Deactive"
                                  disabled
                                  requireTitle
                                />
                                <OverflowMenuItem
                                  itemText="Duplicate/Edit"
                                  onClick={(e) => {
                                    debugger;
                                    router.push(`/pipelines/${item?.name}`);
                                  }}
                                />
                                <OverflowMenuItem
                                  hasDivider
                                  isDelete
                                  itemText="Delete"
                                  onClick={(e) => {
                                    axios
                                      .delete(
                                        Address() +
                                          '/api/v1/pipelines?name=' +
                                          item.name,
                                        {
                                          headers: {
                                            Authorization: `Bearer ${localStorage.getItem(
                                              'token'
                                            )}`,
                                          },
                                        }
                                      )
                                      .then(function (response) {
                                        if (response.status !== 200) {
                                          setErrorMessage(
                                            'Unexpected response from server'
                                          );
                                          return;
                                        }
                                        syncPipelines();
                                      })
                                      .catch(function (error) {
                                        if (error.response) {
                                          if (error.response.status === 401) {
                                            window.location.href = '/auth';
                                            return;
                                          }
                                          setErrorMessage(
                                            error.response.data.message
                                          );
                                        } else {
                                          setErrorMessage('Unknown Error');
                                          console.error(error);
                                        }
                                      });
                                  }}
                                />
                              </OverflowMenu>
                            </div>
                          </StructuredListCell>
                        </StructuredListRow>
                      ))}
                  </StructuredListBody>
                </StructuredListWrapper>
              ) : (
                <p>No pipelines</p>
              )}
            </Column>
          </Grid>
        </Content>
      </div>
    </>
  );
}
