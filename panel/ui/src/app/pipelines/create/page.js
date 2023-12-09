'use client';

import {
  Breadcrumb,
  BreadcrumbItem,
  Column,
  Content,
  Grid,
  Search,
  Theme,
  TextInput,
  NumberInput,
  ContainedList,
  Tag, ExpandableTile, TileAboveTheFoldContent, TileBelowTheFoldContent, InlineNotification, FluidForm,
} from '@carbon/react';
import MainHeader from '@/components/MainHeader/MainHeader';
import React, { useEffect, useRef, useState } from 'react';
import StickyBox from 'react-sticky-box';
import axios from 'axios';
import { Address } from '@/settings';

export default function CreatePipelinePage() {
  const [errorMessage, setErrorMessage] = useState('');

  const [draggingProcessorIndex, setDraggingProcessorIndex] = useState(-1);

  const [usedProcessors, setUsedProcessors] = useState([]);

  const [processors, setProcessors] = useState([]);

  useEffect(() => {
    axios.get(Address + '/api/v1/processors', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('token')}`,
      },
    })
      .then(function(response) {
        if (response.status !== 200) {
          setErrorMessage('Unexpected response from server');
          return;
        }
        console.log(response.data);
        setProcessors(response.data);
      })
      .catch(function(error) {
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
  }, []);

  const draggingItem = useRef();
  const dragOverItem = useRef();

  const handleDragStart = (e, position, source) => {
    draggingItem.current = position;
    draggingItem.source = source;
  };

  const handleDragEnter = (e, position, destination) => {
    dragOverItem.destination = destination;
    dragOverItem.current = position;
  };

  const handleDragEnd = (e) => {
    if (dragOverItem.destination === 'processors') {
      return;
    }

    const listCopy = [...usedProcessors];

    let draggingItemContent;
    if (draggingItem.source === 'processors') {
      draggingItemContent = processors[draggingItem.current];
      if (dragOverItem.current + 1 === listCopy.length) {
        listCopy.push(draggingItemContent);
      } else {
        listCopy.splice(dragOverItem.current, 0, draggingItemContent);
      }
    } else {
      draggingItemContent = listCopy[draggingItem.current];
      listCopy.splice(draggingItem.current, 1);
      listCopy.splice(dragOverItem.current, 0, draggingItemContent);
    }


    draggingItem.current = null;
    dragOverItem.current = null;
    setUsedProcessors(listCopy);
  };

  const categoryToColor = (category) => {
    switch (category) {
      case 'Communication':
        return 'green';
      case 'Flow':
        return 'blue';
      case 'InputOutput':
        return 'red';
      default:
        return '';
    }
  };

  return (
    <>
      {localStorage.getItem('token') ? (
        <div>
          <Theme theme='g100'>
            <MainHeader />
          </Theme>
          <Content>
            <Grid className='create-page' fullWidth>
              <Column
                lg={16}
                md={8}
                sm={4}
                className='create-page_header'
                style={{ marginBottom: '48px' }}
              >
                <Breadcrumb>
                  <BreadcrumbItem>
                    <a href='/'>Home</a>
                  </BreadcrumbItem>
                  <BreadcrumbItem>
                    <a href='/pipelines'>Pipelines</a>
                  </BreadcrumbItem>
                </Breadcrumb>
                <h1 className='create-page__heading'>Create a new pipeline</h1>
              </Column>

              <Column md={4} lg={{ span: 7, offset: 1 }} sm={4}>
                <div
                  style={{
                    paddingBottom: '30px',
                    paddingTop: '30px',
                  }}
                >
                  {
                    usedProcessors.length ?
                      usedProcessors.map((item, index) => (
                        <div onDragStart={(e) => handleDragStart(e, index, 'used_processors')}
                             onDragOver={(e) => e.preventDefault()}
                             onDragEnter={(e) => handleDragEnter(e, index, 'used_processors')}
                             onDragEnd={handleDragEnd}
                             key={index}
                             draggable>
                          <div
                            style={{
                              maxWidth: '400px',
                              padding: '20px 20px',
                              borderTopRightRadius: '10px',
                              borderTopLeftRadius: '10px',
                              background: '#000',
                            }}
                          >
                            <p
                              style={{
                                fontSize: '15px',
                                fontWeight: 'bold',
                                color: 'white',
                              }}
                            >{item.name}</p>
                          </div>
                          {
                            item.arguments.length ? <div
                              style={{
                                maxWidth: '400px',
                                borderBottomRightRadius: '10px',
                                borderBottomLeftRadius: '10px',
                                paddingTop: '20px',
                                paddingBottom: '10px',
                                background: '#f4f4f4',
                                marginBottom: '30px',
                              }}>
                              {
                                item.arguments ? item.arguments.map((arg, index) => (
                                  <div style={{
                                    marginBottom: '20px',
                                    marginLeft: '20px',
                                    marginRight: '20px',
                                  }}>
                                    {
                                      arg.type === 'Text' ?
                                        (
                                          <TextInput type='text' required={arg.required} placeholder={arg.name}
                                                     defaultValue={arg.default} />
                                        ) : arg.type === 'Numeric' ?
                                          (
                                            <NumberInput type='number' required={arg.required} defaultValue={arg.default}
                                                         placeholder={arg.name} />
                                          ) :
                                          <p>'{arg.name}' must be filled by
                                            processors</p>
                                    }

                                    {/*{arg.required ? <p style={{ paddingTop: '3px', color: 'red' }}>*</p> : null}*/}
                                    {/*<p style={{ fontWeight: 'bold', padding: '3px 10px 0 10px' }}>{arg.type}</p>*/}
                                    {/*<p*/}
                                    {/*  style={{ paddingTop: '3px' }}>{arg.name !== 'processable.body' ? arg.name : ''}</p>*/}
                                  </div>
                                )) : null
                              }
                            </div> : <p>No arguments</p>
                          }
                        </div>
                      )) :
                      <div
                        onDragOver={() => {
                          dragOverItem.destination = 'used_processors';
                        }}
                        style={
                          {
                            maxWidth: '400px',
                            height: '100px',
                            border: '2px dashed gray',
                            borderColor: '#d2d2d2',
                            borderRadius: '5px',
                            display: 'flex',
                            alignItems: 'center',
                            justifyContent: 'center',
                          }
                        }
                      >
                        <p>Drop a processor here</p>
                      </div>
                  }
                </div>
              </Column>

              <Column md={4} lg={{ span: 6, offset: 8 }} sm={4}>
                <StickyBox offsetTop={100} offsetBottom={20}>
                  <ContainedList label='Processors' kind='on-page' action={''}>
                    {/*value={searchTerm}*/}
                    {/*onChange={handleChange}*/}
                    <Search placeholder='Filter'
                            closeButtonLabelText='Clear search input' size='lg' labelText='Filter search' />
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
                    {
                      processors &&
                      processors.map((item, index) => (
                        <div
                          onDragStart={(e) => {
                            setDraggingProcessorIndex(index);
                            handleDragStart(e, index, 'processors');
                          }}
                          onDragOver={(e) => e.preventDefault()}
                          onDragEnter={(e) => handleDragEnter(e, index, 'processors')}
                          onDragEnd={handleDragEnd}
                          key={index}
                          draggable
                        >
                          <ExpandableTile
                            style={{
                              paddingLeft: '20px',
                              marginTop: '10px',
                              marginBottom: '10px',
                            }}
                            tileCollapsedIconText='Details'
                            tileExpandedIconText='Details'>
                            <TileAboveTheFoldContent>
                              <div>
                                <h5>{item.name}</h5>
                                <p style={{
                                  fontSize: '14px',
                                }}>{item.description}</p>
                                <Tag type={categoryToColor(item.category)}>
                                  {item.category}
                                </Tag>
                              </div>
                            </TileAboveTheFoldContent>
                            <TileBelowTheFoldContent>

                              <div style={{ marginTop: '10px' }}>
                                <p style={{ fontWeight: 'bold', fontSize: '18px' }}>Arguments</p>
                                {
                                  item.arguments ? item.arguments.map((arg, index) => (
                                    <div style={{ display: 'flex' }}>
                                      {arg.required ? <p style={{ paddingTop: '3px', color: 'red' }}>*</p> : null}
                                      <p style={{ fontWeight: 'bold', padding: '3px 10px 0 10px' }}>{arg.type}</p>
                                      <p
                                        style={{ paddingTop: '3px' }}>{arg.name !== 'processable.body' ? arg.name : ''}</p>
                                    </div>
                                  )) : <p>No arguments</p>
                                }
                              </div>

                              <div style={{ marginTop: '10px' }}>
                                <p style={{ fontWeight: 'bold', fontSize: '18px' }}>Input</p>
                                {
                                  item.input ? item.input.map((arg, index) => (
                                    <div style={{ display: 'flex' }}>
                                      {arg.required ? <p style={{ paddingTop: '3px', color: 'red' }}>*</p> : null}
                                      <p style={{ fontWeight: 'bold', padding: '3px 10px 0 10px' }}>{arg.type}</p>
                                      <p
                                        style={{ paddingTop: '3px' }}>{arg.name !== 'processable.body' ? arg.name : ''}</p>
                                    </div>
                                  )) : <p>No input</p>
                                }
                              </div>

                              <div style={{ marginTop: '10px' }}>
                                <p style={{ fontWeight: 'bold', fontSize: '18px' }}>Output</p>
                                {
                                  item.output ? item.output.map((arg, index) => (
                                    <div style={{ display: 'flex' }}>
                                      {arg.required ? <p style={{ paddingTop: '3px', color: 'red' }}>*</p> : null}
                                      <p style={{ fontWeight: 'bold', padding: '3px 10px 0 10px' }}>{arg.type}</p>
                                      <p
                                        style={{ paddingTop: '3px' }}>{arg.name !== 'processable.body' ? arg.name : ''}</p>
                                    </div>
                                  )) : <p>No output</p>
                                }
                              </div>

                            </TileBelowTheFoldContent>
                          </ExpandableTile>
                        </div>
                      ))
                    }
                  </ContainedList>
                </StickyBox>
              </Column>

            </Grid>
          </Content>
        </div>
      ) : (
        window.location.assign('/auth')
      )}
    </>
  );
}
