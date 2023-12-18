'use client';

import {
  Breadcrumb,
  BreadcrumbItem,
  Button,
  Column,
  ContainedList,
  Content,
  ExpandableTile,
  Grid,
  InlineNotification,
  Modal,
  Search,
  Tag,
  TextInput,
  Theme,
  TileAboveTheFoldContent,
  TileBelowTheFoldContent,
} from '@carbon/react';
import MainHeader from '@/components/MainHeader/MainHeader';
import React, { useEffect, useRef, useState } from 'react';
import StickyBox from 'react-sticky-box';
import axios from 'axios';
import { Address } from '@/settings';
import { CloseLarge } from '@carbon/icons-react';
import DraggableItems from '@/components/Piipeline/DraggableItems';
import { v4 as uuidv4 } from 'uuid';
import DraggedItems from '@/components/Piipeline/DraggedItems';

// Category color
export const categoryToColor = (category) => {
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

export default function CreatePipelinePage() {
  const [errorMessage, setErrorMessage] = useState('');
  const [errorMessageForSavePipeline, setErrorMessageForSavePipeline] =
    useState('');

  const [usedProcessors, setUsedProcessors] = useState([]);

  const [processors, setProcessors] = useState([]);

  const [pipelineName, setPipelineName] = useState('');
  const [openPipeline, setOpenPipeline] = useState(true);

  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState([]);

  useEffect(() => {
    if (Boolean(searchTerm) && Boolean(processors?.length)) {
      const results = processors.filter((listItem) =>
        listItem.name.toLowerCase().includes(searchTerm.toLowerCase())
      );
      setSearchResults(results);
    }
  }, [searchTerm]);

  useEffect(() => {
    axios
      .get(Address() + '/api/v1/processors', {
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token')}`,
        },
      })
      .then(function (response) {
        if (response.status !== 200) {
          setErrorMessage('Unexpected response from server');
          return;
        }
        const newGenerationOfData = response.data.map((item) => {
          return { ...item, id: uuidv4() };
        });
        setProcessors(newGenerationOfData);
        setSearchResults(newGenerationOfData);
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
  }, []);

  const draggingItem = useRef();
  const dragOverItem = useRef();

  const handleDragStart = (e, position, source) => {
    draggingItem.current = position;
    draggingItem.source = source;
  };

  //  ---------- new ------------
  const onDragStart = ({ id, event, position, source }) => {
    debugger;
    draggingItem.current = position;
    draggingItem.source = source;
  };
  const onDragEnter = ({ id, event, position, destination }) => {
    dragOverItem.destination = destination;
    dragOverItem.current = position;
  };

  // -------------------------

  const handleDragEnter = (e, position, destination) => {
    dragOverItem.destination = destination;
    dragOverItem.current = position;
  };

  const handleDragEnd = ({ event, id }) => {
    if (dragOverItem.destination === 'processors') {
      return;
    }

    const listCopy = [...usedProcessors];

    let draggingItemContent;
    if (draggingItem.source === 'processors') {
      draggingItemContent = processors[draggingItem.current];

      // Create a deep copy of draggingItemContent
      draggingItemContent = JSON.parse(JSON.stringify(draggingItemContent));

      draggingItemContent.arguments.forEach((arg) => {
        if (arg.type === 'Text') {
          arg.value = '';
        } else if (arg.type === 'Numeric') {
          arg.value = 0;
        } else {
          console.log('handleDragEnd: Type not found.', arg.type);
        }
        if (arg.default) {
          arg.value = arg.default;
        }
      });

      listCopy.splice(dragOverItem.current + 1, 0, draggingItemContent);
    } else {
      const draggedItem = listCopy[draggingItem.current];

      // Remove the dragged item
      listCopy.splice(draggingItem.current, 1);

      // Insert the dragged item at the new position
      listCopy.splice(dragOverItem.current, 0, draggedItem);
    }

    draggingItem.current = null;
    dragOverItem.current = null;
    setUsedProcessors(listCopy);
  };

  if (typeof window !== 'undefined')
    if (!localStorage.getItem('token')) window.location.assign('/auth');

  return (
    <>
      <div>
        <Theme theme="g100">
          <MainHeader />
        </Theme>
        <Content>
          {openPipeline ? (
            <Modal
              open
              preventCloseOnClickOutside={true}
              isFullWidth
              modalHeading="Create a new pipeline"
              modalLabel="Pipeline information"
              primaryButtonText="Continue"
              secondaryButtonText="Back to pipelines"
              onRequestSubmit={(event) => {
                if (pipelineName) {
                  setOpenPipeline(false);
                }
              }}
              onRequestClose={(e) => {
                window.location.href = '/pipelines';
              }}
            >
              <div
                style={{
                  padding: '20px',
                }}
              >
                <TextInput
                  required={true}
                  data-modal-primary-focus
                  id="text-input-1"
                  labelText="Pipeline name"
                  placeholder="e.g. LOGIN_USER__V1"
                  style={{
                    marginBottom: '1rem',
                  }}
                  onChange={(event) => setPipelineName(event.target.value)}
                />
              </div>
            </Modal>
          ) : null}

          <Grid className="create-page" fullWidth>
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
                <BreadcrumbItem>
                  <a href="/pipelines">Pipelines</a>
                </BreadcrumbItem>
              </Breadcrumb>
              <Grid fullWidth>
                <Column md={4} lg={{ span: 7, offset: 0 }} sm={4}>
                  <h1 className="create-page__heading">
                    Create a new pipeline
                  </h1>
                </Column>
                <Column md={4} lg={{ span: 1, offset: 13 }} sm={4}>
                  <Button
                    onClick={(event) => {
                      axios
                        .post(
                          Address() + '/api/v1/pipelines',
                          {
                            name: pipelineName,
                            processors: usedProcessors,
                          },
                          {
                            headers: {
                              Authorization: `Bearer ${localStorage.getItem(
                                'token'
                              )}`,
                            },
                          }
                        )
                        .then(function (response) {
                          if (response.status !== 201) {
                            setErrorMessageForSavePipeline(
                              'Unexpected response from server'
                            );
                            return;
                          }
                          window.location.href = '/pipelines';
                        })
                        .catch(function (error) {
                          if (error.response) {
                            setErrorMessageForSavePipeline(
                              error.response.data.message
                            );
                          } else {
                            setErrorMessageForSavePipeline('Unknown Error');
                          }
                        });
                    }}
                  >
                    Save {pipelineName}
                  </Button>
                </Column>
              </Grid>
            </Column>

            <Column md={4} lg={{ span: 7, offset: 1 }} sm={4}>
              {errorMessageForSavePipeline !== '' && (
                <InlineNotification
                  aria-label="closes notification"
                  kind="error"
                  statusIconDescription="notification"
                  subtitle={errorMessageForSavePipeline}
                  onClose={() => {
                    setErrorMessageForSavePipeline('');
                  }}
                  style={{ marginBottom: '16px', maxWidth: '500px' }}
                />
              )}
              {/* Drop zone */}
              <div
                style={{
                  paddingBottom: '30px',
                  paddingTop: '30px',
                }}
              >
                <DraggedItems
                  usedProcessors={usedProcessors}
                  onDragStart={onDragStart}
                  onDragEnter={onDragEnter}
                  onDragEnd={handleDragEnd}
                  setUsedProcessors={setUsedProcessors}
                />

                {!Boolean(usedProcessors?.length) && (
                  <div
                    onDragOver={() => {
                      dragOverItem.destination = 'used_processors';
                    }}
                    style={{
                      maxWidth: '400px',
                      height: '100px',
                      border: '2px dashed gray',
                      borderColor: '#d2d2d2',
                      borderRadius: '5px',
                      display: 'flex',
                      alignItems: 'center',
                      justifyContent: 'center',
                    }}
                  >
                    <p>Drop a processor here</p>
                  </div>
                )}
              </div>
            </Column>

            {/* Draggable Iitems */}
            <Column md={4} lg={{ span: 6, offset: 8 }} sm={4}>
              <StickyBox
                style={{ height: '20px' }}
                offsetTop={100}
                offsetBottom={20}
              >
                <ContainedList label="Processors" kind="on-page" action={''}>
                  <Search
                    placeholder="Filter"
                    closeButtonLabelText="Clear search input"
                    size="lg"
                    labelText="Filter search"
                    value={searchTerm}
                    onChange={(e) => {
                      setSearchTerm(e.target.value);
                    }}
                  />
                  {/* {errorMessage !== '' && (
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
                  )} */}
                  <div
                    style={{
                      height: '80vh',
                      overflowY: 'scroll',
                    }}
                  >
                    <DraggableItems
                      searchResults={searchResults}
                      onDragStart={onDragStart}
                      onDragEnter={onDragEnter}
                      handleDragEnd={handleDragEnd}
                    />
                    {/* <div
                          onDragStart={(e) => {
                            handleDragStart(e, index, 'processors');
                          }}
                          onDragOver={(e) => e.preventDefault()}
                          onDragEnter={(e) =>
                            handleDragEnter(e, index, 'processors')
                          }
                          onDragEnd={handleDragEnd}
                          key={index}
                          draggable>
                          <ExpandableTile
                            style={{
                              paddingLeft: '20px',
                              marginTop: '10px',
                              marginBottom: '10px',
                            }}
                            tileCollapsedIconText="Details"
                            tileExpandedIconText="Details">
                            <TileAboveTheFoldContent>
                              <div>
                                <h5>{item.name}</h5>
                                <p
                                  style={{
                                    fontSize: '14px',
                                  }}>
                                  {item.description}
                                </p>
                                <Tag type={categoryToColor(item.category)}>
                                  {item.category}
                                </Tag>
                              </div>
                            </TileAboveTheFoldContent>
                            <TileBelowTheFoldContent>
                              <div style={{ marginTop: '10px' }}>
                                <p
                                  style={{
                                    fontWeight: 'bold',
                                    fontSize: '18px',
                                  }}>
                                  Arguments
                                </p>
                                {item.arguments ? (
                                  item.arguments.map((arg, index) => (
                                    <div
                                      key={arg.name}
                                      style={{ display: 'flex' }}>
                                      {arg.required ? (
                                        <p
                                          style={{
                                            paddingTop: '3px',
                                            color: 'red',
                                          }}>
                                          *
                                        </p>
                                      ) : null}
                                      <p
                                        style={{
                                          fontWeight: 'bold',
                                          padding: '3px 10px 0 10px',
                                        }}>
                                        {arg.type}
                                      </p>
                                      <p style={{ paddingTop: '3px' }}>
                                        {arg.name !== 'processable.body'
                                          ? arg.name
                                          : ''}
                                      </p>
                                    </div>
                                  ))
                                ) : (
                                  <p>No arguments</p>
                                )}
                              </div>

                              <div style={{ marginTop: '10px' }}>
                                <p
                                  style={{
                                    fontWeight: 'bold',
                                    fontSize: '18px',
                                  }}>
                                  Input
                                </p>
                                {item.input ? (
                                  item.input.map((arg, index) => (
                                    <div
                                      key={arg.name}
                                      style={{ display: 'flex' }}>
                                      {arg.required ? (
                                        <p
                                          style={{
                                            paddingTop: '3px',
                                            color: 'red',
                                          }}>
                                          *
                                        </p>
                                      ) : null}
                                      <p
                                        style={{
                                          fontWeight: 'bold',
                                          padding: '3px 10px 0 10px',
                                        }}>
                                        {arg.type}
                                      </p>
                                      <p style={{ paddingTop: '3px' }}>
                                        {arg.name !== 'processable.body'
                                          ? arg.name
                                          : ''}
                                      </p>
                                    </div>
                                  ))
                                ) : (
                                  <p>No input</p>
                                )}
                              </div>

                              <div style={{ marginTop: '10px' }}>
                                <p
                                  style={{
                                    fontWeight: 'bold',
                                    fontSize: '18px',
                                  }}>
                                  Output
                                </p>
                                {item.output ? (
                                  item.output.map((arg, index) => (
                                    <div
                                      key={arg.name}
                                      style={{ display: 'flex' }}>
                                      {arg.required ? (
                                        <p
                                          style={{
                                            paddingTop: '3px',
                                            color: 'red',
                                          }}>
                                          *
                                        </p>
                                      ) : null}
                                      <p
                                        style={{
                                          fontWeight: 'bold',
                                          padding: '3px 10px 0 10px',
                                        }}>
                                        {arg.type}
                                      </p>
                                      <p style={{ paddingTop: '3px' }}>
                                        {arg.name !== 'processable.body'
                                          ? arg.name
                                          : ''}
                                      </p>
                                    </div>
                                  ))
                                ) : (
                                  <p>No output</p>
                                )}
                              </div>
                            </TileBelowTheFoldContent>
                          </ExpandableTile>
                        </div>
                      ))} */}
                  </div>
                </ContainedList>
              </StickyBox>
            </Column>
          </Grid>
        </Content>
      </div>
    </>
  );
}
