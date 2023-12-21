import axios from 'axios';
import { forwardRef, useEffect, useImperativeHandle, useState } from 'react';
import { Address } from '../../../settings';
import { Draggable } from 'react-beautiful-dnd';
import {
  ExpandableTile,
  Tag,
  Tile,
  TileAboveTheFoldContent,
  TileBelowTheFoldContent,
} from '@carbon/react';
import uuidv4 from '../../../utils/uuidv4';
import _mock from '../../../_mock/proccessors.json';

// ---------------------
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
//-----------------------

const ProcessorList = forwardRef(
  ({ setErrorMessage, searchTerm, setSearchTerm }, ref) => {
    const [processors, setProcessors] = useState([]);
    const [searchResults, setSearchResults] = useState([]);
    useImperativeHandle(
      ref,
      () => {
        return { searchResults, processors, setProcessors };
      },
      [processors]
    );

    useEffect(() => {
      const data = _mock;
      debugger;
      setProcessors(data);
      setSearchResults(data);
      // axios
      //   .get(Address() + '/api/v1/processors', {
      //     headers: {
      //       Authorization: `Bearer ${localStorage.getItem('token')}`,
      //     },
      //   })
      //   .then(function (response) {
      //     if (response.status !== 200) {
      //       setErrorMessage('Unexpected response from server');
      //       return;
      //     }
      //     const newData = response.data.map((item) => {
      //       return { ...item, id: uuidv4() };
      //     });
      //     debugger;
      //     setProcessors(newData);
      //     setSearchResults(newData);
      //   })
      //   .catch(function (error) {
      //     if (error.response) {
      //       if (error.response.status === 401) {
      //         window.location.href = '/auth';
      //         return;
      //       }
      //       setErrorMessage(error.response.data.message);
      //     } else {
      //       setErrorMessage('Unknown Error');
      //       console.error(error);
      //     }
      //   });
    }, []);

    // useEffect(() => {
    //   const results = processors.filter((listItem) =>
    //     listItem.name.toLowerCase().includes(searchTerm.toLowerCase())
    //   );
    //   setSearchResults(results);
    // }, [searchTerm]);

    return (
      <>
        {' '}
        {searchResults &&
          searchResults.map((item, index) => (
            <Draggable
              key={item?.id}
              draggableId={JSON.stringify(item)}
              index={index}
            >
              {(provided) => (
                <div
                  key={item?.id}
                  style={{
                    paddingLeft: '20px',
                    marginTop: '10px',
                    marginBottom: '10px',
                  }}
                >
                  <Tile
                    {...provided.draggableProps}
                    ref={provided.innerRef}
                    {...provided.dragHandleProps}
                  >
                    <ExpandableTile
                      style={{
                        paddingLeft: '20px',
                        marginTop: '10px',
                        marginBottom: '10px',
                      }}
                      tileCollapsedIconText="Details"
                      tileExpandedIconText="Details"
                    >
                      <TileAboveTheFoldContent>
                        <div>
                          <h5>{item.name}</h5>
                          <p
                            style={{
                              fontSize: '14px',
                            }}
                          >
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
                            }}
                          >
                            Arguments
                          </p>
                          {item.arguments ? (
                            item.arguments.map((arg, index) => (
                              <div
                                key={arg.name}
                                style={{
                                  display: 'flex',
                                }}
                              >
                                {arg.required ? (
                                  <p
                                    style={{
                                      paddingTop: '3px',
                                      color: 'red',
                                    }}
                                  >
                                    *
                                  </p>
                                ) : null}
                                <p
                                  style={{
                                    fontWeight: 'bold',
                                    padding: '3px 10px 0 10px',
                                  }}
                                >
                                  {arg.type}
                                </p>
                                <p
                                  style={{
                                    paddingTop: '3px',
                                  }}
                                >
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
                            }}
                          >
                            Input
                          </p>
                          {item.input ? (
                            item.input.map((arg, index) => (
                              <div
                                key={arg.name}
                                style={{
                                  display: 'flex',
                                }}
                              >
                                {arg.required ? (
                                  <p
                                    style={{
                                      paddingTop: '3px',
                                      color: 'red',
                                    }}
                                  >
                                    *
                                  </p>
                                ) : null}
                                <p
                                  style={{
                                    fontWeight: 'bold',
                                    padding: '3px 10px 0 10px',
                                  }}
                                >
                                  {arg.type}
                                </p>
                                <p
                                  style={{
                                    paddingTop: '3px',
                                  }}
                                >
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
                            }}
                          >
                            Output
                          </p>
                          {item.output ? (
                            item.output.map((arg, index) => (
                              <div
                                key={arg.name}
                                style={{
                                  display: 'flex',
                                }}
                              >
                                {arg.required ? (
                                  <p
                                    style={{
                                      paddingTop: '3px',
                                      color: 'red',
                                    }}
                                  >
                                    *
                                  </p>
                                ) : null}
                                <p
                                  style={{
                                    fontWeight: 'bold',
                                    padding: '3px 10px 0 10px',
                                  }}
                                >
                                  {arg.type}
                                </p>
                                <p
                                  style={{
                                    paddingTop: '3px',
                                  }}
                                >
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
                  </Tile>
                </div>
              )}
            </Draggable>
          ))}
      </>
    );
  }
);

export default ProcessorList;
